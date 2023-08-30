package auth

import (
	"errors"
	"time"

	"github.com/shobky/giggs/api/schema"
	"github.com/shobky/giggs/config"

	"github.com/golang-jwt/jwt"

	"context"

	"google.golang.org/api/idtoken"
)

var TOKENEXP = "100h"

// Generate generates the jwt token based on payload
func GenerateToken(payload *schema.TokenPayload) (string, error) {
	v, err := time.ParseDuration(TOKENEXP)

	if err != nil {
		return "", err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(v).Unix(),
		"ID":    payload.ID,
		"email": payload.Email,
	})

	token, err := t.SignedString([]byte(config.Config.JWTTokenSecret))

	if err != nil {
		return "", err
	}

	return token, nil
}

func parseToken(token string) (*jwt.Token, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Can't parse token,: unexpected signing method")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Config.JWTTokenSecret), nil
	})
}

// Verify verifies the jwt token against the secret
func Verify(token string) (*schema.TokenPayload, error) {
	parsed, err := parseToken(token)

	if err != nil {
		return nil, err
	}

	// Parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	// Getting ID, it's an interface{} so I need to cast it to uint
	id, ok := claims["ID"].(float64)
	email, ok := claims["email"].(string)

	if !ok {
		return nil, errors.New("something went wrong")
	}

	return &schema.TokenPayload{
		ID:    uint(id),
		Email: email,
	}, nil
}

func ParseIdtoken(tokenString string) (*idtoken.Payload, error) {

	tokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		return nil, err
	}

	payload, err := tokenValidator.Validate(context.Background(), tokenString, config.Config.GoogleClientID)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
