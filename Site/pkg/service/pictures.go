package service

import (
	"site/pkg/repository"
	"strings"
)

type PicturesService struct {
	repo repository.Pictures
}

func NewPicturesService(repo repository.Pictures) *PicturesService {
	return &PicturesService{repo: repo}
}
func (s *PicturesService) GetUserName(id int) (string, error) {
	return s.repo.GetUserName(id)
}

func (s *PicturesService) GetNewImages(lastImageId int) (urlsHtml []string, err error) {
	return s.repo.GetNewImages(lastImageId)
}

func (s *PicturesService) GetImagePromts(imageUrl string, userId int) (string, error) {
	return s.repo.GetImagePromts(imageUrl, userId)
}

func (s *PicturesService) SearchImages(promt string, lastImageId int) (urlsHtml []string, err error) {
	result := make([]string, 0)
	word := ""

	//Use ASCII table to split if is not letter, number, special symbols : '+', '-', '(', ')'
	strings.ToLower(promt)
	for _, val := range promt {
		if (val < 97 || val > 122) && (val < 48 || val > 57) && val != 45 && val != 43 && val != 41 && val != 40 {
			if word != "" {
				result = append(result, word)
				word = ""
			}

		} else {
			word += string(val)
		}
	}
	if word != "" {
		result = append(result, word)
	}

	return s.repo.SearchImages(word, lastImageId)
}

func (s *PicturesService) AddLike(imageUrl string, userId int, countLike int) (string, error) {
	return s.repo.AddLike(imageUrl, userId, countLike)
}
