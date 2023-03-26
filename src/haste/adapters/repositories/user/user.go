package user

import (
	"fmt"
	"haste/adapters/repositories"

	repositoryPort "haste/core/ports/repositories"
)

type UserRepository struct {
	repositories.BaseRepository
}

func (r *UserRepository) GetAllUsersInDB() []*repositoryPort.User {
	var result []*repositoryPort.User

	for i := 1; i <= 5; i++ {
		result = append(result, &repositoryPort.User{ID: fmt.Sprintf("%d", i), Name: "asdf"})
	}
	return result
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
