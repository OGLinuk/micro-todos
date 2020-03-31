package main

import (
	"fmt"
	"net/http"

	todopb "github.com/OGLinuk/micro-todos/retrieve/proto"
	"github.com/gin-gonic/gin"
)

func RetrieveAllHandler(ctx *gin.Context) {
	req := &todopb.RetrieveAllRequest{Token: "test"}
	if resp, err := Client.RetrieveAll(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(resp.Tasks),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
