package utils

import (
	"encoding/json"
	"errors"
	"time"

	"online-store/httpserver/models"

	"github.com/golang-jwt/jwt"
)

type AuthHelper interface {
	VerifyToken(token string) (bool, interface{}, error)
	GenerateToken(user *models.User) (string, string, error)
	JwtClaimsToUser(jwt.MapClaims) models.User
}

type authHelper struct {
	JWT_SECRET_KEY string
}

func NewAuthHelper(jwtSecret string) *authHelper {
	return &authHelper{
		JWT_SECRET_KEY: jwtSecret,
	}
}

func (s *authHelper) VerifyToken(accessToken string) (bool, interface{}, error) {
	jwtToken, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		method, isRsa := t.Method.(*jwt.SigningMethodHMAC)
		if !isRsa {
			return nil, errors.New("invalid algorithm")
		}
		if method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid algorithm")
		}

		return []byte(s.JWT_SECRET_KEY), nil
	})

	if jwtToken == nil {
		return false, nil, errors.New("invalid token")
	}

	if err != nil {
		validation, _ := err.(*jwt.ValidationError)
		if validation.Errors == jwt.ValidationErrorExpired {
			return false, nil, errors.New("token expired")
		}
	}

	_, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok || !jwtToken.Valid {
		return false, nil, errors.New("invalid token")
	}

	return true, jwtToken.Claims.(jwt.MapClaims), nil

}

func (s *authHelper) GenerateToken(user *models.User) (string, string, error) {
	const ttlAccessToken = 24 * time.Hour
	const ttlRefreshToken = (24 * 7) * time.Hour

	var userMap map[string]interface{}
	data, err := json.Marshal(user)
	if err != nil {
		return "", "", err
	}

	json.Unmarshal(data, &userMap)

	accessClaims, refreshClaims := jwt.MapClaims{
		"data": userMap,
		"exp":  time.Now().UTC().Add(ttlAccessToken).Unix(),
	}, jwt.MapClaims{
		"data": userMap,
		"exp":  time.Now().UTC().Add(ttlRefreshToken).Unix(),
	}
	var secretKeyByte = []byte(s.JWT_SECRET_KEY)

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secretKeyByte)

	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretKeyByte)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *authHelper) JwtClaimsToUser(claims jwt.MapClaims) models.User {
	data := claims["data"].(map[string]interface{})
	user := models.User{
		Name:     data["name"].(string),
		Password: data["password"].(string),
		Email:    data["email"].(string),
	}
	return user
}
