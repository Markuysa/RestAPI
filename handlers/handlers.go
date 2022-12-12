package handlers

import repos "RestAPI/repositories"

func GetAll() []string {
	return repos.GetAllUsers()
}
func GetByID(ID int) string {
	return repos.GetUserByID(ID)
}
