package services

import (
	repositories "authservice/repositories/auth"
)

type ServicePing interface {
	PingService() string
}

type servicePing struct {
	repository repositories.RepositoryPing
}

func NewServicePing(repository repositories.RepositoryPing) *servicePing {
	return &servicePing{repository: repository}
}

func (s *servicePing) PingService() string {
	res := s.repository.PingRepository()
	return res
}
