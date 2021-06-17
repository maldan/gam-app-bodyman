package caloryman

import "time"

type FoodComponent struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Food struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	Fat          string `json:"fat"`
	Protein      string `json:"protein"`
	Carbohydrate string `json:"carbohydrate"`

	MineralList []FoodComponent `json:"mineralList"`
	VitaminList []FoodComponent `json:"vitaminList"`
	OtherList   []FoodComponent `json:"otherList"`
}

type Eat struct {
	Id        string    `json:"id"`
	ProductId string    `json:"name"`
	Amount    string    `json:"amount"`
	Created   time.Time `json:"created"`
}
