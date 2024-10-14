package models

type Restaurant struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	Cuisine string `json:"Cuisine"`
	Address string `json:"Address"`
}

type Dish struct {
	Name        string
	Description string
	Restaurant  Restaurant
}

type Review struct {
	Comment string
	Dish    Dish
}

type Ranking struct {
	Previous Review
	Next     Review
	Review   Review
}
