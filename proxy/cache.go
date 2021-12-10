package main

type UserCache struct {
	stack    []User
	capacity int
}

func NewUserCache(capacity int) (cache *UserCache, err error) {

	return &UserCache{capacity: capacity}, nil
}

func (cache *UserCache) Get(id int) (user *User, err error) {
	for _, val := range cache.stack {
		if id == val.ID {
			return &val, nil
		}
	}
	return nil, nil
}

func (cache *UserCache) Put(user *User) (err error) {
	if len(cache.stack) >= cache.capacity {
		cache.stack = append(cache.stack[1:], *user)
	} else {
		cache.stack = append(cache.stack, *user)
	}
	return nil
}
