package http

import (
	"github.com/gin-gonic/gin"
	"github.com/juanmabm/focusio-core/management/domain"
)

type FocusCatalogItemHandler struct{}

func NewFocusCatalogItemHandler() *FocusCatalogItemHandler {
	return &FocusCatalogItemHandler{}
}

func (fh *FocusCatalogItemHandler) findByName(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &domain.FocusCatalogItem{})
}

func (fh *FocusCatalogItemHandler) findAll(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &domain.FocusCatalogItem{})
}

func (fh *FocusCatalogItemHandler) create(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &domain.FocusCatalogItem{})
}

func (fh *FocusCatalogItemHandler) update(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &domain.FocusCatalogItem{})
}

func (fh *FocusCatalogItemHandler) delete(c *gin.Context) {
	// TODO: Implement
}
