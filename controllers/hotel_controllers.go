package controllers

import (
	"net/http"

	"github.com/haminhtoan123/gohotel/models/models"
	"github.com/haminhtoan123/gohotel/repo/repo"

	"github.com/gin-gonic/gin"
)

type HotelController struct {
	repo repo.HotelRepository
}

func NewHotelController(repo repo.HotelRepository) *HotelController {
	return &HotelController{repo: repo}
}

func (c *HotelController) AddHotel(ctx *gin.Context) {
	var hotel models.Hotel
	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.repo.Add(&hotel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, hotel)
}

func (c *HotelController) UpdateHotel(ctx *gin.Context) {
	name := ctx.Param("name")

	var hotel models.Hotel
	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotel.Name = name

	if err := c.repo.Update(&hotel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, hotel)
}

func (c *HotelController) GetHotel(ctx *gin.Context) {
	name := ctx.Param("name")

	hotel, err := c.repo.Get(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, hotel)
}
