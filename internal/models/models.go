package models

var IDs int

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

var Users = []*User{}

func ListUsers() []*User {
	return Users
}

func GetUser(id int) *User {
	for _, user := range Users {
		if user.ID == id {
			return user
		}
	}
	return nil
}

func UpdateUser(id int, userUpdate User) *User {
	for i, user := range Users {
		if user.ID == id {
			Users[i] = &userUpdate
			return user
		}
	}
	return nil
}

func StoreUser(user User) {
	Users = append(Users, &user)
}

func DeleteUser(id int) *User {
	for _, user := range Users {
		for i, v := range user.Friends {
			if v == id {
				user.Friends = append(user.Friends[:i], (user.Friends)[i+1:]...)
			}
		}
	}
	for i, user := range Users {
		if user.ID == id {
			Users = append(Users[:i], (Users)[i+1:]...)
			return &User{}
		}
	}
	return nil
}
