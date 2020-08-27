package handler

import (
	"github.com/fthyilmz/workshop-go.git/app/model"
	"github.com/fthyilmz/workshop-go.git/app/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Total price of Furniture
// @Summary Total price of Furniture
// @ID furniture_total_list
// @Accept  json
// @Produce  json
// @Tags Total
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Security Bearer
// @Router /total/furniture [get]
func (d *DashboardHandler) TotalFurniture(c *gin.Context) {

	totalRoom := repository.GetFurnitureRepository().TotalFurniture()

	c.JSON(http.StatusOK, totalRoom)
}

// Total price of Furniture by apartmentId
// @Summary Total price of Furniture by apartmentId
// @ID furniture_total_list_by_apartment_id
// @Accept  json
// @Produce  json
// @Tags Total
// @Param apartmentId path string true "Apartment Id"
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Security Bearer
// @Router /total/furniture/{apartmentId} [get]
func (d *DashboardHandler) TotalFurnitureOfApartment(c *gin.Context) {

	apartmentId := c.Params.ByName("apartmentId")

	rooms, err := repository.GetRoomRepository().GetRoomByApartmentId(apartmentId)

	if len(rooms) <= 0 || err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Messages: []string{model.ErrNotFound.Error()}})
		return
	}

	furnitures, err := repository.GetRoomRepository().TotalFurnitureOfApartment(rooms)

	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Messages: []string{model.ErrNotFound.Error()}})
		return
	}

	c.JSON(http.StatusOK, furnitures)
}
