package repository

type IRepository interface {
	UserRepository() IUserRepository
}

type repositoryImpl struct {
	userRepository IUserRepository
}

func NewRepository() IRepository {
	return &repositoryImpl{
		userRepository: NewUserRepository(),
	}
}

func (u *repositoryImpl) UserRepository() IUserRepository {
	return u.userRepository
}
