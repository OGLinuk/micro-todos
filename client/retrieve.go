package main

import (
	"fmt"
	"net/http"

	todopb "github.com/OGLinuk/micro-todos/proto"
	"github.com/gin-gonic/gin"
)

func RetrieveHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	req := &todopb.RetrieveRequest{Token: "test", Id: id}
	if resp, err := Client.Retrieve(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(resp.Task),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
