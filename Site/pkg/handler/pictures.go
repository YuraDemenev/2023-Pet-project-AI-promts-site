package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type lastImageIdData struct {
	Id string `json:"lastImageId"`
}

// For check exist image in redis or no
func (h *Handler) GetBytes(imageNums []string, imageBytes *[]string) {
	//checking does imageNums exist in redis
	var wg sync.WaitGroup
	for i, num := range imageNums {
		if num == "" {
			continue
		}
		wg.Add(1)
		go func(i int, num string) {

			defer wg.Done()
			// localCheck := h.postCache.CheckExist(num)
			// //if not exist add bytes slice to redis
			// if !localCheck {
			// 	//open image
			// 	file, err := os.Open("../static/images/lowQuality/" + num)
			// 	if err != nil {
			// 		logrus.Errorf("Cant open file: %s", num)
			// 		return
			// 	}

			// 	//Decode image into a go image object
			// 	img, _, err := image.Decode(file)
			// 	if err != nil {
			// 		logrus.Errorf("Cant decode file: %s", num)
			// 		return
			// 	}

			// 	//encode img to slice of bytes
			// 	buf := new(bytes.Buffer)
			// 	err = jpeg.Encode(buf, img, nil)
			// 	if err != nil {
			// 		logrus.Errorf("Cant convert to bytes slice file: %s", num)
			// 		return
			// 	}
			// 	imageBytes := buf.Bytes()
			// 	stringBytes := base64.StdEncoding.EncodeToString(imageBytes)

			// 	h.postCache.SetImageBytes(num, stringBytes)
			// }
			(*imageBytes)[i] = h.postCache.GetImageBytes(num)
		}(i, num)

	}
	wg.Wait()

}

// Function for home page when you scroll down
func (h *Handler) watchPicturePost(c *gin.Context) {
	//Get lastImageId from header
	lastImageId := c.Request.Header.Get("lastImageId")
	lastImageIdInt, _ := strconv.Atoi(lastImageId)

	//Get new urls
	urls, imageNums, err := h.service.Pictures.GetNewImages(lastImageIdInt, h.postCache)

	imageBytes := make([]string, len(urls))
	h.GetBytes(imageNums, &imageBytes)

	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
	}

	//Add bytes in img src {{.Bytes}}
	var wg sync.WaitGroup
	for i := range imageBytes {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			t, err := template.New("Bytes").Parse(urls[i])
			if err != nil {
				generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
			}

			dataBytes := map[string]string{
				"Bytes": imageBytes[i],
			}

			builder := &strings.Builder{}
			err = t.Execute(builder, dataBytes)
			if err != nil {
				generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
			}

			urls[i] = builder.String()
		}(i)
	}
	wg.Wait()

	for _, htmlStr := range urls {
		if htmlStr == "" {
			break
		}
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(c.Writer, nil)
	}

}

// Function for home page when you just click Home
func (h *Handler) watchPictureGet(c *gin.Context) {
	//Get user user name
	idData, _ := c.Get(userCtx)
	id, check := idData.(int)
	if !check {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", nil, *&c)
		return
	}
	userName, err := h.service.Pictures.GetUserName(id)
	if err != nil {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", err, *&c)
		return
	}
	tmpl, err := template.ParseFiles("../templates/pictures.html")

	if err != nil {
		logrus.Fatalf("Problem parse pictures.html :%s", err.Error())
		return
	}

	//Get Urls

	urls, imageNums, err := h.service.Pictures.GetNewImages(0, h.postCache)
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
	}

	imageBytes := make([]string, len(urls))
	//Check here if image exist in redis. if no, add in redis
	h.GetBytes(imageNums, &imageBytes)
	urslHtml := make([]template.HTML, len(urls))

	//Sync execute bytes to img src
	var wg sync.WaitGroup
	for i := range imageBytes {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			t, err := template.New("Bytes").Parse(urls[i])
			if err != nil {
				generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
			}

			dataBytes := map[string]string{
				"Bytes": imageBytes[i],
			}

			builder := &strings.Builder{}
			err = t.Execute(builder, dataBytes)
			if err != nil {
				generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
			}

			urslHtml[i] = template.HTML(builder.String())
		}(i)
	}
	wg.Wait()

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

	//If we admin we have got another logic
	if id == adminId {
		//Get image url
		url := c.Request.Header.Get("url")
		htmlStr, err := h.service.Profile.GetApproveButtons(url)
		if err != nil {
			generateErrorAller(http.StatusBadGateway, "DataBase fail", "please try again", err, *&c)
		}
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(c.Writer, nil)

	} else {

		//Get image url
		url := c.Request.Header.Get("url")
		htmlStr, err := h.service.Pictures.GetImagePromts(url, idInt)
		if err != nil {
			generateErrorAller(http.StatusBadGateway, "DataBase fail", "please try again", err, *&c)
		}
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(c.Writer, nil)
	}

}

// Function for search when you just use search
func (h *Handler) searchGet(c *gin.Context) {
	//Get UserName
	idData, _ := c.Get(userCtx)
	id, check := idData.(int)
	if !check {
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
		c.Redirect(http.StatusMovedPermanently, "https://imagepromts.ru/pictures/")
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

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(c.Writer, nil)

		return
	}

}
