package repository

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"site/pkg/cache"
	"strings"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PicturesPostgres struct {
	db *sqlx.DB
}

// var (
// 	postCache cache.CacheImages
// )

const (
	countOfGetImages = 22
	likeEmty         = `
	<svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
		<path d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"></path>
	</svg>`
	likeFill = `
	<svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="currentColor" class="bi bi-hand-thumbs-up-fill" viewBox="0 0 16 16">
		<path d="M6.956 1.745C7.021.81 7.908.087 8.864.325l.261.066c.463.116.874.456 1.012.965.22.816.533 2.511.062 4.51a9.84 9.84 0 0 1 .443-.051c.713-.065 1.669-.072 2.516.21.518.173.994.681 1.2 1.273.184.532.16 1.162-.234 1.733.058.119.103.242.138.363.077.27.113.567.113.856 0 .289-.036.586-.113.856-.039.135-.09.273-.16.404.169.387.107.819-.003 1.148a3.163 3.163 0 0 1-.488.901c.054.152.076.312.076.465 0 .305-.089.625-.253.912C13.1 15.522 12.437 16 11.5 16H8c-.605 0-1.07-.081-1.466-.218a4.82 4.82 0 0 1-.97-.484l-.048-.03c-.504-.307-.999-.609-2.068-.722C2.682 14.464 2 13.846 2 13V9c0-.85.685-1.432 1.357-1.615.849-.232 1.574-.787 2.132-1.41.56-.627.914-1.28 1.039-1.639.199-.575.356-1.539.428-2.59z"/>
	</svg>`
)

func NewPicturesPostgres(db *sqlx.DB) *PicturesPostgres {
	return &PicturesPostgres{db: db}
}

// Help function that create a div for other functions. Get urls and return div data
func getStandartDivForImages(urls []string, lastImageId int, countRows int, urlForPost string, promt string, countOfImages int, target string) (urlsHtml []string, err error) {

	urlsHtml = make([]string, len(urls))
	//<img id="picture" src='../static/images/lowQuality/%s'

	//For last we need to add hx-trigger='revealed' for load new image when user scroll down
	for i, str := range urls {
		locStr := ""
		//if i == len(urls)-1 && len(urls) == countOfImages && lastImageId+countOfImages < countRows
		if i == len(urls)-1 && len(urls) == countOfImages && lastImageId+countOfImages < countRows {
			page := (lastImageId + countOfImages) / countOfImages
			locStr = fmt.Sprintf(`
			<div class='blur-load' style='background-image: url(../static/images/20pxImage/%s)'>
				<a hx-post="/pictures/info=%s" hx-headers='{"url": "%s"}'  hx-target='#overlay'>
					<img id="picture" src='../static/images/lowQuality/%s' loading='lazy' hx-post='http://localhost:8080/%s?page=%d' 
					hx-trigger='revealed' hx-swap='beforebegin' hx-headers='{"promt": "%s", "lastImageId": "%d"}' hx-target='#%s'/>
				</a>
			</div>`, str, str, str, str, urlForPost, page, promt, lastImageId+countOfImages, target)

		} else if str == "" {
			break

		} else {
			locStr = fmt.Sprintf(`
		<div class='blur-load' style='background-image: url(../static/images/20pxImage/%s)'>
			<a hx-post="/pictures/info=%s" hx-headers='{"url":"%s"}'  hx-target='#overlay'>
				<img id="picture" src='../static/images/lowQuality/%s' loading='lazy'>
			</a>
		</div>`, str, str, str, str)
		}
		urlsHtml[i] = locStr

	}

	return urlsHtml, nil
}

// Help function that create a div for other functions. Get urls and return div data
func getDivForImages(urls []string, lastImageId int, countRows int, urlForPost string, promt string, countOfImages int, target string) (urlsHtml []string, err error) {

	urlsHtml = make([]string, len(urls))
	//<img id="picture" src='../static/images/lowQuality/%s'

	//For last we need to add hx-trigger='revealed' for load new image when user scroll down
	for i, str := range urls {
		locStr := ""
		//if i == len(urls)-1 && len(urls) == countOfImages && lastImageId+countOfImages < countRows
		if i == len(urls)-1 && len(urls) == countOfImages && lastImageId+countOfImages < countRows {
			page := (lastImageId + countOfImages) / countOfImages
			locStr = fmt.Sprintf(`
			<div class='blur-load' style='background-image: url(../static/images/20pxImage/%s)'>
				<a hx-post="/pictures/info=%s" hx-headers='{"url": "%s"}'  hx-target='#overlay'>
					<img id="picture" src='data:image/jpg;base64,{{.Bytes}}' loading='lazy' hx-post='http://localhost:8080/%s?page=%d' 
					hx-trigger='revealed' hx-swap='beforebegin' hx-headers='{"promt": "%s", "lastImageId": "%d"}' hx-target='#%s'/>
				</a>
			</div>`, str, str, str, urlForPost, page, promt, lastImageId+countOfImages, target)

		} else if str == "" {
			break

		} else {
			locStr = fmt.Sprintf(`
		<div class='blur-load' style='background-image: url(../static/images/20pxImage/%s)'>
			<a hx-post="/pictures/info=%s" hx-headers='{"url":"%s"}'  hx-target='#overlay'>
				<img id="picture" src='data:image/jpg;base64,{{.Bytes}}' loading='lazy'>
			</a>
		</div>`, str, str, str)
		}
		urlsHtml[i] = locStr

	}

	return urlsHtml, nil
}

// Get userName
func (r *PicturesPostgres) GetUserName(id int) (string, error) {
	var username string
	query_ := fmt.Sprintf("SELECT username FROM %s WHERE id=$1", usersTable)
	row := r.db.QueryRow(query_, id)
	if err := row.Scan(&username); err != nil {
		return "Anonym", err
	}

	return username, nil
}

// Func that we use when user scroll down at home page
func (r *PicturesPostgres) GetNewImages(lastImageId int, postCache cache.CacheImages) (urlsHtml []string, imageNums []string, err error) {
	urls := make([]string, countOfGetImages)

	//Get count of rows for stop add hx-trigger='revealed' => stop load images
	var countRows int
	query_ := fmt.Sprintf("SELECT count(image_url) from %s", imagesTable)
	row := r.db.QueryRow(query_)
	if err := row.Scan(&countRows); err != nil {
		return nil, nil, err
	}

	//If it is start of load /pictures
	if lastImageId == 0 {
		query_ = fmt.Sprintf("SELECT image_url from %s ORDER BY like_count DESC,image_id LIMIT %d", imagesTable, countOfGetImages)

	} else {
		query_ = fmt.Sprintf("SELECT image_url from %s ORDER BY like_count DESC,image_id OFFSET %d LIMIT %d", imagesTable, lastImageId, countOfGetImages)
	}

	//SQL query
	rows, err := r.db.Query(query_)
	if err != nil {
		return nil, nil, err
	}
	i := 0
	for rows.Next() {
		//Get all urls
		var url string
		if err := rows.Scan(&url); err != nil {
			return nil, nil, err
		}
		urls[i] = url
		i++
	}

	urlsHtml, err = getDivForImages(urls, lastImageId, countRows, "pictures/", "", countOfGetImages, "grid")
	if err != nil {
		return nil, nil, err
	}

	//check does we have this image in redis
	//if no we save urls in urlsRedis for add in Redis
	urlsRedis := make([]string, len(urlsHtml))
	lastId := 0

	//../static/images/20pxImage/202310011511588352.jpg
	for _, v := range urls {
		if v == "" {
			break
		}
		check := postCache.CheckExist(v)
		if !check {
			urlsRedis[lastId] = v
			lastId++
		}
	}

	var wg sync.WaitGroup
	for _, v := range urlsRedis {
		if v != "" {
			wg.Add(1)

			go func(url string) {
				defer wg.Done()
				//open image
				file, err := os.ReadFile("../static/images/lowQuality/" + url)
				if err != nil {
					logrus.Errorf("Cant open file: %s", url)
					return
				}

				//Decode image into a go image object
				img, _, err := image.Decode(bytes.NewBuffer(file))
				if err != nil {
					logrus.Errorf("Cant decode file: %s", url)
					return
				}

				//encode img to slice of bytes
				var buf bytes.Buffer
				err = jpeg.Encode(&buf, img, nil)
				if err != nil {
					logrus.Errorf("Cant convert to bytes slice file: %s", url)
					return
				}
				imageBytes := buf.Bytes()

				stringBytes := base64.StdEncoding.EncodeToString(imageBytes)

				postCache.SetImageBytes(url, stringBytes)

				// sliceBytes, err := os.ReadFile("../static/images/lowQuality/" + url)
				// if err != nil {
				// 	logrus.Errorf("Cant open file: %s", url)
				// 	return
				// }

				// stringBytes := base64.StdEncoding.EncodeToString(sliceBytes)
				// postCache.SetImageBytes(url, stringBytes)

			}(v)
		}
	}
	wg.Wait()

	return urlsHtml, urls, nil
}

// Function for show user promts when user click on image
func (r *PicturesPostgres) GetImagePromts(imageUrl string, userId int) (string, error) {

	var countLikes int
	var promt string
	// query_ := fmt.Sprintf(`
	// SELECT promts.title, images.like_count
	// FROM %s,%s
	// WHERE promts.image_url LIKE '%s'
	// AND promts.image_url = images.image_url`, promtsTable, imagesTable, imageUrl)
	query_ := fmt.Sprintf(`
	SELECT title, like_count 
	FROM %s 
	WHERE image_url = '%s'`, imagesTable, imageUrl)

	likeSVG := ""

	//SQL query
	row := r.db.QueryRow(query_)
	if err := row.Scan(&promt, &countLikes); err != nil {
		return "", err
	}

	//Get info if user click like on image
	check, err := r.checkIfLike(imageUrl, userId)
	if err != nil {
		return "", err
	}

	//if check == false it means that user never like it or like and after delete like
	// and we use empty icon else we use fill icon
	if !check {
		likeSVG = likeEmty
	} else {
		likeSVG = likeFill
	}

	//SQL query
	// rows, err := r.db.Query(query_)
	// if err != nil {
	// 	return "", err
	// }

	// for rows.Next() {
	// 	//Get all promts
	// 	var promt string
	// 	if err := rows.Scan(&promt, &countLikes); err != nil {
	// 		return "", err
	// 	}
	// 	promts += promt + " "
	// }

	htmlStr := fmt.Sprintf(`
	<div id="overDiv">
		<style>
			body {overflow-y: hidden;}
		</style>
		<div id="gridInOverDiv">
			<div id="imageContainerInOverDiv">
				<img id="imgInOverDiv" src='../static/images/highQuality/%s'/>
			</div>
			<div id="textContainerInOverDiv">
				
				<p>Promts:<br />%s</p>
			</div>
			<div id="countLikesContainerInOverDiv" >
				<p>Count likes:<br />%d likes</p>
				<a hx-headers='{"url":"%s","countLike":"%d"}' hx-post="/pictures/addLike" hx-target="#countLikesContainerInOverDiv" >
					%s
              	</a>
				<div id="test"></div>
			</div>
			<div onclick="deleteOverDiv()">
				<a id="xIconInOverDiv">
					<svg  xmlns="http://www.w3.org/2000/svg" width="50" height="50" fill="currentColor" class="bi bi-x" viewBox="0 0 16 16">
						<path d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z"/>
					</svg>
				</a>
			</div>
			<div id="errorsAlertInOverDiv"></div>
		</div>
	</div>`, imageUrl, promt, countLikes, imageUrl, countLikes, likeSVG)

	return htmlStr, nil
}

// Func for search images
func (r *PicturesPostgres) SearchImages(promt string, lastImageId int) (urlsHtml []string, err error) {
	promtSlice := strings.Split(promt, " ")
	urlsSlice := make([]string, 0)

	//Get count of rows for stop add hx-trigger='revealed' => stop load images
	var countRows int
	query_ := fmt.Sprintf("SELECT count(image_url) from %s", imagesTable)
	row := r.db.QueryRow(query_)
	if err := row.Scan(&countRows); err != nil {
		return nil, err
	}

	//If it is start of load /pictures
	for _, val := range promtSlice {
		query_ := fmt.Sprintf("SELECT %s.image_url FROM %s,%s WHERE promts.title like '%s' AND promts.image_url = images.image_url ORDER BY images.like_count DESC,images.image_id OFFSET %d LIMIT %d",
			promtsTable, promtsTable, imagesTable, val, lastImageId, countOfGetImages)

		rows, err := r.db.Query(query_)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			//Get all urls
			var url string
			if err := rows.Scan(&url); err != nil {
				return nil, err
			}

			urlsSlice = append(urlsSlice, url)
		}
	}

	result, err := getStandartDivForImages(urlsSlice, lastImageId, countRows, "pictures/search", promt, countOfGetImages, "grid")
	if err != nil {
		return nil, err
	}

	return result, nil

}

// Func for check if user like it image
func (r *PicturesPostgres) checkIfLike(imageUrl string, userId int) (bool, error) {
	var check bool

	query_ := fmt.Sprintf(`
	SELECT like_check from %s 
	WHERE user_id = %d and
	image_url LIKE '%s'`, likesTable, userId, imageUrl)

	rows, err := r.db.Query(query_)
	if err != nil {
		return false, err
	}

	for rows.Next() {
		if err := rows.Scan(&check); err != nil {
			return false, err
		}
	}

	return check, nil
}

// Func start when user click on image and after click on icon like
func (r *PicturesPostgres) AddLike(imageUrl string, userId int, countLike int) (string, error) {
	result := ""
	//Get info like user this image
	check, err := r.checkIfLike(imageUrl, userId)
	if err != nil {
		return "", err
	}
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	//Check = false that mean user not like this image, and we need add rows in DB
	//for save like on image
	//Else we need to delete rows in db because no diffrent if user like it and after
	//delete like or never like it
	if !check {

		//Add info in likes table
		query_ := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2, $3)", likesTable)
		// r.db.QueryRow(query_, userId, imageUrl, true)
		_, err = tx.Exec(query_, userId, imageUrl, true)
		if err != nil {
			tx.Rollback()
			return "", err
		}
		tx.Commit()

		//Update count of like in images
		query_ = fmt.Sprintf(`
		UPDATE %s SET like_count = like_count + 1 
		WHERE image_url LIKE '%s'`, imagesTable, imageUrl)
		// r.db.Query(query_)
		_, err := r.db.Exec(query_)
		if err != nil {
			return "", err
		}

		result = fmt.Sprintf(`
		<p>Count likes:<br />%d likes</p>
		<a hx-headers='{"url":"%s","countLike":%d}' hx-post="/pictures/addLike" hx-target="#countLikesContainerInOverDiv">
			%s
	  	</a>`, countLike+1, imageUrl, countLike+1, likeFill)

	} else {

		//Update count of like in images
		query_ := fmt.Sprintf(`
				UPDATE %s SET like_count = like_count -1 
				WHERE image_url LIKE '%s'`, imagesTable, imageUrl)
		// r.db.Query(query_)
		_, err := r.db.Exec(query_)
		if err != nil {
			return "", err
		}

		//DELETE a row from table with likes
		query_ = fmt.Sprintf(`
		DELETE FROM %s
		WHERE user_id = %d AND
		image_url LIKE '%s'`, likesTable, userId, imageUrl)
		// r.db.Query(query_)
		_, err = r.db.Exec(query_)
		if err != nil {
			return "", err
		}

		result = fmt.Sprintf(`
		<p>Count likes:<br />%d likes</p>
		<a hx-headers='{"url":"%s","countLike":%d}' hx-post="/pictures/addLike" hx-target="#countLikesContainerInOverDiv">
			%s
	  	</a>`, countLike-1, imageUrl, countLike-1, likeEmty)
	}

	return result, nil
}
