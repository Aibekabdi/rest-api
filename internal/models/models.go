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

type UserModel struct {
	Id        int    `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}
