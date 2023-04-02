package user

import (
	"context"
	"fmt"
	"haste/adapters/repositories"
	db "haste/infra/driven/database/sqlc"

	repositoryPort "haste/core/ports/repositories"
)

type UserRepository struct {
	repositories.BaseRepository
}

func (r *UserRepository) CreatelUsersInDB(ctx context.Context, arg db.CreateUserParams) (repositoryPort.User, error) {
	user, err := repositories.TaskQueries.CreateUser(ctx, arg)
	return repositoryPort.User(user), err
}

func (r *UserRepository) GetAllUsersInDB(ctx context.Context) ([]repositoryPort.User, error) {
	var result []repositoryPort.User
	fmt.Println(ctx)
	return result, nil
}

func init() {
	fmt.Println("user repository---")
	repositories.RepositoryMap["User"] = func(br *repositories.BaseRepository) interface{} {
		r := &UserRepository{
			BaseRepository: *br,
		}
		return repositoryPort.UserPort(r)
	}
}
