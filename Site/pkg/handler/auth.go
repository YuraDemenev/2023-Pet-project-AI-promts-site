package handler

import (
	"fmt"
	"html/template"
	"net/http"
	site "site/pkg/elements"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// FUNCTION FOR RETURN ERROR ALERT THROW HTMX
func generateErrorAller(errorStatus int, errorTitle string, errorMessage string, err error, c *gin.Context) {
	if err != nil {
		logrus.Infof("ErrorStatus : %d | Error Title :%s | Error Message :%s | Error :%s", errorStatus, errorTitle, errorMessage, err.Error())
	} else {
		logrus.Infof("ErrorStatus : %d | Error Title :%s | Error Message :%s | Error :nill", errorStatus, errorTitle, errorMessage)
	}

	htmlStr := fmt.Sprintf("<div class='alert alert-danger' role='alert'>%s! %s</div>", errorTitle, errorMessage)
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(c.Writer, nil)
}

// POST FUNCTIONS
func (h *Handler) signUp(c *gin.Context) {
	username := c.Request.PostFormValue("inputUserName")
	password := c.Request.PostFormValue("inputPassword")
	reapeatPassword := c.Request.PostFormValue("repeatPassword")

	if password != reapeatPassword {
		generateErrorAller(http.StatusBadRequest, "Registration Failed", "Incorrectly entered repeated password", nil, c)
		return
	}
	if password == "" {
		generateErrorAller(http.StatusBadRequest, "Registration Failed", "Password is empty", nil, c)
		return
	}
	if username == "" {
		generateErrorAller(http.StatusBadRequest, "Registration Failed", "Username is empty", nil, c)
		return
	}
	if len(username) > 75 {
		generateErrorAller(http.StatusBadRequest, "Registration Failed", "Username is to long", nil, c)
		return
	}
	if len(password) > 500 {
		generateErrorAller(http.StatusBadRequest, "Registration Failed", "Password is to long", nil, c)
		return
	}

	var locUser site.User
	locUser.Password = password
	locUser.UserName = username

	id, err := h.service.Authorization.CreateUser(locUser)
	if err != nil {
		generateErrorAller(http.StatusBadRequest, "Registration Failed", "This user is registred", nil, c)
		return
	}
	c.Set(userCtx, id)

	htmlStr := "<div class='alert alert-success' role='alert'>Succes registration</div>"
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(c.Writer, nil)
	c.HTML(http.StatusOK, "", gin.H{
		"Message": "Succes registration",
	})

}

func (h *Handler) signIn(c *gin.Context) {
	username := c.PostForm("inputUserName")
	password := c.PostForm("inputPassword")

	if password == "" {
		generateErrorAller(http.StatusBadRequest, "Authorization Failed", "Password is empty", nil, c)
		return
	}
	if username == "" {
		generateErrorAller(http.StatusBadRequest, "Authorization Failed", "Username is empty", nil, c)
		return
	}

	token, err := h.service.Authorization.GenerateToken(username, password)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			generateErrorAller(http.StatusInternalServerError, "Authorization Failed", "No such user", err, c)
		}
		//newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie("", "", -1, "/", h.url, false, true) // Empty name deletes all cookies

	//get correct url. in yml like: `http://localhost:8080` we need `localhost`
	urls := strings.Split(h.url, "/")
	url := strings.Split(urls[len(urls)-1], ":")

	//Save jwt token
	c.SetCookie("token", token, int(time.Second*3600*2), "/", url[0], false, true)
	//c.Set("Authorization", token)
	c.Header("HX-Redirect", "/pictures/")

}

// GET FUNCIONS
func (handler *Handler) registration(c *gin.Context) {
	c.HTML(http.StatusOK, "registration.html", gin.H{
		"ErrorTitle":   "",
		"ErrorMessage": "",
		"URL":          handler.url,
	})
}

func (handler *Handler) authorization(c *gin.Context) {
	c.HTML(http.StatusOK, "authorization.html", gin.H{
		"ErrorTitle":   "",
		"ErrorMessage": "",
		"URL":          handler.url,
	})
}
