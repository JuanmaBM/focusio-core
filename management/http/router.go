package http

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine, fah *FocusAppHandler, fch *FocusCatalogItemHandler) {

	r.GET("/app", fah.findAll)
	r.GET("/app/:name", fah.findByName)
	r.POST("/app", fah.create)
	r.PUT("/app/:name", fah.update)
	r.DELETE("/app", fah.delete)

	r.GET("/catalog", fch.findAll)
	r.GET("/catalog/:name", fch.findByName)
	r.POST("/catalog", fch.create)
	r.PUT("/catalog/:name", fch.update)
	r.DELETE("/catalog", fch.delete)
}
