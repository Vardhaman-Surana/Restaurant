package models

type Admin struct{
	Name string `json:"Name"`
	Password string	`json:"Password"`
	IsSuper int	`json:"IsSuper"`
}