package http

import (
	"github.com/gin-gonic/gin"
	"github.com/juanmabm/focusio-core/management/domain"
)

type FocusAppHandler struct{}

func NewFocusAppHandler() *FocusAppHandler {
	return &FocusAppHandler{}
}

func (fh *FocusAppHandler) findByName(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &domain.FocusApp{})
}

func (fh *FocusAppHandler) findAll(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, []domain.FocusApp{})
}

func (fh *FocusAppHandler) create(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &domain.FocusApp{})
}

func (fh *FocusAppHandler) update(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &domain.FocusApp{})
}

func (fh *FocusAppHandler) delete(c *gin.Context) {
	// TODO: Implement
}
