package model

import (
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint      `gorm:"primaryKey;" json:"id"`
	Name         string    `json:"name"`
	CandidateID  string    `json:"candidate_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	ModifiedAt   time.Time `json:"modified_at"`
	ModifiedBy   string    `json:"modified_by"`
}

func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	u.PasswordHash = string(bytes)
}

func (u *User) GenerateToken() (string, error) {
	var (
		jwtKey = os.Getenv("JWT_KEY")
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      u.Email,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"phone":      u.Phone,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}
