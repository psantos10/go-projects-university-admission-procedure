package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type Applicant struct {
	firstName     string
	lastName      string
	physics       float64
	chemistry     float64
	math          float64
	computer      float64
	admissionExam float64
	firstOption   string
	secondOption  string
	thirdOption   string
}

func sortApplicantsByDepartment(applicants []Applicant, department string) {
	sort.Slice(applicants, func(i, j int) bool {
		scoreI, scoreJ := getScore(applicants[i], department), getScore(applicants[j], department)
		if scoreI != scoreJ {
			return scoreI > scoreJ
		}
		if applicants[i].firstName != applicants[j].firstName {
			return applicants[i].firstName < applicants[j].firstName
		}
		return applicants[i].lastName < applicants[j].lastName
	})
}

func getScore(applicant Applicant, department string) float64 {
	var score float64
	switch department {
	case "Physics":
		score = (applicant.physics + applicant.math) / 2.0
	case "Chemistry":
		score = applicant.chemistry
	case "Mathematics":
		score = applicant.math
	case "Engineering":
		score = (applicant.computer + applicant.math) / 2.0
	case "Biotech":
		score = (applicant.chemistry + applicant.physics) / 2.0
	default:
		return applicant.admissionExam
	}

	return math.Max(score, applicant.admissionExam)
}

func sortApplicantsByPriority(applicants []Applicant, priorityFunc func(Applicant) string) {
	sort.Slice(applicants, func(i, j int) bool {
		departmentI, departmentJ := priorityFunc(applicants[i]), priorityFunc(applicants[j])
		scoreI, scoreJ := getScore(applicants[i], departmentI), getScore(applicants[j], departmentJ)
		if scoreI != scoreJ {
			return scoreI > scoreJ
		}
		if applicants[i].firstName != applicants[j].firstName {
			return applicants[i].firstName < applicants[j].firstName
		}
		return applicants[i].lastName < applicants[j].lastName
	})
}

func uniqueApplicants(applicants []Applicant, departments map[string][]Applicant) []Applicant {
	filtered := make([]Applicant, 0, len(applicants))
	for _, applicant := range applicants {
		found := false
		for _, dept := range departments {
			for _, a := range dept {
				if a.firstName == applicant.firstName && a.lastName == applicant.lastName {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			filtered = append(filtered, applicant)
		}
	}
	return filtered
}

func main() {
	var maxNumberOfStudentsPerDepartment int
	fmt.Scan(&maxNumberOfStudentsPerDepartment)

	applicantsFile, err := os.Open("applicants.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer applicantsFile.Close()

	scanner := bufio.NewScanner(applicantsFile)
	var applicants []Applicant

	for scanner.Scan() {
		var applicant Applicant
		fmt.Sscan(scanner.Text(), &applicant.firstName, &applicant.lastName, &applicant.physics, &applicant.chemistry, &applicant.math, &applicant.computer, &applicant.admissionExam, &applicant.firstOption, &applicant.secondOption, &applicant.thirdOption)
		applicants = append(applicants, applicant)
	}

	departments := map[string][]Applicant{
		"Biotech":     {},
		"Chemistry":   {},
		"Engineering": {},
		"Mathematics": {},
		"Physics":     {},
	}

	for _, priorityFunc := range []func(applicant Applicant) string{
		func(applicant Applicant) string { return applicant.firstOption },
		func(applicant Applicant) string { return applicant.secondOption },
		func(applicant Applicant) string { return applicant.thirdOption },
	} {
		sortApplicantsByPriority(applicants, priorityFunc)

		var remainingApplicants []Applicant

		for _, applicant := range applicants {
			department := priorityFunc(applicant)
			if len(departments[department]) < maxNumberOfStudentsPerDepartment {
				departments[department] = append(departments[department], applicant)
			} else {
				remainingApplicants = append(remainingApplicants, applicant)
			}
		}

		applicants = uniqueApplicants(applicants, departments)
	}

	for _, department := range []string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"} {
		admittedApplicants := departments[department]
		sortApplicantsByDepartment(admittedApplicants, department)

		filename := fmt.Sprintf("%s.txt", strings.ToLower(department))
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Failed to create file: %s", err)
		}

		for _, applicant := range admittedApplicants {
			_, err := fmt.Fprintf(file, "%s %s %.2f\n", applicant.firstName, applicant.lastName, getScore(applicant, department))
			if err != nil {
				log.Fatalf("Failed to write to file: %s", err)
			}
		}
		file.Close()
	}
}
