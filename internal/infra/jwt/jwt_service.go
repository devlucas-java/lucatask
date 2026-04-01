package jwt

import (
	"errors"

	"github.com/devlucas-java/lucatask/internal/domain"
	"github.com/go-chi/jwtauth"
)

type JwtService struct {
	auth *jwtauth.JWTAuth
}

func NewJwtService(secret string) *JwtService {
	j := jwtauth.New("HS256", []byte(secret), nil)
	return &JwtService{auth: j}
}

func (j *JwtService) GenerateToken(user *domain.User) (string, error) {
	_, token, err := j.auth.Encode(map[string]interface{}{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"role":    user.Role,
	})

	return token, err
}

func (j *JwtService) Validate(tokenString string) (map[string]interface{}, error) {
	token, err := jwtauth.VerifyToken(j.auth, tokenString)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, errors.New("invalid token")
	}

	claims := token.PrivateClaims()

	return claims, nil
}
