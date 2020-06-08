package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/blrobin2/go-local-api/models"
	"github.com/blrobin2/go-local-api/dtos"
)

// GET /contracts
// Get all contracts
func FindContracts(c *gin.Context) {
	var contracts []models.Contract
	models.DB.Find(&contracts)

	c.JSON(http.StatusOK, gin.H{"data": contracts})
}

// POST /contracts
// Create new contract
func CreateContract(c *gin.Context) {
	var input DTOs.CreateContractInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contract := models.Contract{
		ContractNumber: input.ContractNumber,
		CustomerName: input.CustomerName,
		StartDate: input.StartDate,
		EndDate: input.EndDate,
		Balance: input.Balance,
	}
	models.DB.Create(&contract)

	c.JSON(http.StatusOK, gin.H{"data": contract})
}

// GET /contracts/:id
// Find a contract
func FindContract(c *gin.Context) {
	var contract models.Contract

	if err := models.DB.Where("id = ?", c.Param("id")).First(&contract).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contract not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contract})
}

// PATCH /contracts/:id
// Update a contract
func UpdateContract(c *gin.Context) {
	var contract models.Contract
	if err := models.DB.Where("id = ?", c.Param("id")).First(&contract).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contract not found"})
		return
	}

	var input DTOs.UpdateContractInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&contract).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": contract})
}

// DELETE /contracts/:id
// Delete a contract
func DeleteContract(c *gin.Context) {
	var contract models.Contract
	if err := models.DB.Where("id = ?", c.Param("id")).First(&contract).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contract not found"})
		return
	}

	models.DB.Delete(&contract)

	c.JSON(http.StatusOK, gin.H{"data": true})
}