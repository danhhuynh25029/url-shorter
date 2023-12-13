package repository

type ILinkRepository interface {
	Insert() error
	Get() error
}

func NewLinkRepository() ILinkRepository {
	return &linkRepositoryImpl{}
}

type linkRepositoryImpl struct {
}

func (l *linkRepositoryImpl) Insert() error {
	return nil
}

func (l *linkRepositoryImpl) Get() error {
	return nil
}
