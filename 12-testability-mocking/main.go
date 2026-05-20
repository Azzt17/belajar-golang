package main

import (
	"errors"
)

var (
	ErrUserNotFound    = errors.New("user tidak ditemukan")
	ErrWrongPassword   = errors.New("password salah")
	ErrDatabaseTimeout = errors.New("database timeout")
)

type User struct {
	ID       int
	Email    string
	Password string
}

// Interface - Kunci dari Testability
type UserRepository interface {
	CariUserBerdasarkanEmail(email string) (User, error)
}

// SERVICE
type AuthService struct {
	repo UserRepository // Ketergantungan disembunyikan di sini
}

// Constructor untuk inject dependensi dari luar
func NewAuthService(r UserRepository) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) Login(email, password string) (string, error) {
	// Memanggil fungsi dari interface, bukan dari database langsung
	user, err := s.repo.CariUserBerdasarkanEmail(email)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", ErrWrongPassword
	}

	return "token_rahasia_jwt", nil
}
