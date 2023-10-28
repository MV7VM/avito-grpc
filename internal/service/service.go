package service

import (
	"avito/common"
	"fmt"
	"regexp"
)

type repository interface {
	CreateLink(msg string) (string, error)
	GetLink(id string) (string, error)
}

type Service struct {
	repository
}

func New(repository repository) *Service {
	return &Service{repository}
}

func (s *Service) Create(msg string) (string, error) {
	if ok := Validate(msg); !ok {
		fmt.Println("Validate - ", msg, ok)
		return "", common.ErrNotUrl
	} else {
		return s.repository.CreateLink(msg)
	}
}

func (s *Service) Get(id string) (string, error) {
	return s.repository.GetLink(id)
}

func Validate(msg string) bool {
	// Регулярное выражение для URL
	regex := regexp.MustCompile("^(http(s)?://)?(www.)?[a-zA-Z0-9.-]+.[a-zA-Z]{2,6}(/[a-zA-Z0-9.-]*)*$")
	return regex.MatchString(msg)
}
