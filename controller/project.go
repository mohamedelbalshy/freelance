package controller

import (
	"fmt"
	"freelance/helper"
	"freelance/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProjects(ctx *gin.Context) {
	projects := model.GetAllProjects()

	ctx.JSON(http.StatusOK, gin.H{"projects": projects})
}

func CreateProject(ctx *gin.Context) {
	var input model.Project

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(ctx)

	fmt.Println(user)
	if err != nil {
		log.Fatal("Error Getting User from jwt token")
	}

	project := model.Project{
		Title:       input.Title,
		Description: input.Description,
		Cost:        input.Cost,
		Duration:    input.Duration,
		ProjectType: input.ProjectType,
		OwnerId:     user.ID,
		Owner:       user,
	}

	savedProject, err := project.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"project": savedProject})

}
