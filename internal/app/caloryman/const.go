package caloryman

import "time"

type Component struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Food struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	Fat          string `json:"fat"`
	Protein      string `json:"protein"`
	Carbohydrate string `json:"carbohydrate"`

	ComponentIdList []string `json:"componentIdList"`
}

type Eat struct {
	Id        string    `json:"id"`
	ProductId string    `json:"name"`
	Amount    string    `json:"amount"`
	Created   time.Time `json:"created"`
}

type IdArgs struct {
	Id string
}
