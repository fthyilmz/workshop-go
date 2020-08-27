package handler

import (
	"github.com/fthyilmz/workshop-go.git/app/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Room list
// @Summary List of room
// @ID room_list
// @Accept  json
// @Produce  json
// @Tags Room
// @Success 200 {array} model.Room
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Security Bearer
// @Router /room [get]
func (h *Handler) GetRoomList(c *gin.Context) {
	rooms, _ := repository.GetRoomRepository().AllRoom()

	c.JSON(http.StatusOK, rooms)
}
