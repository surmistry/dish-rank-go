package models

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
