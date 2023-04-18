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

// Create a new list
// @Summary Create a new list
// @Description Create a new list with the specified name
// @Tags List
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param list body appModel.NewList true "List object to create"
// @Success 201 {object} appModel.List
// @Failure 400 {object} commonModels.ErrorResponse
// @Failure 401 {object} commonModels.ErrorResponse
// @Failure 409 {object} commonModels.ErrorResponse
// @Router /list [post]
func (app *ListHandler) Create(c *gin.Context) {
	var list appModel.NewList
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := c.GetInt64("userId")
	createdList, err := app.service.CreateList(&list, userId)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdList)
}

// GetById Get list by ID.
// @Summary Get list by ID
// @Description Get a list using the provided ID
// @Tags List
// @Param id path string true "List ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} appModel.ListWithItems
// @Failure 401 {object} commonModels.ErrorResponse
// @Failure 404 {object} commonModels.ErrorResponse
// @Router /list/{id} [get]
func (app *ListHandler) GetById(c *gin.Context) {
	listId := c.Param("id")
	list, err := app.service.GetListByID(listId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "list not found"})
		return
	}
	c.JSON(http.StatusOK, list)
}

// GetAll Get all lists.
// @Summary Get all lists
// @Description Get all lists
// @Tags List
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []appModel.List
// @Failure 401 {object} commonModels.ErrorResponse
// @Failure 500 {object} commonModels.ErrorResponse
// @Router /list [get]
func (app *ListHandler) GetAll(c *gin.Context) {
	userId := c.GetInt64("userId")
	lists, err := app.service.GetLists(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lists)
}

// Update list by ID.
// @Summary update list by ID
// @Description update a list using the provided ID
// @Tags List
// @Param id path string true "List ID"
// @Param list body appModel.NewList true "List object to update"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} appModel.List
// @Failure 400 {object} commonModels.ErrorResponse
// @Failure 401 {object} commonModels.ErrorResponse
// @Failure 404 {object} commonModels.ErrorResponse
// @Failure 500 {object} commonModels.ErrorResponse
// @Router /list/{id} [put]
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

// Delete list by ID.
// @Summary delete list by ID
// @Description delete a list using the provided ID
// @Tags List
// @Param id path string true "List ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} commonModels.OkResponse
// @Failure 400 {object} commonModels.ErrorResponse
// @Failure 401 {object} commonModels.ErrorResponse
// @Router /list/{id} [delete]
func (app *ListHandler) Delete(c *gin.Context) {
	listId := c.Param("id")
	if err := app.service.DeleteList(listId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "list removed!"})
}
