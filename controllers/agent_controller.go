package controllers

import (
	"net/http"

	"app.myriadflow.com/db"
	"app.myriadflow.com/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateAgent(c *gin.Context) {
	var agent models.Agent
	if err := c.ShouldBindJSON(&agent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	agent.ID = uuid.New()
	db.DB.Create(&agent)
	c.JSON(http.StatusCreated, agent)
}

func GetAgents(c *gin.Context) {
	var agents []models.Agent
	db.DB.Find(&agents)
	c.JSON(http.StatusOK, agents)
}

func GetAgentByID(c *gin.Context) {
	id := c.Param("id")
	var agent models.Agent
	if err := db.DB.First(&agent, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
		return
	}
	c.JSON(http.StatusOK, agent)
}

func UpdateAgent(c *gin.Context) {
	id := c.Param("id")
	var agent models.Agent
	if err := db.DB.First(&agent, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
		return
	}

	if err := c.ShouldBindJSON(&agent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Save(&agent)
	c.JSON(http.StatusOK, agent)
}

func DeleteAgent(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Agent{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete agent"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Agent deleted successfully"})
}
