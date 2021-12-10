package main

import "errors"

type User struct {
	ID   int
	Name string
}

type UserDao struct {
	users []User
}

func (dao *UserDao) FindByID(id int) (user *User, err error) {
	for _, val := range dao.users {
		if id == val.ID {
			return &val, nil
		}
	}
	return nil, nil
}

func (dao *UserDao) Create(user *User) (err error) {
	for _, val := range dao.users {
		if user.ID == val.ID {
			return errors.New("id is duplicate")
		}
	}
	dao.users = append(dao.users, *user)
	return nil
}
