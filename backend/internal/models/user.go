package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

// Create User
func (u *User) Create() (err error) {
	// Preprocess the User by Hashing the Password
	u.Password, err = hashPassword(u.Password)
	if err != nil {
		return err
	}

	// Sanitise username from html tags and spaces to prevent attacks
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	// Create the user record
	err = DB.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(uid uint) (u User, err error) {
	if err = DB.First(&u, uid).Error; err != nil {
		return u, errors.New("user not found")
	}

	// Sanitise the user struct before sending it back to the client
	u.PrepareGive()
	return u, nil
}

// PrepareGive removes the password from the user struct before sending it to
// back to the client
func (u *User) PrepareGive() {
	u.Password = ""
}

func CheckUserExists(username string) (bool, error) {
	var count int
	err := DB.Model(User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// LoginCheck checks if the username and password are correct then returns
// a jwt token which is used for all authenticated requests.
func LoginCheck(username string, password string) (token string, err error) {
	u := User{}

	// Fetch the user from the database by username
	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return "", err
	}

	// Verify the password
	err = verifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	// Generate jwt token
	token, err = GenerateToken(u)
	if err != nil {
		return "", err
	}

	return token, nil

}

// HashPassword generates a hashed password from a plaintext password using
// bcrypt, this allows us to store the hashed password in the database.
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword compares the password with the hashed password
func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
