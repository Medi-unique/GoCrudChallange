package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/person_crud/models"
	"example.com/person_crud/data"
)

func GetAllPersons(c *gin.Context) {
	fmt.Println(data.Data)
	c.JSON(http.StatusOK, data.GetAllPersons())
}

func GetPerson(c *gin.Context) {

	id := c.Param("id")
	person, err := data.GetPerson(id)
	if err == nil {
		c.IndentedJSON(http.StatusFound, gin.H{"message": "Found", "data": person})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

// Done
func UpdatePerson(c *gin.Context) {
	var person models.Person
	id := c.Param("id")
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	val, err := data.UpdatePerson(person, id)
	if err == nil {
		fmt.Println("data found")
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "Added", "data": val})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	person, err := data.DeletePerson(id)
	if err == nil {
		fmt.Println("data found")
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "successfully Deleted", "data": person})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func PostPerson(c *gin.Context) {

	var task models.Person
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	val, err := data.AddPerson(task)
	if err == nil {
		fmt.Println("data found")
		c.IndentedJSON(http.StatusOK, gin.H{"message": "successfully ADD", "data": val})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

}
