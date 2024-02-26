package dto

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
	Email       string `json:"email"`
	Phone       string `json:"phone_number"`
	Address     string `json:"adress"`
	Password    string `json:"password"`
}
