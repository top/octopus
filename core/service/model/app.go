package model

type App struct {
	Model
	From string `json:"from" form:"from"`
	To   string `json:"to" form:"to"`
}
