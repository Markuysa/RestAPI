package repositories

import "fmt"

var users = []string{
	"Masha", "Sasha", "Sasha", "Sasha", "Sasha", "Sasha", "Sasha", "Sasha",
}

func GetAllUsers() []string {
	return users
}

func GetUserByID(ID int) string {
	fmt.Println(ID)
	if ID > len(users) {
		return "Not found"
	}
	return users[ID]
}
