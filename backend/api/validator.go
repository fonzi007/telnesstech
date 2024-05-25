package api

import (
	"strings"
	"telnesstech/errors"
	"telnesstech/model"
)

var (
	gradeEnums = []string{"0", "1", "2", "3", "4", "5"}
)

type Validator struct{}

func (v *Validator) validateFilterRequest(request *model.PhoneNumberSearchRequest) error {
	errorMessage := ""
	valid := isValidGrade(request.FilterBy.Grade)
	if valid != true {
		errorMessage = errorMessage + "\ngrade should be between 0 and 5"
	}
	valid = isValidType(request.FilterBy.Type)
	if valid != true {
		errorMessage = errorMessage + "\ntype should be CELL or FIX"
	}
	if errorMessage != "" {
		return errors.RequestInvalidError("Invalid request: " + errorMessage)
	}
	return nil
}

func isValidGrade(grade string) bool {
	if grade == "" {
		return true
	}
	for _, gradeEnum := range gradeEnums {
		if grade == gradeEnum {
			return true
		}
	}
	return false
}

func isValidType(requestType string) bool {
	if requestType == "" {
		return true
	}
	if strings.ToUpper(requestType) == "CELL" || strings.ToUpper(requestType) == "FIX" {
		return true
	}
	return false
}
