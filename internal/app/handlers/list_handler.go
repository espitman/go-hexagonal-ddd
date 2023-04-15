package handlers

import (
	appModel "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/models"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListHandler struct {
	service appServices.ListService
}

func NewListHandler(service appServices.ListService) *ListHandler {
	return &ListHandler{service}
}

func (app *ListHandler) Create(c *gin.Context) {
	var list appModel.NewList
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list.UserId = c.GetInt64("userId")
	createdList, err := app.service.CreateList(&list)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdList)
}

func (app *ListHandler) GetById(c *gin.Context) {
	listId := c.Param("id")
	list, err := app.service.GetListByID(listId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "list not found"})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (app *ListHandler) GetAll(c *gin.Context) {
	lists, err := app.service.GetLists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lists)
}

func (app *ListHandler) Update(c *gin.Context) {
	listId := c.Param("id")

	var list appModel.NewList
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedList, err := app.service.UpdateList(listId, &list)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedList)
}

func (app *ListHandler) Delete(c *gin.Context) {
	listId := c.Param("id")
	if err := app.service.DeleteList(listId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "list removed!")
}
