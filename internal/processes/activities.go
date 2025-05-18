package processes

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	repo "github.com/go-microfrontend/auth/internal/repository"
)

type Repo interface {
	CreateUser(ctx context.Context, arg repo.CreateUserParams) (repo.User, error)
	GetUserByEmail(ctx context.Context, email string) (repo.User, error)
}

type Activities struct {
	repo Repo
}

func New(repo Repo) *Activities {
	return &Activities{repo: repo}
}

type userInput struct {
	Email    string
	Password string
}

func (a *Activities) CreateUser(ctx context.Context, input userInput) (repo.User, error) {
	b, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	arg := repo.CreateUserParams{
		Email:        input.Email,
		PasswordHash: string(b),
	}
	return a.repo.CreateUser(ctx, arg)
}

func (a *Activities) GetUserByEmail(ctx context.Context, email string) (repo.User, error) {
	return a.repo.GetUserByEmail(ctx, email)
}

type hashInput struct {
	Hash     string
	Password string
}

func (a *Activities) CheckHash(ctx context.Context, input hashInput) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(input.Hash), []byte(input.Password))
	if err != nil {
		return false, nil
	}

	return true, nil
}
