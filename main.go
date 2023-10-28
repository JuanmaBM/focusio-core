package focusiocore

import (
	"github.com/gin-gonic/gin"
	ihttp "github.com/juanmabm/focusiocore/management/http"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	h := ihttp.NewHandler()
}
