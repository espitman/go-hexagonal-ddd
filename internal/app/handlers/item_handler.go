package handlers

import (
	appModel "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/models"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ItemHandler struct {
	service appServices.ItemService
}

func NewItemHandler(service appServices.ItemService) *ItemHandler {
	return &ItemHandler{service}
}

// Create a new Item
// @Summary Add a new item to list
// @Description Add a new item to list
// @Tags Item
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param list body appModel.NewItem true "Item object to Add"
// @Success 201 {object} appModel.Item
// @Failure 400 {object} commonModels.ErrorResponse
// @Failure 401 {object} commonModels.ErrorResponse
// @Router /item [post]
func (app *ItemHandler) Create(c *gin.Context) {
	var item appModel.NewItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdItem, err := app.service.CreateItem(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdItem)
}

// Delete item by ID.
// @Summary remove Item from the list by ID
// @Description remove Item from the list by ID
// @Tags Item
// @Param id path string true "Item ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} commonModels.OkResponse
// @Failure 400 {object} commonModels.ErrorResponse
// @Failure 401 {object} commonModels.ErrorResponse
// @Router /item/{id} [delete]
func (app *ItemHandler) Delete(c *gin.Context) {
	itemId := c.Param("id")
	if err := app.service.DeleteItem(itemId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "item removed!")
}
