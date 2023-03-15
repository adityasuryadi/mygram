package usecase

import (
	"mygram/commons/exceptions"
	domains "mygram/domains/user"
	userEntities "mygram/domains/user/entity"
	"mygram/domains/user/model"
	hashing "mygram/infrastructures/hash"
)

func NewUserUseCase(repository domains.UserRepository) domains.UserUsecase {
	return &UserUseCaseImpl{
		repository: repository,
	}
}

type UserUseCaseImpl struct {
	repository domains.UserRepository
}

// RegisterUser implements domains.UserUsecase
func (usecase *UserUseCaseImpl) RegisterUser(request model.RegisterUserRequest) {
	user := userEntities.User{
		UserName:  request.Username,
		Email:     request.Email,
		Password:  hashing.GetHash([]byte(request.Password)),
		Age:       request.Age,
	}
	err:=usecase.repository.Insert(user)
	if err != nil {
		exceptions.PanicIfNeeded(err)
	}
}

func (usecase *UserUseCaseImpl) FetchUserLogin() {
	panic("implement me")
}
