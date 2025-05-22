package models

import (
	"time"
	"fmt"

	"github.com/google/uuid"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/config"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
)

type User struct {
	ID         uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	FirstName  string      `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName   string      `gorm:"type:varchar(100);not null" json:"last_name"`
	Avatar     *string     `gorm:"type:varchar(255)" json:"avatar"`
	Email      *string     `gorm:"type:varchar(100);unique" json:"email"`
	Role     string      `gorm:"type:varchar(50);default:'user'" json:"role"`
	Password   string      `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt  time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  *time.Time  `gorm:"index" json:"deleted_at"`
	Portfolios []Portfolio `gorm:"foreignKey:UserID;references:ID" json:"portfolios,omitempty"`
}

func AutoMigrate() error {
    db := config.LoadDB()
    
    if err := db.AutoMigrate(&User{}); err != nil {
        return fmt.Errorf("failed to migrate User: %v", err)
    }
    
    if err := db.AutoMigrate(&Portfolio{}); err != nil {
        return fmt.Errorf("failed to migrate Portfolio: %v", err)
    }

	if err := db.AutoMigrate(&Asset{}); err != nil {
		return fmt.Errorf("failed to migrate assets : %v", err)
	}
    
    return nil
}

func Validate (u *User) error {
	if u.FirstName == "" {
		return &errors.ValidationError{Message: "First name is required"}
	}
	if u.LastName == "" {
		return &errors.ValidationError{Message: "Last name is required"}
	}
	if u.Email == nil || *u.Email == "" {
		return &errors.ValidationError{Message: "Email is required"}
	}
	if u.Role == "" {
		return &errors.ValidationError{Message: "Role is required"}
	}
	if u.Password == "" {
		return &errors.ValidationError{Message: "Password is required"}
	}
	return nil
}

func SaveUserToDB (u *User) error {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	db:= config.LoadDB()	
	go AutoMigrate()
	
	var existingUser User
	if err := db.Where("email = ?", u.Email).First(&existingUser).Error; err == nil {
		return &errors.DatabaseError{Message: fmt.Sprintf("User with email %s already exists", *u.Email)}
	}
	if err := db.Create(u).Error; err != nil {
		fmt.Printf("Error occurred: %v", err);
		return &errors.DatabaseError{Message: "Error saving user to database", Err: err}
	}

	return nil
}

func GetUserByEmail(email string) (*User, error) {
	db := config.LoadDB()
	var user User
	if email == "" {
		return nil, fmt.Errorf("Email is required");
	}
	if err := db.First(&user, "email = ?", email).Error; err != nil {
		fmt.Printf("Error Occurred Getting Email: %v", err)
		return nil, fmt.Errorf("Error Occurred Getting User Email")
	}
	return &user, nil
}

func GetUserByID(userID string) (*User, error) {
	db := config.LoadDB()
	var user User

	if userID == "" {
		return nil, fmt.Errorf("ID is required")
	}

	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, fmt.Errorf("User with ID %s does not exist", userID)
	}

	return &user, nil
}

func GetAllUsers() ([]User, error) {
	db := config.LoadDB()
	var users []User
	if err := db.Where("role = ?", "user").Find(&users).Error; err != nil {
		return nil, fmt.Errorf("Error getting users with role 'user': %v", err)
	}
	return users, nil
}

func UpdateUser(user *User) error {
	db := config.LoadDB()
	user.UpdatedAt = time.Now()
	if user.ID == uuid.Nil {
		return fmt.Errorf("User ID is required for update")
	}

	var existingUser User
	if err := db.First(&existingUser, "id = ?", user.ID).Error; err != nil {
		return fmt.Errorf("User with ID %s does not exist", user.ID)
	}

	if err := db.Save(user).Error; err != nil {
		return fmt.Errorf("Error updating user: %v", err)
	}
	return nil
}

func DeleteUser(userID string) error {
	db := config.LoadDB()
	if userID == "" {
		return fmt.Errorf("User ID is required for deletion")
	}

	var user User
	if err := db.First(&user, "id = ?", userID).Error; err != nil {
		return fmt.Errorf("User with ID %s does not exist", userID)
	}

	if err := db.Delete(&user).Error; err != nil {
		fmt.Printf("Error occurred while trying to delete User: %v", err)
		return fmt.Errorf("Error deleting user")
	}
	return nil
}

func GetUserByRole(role string) ([]User, error) {
	db := config.LoadDB()
	var users []User
	if role == "" {
		return nil, fmt.Errorf("Role is required")
	}
	if err := db.Where("role = ?", role).Find(&users).Error; err != nil {
		fmt.Printf("Error getting users by role: %v", err)
		return nil, fmt.Errorf("Error getting users by role")
	}
	return users, nil
}

func GetPortfolioForUser(userID uuid.UUID) (*User, error) {
	db := config.LoadDB()
	go AutoMigrate()

	if userID == uuid.Nil {
		return nil, fmt.Errorf("User ID is required for showing Portfolio")
	}

	var user User
	if err := db.Preload("Portfolios").Preload("Portfolios.Assets").First(&user, "id = ?", userID).Error; err != nil {
		fmt.Printf("Error getting user with ID %s: %v", userID, err)
		return nil, fmt.Errorf("Error getting user with ID %s", userID)
	}

	for i := range user.Portfolios {
		user.Portfolios[i].UserID = user.ID
	}

	return &user, nil
}