package api

import (
	"blog/server/model/response"
	"blog/server/service"

	"github.com/gin-gonic/gin"
)

// FindType : request the amount of each type
func FindType(c *gin.Context) {
	var blogType service.BlogType
	result, err := blogType.FindAllTypeCount()
	
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}
	res := response.Response{
		Data: result,
	}
	res.Json(c)
}