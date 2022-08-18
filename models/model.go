package models

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Telefone    int64  `json:"telefone"`
	Email       string `json:"email"`
	DataCriação string `json:"data_criacao"`
}
