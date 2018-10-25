package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rodolfo-r/socialnet"
	"golang.org/x/crypto/bcrypt"
)

// UserAuth requires a UserStorage to retrieve stored passwords.
type UserAuth struct {
	store     socialnet.UserStorage
	address   string
	signature string
}

// New returns a UserAuth with a user storage, a web address (jwt issuer) and a jwt signature.
func New(store socialnet.UserStorage, address, signature string) *UserAuth {
	return &UserAuth{store, address, signature}
}

// Claims describe jwt format
type Claims struct {
	Username string `json:"usn"`
	jwt.StandardClaims
}

// CreateToken creates a signed jwt token signed with the userAuth signature.
func (userAuth UserAuth) CreateToken(username string) (string, error) {
	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    userAuth.address,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(userAuth.signature))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ValidateToken validates a jwt token.
func (userAuth UserAuth) ValidateToken(tokenStr string) (username string, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(userAuth.signature), nil
	})

	if err != nil {
		return "", errors.New("failed to parse token: " + err.Error())
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return "", errors.New("invalid auth token")
	}

	if claims.Username == "" {
		return "", errors.New("invalid auth token")
	}

	return claims.Username, nil
}

// ValidateCreds compares the given and stored passwords using bcrypt.
func (userAuth UserAuth) ValidateCreds(username, password string) error {
	usr, err := userAuth.store.Read(username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	return err
}
