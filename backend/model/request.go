package model

type PhoneNumberSearchRequest struct {
	Suffix    string        `json:"suffix"`
	Prefix    string        `json:"prefix"`
	Substring string        `json:"substring"`
	FilterBy  FilterRequest `json:"filter_by"`
}

type FilterRequest struct {
	Grade string `json:"grade"`
	Type  string `json:"type"`
}
