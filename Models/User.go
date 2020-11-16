package Models

import "github.com/brianvoe/gofakeit"

type User struct {
	ID       string                `fake:"{uuid}"`
	Name     string                `fake:"{firstname}"`
	LastName string                `fake:"{lastname}"`
	Email    string                `fake:"{email}"`
	Gender   string                `fake:"{gender}"`
	Phone    string                `fake:"{phone}"`
	Adress   *gofakeit.AddressInfo `fake:"{AdressInfo}"`
}
