package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type lastImageIdData struct {
	Id string `json:"lastImageId"`
}

// func (h *Handler) getUserNameLocal(c *gin.Context) string {
// 	idData, _ := c.Get(userCtx)
// 	id, check := idData.(int)
// 	if check == false {
// 		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", nil, *&c)
// 		return ""
// 	}
// 	userName, err := h.service.Pictures.GetUserName(id)
// 	if err != nil {
// 		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", err, *&c)
// 		return ""
// 	}
// 	return userName
// }

// Function for home page when you scroll down
func (h *Handler) watchPicturePost(c *gin.Context) {
	//Get lastImageId from header
	lastImageId := c.Request.Header.Get("lastImageId")
	lastImageIdInt, _ := strconv.Atoi(lastImageId)

	//Get new urls
	urls, err := h.service.Pictures.GetNewImages(lastImageIdInt)
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
	}
	for _, htmlStr := range urls {
		if htmlStr == "" {
			break
		}
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(c.Writer, nil)
	}

	return
}

// Function for home page when you just click Home
func (h *Handler) watchPictureGet(c *gin.Context) {
	//Get user user name
	idData, _ := c.Get(userCtx)
	id, check := idData.(int)
	if check == false {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", nil, *&c)
		return
	}
	userName, err := h.service.Pictures.GetUserName(id)
	if err != nil {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", err, *&c)
		return
	}
	tmpl, _ := template.ParseFiles("../templates/pictures.html")

	//Get Urls
	urls, err := h.service.Pictures.GetNewImages(0)
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
	}
	urslHtml := make([]template.HTML, len(urls))
	for i := range urls {
		urslHtml[i] = template.HTML(urls[i])
	}

	data := map[string]interface{}{
		"Username": userName,
		"Urls":     urslHtml,
	}

	tmpl.Execute(c.Writer, data)

}

// Function for show promts when you click to image
func (h *Handler) showPromts(c *gin.Context) {
	//Get user id
	id, _ := c.Get(userCtx)
	idInt, ok := id.(int)
	if !ok {
		generateErrorAller(http.StatusInternalServerError, "Server error", "Please try again later", nil, *&c)
		return
	}

	//Get image url
	url := c.Request.Header.Get("url")
	htmlStr, err := h.service.Pictures.GetImagePromts(url, idInt)
	if err != nil {
		generateErrorAller(http.StatusBadGateway, "DataBase fail", "please try again", err, *&c)
	}
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(c.Writer, nil)
	return
}

// Function for search when you just use search
func (h *Handler) searchGet(c *gin.Context) {
	//Get UserName
	idData, _ := c.Get(userCtx)
	id, check := idData.(int)
	if check == false {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", nil, *&c)
		return
	}
	userName, err := h.service.Pictures.GetUserName(id)
	if err != nil {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", err, *&c)
		return
	}

	//If promt null return to main page
	promt := c.Request.FormValue("promt")
	if promt == "" {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/pictures/")
	}

	//Get urls
	urls, err := h.service.Pictures.SearchImages(promt, 0)

	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
		return
	}

	//FOR SEARCH WE RENDER A NEW PAGE AND RENDER IMAGES AT THIS PAGE
	tmpl, _ := template.ParseFiles("../templates/search.html")

	urslHtml := make([]template.HTML, len(urls))
	for i := range urls {
		urslHtml[i] = template.HTML(urls[i])
	}

	data := map[string]interface{}{
		"Username": userName,
		"Urls":     urslHtml,
	}

	tmpl.Execute(c.Writer, data)
}

// Function for search when you scroll down
func (h *Handler) searchPost(c *gin.Context) {
	//Get lastImageId from header
	lastImageId := c.Request.Header.Get("lastImageId")
	lastImageIdInt, _ := strconv.Atoi(lastImageId)
	promt := c.Request.Header.Get("promt")

	//Get urls
	urls, err := h.service.Pictures.SearchImages(promt, lastImageIdInt)

	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
		return
	}

	for _, htmlStr := range urls {
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(c.Writer, nil)
	}

	return
}

// For add or delete like
func (h *Handler) addLikePost(c *gin.Context) {
	//Get user id
	id, _ := c.Get(userCtx)
	idInt, ok := id.(int)

	countLike := c.Request.Header.Get("countLike")
	countLikeInt, err := strconv.Atoi(countLike)
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Server error", "Please try again later", nil, *&c)
		return
	}

	//Check if user is anonym
	if !ok {
		generateErrorAller(http.StatusInternalServerError, "Server error", "Please try again later", nil, *&c)
		return
	}

	if idInt == anonymId {
		generateErrorAller(http.StatusBadRequest, "You a not log in", "Please log in before like it", nil, *&c)
		return

	} else {
		imageUrl := c.Request.Header.Get("url")
		htmlStr, err := h.service.Pictures.AddLike(imageUrl, idInt, countLikeInt)
		if err != nil {
			generateErrorAller(http.StatusInternalServerError, "Server error", "Please try again later", nil, *&c)
			return
		}

		// c.Writer.WriteHeader(302)
		// c.Header("test", "test")
		// c.HTML(http.StatusOK,"",gin.H{"stringSVG": htmlStr})
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(c.Writer, nil)

		return
	}

}
