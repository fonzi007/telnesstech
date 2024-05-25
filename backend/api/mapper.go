package api

import "telnesstech/model"

func mapToResponse(in []model.PhoneNumber) []model.PhoneNumberResponse {
	response := make([]model.PhoneNumberResponse, 0)
	for _, number := range in {
		response = append(response, ConvertToResponse(number))
	}
	return response
}

func ConvertToResponse(in model.PhoneNumber) model.PhoneNumberResponse {
	return model.PhoneNumberResponse{
		Msisdn:     in.Msisdn,
		Grade:      in.Grade,
		Type:       in.Type,
		ReservedAt: in.ReservedAt,
		ExpiresAt:  in.ExpiresAt,
	}
}
