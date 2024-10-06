package jwtx

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kratos-application/api/gen/ecode"
	"kratos-application/core/responsex"
)

func GinSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &responsex.ResponseModel{
		Code: ecode.ECode_Success,
		Data: data,
	})
}

func GinError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, responsex.ResponseError(err))
}
