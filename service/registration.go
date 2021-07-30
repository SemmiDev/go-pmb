package service

import (
	"github.com/google/uuid"
	"go-clean/entity"
	"go-clean/model"
	util2 "go-clean/util"
	"log"
	"time"
)

type service struct {
	RegistrationRepository entity.RegistrationRepository
}

func NewRegistrationService(rp *entity.RegistrationRepository) entity.RegistrationService {
	return &service{
		RegistrationRepository: *rp,
	}
}

func (s *service) Create(request *model.RegistrationRequest, program entity.Program) (*model.RegistrationResponse, error) {
	username, _ := util2.Hash(uuid.NewString())
	password, _ := util2.Hash(uuid.NewString())

	register := entity.Registration{
		ID:        uuid.NewString(),
		Name:      request.Name,
		Email:     request.Email,
		Username:  username,
		Password:  password,
		Kind:      program,
		CreatedAt: time.Now().Unix(),
	}

	err := s.RegistrationRepository.Insert(&register)
	if err != nil {
		log.Printf("Service.Create: %v", err.Error())
		return nil, err
	}

	response := model.RegistrationResponse{
		Username: username,
		Password: password,
	}

	return &response, nil
}
