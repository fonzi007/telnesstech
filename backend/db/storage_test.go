package db

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"telnesstech/model"
	"testing"
	"time"
)

func Test_SearchBy(t *testing.T) {
	db, _ := sql.Open("mysql", "root:changeme@tcp(localhost:3306)/testDB?parseTime=true")
	defer db.Close()
	storage, _ := NewStorage(db)

	tc := []struct {
		In    *model.PhoneNumberSearchRequest
		Out   []model.PhoneNumber
		Error error
	}{
		{
			In: &model.PhoneNumberSearchRequest{
				Prefix: "+336",
			},
			Out: []model.PhoneNumber{
				{
					Msisdn:     "+33652412359",
					Grade:      "0",
					Type:       "CELL",
					ReservedAt: nil,
					ExpiresAt:  nil,
				},
			},
		},
		{
			In: &model.PhoneNumberSearchRequest{
				FilterBy: model.FilterRequest{
					Type: "FIX",
				},
			},
			Out: []model.PhoneNumber{
				{
					Msisdn:     "0152412359",
					Grade:      "2",
					Type:       "FIX",
					ReservedAt: timePtr(time.Now().UTC()),
					ExpiresAt:  timePtr(time.Now().UTC()),
				},
			},
		},
        {
            In: &model.PhoneNumberSearchRequest{
                Suffix: "59"
            },
            Out: []model.PhoneNumber{
                {
                    Msisdn:     "+33652412359",
                    Grade:      "0",
                    Type:       "CELL",
                    ReservedAt: nil,
                    ExpiresAt:  nil,
                },
                {
                    Msisdn:     "0152412359",
                    Grade:      "2",
                    Type:       "FIX",
                    ReservedAt: timePtr(time.Now().UTC()),
                    ExpiresAt:  timePtr(time.Now().UTC()),
                },
                {
                    Msisdn:     "0652412359",
                    Grade:      "1",
                    Type:       "CELL",
                    ReservedAt: timePtr(time.Now().UTC()),
                    ExpiresAt:  timePtr(time.Now().UTC()),
                },
            },
        },
	}

	for _, testCase := range tc {
		out, _ := storage.FindBy(*testCase.In)
		assert.Equal(t, testCase.Out, out)
	}
}

func timePtr(t time.Time) *time.Time {
   timeTruncated := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
   return &timeTruncated

}