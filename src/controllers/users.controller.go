package usersController

import (
	userModel "crud-golang/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var body createUserRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := userModel.Create(body.Name, body.Email, body.Password)
	userModel.Save(user)

	ctx.Status(http.StatusCreated)
}

func Update(ctx *gin.Context) {
	var body updateUserRequest
	idStr := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idNumeric, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := userModel.FindById(uint(idNumeric))

	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	user.Name = body.Name
	user.Email = body.Email
	user.Password = body.Password

	userModel.Save(user)

	ctx.Status(http.StatusOK)
}

func FindAllPaginated(ctx *gin.Context) {
	var queryParams findAllPaginatedRequest

	if err := ctx.ShouldBind(&queryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, err := strconv.ParseUint(queryParams.Page, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	limit, err := strconv.ParseUint(queryParams.Limit, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := userModel.FindAllPaginated(uint(page), uint(limit))

	ctx.JSON(http.StatusOK, &result)
}

func FindById(ctx *gin.Context) {
	idStr := ctx.Param("id")

	idNumeric, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := userModel.FindById(uint(idNumeric))

	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, &user)
}

func Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")

	idNumeric, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userModel.Delete(uint(idNumeric))

	ctx.Status(http.StatusOK)
}
