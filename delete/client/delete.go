package main

import (
	"fmt"
	"net/http"

	todopb "github.com/OGLinuk/micro-todos/delete/proto"
	"github.com/gin-gonic/gin"
)

func DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	req := &todopb.DeleteRequest{Token: "test", Id: id}
	if resp, err := Client.Delete(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(resp.Success),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
