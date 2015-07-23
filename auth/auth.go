package auth

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/patelndipen/AP1/settings"
	"golang.org/x/crypto/bcrypt"
)

const (
	JWTLife     = 72
	StoreOffset = 60
)

type Token struct {
	SignedToken string `json:"token"`
}

type AuthContext struct {
	UserID     string
	Exp        time.Time
	TokenStore TokenStoreServices
}

func (ac *AuthContext) Login(enteredPassword, hashedPassword string) *Token {
	if !authenticate(hashedPassword, enteredPassword) {
		return nil // Returns "" if an incorrect password is submitted
	}

	return setTokenClaims(ac.UserID)
}

func (ac *AuthContext) Logout(signedToken string) {

	storeTime := int(ac.Exp.Sub(time.Now()).Seconds() + StoreOffset)

	ac.TokenStore.StoreToken(ac.UserID, signedToken, storeTime)

}

func (ac *AuthContext) RefreshToken() *Token {
	return setTokenClaims(ac.UserID)
}

func authenticate(hashedPassword, enteredPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword)) == nil
}

func setTokenClaims(userID string) *Token {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["exp"] = time.Now().Add(time.Hour * time.Duration(JWTLife))
	token.Claims["sub"] = userID

	signedToken, err := token.SignedString(settings.GetPrivateKey())
	if err != nil {
		log.Fatal(err)
	}
	return &Token{SignedToken: signedToken}
}
