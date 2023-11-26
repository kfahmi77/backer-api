package auth

import "github.com/golang-jwt/jwt"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SecretKey = []byte("BWABACKER_s3cret")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
