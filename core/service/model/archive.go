package model

type Archive struct {
	Model
	Path string `json:"path"`
	Type string `json:"type"`
}
