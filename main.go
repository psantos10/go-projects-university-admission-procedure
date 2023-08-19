package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Applicant struct {
	firstName    string
	lastName     string
	gpa          float64
	firstOption  string
	secondOption string
	thirdOption  string
}

func sortApplicants(applicants []Applicant) {
	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].gpa != applicants[j].gpa {
			return applicants[i].gpa > applicants[j].gpa
		}
		if applicants[i].firstName != applicants[j].firstName {
			return applicants[i].firstName < applicants[j].firstName
		}
		return applicants[i].lastName < applicants[j].lastName
	})
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
		fmt.Sscan(scanner.Text(), &applicant.firstName, &applicant.lastName, &applicant.gpa, &applicant.firstOption, &applicant.secondOption, &applicant.thirdOption)
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
		sortApplicants(applicants)

		var remainingApplicants []Applicant

		for _, applicant := range applicants {
			department := priorityFunc(applicant)
			if len(departments[department]) < maxNumberOfStudentsPerDepartment {
				departments[department] = append(departments[department], applicant)
			} else {
				remainingApplicants = append(remainingApplicants, applicant)
			}
		}

		applicants = remainingApplicants
	}

	for _, department := range []string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"} {
		fmt.Println(department)

		admittedApplicants := departments[department]
		sortApplicants(admittedApplicants)

		for _, applicant := range admittedApplicants {
			fmt.Printf("%s %s %.2f\n", applicant.firstName, applicant.lastName, applicant.gpa)
		}
		fmt.Println()
	}
}
