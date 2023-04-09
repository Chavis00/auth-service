package services

import (
	repositories "authservice/repositories/auth"

	"authservice/models"

	"authservice/schemas"
)

type ServiceRegister interface {
	RegisterService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type serviceRegister struct {
	repository repositories.RepositoryRegister
}

func NewServiceRegister(repository repositories.RepositoryRegister) *serviceRegister {
	return &serviceRegister{repository: repository}
}

func (s *serviceRegister) RegisterService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {
	var schema schemas.SchemaAuth
	schema.Fullname = input.Fullname
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.repository.RegisterRepository(&schema)
	if err != (schemas.SchemaDatabaseError{}) {
		return nil, err
	}
	return res, schemas.SchemaDatabaseError{}
}
