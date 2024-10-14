package models

type Restaurant struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	Cuisine string `json:"Cuisine"`
	Address string `json:"Address"`
}
