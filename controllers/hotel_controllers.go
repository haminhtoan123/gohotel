package controllers

import (
	"net/http"

	models "github.com/haminhtoan123/gohotel/models"
	repo "github.com/haminhtoan123/gohotel/repo"

	"github.com/gin-gonic/gin"
)

type HotelController struct {
	repo repo.HotelRepository
}

func NewHotelController(repo repo.HotelRepository) *HotelController {
	return &HotelController{repo: repo}
}

// AddHotel godoc
// @Summary Add a hotel
// @Description Add a hotel to the database
// @Accept json
// @Produce json
// @Param hotel body models.Hotel true "Hotel"
// @Success 201 {object} models.Hotel
// @Router /hotels [post]
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

// UpdateHotel godoc
// @Summary Update a hotel
// @Description Update a hotel in the database
// @Accept json
// @Produce json
// @Param name path string true "Hotel Name"
// @Param hotel body models.Hotel true "Hotel"
// @Success 200 {object} models.Hotel
// @Router /hotels/{name} [put]
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

// GetHotel godoc
// @Summary Get a hotel by name
// @Description Get a hotel from the database by its name
// @Produce json
// @Param name path string true "Hotel Name"
// @Success 200 {object} models.Hotel
// @Router /hotels/{name} [get]
func (c *HotelController) GetHotel(ctx *gin.Context) {
	name := ctx.Param("name")

	hotel, err := c.repo.Get(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, hotel)
}
