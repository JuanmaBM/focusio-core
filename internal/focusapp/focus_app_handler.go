package focusapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanmabm/focusio-core/internal/entity"
	"github.com/juanmabm/focusio-core/internal/focuscatalog"
	"github.com/juanmabm/focusio-core/pkg/argocdclient"
)

type FocusAppHandler struct {
	repository        FocusAppRepository
	catalogRepository focuscatalog.FocusCatalogItemRepository
	argoClient        argocdclient.ArgoCDClient
}

func RegisterHandlers(ge *gin.Engine, r FocusAppRepository, fcr focuscatalog.FocusCatalogItemRepository, ac argocdclient.ArgoCDClient) {

	handler := FocusAppHandler{r, fcr, ac}

	ge.GET("/app", handler.findAll)
	ge.POST("/app", handler.create)
	ge.GET("/app/:name", handler.findByName)
	ge.PUT("/app/:name", handler.update)
	ge.DELETE("/app/:name", handler.delete)
}

func (fh FocusAppHandler) findByName(c *gin.Context) {
	name := c.Param("name")

	if app, err := fh.repository.FindByName(name); err == nil {
		c.JSON(http.StatusOK, app)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (fh FocusAppHandler) findAll(c *gin.Context) {
	c.JSON(200, fh.repository.FindAll())
}

func (fh FocusAppHandler) create(c *gin.Context) {
	var app entity.FocusApp

	if bindErr := c.ShouldBind(&app); bindErr != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": bindErr.Error()})
		return
	}

	if a, _ := fh.repository.FindByName(app.Name); a.Name == app.Name {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "The application " + app.Name + " already exists"})
		return
	}

	if _, err := fh.catalogRepository.FindByName(app.FocusCatalogItem); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "The CatalogItem " + app.FocusCatalogItem + " doesn't exists"})
		return
	}

	if bdErr := fh.repository.Insert(app); bdErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": bdErr.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

func (fh FocusAppHandler) update(c *gin.Context) {
	var app entity.FocusApp
	name := c.Param("name")

	if bindErr := c.ShouldBind(&app); bindErr != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": bindErr.Error()})
		return
	}

	if _, err := fh.repository.FindByName(name); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := fh.repository.Update(name, &app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

func (fh FocusAppHandler) delete(c *gin.Context) {
	name := c.Param("name")
	fh.repository.Delete(name)
	c.JSON(http.StatusOK, nil)
}
