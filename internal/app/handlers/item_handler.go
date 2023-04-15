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

func (app *ItemHandler) Create(c *gin.Context) {
	var item appModel.NewItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdItem, err := app.service.CreateItem(&item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdItem)
}

func (app *ItemHandler) GetItemsByListId(c *gin.Context) {
	listId := c.Param("id")
	item, err := app.service.GetItemsByListID(listId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (app *ItemHandler) Delete(c *gin.Context) {
	itemId := c.Param("id")
	if err := app.service.DeleteItem(itemId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "item removed!")
}
