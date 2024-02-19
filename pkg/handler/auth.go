package handler

import (
	"net/http"

	todo "github.com/delapaska/university-storage"
	"github.com/gin-gonic/gin"
)

var Token string

func (h *Handler) signUp(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "sign_up.html", nil)
		return
	}
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		name := c.PostForm("name")
		password := c.PostForm("password")
		//var input todo.User

		input := todo.User{
			Username: username,
			Name:     name,
			Password: password,
		}
		/*
			var user todo.User
			if err := c.BindJSON(&user); err != nil {
				newErrorResponse(c, http.StatusBadRequest, err.Error())
				return
			}
		*/
		_, err := h.services.Authorization.CreateUser(input)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		//c.Redirect(http.StatusOK, "auth//sign-in")
		//	c.JSON(http.StatusOK, map[string]interface{}{
		////		"id": id,
		//	})
		c.Redirect(http.StatusSeeOther, "/auth/sign-in")
		//c.Redirect(http.StatusSeeOther, "/auth/sign-in")
	}
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "sign_in.html", nil)
		return
	}
	username := c.PostForm("username")

	password := c.PostForm("password")
	//var user todo.User

	input := signInInput{
		Username: username,

		Password: password,
	}
	/*
		if err := c.BindJSON(&user); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	*/
	Token, _ = h.services.Authorization.GenerateToken(input.Username, input.Password)
	/*
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	*/
	c.Redirect(http.StatusSeeOther, "/api/main")
	/*
		c.JSON(http.StatusOK, map[string]interface{}{
			"token": Token,
		})
	*/
}

func (h *Handler) loadMainPage(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "api.html", nil)
		return
	}
}
