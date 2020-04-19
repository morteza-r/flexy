package main

type User struct {
	Id      float64 `json:"id"`
	Name    string  `json:"name"`
	Family  string  `json:"family"`
	Active  bool    `json:"active"`
	Account Account `json:"account"`
}

type Account struct {
	Name   string   `json:"name"`
	Credit int64    `json:"credit"`
	Tags   []string `json:"tags"`
}
