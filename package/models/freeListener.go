package models

import "gorm.io/gorm"

type FreeListenerCore struct {
	UserCore
}

type FreeListenerDB struct {
	gorm.Model
	UserDB
}

type FreeListenerHttp struct {
	UserHttp
}

func (em *FreeListenerDB) ToCore() *FreeListenerCore {
	return &FreeListenerCore{
		UserCore{
			Email:      em.Email,
			Password:   em.Password,
			Role:       Role(em.Role),
			Nickname:   em.Nickname,
			Firstname:  em.Firstname,
			Lastname:   em.Lastname,
			Middlename: em.Middlename,
			CreatedAt:  em.CreatedAt.String(),
		},
	}
}

func (em *FreeListenerDB) FromCore(freeListener *FreeListenerCore) {
	em.Email = freeListener.Email
	em.Password = freeListener.Password
	em.Role = uint(freeListener.Role)
	em.Nickname = freeListener.Nickname
	em.Firstname = freeListener.Firstname
	em.Lastname = freeListener.Lastname
	em.Middlename = freeListener.Middlename
}

func (ht *FreeListenerHttp) ToCore() *FreeListenerCore {
	return &FreeListenerCore{
		UserCore{
			Email:      ht.Email,
			Password:   ht.Password,
			Role:       Role(ht.Role),
			Nickname:   ht.Nickname,
			Firstname:  ht.Firstname,
			Lastname:   ht.Lastname,
			Middlename: ht.Middlename,
		},
	}
}

func (ht *FreeListenerHttp) FromCore(freeLister *FreeListenerCore) {
	ht.CreatedAt = freeLister.CreatedAt
	ht.Email = freeLister.Email
	ht.Password = freeLister.Password
	ht.Role = uint(freeLister.Role)
	ht.Nickname = freeLister.Nickname
	ht.Firstname = freeLister.Firstname
	ht.Lastname = freeLister.Lastname
	ht.Middlename = freeLister.Middlename
}
