package services

import (
	"authservice/models"
	repositories "authservice/repositories/auth"
	"authservice/schemas"
)

type ServiceActivation interface {
	ActivationService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type serviceActivation struct {
	repository repositories.RepositoryActivation
}

func NewServiceActivation(repository repositories.RepositoryActivation) *serviceActivation {
	return &serviceActivation{repository: repository}
}

func (s *serviceActivation) ActivationService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

	var schema schemas.SchemaAuth
	schema.Email = input.Email
	schema.Active = input.Active
	schema.Token = input.Token

	res, err := s.repository.ActivationRepository(&schema)
	return res, err
}
