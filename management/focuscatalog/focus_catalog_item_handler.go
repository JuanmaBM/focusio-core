package focuscatalog

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanmabm/focusio-core/management/entity"
)

type FocusCatalogItemHandler struct {
	repository FocusCatalogItemRepository
}

func RegisterHandlers(ge *gin.Engine, r FocusCatalogItemRepository) {

	handler := FocusCatalogItemHandler{r}

	ge.GET("/catalog", handler.findAll)
	ge.POST("/catalog", handler.create)
	ge.GET("/catalog/:name", handler.findByName)
	ge.PUT("/catalog/:name", handler.update)
	ge.DELETE("/catalog/:name", handler.delete)
}

func (fh FocusCatalogItemHandler) findByName(c *gin.Context) {
	name := c.Param("name")

	if app, err := fh.repository.findByName(name); err == nil {
		c.JSON(http.StatusOK, app)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (fh FocusCatalogItemHandler) findAll(c *gin.Context) {
	c.JSON(200, fh.repository.findAll())
}

func (fh FocusCatalogItemHandler) create(c *gin.Context) {
	var item entity.FocusCatalogItem

	if bindErr := c.ShouldBind(&item); bindErr != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": bindErr.Error()})
		return
	}

	if a, _ := fh.repository.findByName(item.Name); a.Name == item.Name {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "The application " + item.Name + " already exists"})
		return
	}

	if bdErr := fh.repository.insert(item); bdErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": bdErr.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (fh FocusCatalogItemHandler) update(c *gin.Context) {
	var item entity.FocusCatalogItem
	name := c.Param("name")

	if bindErr := c.ShouldBind(&item); bindErr != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": bindErr.Error()})
		return
	}

	if _, err := fh.repository.findByName(name); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := fh.repository.update(name, item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (fh FocusCatalogItemHandler) delete(c *gin.Context) {
	name := c.Param("name")
	fh.repository.delete(name)
	c.JSON(http.StatusOK, nil)
}
