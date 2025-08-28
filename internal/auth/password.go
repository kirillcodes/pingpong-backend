package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b), err
}

func CheckPasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// CompareHashAndPassword - возвращает nil, если совпало!
// GenerateFromPassword - функция, которая генерирует хэш пароль с алгоритма пакета
