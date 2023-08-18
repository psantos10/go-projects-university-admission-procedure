package main

import (
	"fmt"
	"sort"
)

type Applicant struct {
	firstName string
	lastName  string
	gpa       float64
}

func main() {
	var totalNumberOfApplicants int
	fmt.Scan(&totalNumberOfApplicants)

	var numberOfApplicantsToBeAccepted int
	fmt.Scan(&numberOfApplicantsToBeAccepted)

	var applicants []Applicant
	for i := 0; i < totalNumberOfApplicants; i++ {
		var applicant Applicant

		fmt.Scan(&applicant.firstName, &applicant.lastName, &applicant.gpa)

		applicants = append(applicants, applicant)
	}

	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].gpa == applicants[j].gpa {
			if applicants[i].firstName == applicants[j].firstName {
				return applicants[i].lastName < applicants[j].lastName
			}

			return applicants[i].firstName < applicants[j].firstName
		}

		return applicants[i].gpa > applicants[j].gpa
	})

	fmt.Println("Successful applicants:")
	for i := 0; i < numberOfApplicantsToBeAccepted; i++ {
		fmt.Println(applicants[i].firstName, applicants[i].lastName)
	}
}
