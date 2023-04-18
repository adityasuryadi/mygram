package model

type UpdateRequestRole struct {
	Name        string `json:"name"`
	Permissions []int  `json:"permissions"`
}

type ResponseRole struct {
	Name        string `json:"name"`
	Permissions []int  `json:"permissions"`
}