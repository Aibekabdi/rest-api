package models

type SubstrModel struct {
	Text string `json:"text"`
}

type Email []struct {
	Email string `json:"Email"`
}

type CounterModel struct {
	Num int `json:"num"`
}
