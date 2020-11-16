package Services

import (
	"fakeme/Models"
	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func GetPasswords(c *gin.Context) {
	minpasslength := 10
	maxpasslength := 20
	var passwords []Models.Auth
	counter := c.Query("count")
	var counterint int
	var err error

	if counter != "" {
		counterint, err = strconv.Atoi(counter)
		if err != nil {
			log.Fatal(err)
		}
		passwords = make([]Models.Auth, counterint)
	} else {
		counterint = 100
		passwords = make([]Models.Auth, 100)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < counterint; i++ {
		//passwords[i].Username = gofakeit.Username()
		passwords[i].Password = gofakeit.Password(true, true, true, true, false, rand.Intn(maxpasslength-minpasslength)+minpasslength)
	}
	c.JSON(200, passwords)

}
