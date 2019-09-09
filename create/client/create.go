package main

import (
	"fmt"
	"net/http"
	"strconv"

	todopb "github.com/OGLinuk/micro-todos/create/proto"
	"github.com/gin-gonic/gin"
)

func CreateHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	description := ctx.Param("desc")
	priority, err := strconv.ParseInt(ctx.Param("priority"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Priority Parameter"})
		return
	}

	task := &todopb.Task{
		Name:        name,
		Description: description,
		Priority:    int32(priority),
	}

	req := &todopb.CreateRequest{Token: "test", Task: task}
	if resp, err := Client.Create(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(resp.Task),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
