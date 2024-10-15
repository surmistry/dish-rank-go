package models

type Dish struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	RestaurantId  string `json:"RestaurantId"`
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
