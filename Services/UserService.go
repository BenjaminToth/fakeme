package Services

import (
	"fakeme/Models"
	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func GetUsers(c *gin.Context) {

	var users []Models.User
	counter := c.Query("count")
	var counterint int
	var err error

	if counter != "" {
		counterint, err = strconv.Atoi(counter)
		if err != nil {
			log.Fatal(err)
		}
		users = make([]Models.User, counterint)
	} else {
		counterint = 100
		users = make([]Models.User, 100)
	}

	for i := 0; i < counterint; i++ {
		users[i].ID = gofakeit.UUID()
		users[i].Name = gofakeit.Name()
		users[i].LastName = gofakeit.LastName()
		users[i].Email = gofakeit.Email()
		users[i].Gender = gofakeit.Gender()
		users[i].Phone = gofakeit.Phone()
		users[i].Adress = gofakeit.Address()
	}
	c.JSON(200, users)

}
