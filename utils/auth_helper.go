package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"online-store/httpserver/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthHelper interface {
	VerifyToken(token string) (bool, interface{}, error)
	GenerateToken(user *models.User) (string, string, error)
	JwtClaimsToUser(jwt.MapClaims) models.User
	ValidateToken() gin.HandlerFunc
}

type authHelper struct {
	JWT_SECRET_KEY string
}

func NewAuthHelper() *authHelper {
	return &authHelper{
		JWT_SECRET_KEY: os.Getenv("JWT_SECRET_KEY"),
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

	claims, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok || !jwtToken.Valid {
		return false, nil, errors.New("invalid token")
	}

	return true, claims, nil
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

func (a *authHelper) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, "Authorization header is required")
			ctx.Abort()
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(a.JWT_SECRET_KEY), nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "Invalid token")
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			_, ok := claims["data"].(map[string]interface{})["id"]
			if !ok {
				ctx.JSON(http.StatusUnauthorized, "Invalid token data")
				ctx.Abort()
				return
			}

			// Use a type assertion that can handle both float64 and string
			// switch id := userID.(type) {
			// case float64:
			// 	ctx.Set("userID", uint(id))
			// case string:
			// 	// Assuming the ID might also be a string in some cases
			// 	ctx.Set("userID", id)
			// default:
			// 	ctx.JSON(http.StatusUnauthorized, "Invalid token data")
			// 	ctx.Abort()
			// 	return
			// }
		} else {
			ctx.JSON(http.StatusUnauthorized, "Invalid token")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
