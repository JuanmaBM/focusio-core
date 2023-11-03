package http

import (
	"github.com/gin-gonic/gin"
	"github.com/juanmabm/focusio-core/management/entity"
)

type FocusCatalogItemHandler struct{}

func NewFocusCatalogItemHandler() *FocusCatalogItemHandler {
	return &FocusCatalogItemHandler{}
}

func (fh *FocusCatalogItemHandler) findByName(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &entity.FocusCatalogItem{})
}

func (fh *FocusCatalogItemHandler) findAll(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &entity.FocusCatalogItem{})
}

func (fh *FocusCatalogItemHandler) create(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &entity.FocusCatalogItem{})
}

func (fh *FocusCatalogItemHandler) update(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &entity.FocusCatalogItem{})
}

func (fh *FocusCatalogItemHandler) delete(c *gin.Context) {
	// TODO: Implement
}
