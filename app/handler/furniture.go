package handler

import (
	"encoding/json"
	"github.com/fthyilmz/workshop-go.git/app/model"
	"github.com/fthyilmz/workshop-go.git/app/repository"
	"github.com/fthyilmz/workshop-go.git/app/validation"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// Furniture list
// @Summary List of furniture
// @ID furniture_list
// @Accept  json
// @Produce  json
// @Tags Furniture
// @Success 200 {array} model.Furniture
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Security Bearer
// @Router /furniture [get]
func (h *Handler) All(c *gin.Context) {
	furnitures, _ := repository.GetFurnitureRepository().All()

	c.JSON(http.StatusOK, furnitures)
}

// Get furniture by id
// @Summary Get furniture by id
// @ID furniture_get_by_id
// @Accept  json
// @Produce  json
// @Tags Furniture
// @Param id path int true "Furniture Id"
// @Success 200 {object} model.Furniture
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Security Bearer
// @Router /furniture/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := repository.GetFurnitureRepository().ById(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// Add furniture
// @Summary Add furniture
// @ID add_furniture
// @Accept  json
// @Produce  json
// @Tags Furniture
// @Param body body model.FurnitureForm true "body"
// @Success 200 {object} model.FurnitureForm
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Security Bearer
// @Router /furniture [post]
func (h *Handler) Store(c *gin.Context) {
	furniture := model.NewFurniture()

	err := json.NewDecoder(c.Request.Body).Decode(&furniture)

	if err != nil {
		log.Error("Failed reading the request body %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = validation.Check(furniture)
	if err != nil {
		errors := strings.Split(err.Error(), ";")
		response := model.ErrorResponse{Messages: errors}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	repository.GetFurnitureRepository().Store(furniture)

	c.JSON(http.StatusCreated, furniture)
}

// Update furniture
// @Summary Update furniture
// @ID update_furniture
// @Accept  json
// @Produce  json
// @Tags Furniture
// @Param id path int true "Furniture Id"
// @Param body body model.FurnitureForm true "body"
// @Success 201 {object} model.FurnitureForm
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Security Bearer
// @Router /furniture/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")
	furniture, err := repository.GetFurnitureRepository().ById(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	errs := json.NewDecoder(c.Request.Body).Decode(&furniture)

	if errs != nil {
		log.Error("Failed reading the request body %s", errs)
		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.Error()})
		return
	}

	err = validation.Check(furniture)
	if err != nil {
		errors := strings.Split(err.Error(), ";")
		response := model.ErrorResponse{Messages: errors}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result := repository.GetFurnitureRepository().Update(furniture)

	if !result {
		c.JSON(http.StatusBadRequest, gin.H{"message": ""})
		return
	}

	c.JSON(http.StatusOK, furniture)
}

// Delete furniture
// @Summary Delete furniture
// @ID delete_furniture
// @Accept  json
// @Produce  json
// @Tags Furniture
// @Param id path int true "Furniture Id"
// @Success 204 {string} Token "null"
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Security Bearer
// @Router /furniture/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	furniture, err := repository.GetFurnitureRepository().ById(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	result := repository.GetFurnitureRepository().Delete(furniture)

	if !result {
		c.JSON(http.StatusBadRequest, gin.H{"message": ""})
		return
	}

	c.JSON(http.StatusNoContent, furniture)
}
