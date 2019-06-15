package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoContentResponse(context *gin.Context) {
	context.JSON(http.StatusNoContent, nil)
}
