package api

import (
	"github.com/stretchr/testify/assert"
	"telnesstech/model"
	"testing"
	"time"
)

func Test_convertResponse(t *testing.T) {
	tc := []struct {
		In    *model.PhoneNumber
		Out   model.PhoneNumberResponse
		Error error
	}{
		{
			In: &model.PhoneNumber{
				Msisdn: "+46652412359",
			},
			Out: model.PhoneNumberResponse{
				Msisdn: "+46652412359",
			},
		},
		{
			In: &model.PhoneNumber{
				Msisdn:     "+33652412359",
				Grade:      "2",
				Type:       "CELL",
				ReservedAt: &time.Time{},
			},
			Out: model.PhoneNumberResponse{
				Msisdn:     "+33652412359",
				Grade:      "2",
				Type:       "CELL",
				ReservedAt: &time.Time{},
			},
		},
		{
			In: &model.PhoneNumber{
				Msisdn:     "0652412359",
				Grade:      "3",
				Type:       "CELL",
				ReservedAt: &time.Time{},
				ExpiresAt:  &time.Time{},
			},
			Out: model.PhoneNumberResponse{
				Msisdn:     "0652412359",
				Grade:      "3",
				Type:       "CELL",
				ReservedAt: &time.Time{},
				ExpiresAt:  &time.Time{},
			},
		},
	}

	for _, testCase := range tc {
		out := ConvertToResponse(*testCase.In)
		assert.Equal(t, testCase.Out, out)
	}
}
