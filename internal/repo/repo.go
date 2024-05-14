package repo

import (
	"context"
	"main/internal/entity"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

type Repo interface {
	GetUser(id int, ctx context.Context) (entity.User, error)
	AddUser(user entity.User, ctx context.Context) error
	DeleteUser(id int, ctx context.Context) error
	ChangeUser(user entity.User, ctx context.Context) error
}

type Repository struct {
	conn *pgx.Conn
	log  logrus.Logger
}

func NewRepository(ctx context.Context, l logrus.Logger) *Repository {
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		l.Fatal(err)
		panic(err)
	}
	return &Repository{conn: conn}
}
func (r *Repository) GetUser(id int, ctx context.Context) (entity.User, error) {

	res, err := r.conn.Query(ctx, "SELECT * FROM users")
	if err != nil {
		r.log.Error(err)
		return entity.User{}, err
	}
	for res.Next() {
		var user entity.User
		err = res.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			r.log.Error(err)
			return entity.User{}, err
		}
	}
	return res, nil
}
func (r *Repository) AddUser(user entity.User, ctx context.Context) error {
	_, err := r.conn.Exec(ctx, "INSERT INTO users (name) VALUES ($1)")
	if err != nil {
		r.log.Error(err)
		return err
	}
	return nil
}

func (r *Repository) DeleteUser(id int, ctx context.Context) error {
	_, err := r.conn.Exec(ctx, "DELETE FROM users WHERE id = $1")
	if err != nil {
		r.log.Error(err)
		return err
	}
	return nil
}

func (r *Repository) ChangeUser(user entity.User, ctx context.Context) error {
	_, err := r.conn.Query(ctx, "INSERT INTO names WHERE id = $1")
	if err != nil {
		r.log.Error(err)
		return err
	}
	return nil
}

func (r *Repository) migrate() error {

}
