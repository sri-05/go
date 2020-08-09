package webgo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

var v struct {
	Data []Employee
}

func GetEmployee(c *gin.Context) {
	name := c.Query("name")
	id := c.Query("id")
	//Connect to database and check the result
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_db")
	if err != nil {
		log.Println("Database connection Failure")
		panic(err.Error())

	}
	// defer db.Close()

	result, err := db.Query("SELECT * FROM employee where name=? and id=?", name, id)
	if err != nil {
		log.Println(err.Error())
	}
	var emp Employee
	for result.Next() {

		err := result.Scan(&emp.ID, &emp.Name, &emp.City)
		if err != nil {
			log.Println(err.Error())
		}

	}
	v.Data = append(v.Data, emp)
	result1, _ := json.Marshal(v.Data)
	fmt.Println(string(result1))
	result2 := json.RawMessage(result1)
	c.JSON(200, gin.H{
		"Employee Details": result2,
	})

}
