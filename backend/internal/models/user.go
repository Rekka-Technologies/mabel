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

// VerifyPassword compares the password with the hashed password
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
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
	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	// Generate jwt token
	token, err = GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}

// SaveUser saves a new users state to the database
func (u *User) SaveUser() (*User, error) {
	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {

	// Generate Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	// Sanitise username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	return nil
}
