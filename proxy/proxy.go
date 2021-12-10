package main

import "fmt"

type UserProxy struct {
	UserDao   *UserDao
	UserCache *UserCache
}

func (proxy *UserProxy) FindUserByID(id int) (user *User, err error) {
	user, err = proxy.UserCache.Get(id)
	if err != nil {
		return nil, err
	}
	if user != nil {
		fmt.Println("find from cache")
		return
	}
	user, err = proxy.UserDao.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user != nil {
		err = proxy.UserCache.Put(user)
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("find from dao")
	return
}
