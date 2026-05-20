package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 1. MEMBUAT OBJEK PALSU (MOCK)
type MockUserRepository struct {
	mock.Mock
}

// Memenuhi kontrak UserRepository
func (m *MockUserRepository) CariUserBerdasarkanEmail(email string) (User, error) {
	args := m.Called(email)
	return args.Get(0).(User), args.Error(1)
}

// 2. EKSEKUSI PENGUJIAN
func TestAuthService_Login(t *testing.T) {
	t.Parallel()

	skenarioTes := []struct {
		namaKasus     string
		inputEmail    string
		inputPassword string
		setupMock     func(mockRepo *MockUserRepository)
		expectedToken string
		expectedError error
	}{
		{
			namaKasus:     "Sukses Login",
			inputEmail:    "admin@test.com",
			inputPassword: "password123",
			setupMock: func(mockRepo *MockUserRepository) {
				mockUser := User{ID: 1, Email: "admin@test.com", Password: "password123"}
				// DIKTE: "Jika ditanya email ini, kembalikan user & error nil"
				mockRepo.On("CariUserBerdasarkanEmail", "admin@test.com").Return(mockUser, nil).Once()
			},
			expectedToken: "token_rahasia_jwt",
			expectedError: nil,
		},
		{
			namaKasus:     "Gagal - User Tidak Ditemukan",
			inputEmail:    "hantu@test.com",
			inputPassword: "password123",
			setupMock: func(mockRepo *MockUserRepository) {
				mockRepo.On("CariUserBerdasarkanEmail", "hantu@test.com").Return(User{}, ErrUserNotFound).Once()
			},
			expectedToken: "",
			expectedError: ErrUserNotFound,
		},
		{
			namaKasus:     "Gagal - Database Timeout (Simulasi Serangan/Down)",
			inputEmail:    "admin@test.com",
			inputPassword: "password123",
			setupMock: func(mockRepo *MockUserRepository) {
				mockRepo.On("CariUserBerdasarkanEmail", "admin@test.com").Return(User{}, ErrDatabaseTimeout).Once()
			},
			expectedToken: "",
			expectedError: ErrDatabaseTimeout,
		},
	}

	for _, skenario := range skenarioTes {
		t.Run(skenario.namaKasus, func(t *testing.T) {
			t.Parallel()

			// a. Siapkan Mock Object (Database Palsu)
			mockRepo := new(MockUserRepository)
			skenario.setupMock(mockRepo)

			// b. DEPENDENCY INJECTION: Menyuntikkan database palsu ke Service
			authService := NewAuthService(mockRepo)

			// c. Panggil fungsinya
			token, err := authService.Login(skenario.inputEmail, skenario.inputPassword)

			// d. ASSERTION: Validasi dengan Testify
			if skenario.expectedError != nil {
				assert.ErrorIs(t, err, skenario.expectedError, "Error harus sesuai skenario")
				assert.Empty(t, token, "Token harus kosong jika error")
			} else {
				assert.NoError(t, err, "Seharusnya tidak ada error")
				assert.Equal(t, skenario.expectedToken, token, "Token tidak cocok")
			}

			// e. AUDIT: Memastikan fungsi DB palsu benar-benar dieksekusi
			mockRepo.AssertExpectations(t)
		})
	}
}
