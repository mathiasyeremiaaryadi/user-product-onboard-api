package dto

type UserPostRequestBody struct {
	PersonalNumber string `json:"personalNumber"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Name           string `json:"name"`
}

type UserPutRequestBody struct {
	PersonalNumber string         `json:"personalNumber"`
	Password       string         `json:"password"`
	Email          string         `json:"email"`
	Name           string         `json:"name"`
	Active         bool           `json:"active"`
	Role           map[string]int `json:"role"`
}
