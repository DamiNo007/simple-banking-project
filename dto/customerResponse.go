package dto

type CustomerResponseList []*CustomerResponse

type CustomerResponse struct {
	Id          string `json:"id" xml:"id"`
	Name        string `json:"fullName" xml:"fullName"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode"`
	DateOfBirth string `json:"dateOfBirth" xml:"dateOfBirth"`
	Status      string `json:"status" xml:"status"`
}
