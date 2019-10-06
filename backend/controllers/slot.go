package controllers

import (
	"net/http"

	"github.com/durgaprasad-budhwani/home-automation/backend/models"
	"github.com/durgaprasad-budhwani/home-automation/backend/services"

	"github.com/gin-gonic/gin"
)

type SlotController struct {
	service services.SlotService
}

func NewSlotController(service services.SlotService) SlotController {
	return SlotController{
		service,
	}
}

func (c SlotController) GetSlots(context *gin.Context) {
	schedulers, err := c.service.GetAll(10, 0)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, schedulers)
}

func (c SlotController) AddSlot(context *gin.Context) {
	var scheduler models.Slot
	err := context.BindJSON(&scheduler)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = c.service.Save(&scheduler)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, scheduler)
}

func (c SlotController) UpdateSlot(context *gin.Context) {
	var scheduler models.Slot
	err := context.BindJSON(&scheduler)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = c.service.Update(scheduler.ID, &scheduler)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, scheduler)
}
