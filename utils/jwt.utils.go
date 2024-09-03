package utils

import (
	"os"
	customErrors "phoenixia/errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CreateJWTPasswordReset(username string) (token string, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		err = &customErrors.ServerError
		return
	}
	key := []byte(os.Getenv("JWT_SECRET"))
	jwToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour).Unix(),
		})
	token, err = jwToken.SignedString(key)
	if err != nil {
		token = ""
		err = &customErrors.ServerError
		return
	}
	return
}

func ValidateJWTPasswordReset(tokenString string) (username string, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		err = &customErrors.ServerError
		return
	}
	key := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		err = &customErrors.InvalidToken
		return
	}
	if !token.Valid {
		err = &customErrors.InvalidToken
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				err = &customErrors.InvalidToken
				return
			}
		} else {
			err = &customErrors.InvalidToken
			return
		}
		if user, ok := claims["username"].(string); ok {
			username = user
		} else {
			err = &customErrors.InvalidToken
			return
		}
	} else {
		err = &customErrors.InvalidToken
		return
	}
	return
}

func CreateJWTSession(id uint, username, role string, documents []string) (token string, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		err = &customErrors.ServerError
		return
	}
	key := []byte(os.Getenv("JWT_SECRET"))
	jwToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":        id,
			"username":  username,
			"role":      role,
			"documents": documents,
			"exp":       time.Now().Add(time.Hour).Unix(),
		})
	token, err = jwToken.SignedString(key)
	if err != nil {
		token = ""
		err = &customErrors.ServerError
		return
	}
	return
}

func ValidateJWTSession(tokenString string) (id, username, role string, documents []string, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		err = &customErrors.ServerError
		return
	}
	key := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		err = &customErrors.InvalidToken
		return
	}
	if !token.Valid {
		err = &customErrors.InvalidToken
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				err = &customErrors.InvalidToken
				return
			}
		} else {
			err = &customErrors.InvalidToken
			return
		}
		if idJWT, ok := claims["id"].(string); ok {
			id = idJWT
		} else {
			err = &customErrors.InvalidToken
			return
		}
		if usernameJWT, ok := claims["username"].(string); ok {
			username = usernameJWT
		} else {
			err = &customErrors.InvalidToken
			return
		}
		if roleJWT, ok := claims["role"].(string); ok {
			role = roleJWT
		} else {
			err = &customErrors.InvalidToken
			return
		}
		if documentsJWT, ok := claims["documents"].(string); ok {
			array := strings.Split(documentsJWT, ",")
			documents = array
		} else {
			err = &customErrors.InvalidToken
			return
		}
	} else {
		err = &customErrors.InvalidToken
		return
	}
	return
}
