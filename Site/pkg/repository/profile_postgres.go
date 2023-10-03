package repository

import (
	"fmt"

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
	urlsHtmlImagesUpload, err = getDivForImages(urlsHtmlImagesUpload, lastImageId, countRows, "pictures/profile", "upload",
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
	urlsHtmlImagesLike, err = getDivForImages(urlsHtmlImagesLike, lastImageId, countRows,
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
		urlsHtml, err = getDivForImages(urlsHtml, lastImageId, countRows, "pictures/profile", "upload", countOfGetImagesProfile, "targetUpload")
		if err != nil {
			return nil, err
		}
	} else {
		//We use header {promt:like} for recognize what images we need send
		//when user scroll down in profile
		urlsHtml, err = getDivForImages(urlsHtml, lastImageId, countRows, "pictures/profile", "like", countOfGetImagesProfile, "targetLike")
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
