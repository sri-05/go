package main

import (
	"github.com/sri-05/Webgo/webgo"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func firstPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "firstPageEdit",
	})
}

func createEmployee(c *gin.Context) {
	var name string = c.Param("name")
	id := c.Param("id")
	c.JSON(201, gin.H{
		"message": "Emploee name is " + name + " Employee ID is: " + id, //Employee/Sridhar/1233
	})
}

func updateEmployee(c *gin.Context) {
	c.JSON(202, gin.H{
		"message": "update method employee",
	})
}

func deleteEmployee(c *gin.Context) {
	c.JSON(204, gin.H{
		"message": "delete method employee",
	})
}

func getEmployee() {
	webgo.GetEmployee()
}
func main() {
	r := gin.Default()
	r.GET("/home", firstPage)
	r.GET("/Employee", getEmployee)
	r.POST("/Employee/:name/:id", createEmployee)
	r.PUT("/Employee", updateEmployee)
	r.DELETE("/Employee", deleteEmployee)

	r.Run("127.0.0.1:9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
