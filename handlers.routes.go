package main

import (
	"bin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func showProjectFirst(c *gin.Context) {
	proj := models.FirstProject()
	c.JSON(200, gin.H{
		"project": proj,
	})
}

func showProjectsPage(c *gin.Context) {
	if projects, err := models.GetAllProjects(); err == nil {
		// Вызовем метод HTML из Контекста Gin для обработки шаблона
		c.JSON(200, gin.H{
			"projects": projects,
		})
	} else {
		// Если топика нет, прервём с ошибкой
		c.AbortWithError(http.StatusNotFound, err)
	}
}
