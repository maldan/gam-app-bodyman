package caloryman

import "time"

type Component struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Product struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	Protein      string `json:"protein"`
	Carbohydrate string `json:"carbohydrate"`
	Fat          string `json:"fat"`

	ComponentIdList []string `json:"componentIdList"`
}

type Eat struct {
	Id        string    `json:"id"`
	ProductId string    `json:"productId"`
	Amount    string    `json:"amount"`
	Calory    float64   `json:"calory"`
	Product   Product   `json:"product"`
	Created   time.Time `json:"created"`
}

type Exercise struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	MuscleList []string `json:"muscleList"`
	Tool       string   `json:"tool"`
}

type Training struct {
	Id         string   `json:"id"`
	ExerciseId string   `json:"exerciseId"`
	Exercise   Exercise `json:"exercise"`

	//Title      string   `json:"title"`
	//MuscleList []string `json:"muscleList"`
	//Tool       string   `json:"tool"`

	Reps     int       `json:"reps"`
	Weight   int       `json:"weight"`
	Distance int       `json:"distance"`
	Duration int       `json:"duration"`
	Created  time.Time `json:"created"`
}

type Note struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
}

type Weight struct {
	Id      string    `json:"id"`
	Value   float32   `json:"value"`
	Created time.Time `json:"created"`
}

type IdArgs struct {
	Id string
}

type DateArgs struct {
	Date time.Time
}

type NameArgs struct {
	Name string
}
