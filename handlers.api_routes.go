package main

import (
	"bin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showApiProjectsPage(c *gin.Context) {
	// Проверим существование топика
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

func showApiProjectFirst(c *gin.Context) {
	proj := models.FirstProject()
	c.JSON(200, gin.H{
		"projects": proj,
	})
}
func getApiProject(c *gin.Context) {
	if projectID, err := strconv.Atoi(c.Param("project_id")); err == nil {
		// Проверим существование топика
		if project, err := models.GetProjectByID(projectID); err == nil {
			// Вызовем метод HTML из Контекста Gin для обработки шаблона
			c.JSON(200, gin.H{
				"project": project,
			})
		} else {
			// Если топика нет, прервём с ошибкой
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// При некорректном ID в URL, прервём с ошибкой
		c.AbortWithStatus(http.StatusNotFound)
	}
}
