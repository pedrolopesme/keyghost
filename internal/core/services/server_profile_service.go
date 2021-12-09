package services

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pedrolopesme/keyghost/internal/core/domain"
)

type ServerProfileService struct{}

func NewServerProfileService() *ServerProfileService {
	return &ServerProfileService{}
}

func (s ServerProfileService) LoadServerProfile(path string) (*domain.ServerProfile, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	fileBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var serverProfile domain.ServerProfile
	json.Unmarshal(fileBytes, &serverProfile)
	return &serverProfile, nil
}
