/*
A common way to use JWTs is as OAuth bearer tokens. In this example, an authorization server
creates a JWT at the request of a client and signs it so that it cannot be altered by any other party.
The client will then send this JWT with its request to a REST API. The REST API will verify that the JWT’s
signature matches its payload and header to determine that the JWT is valid. When the REST API has verified
the JWT, it can use the claims to either grant or deny the client’s request.

In simpler terms, you can think of a JWT bearer token as an identity badge to get into a secured building.
The badge comes with special permissions (the claims); that is, it may grant access to only select areas of
the building. The authorization server in this analogy is the reception desk — or the issuer of the badge.
And to verify that the badge is valid, the company logo is printed on it, similar to the signature of the JWT.
If the badge holder attempts to access a restricted area, the permissions on the badge determine whether or
not they can access the area, similar to the claims in a JWT.
*/

package tokenutil

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	secret = "asdfghjkl"
)

type Claims struct {
	USERID int64 `json:"user_id"`
	jwt.StandardClaims
}

func CreateToken(userID int64) (string, error) {
	// Create the Claims
	claims := Claims{
		USERID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeToken(tokenString string) (int64, error) {
	claims := Claims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("Token is invalid")
	}

	return claims.UserID, nil
}
