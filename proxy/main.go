package main

import "fmt"

func main() {
	userDao := &UserDao{}
	err := userDao.Create(&User{
		ID:   1,
		Name: "ann",
	})
	if err != nil {
		panic(err)
	}
	userCache, err := NewUserCache(10)
	if err != nil {
		panic(err)
	}
	userProxy := UserProxy{userDao, userCache}
	user, err := userProxy.FindUserByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	user, err = userProxy.FindUserByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
