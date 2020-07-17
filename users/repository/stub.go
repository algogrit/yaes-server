package repository

import "algogrit.com/yaes-server/entities"

type UserRepoStub struct {
	RetrieveOthersFn func(u entities.User) (users []*entities.User, err error)
	FindByFn         func(username string) (u *entities.User, err error)
	FindByIDFn       func(ID interface{}) (u *entities.User, err error)
	SaveFn           func(u entities.User) (user *entities.User, err error)
}

func (inmemRepo *UserRepoStub) RetrieveOthers(u entities.User) (users []*entities.User, err error) {
	if inmemRepo.RetrieveOthersFn != nil {
		return inmemRepo.RetrieveOthersFn(u)
	}
	return
}

func (inmemRepo *UserRepoStub) FindBy(username string) (u *entities.User, err error) {
	if inmemRepo.FindByFn != nil {
		return inmemRepo.FindByFn(username)
	}
	return
}

func (inmemRepo *UserRepoStub) FindByID(ID interface{}) (u *entities.User, err error) {
	if inmemRepo.FindByIDFn != nil {
		return inmemRepo.FindByIDFn(ID)
	}
	return
}

func (inmemRepo *UserRepoStub) Save(u entities.User) (user *entities.User, err error) {
	if inmemRepo.SaveFn != nil {
		return inmemRepo.SaveFn(u)
	}
	return
}
