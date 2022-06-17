package controllers

import (
	"fmt"
	"gin-apicrud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUsuariosInput struct {
	Folio       string `json:"folio" binding:"required"`
	Nombre      string `json:"nombre" binding:"required"`
	Areas       string `json:"areas" binding:"required"`
	Privilegios string `json:"privilegios"`
}

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

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create user
	//user := models.Usuarios{Folio: input.Folio, Nombre: input.Nombre, Areas: input.Areas, Privilegios: input.Privilegios}

	models.DB.Create(&input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func FindUser(c *gin.Context) {
	var usuario models.Usuarios

	fmt.Println(c.Param("id"))

	if err := models.DB.Where("Folio=?", c.Param("id")).First(&usuario).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": usuario})
}

func UpdateUser(c *gin.Context) {
	var user models.Usuarios
	if err := models.DB.Where("Folio = ? ", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Record not found"})
		return
	}

	//validate input
	var input models.Usuarios
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user models.Usuarios

	if err := models.DB.Where("Folio = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
