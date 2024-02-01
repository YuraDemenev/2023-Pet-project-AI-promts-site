package handler

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	anonymId            = 1
	adminId             = 2
)

func (h *Handler) userIdentity(c *gin.Context) {
	header, _ := c.Cookie("token")
	// if err != nil {
	// 	logrus.Fatalf("Middlware cookie problem :%s", err.Error())
	// 	return
	// }

	if header == "" {
		//You are anonym
		c.Set(userCtx, anonymId)
		return
	}

	//parse token
	userId, err := h.service.Authorization.ParseToken(header)
	if err != nil {
		//You are anonym
		fmt.Printf("Cant parse token, user is anonym :%s", err.Error())
		c.Set(userCtx, anonymId)
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		// fmt.Println("You are Anonym")
		// id = 16
		return -1, errors.New("User id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return -1, errors.New("User id not int")
	}

	return idInt, nil
}
