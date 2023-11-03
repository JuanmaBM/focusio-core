package focusapp

import (
	"github.com/gin-gonic/gin"
	"github.com/juanmabm/focusio-core/management/entity"
)

type FocusAppHandler struct {
	repository FocusAppRepository
}

func RegisterHandlers(ge *gin.Engine, r FocusAppRepository) {

	handler := FocusAppHandler{r}

	ge.GET("/app", handler.findAll)
	ge.GET("/app/:name", handler.findByName)
	ge.POST("/app", handler.create)
	ge.PUT("/app/:name", handler.update)
	ge.DELETE("/app", handler.delete)
}

func (fh FocusAppHandler) findByName(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &entity.FocusApp{})
}

func (fh FocusAppHandler) findAll(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, []entity.FocusApp{})
}

func (fh FocusAppHandler) create(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &entity.FocusApp{})
}

func (fh FocusAppHandler) update(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, &entity.FocusApp{})
}

func (fh FocusAppHandler) delete(c *gin.Context) {
	// TODO: Implement
}
