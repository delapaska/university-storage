package handler

import (
	"fmt"
	"net/http"
	"strconv"

	todo "github.com/delapaska/university-storage"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createProject(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "create.html", nil)
		return
	}
	userId, err := getUserId(c)
	fmt.Println(userId)
	if err != nil {
		return
	}

	directory := c.PostForm("directory")
	title := c.PostForm("title")

	//var input todo.User

	input := todo.ProjectList{
		Title:     directory,
		Directory: title,
	}
	//var input todo.ProjectList
	fmt.Println(input)
	/*
		if err := c.BindJSON(&input); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	*/
	id, err := h.services.ProjectList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllListsResponse struct {
	Data []todo.ProjectList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	fmt.Println(userId)
	if err != nil {
		return
	}

	projects, err := h.services.ProjectList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: projects,
	})
}

func (h *Handler) getProjectById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	project, err := h.services.ProjectList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, project)
}
