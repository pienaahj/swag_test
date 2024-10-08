package model

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{
		ID:       1,
		UserName: "pienaahj",
		FullName: "Hendrik Pienaar",
		Email:    "pienaahj@example.com",
		IsActive: true,
	},
	{
		ID:       2,
		UserName: "pienaahj2",
		FullName: "Hendrik Pienaar2",
		Email:    "pienaahj2@example.com",
		IsActive: true,
	},
}

var countries = []Country{
	{
		ID:   1,
		Name: "South Africa",
	},
	{
		ID:   2,
		Name: "United States",
	},
	{
		ID:   3,
		Name: "Nigeria",
	},
	{
		ID:   4,
		Name: "Ghana",
	},
}

func ListUsers() []User {
	return users
}

func GetUser(username string) (*User, error) {
	for _, user := range users {
		if user.UserName == username {
			return &user, nil
		}
	}
	return &User{}, nil
}

func CreateUser(user User) error {
	users = append(users, user)
	return nil
}

func ListCountries() []Country {
	return countries
}
