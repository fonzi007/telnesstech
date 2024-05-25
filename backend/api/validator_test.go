package api

import (
	"github.com/stretchr/testify/assert"
	"telnesstech/model"
	"testing"
)

func TestValidator(t *testing.T) {
	tc := []struct {
		Name          string
		request       *model.PhoneNumberSearchRequest
		expectedError bool
	}{
		{
			Name: "Should throw on invalid grade",
			request: &model.PhoneNumberSearchRequest{
				FilterBy: model.FilterRequest{
					Grade: "8",
				},
			},
			expectedError: true,
		},
		{
			Name: "Should throw on invalid type",
			request: &model.PhoneNumberSearchRequest{
				FilterBy: model.FilterRequest{
					Type: "invalid type",
				},
			},
			expectedError: true,
		},
		{
			Name: "Should throw on invalid type and invalid grade",
			request: &model.PhoneNumberSearchRequest{
				FilterBy: model.FilterRequest{
					Grade: "8",
					Type:  "invalid type",
				},
			},
			expectedError: true,
		},
		{
			Name: "Should not throw on empty request",
			request: &model.PhoneNumberSearchRequest{
				FilterBy: model.FilterRequest{},
			},
			expectedError: false,
		},
		{
			Name: "Should not throw on valid grade",
			request: &model.PhoneNumberSearchRequest{
				FilterBy: model.FilterRequest{
					Grade: "4",
				},
			},
			expectedError: false,
		},
		{
			Name: "Should not throw on valid type",
			request: &model.PhoneNumberSearchRequest{
				FilterBy: model.FilterRequest{
					Type: "cell",
				},
			},
			expectedError: false,
		},
	}
	validator := &Validator{}
	for _, testCase := range tc {
		t.Run(testCase.Name, func(t *testing.T) {
			err := validator.validateFilterRequest(testCase.request)
			if testCase.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
