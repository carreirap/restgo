package main

type person struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Cpf     string  `json:"cpf"`
	Age     int     `json:"age"`
	Address address `json:"address"`
}

type address struct {
	Street  string `json:"street"`
	Number  int    `json:"number"`
	ZipCode string `json:"zip_code"`
}
