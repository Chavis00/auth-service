package services

import (
	"authservice/models"
	repositories "authservice/repositories/auth"
	"authservice/schemas"
)

type ServiceLogin interface {
	LoginService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type serviceLogin struct {
	repository repositories.RepositoryLogin
}

func NewServiceLogin(repository repositories.RepositoryLogin) *serviceLogin {
	return &serviceLogin{repository: repository}
}

func (s *serviceLogin) LoginService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

	var schema schemas.SchemaAuth
	schema.Email = input.Email
	schema.Password = input.Password
	res, err := s.repository.LoginRepository(&schema)
	return res, err
}
