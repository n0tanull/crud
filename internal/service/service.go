package service

import (
	"context"
	"main/internal/entity"
	"main/internal/repo"

	"github.com/sirupsen/logrus"
)

type Servicer interface {
	GetUser(id int, ctx context.Context) (entity.User, error)
	AddUser(user entity.User, ctx context.Context) error
	DeleteUser(id int, ctx context.Context) error
	ChangeUser(user entity.User, ctx context.Context) error
}

type Serv struct {
	repo.Repo
	log *logrus.Logger
}

func NewServ(repo repo.Repo, l logrus.Logger) Servicer {
	return &Serv{
		Repo: repo,
		log:  &l,
	}
}

func (s *Serv) GetUser(id int, ctx context.Context) (entity.User, error) {
	return s.Repo.GetUser(id, ctx)
}

func (s *Serv) AddUser(user entity.User, ctx context.Context) error {
	return s.Repo.AddUser(user, ctx)
}
func (s *Serv) DeleteUser(id int, ctx context.Context) error {
	return s.Repo.DeleteUser(id, ctx)
}
func (s *Serv) ChangeUser(user entity.User, ctx context.Context) error {
	return s.Repo.ChangeUser(user, ctx)
}
