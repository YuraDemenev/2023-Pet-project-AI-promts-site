package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

const countOfGetImagesProfile = 12

type ProfilePostgres struct {
	db *sqlx.DB
}

func NewProfilePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

// Func that we use when user open profile page
func (r *ProfilePostgres) GetNewImagesProfile(lastImageId int, userId int) (urlsHtmlImagesLike []string, urlsHtmlImagesUpload []string, err error) {
	//Get count of rows for stop add hx-trigger='revealed' => stop load images
	var countRows int
	query_ := fmt.Sprintf("SELECT count(image_url) from %s WHERE user_id=%d", imagesTable, userId)
	row := r.db.QueryRow(query_)
	if err := row.Scan(&countRows); err != nil {
		return nil, nil, err
	}

	//PART FOR IMAGE THAT USER UPLOAD
	query_ = fmt.Sprintf(`
	SELECT image_url from %s
	WHERE user_id=%d
	ORDER BY like_count DESC,image_id 
	OFFSET %d LIMIT %d`, imagesTable, userId, lastImageId, countOfGetImagesProfile)

	var url string
	urlsHtmlImagesUpload, err = r.sqlQueryURL(query_, countOfGetImagesProfile, url)
	if err != nil {
		return nil, nil, err
	}

	//We use header {promt:upload} for recognize what images we need send
	//when user scroll down in profile
	urlsHtmlImagesUpload, err = profileGetDivForImages(urlsHtmlImagesUpload, lastImageId, countRows, "pictures/profile", "upload",
		countOfGetImagesProfile, "targetUpload")
	if err != nil {
		return nil, nil, err
	}

	//PART FOR IMAGE THAT USER LIKE
	query_ = fmt.Sprintf(`
	SELECT images.image_url from %s,%s
	WHERE likes.like_check = true AND
	likes.user_id = %d AND
	likes.image_url= images.image_url
	ORDER BY images.like_count DESC,images.image_id 
	OFFSET %d LIMIT %d
	`, imagesTable, likesTable, userId, lastImageId, countOfGetImagesProfile)

	urlsHtmlImagesLike, err = r.sqlQueryURL(query_, countOfGetImagesProfile, url)
	if err != nil {
		return nil, nil, err
	}

	//We use header {promt:like} for recognize what images we need send
	//when user scroll down in profile
	urlsHtmlImagesLike, err = profileGetDivForImages(urlsHtmlImagesLike, lastImageId, countRows,
		"pictures/profile", "like", countOfGetImagesProfile, "targetLike")
	if err != nil {
		return nil, nil, err
	}

	return urlsHtmlImagesLike, urlsHtmlImagesUpload, nil
}

// Func that we use when user scroll down at profile page
func (r *ProfilePostgres) GetNewImages(lastImageId int, userId int, check string) (urlsHtml []string, err error) {
	//Get count of rows for stop add hx-trigger='revealed' => stop load images
	var countRows int
	query_ := fmt.Sprintf("SELECT count(image_url) from %s WHERE user_id=%d", imagesTable, userId)
	row := r.db.QueryRow(query_)
	if err := row.Scan(&countRows); err != nil {
		return nil, err
	}

	//query if we need send upload pictures
	if check == "upload" {
		query_ = fmt.Sprintf(`
		SELECT image_url from %s
		WHERE user_id=%d
		ORDER BY like_count DESC,image_id 
		OFFSET %d LIMIT %d`, imagesTable, userId, lastImageId, countOfGetImagesProfile)
	} else {
		//query if we need send like pictures
		query_ = fmt.Sprintf(`
		SELECT images.image_url from %s,%s
		WHERE likes.like_check = true AND
		likes.user_id = %d AND
		likes.image_url= images.image_url
		ORDER BY images.like_count DESC,images.image_id 
		OFFSET %d LIMIT %d
		`, imagesTable, likesTable, userId, lastImageId, countOfGetImagesProfile)
	}

	var url string
	urlsHtml, err = r.sqlQueryURL(query_, countOfGetImagesProfile, url)
	if err != nil {
		return nil, err
	}

	if check == "upload" {
		//We use header {promt:upload} for recognize what images we need send
		//when user scroll down in profile
		urlsHtml, err = profileGetDivForImages(urlsHtml, lastImageId, countRows, "pictures/profile/", "upload", countOfGetImagesProfile, "targetUpload")
		if err != nil {
			return nil, err
		}
	} else {
		//We use header {promt:like} for recognize what images we need send
		//when user scroll down in profile
		urlsHtml, err = profileGetDivForImages(urlsHtml, lastImageId, countRows, "pictures/profile/", "like", countOfGetImagesProfile, "targetLike")
		if err != nil {
			return nil, err
		}
	}

	return urlsHtml, nil

}

// Get id return count likes,image that user uploaded and count likes on his images
func (r *ProfilePostgres) GetUserInfo(userId int) (countLikes int, countUploaded int, CountLikesOnUploaded int, err error) {
	//Get count likes
	query_ := fmt.Sprintf("SELECT count(image_url) FROM %s WHERE user_id = %d", likesTable, userId)
	row := r.db.QueryRow(query_)
	if err := row.Scan(&countLikes); err != nil {
		return 0, 0, 0, err
	}

	//Check does also user have got upload images
	var check bool
	query_ = fmt.Sprintf(`
	SELECT EXISTS (
		SELECT 1
		FROM %s
		WHERE user_id = %d
	  )`, imagesTable, userId)

	row = r.db.QueryRow(query_)
	if err := row.Scan(&check); err != nil {
		return 0, 0, 0, err
	}

	//If true get other info
	if check {

		//Get count upload
		query_ = fmt.Sprintf("SELECT count(image_url) FROM %s WHERE user_id = %d", imagesTable, userId)
		row = r.db.QueryRow(query_)
		if err := row.Scan(&countUploaded); err != nil {
			return 0, 0, 0, err
		}

		//Get count likes on user uploaded images
		query_ = fmt.Sprintf("SELECT SUM(like_count) FROM %s WHERE user_id = %d", imagesTable, userId)
		row = r.db.QueryRow(query_)
		if err := row.Scan(&CountLikesOnUploaded); err != nil {
			return 0, 0, 0, err
		}
	}

	return countLikes, countUploaded, CountLikesOnUploaded, nil
}

// Help func to do sql query and get urls
func (r *ProfilePostgres) sqlQueryURL(query_ string, lengthOfSlice int, url string) (htmlSlice []string, err error) {
	htmlSlice = make([]string, lengthOfSlice)
	//SQL query
	rows, err := r.db.Query(query_)
	if err != nil {
		return nil, err
	}
	i := 0
	for rows.Next() {
		//Get all urls
		if err := rows.Scan(&url); err != nil {
			return nil, err
		}
		htmlSlice[i] = url
		i++
	}

	return htmlSlice, nil

}

func (r *ProfilePostgres) GetConsiderationImages(lastImageId int) (urlsHtml []string, err error) {
	//Get count of rows for stop add hx-trigger='revealed' => stop load images
	var countRows int
	query_ := fmt.Sprintf("SELECT count(image_url) from %s ", considerationTable)
	row := r.db.QueryRow(query_)
	if err := row.Scan(&countRows); err != nil {
		return nil, err
	}

	query_ = fmt.Sprintf(`
		SELECT image_url from %s
		OFFSET %d LIMIT %d`, considerationTable, lastImageId, countOfGetImagesProfile)

	var url string
	urlsHtml, err = r.sqlQueryURL(query_, countOfGetImagesProfile, url)
	if err != nil {
		return nil, err
	}

	urlsHtml, err = profileGetDivForImages(urlsHtml, lastImageId, countRows, "pictures/profile", "Admin", countOfGetImagesProfile, "targetApprove")
	if err != nil {
		return nil, err
	}

	return urlsHtml, nil
}

func (r *ProfilePostgres) GetAdminInfo() (countImages int, err error) {
	var countLikes int
	//Get count likes
	query_ := fmt.Sprintf("SELECT count(image_url) FROM %s ", considerationTable)
	row := r.db.QueryRow(query_)
	if err := row.Scan(&countLikes); err != nil {
		return 0, err
	}

	return countLikes, nil
}

// This function for admin for approve
// or disaprove images at his profile
func (r *ProfilePostgres) GetApproveButtons(imageUrl string) (HtmlStr string, err error) {
	var promt string
	query_ := fmt.Sprintf(`
	SELECT title from %s
	WHERE image_url = '%s'`, considerationTable, imageUrl)

	row := r.db.QueryRow(query_)
	if err := row.Scan(&promt); err != nil {
		return "", err
	}

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

			<div onclick="deleteOverDiv();">
				<a id="xIconInOverDiv">
					<svg  xmlns="http://www.w3.org/2000/svg" width="50" height="50" fill="currentColor" class="bi bi-x" viewBox="0 0 16 16">
						<path d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z"/>
					</svg>
				</a>
			</div>
			
		</div>

		<div id="buttonsInOverDiv">
	
			<button id="considerButtonOverDiv" type="button" class="btn btn-primary" 
			hx-headers='{"status": "Yes", "url":"%s"}'hx-post="/pictures/profile/consider" 
			hx-target='#errorsAlertInOverDiv'>
				<span class="spinner-border spinner-border-sm htmx-indicator" id="spinner" role="status"
					aria-hidden="true"></span>
				Yes
			</button>
				
			<button id="considerButtonOverDiv" type="button" class="btn btn-danger" 
			hx-headers='{"status": "No", "url":"%s"}' hx-post="/pictures/profile/consider" 
			hx-target='#errorsAlertInOverDiv'>
				<span class="spinner-border spinner-border-sm htmx-indicator" id="spinner" role="status"
					aria-hidden="true"></span>
				No
			</button>	

			<div id="errorsAlertInOverDiv"></div>
		</div>

		

	</div>`, imageUrl, promt, imageUrl, imageUrl)

	return htmlStr, nil
}

// For admin to delete or add in main DB image
func (r *ProfilePostgres) ConsiderImageAdmin(url string, status string) (string, error) {
	//If yes we save info in images table
	promt := ""
	if status == "Yes" {
		//First get promt to save in images
		query_ := fmt.Sprintf(`
		SELECT title from %s
		WHERE image_url = '%s'`, considerationTable, url)

		row := r.db.QueryRow(query_)
		if err := row.Scan(&promt); err != nil {
			return "", err
		}

		//Get user id that add image
		var id int
		query_ = fmt.Sprintf(`
		SELECT user_id from %s
		WHERE image_url = '%s'`, considerationTable, url)

		row = r.db.QueryRow(query_)
		if err := row.Scan(&id); err != nil {
			return "", err
		}

		//Add in image table
		tx, err := r.db.Begin()
		if err != nil {
			tx.Rollback()
			return "", err
		}
		query := fmt.Sprintf(`
		INSERT INTO %s (user_id, image_url,like_count,title) 
		VALUES ($1, $2, $3, $4)`, imagesTable)

		_, err = tx.Exec(query, id, url, 0, promt)
		if err != nil {
			tx.Rollback()
			return "", err
		}
		tx.Commit()

	} else {
		pathHigh := fmt.Sprintf("../static/images/highQuality/%s", url)
		pathLow := fmt.Sprintf("../static/images/lowQuality/%s", url)
		path20px := fmt.Sprintf("../static/images/20pxImage/%s", url)

		err := os.Remove(pathHigh)
		if err != nil {
			return "", err
		}

		err = os.Remove(pathLow)
		if err != nil {
			return "", err
		}

		err = os.Remove(path20px)
		if err != nil {
			return "", err
		}
	}
	//If no or yes we alse need to delete
	//From consideration tabel data
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return "", err
	}
	query := fmt.Sprintf(`
	DELETE FROM %s 
	WHERE image_url = '%s'
	`, considerationTable, url)

	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()

	return promt, nil
}

func (r *ProfilePostgres) GetNewImagesAdmin(lastImageId int, userId int) (urlsHtml []string, err error) {
	//Get count of rows for stop add hx-trigger='revealed' => stop load images
	var countRows int
	query_ := fmt.Sprintf("SELECT count(image_url) from %s ", considerationTable)
	row := r.db.QueryRow(query_)
	if err := row.Scan(&countRows); err != nil {
		return nil, err
	}

	//SQL query
	query_ = fmt.Sprintf(`
		SELECT image_url from %s
		ORDER BY id 
		OFFSET %d LIMIT %d`, considerationTable, lastImageId, countOfGetImagesProfile)

	var url string
	urlsHtml, err = r.sqlQueryURL(query_, countOfGetImagesProfile, url)
	if err != nil {
		return nil, err
	}

	//We use header {promt:upload} for recognize what images we need send
	//when user scroll down in profile
	urlsHtml, err = profileGetDivForImages(urlsHtml, lastImageId, countRows, "pictures/profile", "upload", countOfGetImagesProfile, "targetUpload")
	if err != nil {
		return nil, err
	}

	return urlsHtml, nil
}

func profileGetDivForImages(urls []string, lastImageId int, countRows int, urlForPost string, promt string, countOfImages int, target string) (urlsHtml []string, err error) {

	urlsHtml = make([]string, len(urls))
	//<img id="picture" src='../static/images/lowQuality/%s'

	//For last we need to add hx-trigger='revealed' for load new image when user scroll down
	for i, str := range urls {
		locStr := ""
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
