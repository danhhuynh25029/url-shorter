package repository

type IUserRepository interface {
	Insert() error
	Get() error
}

type userRepositoryImpl struct {
}

func NewUserRepository() IUserRepository {
	return &userRepositoryImpl{}
}

func (u *userRepositoryImpl) Insert() error {
	return nil
}

func (u *userRepositoryImpl) Get() error {
	return nil
}
