package model

type Student struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   int `json:"age"`
	Grade int `json:"grade"`
}