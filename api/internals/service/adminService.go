package service

import(
	"fmt"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
)

func CreateUser(user *models.User) (*models.User, error) {
	fmt.Println("Creating a new user in the User Service Layer")
	fmt.Println("---------------------------------------------")

	if err := models.Validate(user); err != nil {
		return nil, fmt.Errorf("user validation error: %v", err)
	}

	if err := models.SaveUserToDB(user); err != nil {
		return nil, fmt.Errorf("error saving user to database: %v", err)
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error fetching user by email: %v", err)
	}
	return user, nil
}

func GetUserByID(id string) (*models.User, error) {
	user, err := models.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("error fetching user by id: %v", err)
	}
	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	users, err := models.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("error fetching all users: %v", err)
	}
	return users, nil
}

func UpdateUser(user *models.User) (*models.User, error) {
	if err := models.UpdateUser(user); err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}
	return user, nil
}

func DeleteUser(id string) error {
	if err := models.DeleteUser(id); err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}