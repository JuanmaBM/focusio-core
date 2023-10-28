package http

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine, h *Handler) {

	r.GET("/", h.Greetings)
}
