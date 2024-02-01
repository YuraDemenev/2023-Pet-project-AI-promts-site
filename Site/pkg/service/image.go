package service

import (
	"math/rand"
	"site/pkg/repository"
	"strconv"
	"strings"
	"time"
)

type ImageService struct {
	repo repository.Image
}

func NewImageService(repo repository.Image) *ImageService {
	return &ImageService{repo: repo}
}

func (s *ImageService) SaveImageLink(id int, imageLink string) error {
	error_ := s.repo.SaveLink(id, imageLink)
	return error_
}

func (s *ImageService) SavePromtsToPromts(promts []string, imageLink string) error {
	error_ := s.repo.SavePromtsToPromts(promts, imageLink)
	return error_
}

func (s *ImageService) SavePromtsToConsideration(promts string, imageUrl string, userId int) error {
	error_ := s.repo.SavePromtsToConsideration(promts, imageUrl, userId)
	return error_
}

func (s *ImageService) GetPromts(prompt string) []string {
	result := make([]string, 0)
	word := ""
	dictionary := make(map[string]int, 0)
	dictionaryTrashWord := map[string]int{
		"and": 0,
		"or":  0,
		"of":  0,
		"a":   0,
	}

	//Use ASCII table to split if is not letter, number, special symbols : '+', '-', '(', ')'
	lower := strings.ToLower(prompt)
	for _, val := range lower {
		if (val < 97 || val > 122) && (val < 48 || val > 57) && val != 45 && val != 43 && val != 41 && val != 40 {
			if word != "" {
				_, ok := dictionary[word]
				_, ok1 := dictionaryTrashWord[word]
				if !ok && !ok1 {
					dictionary[word] = 0
					result = append(result, word)
				}
				word = ""
			}

		} else {
			word += string(val)
		}
	}
	if word != "" {
		result = append(result, word)
	}

	return result
}

func (s *ImageService) GenerateNewImageName() string {
	t := time.Now()
	randomInt := rand.Intn(10001)
	newName := t.Format("20060102150405") + strconv.Itoa(randomInt) + ".jpg"

	return newName
}
