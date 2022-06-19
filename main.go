package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	connection *Repository
)

func main() {
	router := gin.Default()
	router.GET("/persons", getPersons)
	router.GET("/person/:id", getPersonByID)
	router.POST("/person", postPerson)
	router.DELETE("/person/:id", delPersonByID)
	router.PATCH("/person/:id", patchId)
	connection = connect()
	router.Run("localhost:8080")
}

var persons = []person{
	{ID: "1", Name: "Paulo B Carreira", Cpf: "031.400.100-10", Age: 56,
		Address: address{Street: "test", Number: 10, ZipCode: "81.070-310"},
	},
	{ID: "2", Name: "Paulo B Carreira", Cpf: "031.400.100-10", Age: 56},
	{ID: "3", Name: "Paulo B Carreira", Cpf: "031.400.100-10", Age: 56},
}

func getPersons(c *gin.Context) {
	persons, err := connection.queryAll()
	if err == nil {
		c.IndentedJSON(http.StatusOK, persons)
	} else {
		c.IndentedJSON(http.StatusOK, err)
	}

}

func postPerson(c *gin.Context) {

	var newPerson person

	if error := c.BindJSON(&newPerson); error != nil {
		panic(error)
	}

	id, err := connection.insert(newPerson)
	if err == nil {
		newPerson.ID = strconv.FormatInt(id, 10)
		c.IndentedJSON(http.StatusCreated, newPerson)
	} else {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

}

func patchId(c *gin.Context) {

	var patchPerson person

	if error := c.BindJSON(&patchPerson); error != nil {
		panic(error)
	}

	flag, err := connection.update(patchPerson)
	if err == nil && flag {
		c.IndentedJSON(http.StatusOK, patchPerson)
	} else {
		panic(err)
		//c.IndentedJSON(http.StatusBadRequest, err)
	}

}

func getPersonByID(c *gin.Context) {
	var id int64
	fmt.Sscan(c.Param("id"), &id)

	p, err := connection.query(id)
	if err == nil {
		c.IndentedJSON(http.StatusOK, p)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
}

func delPersonByID(c *gin.Context) {
	var id int64
	fmt.Sscan(c.Param("id"), &id)

	rowsAffected, err := connection.delete(id)
	if err == nil && rowsAffected == 1 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Person " + strconv.FormatInt(id, 10) + " deleted"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}

}
