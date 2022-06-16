package controllers

import (
	"gin-apicrud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//obtener todos los usuarios
func FindUsers(c *gin.Context) {
	var usuarios []models.Usuarios

	models.DB.Find(&usuarios)

	c.JSON(http.StatusOK, gin.H{"data": usuarios})
}

//crear un nuevo usuario
func CreateUser(c *gin.Context) {
	//validate input
	var input models.Usuarios
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create user
	user := models.Usuarios{Folio: input.Folio, Nombre: input.Nombre, Vigencia: input.Vigencia, Areas: input.Areas, Privilegios: input.Privilegios}

	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
