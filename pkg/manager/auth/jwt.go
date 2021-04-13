package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	cacheconf "github.com/IacopoMelani/Go-Starter-Project/pkg/cache_config"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// MARK: Authenticable, JwtAuthenticableClaims & constructor

// Authenticable - Defines an interface to implements jwt custom claims
type Authenticable interface {
	getJwtIdentifier() string
	getJwtCustomClaims() map[string]interface{}
}

// JwtAuthenticableClaims - Defines a custom claims using Authenticable
type JwtAuthenticableClaims struct {
	jwt.StandardClaims
	CustomClaims map[string]interface{}
}

// NewJwtAuthenticableClaims - Returns an instance of JwtAuthenticableClaims ,exp rapresent the minute before token expires
func NewJwtAuthenticableClaims(a Authenticable, exp int) *JwtAuthenticableClaims {

	customClaims := a.getJwtCustomClaims()

	return &JwtAuthenticableClaims{
		CustomClaims: customClaims,
		StandardClaims: jwt.StandardClaims{
			Subject:   a.getJwtIdentifier(),
			Issuer:    cacheconf.GetDefaultCacheConfig().AppName,
			ExpiresAt: time.Now().UTC().Add(time.Duration(exp) * time.Minute).Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
	}
}

// NewJwtStandardClaims - Returns a new standard jwt claims, exp rapresent the minute before token expires
func NewJwtStandardClaims(a Authenticable, exp int) *jwt.StandardClaims {
	return &jwt.StandardClaims{
		Subject:   a.getJwtIdentifier(),
		Issuer:    cacheconf.GetDefaultCacheConfig().AppName,
		ExpiresAt: time.Now().UTC().Add(time.Duration(exp) * time.Minute).Unix(),
		IssuedAt:  time.Now().UTC().Unix(),
	}
}

// MARK: Unxported funcs

// token - Return s a jwt.Token instance from a Claims interface
func token(claims jwt.Claims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

// verifyTokenString - verifies a token string, if succeded return a *jwt.Token instance otherwise error
func verifyTokenString(tokenString string) (*jwt.Token, error) {
	return verifyTokenStringWithClaims(tokenString, &JwtAuthenticableClaims{})
}

// verifyTokenStringWithClaims - verifies a token string using a custom claims, if succeded return a *jwt.Token instance otherwise error
func verifyTokenStringWithClaims(tokenString string, claims jwt.Claims) (*jwt.Token, error) {

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		conf := cacheconf.GetDefaultCacheConfig()

		return []byte(conf.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// MARK: Exported funcs

// AddTokenToHeader - Appends token string to header as Bearer token
func AddTokenToHeader(c echo.Context, token string) {
	c.Response().Header().Set("Authorization", "Bearer "+token)
}

// CheckToken - Check if the request have a valid jwt token
func CheckToken(request *http.Request) (*jwt.Token, error) {

	token, err := verifyTokenString(ExtractToken(request))
	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(*JwtAuthenticableClaims); !ok {
		return nil, errors.New("token not valid")
	}

	return token, nil
}

// ExtractToken - Extracts token string from an http request
func ExtractToken(request *http.Request) string {
	bearerToken := request.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// GenerateTokens - Generate a tokens pair, "access_token" and "refresh_token"
func GenerateTokens(a Authenticable) (map[string]string, error) {

	tokens := make(map[string]string)

	conf := cacheconf.GetDefaultCacheConfig()

	accessToken := token(NewJwtAuthenticableClaims(a, cacheconf.GetDefaultCacheConfig().JwtTTL))
	at, err := accessToken.SignedString([]byte(conf.JwtSecret))
	if err != nil {
		return nil, err
	}
	tokens["access_token"] = at

	refreshToken := token(NewJwtStandardClaims(a, cacheconf.GetDefaultCacheConfig().JwtRefreshTTL))
	rt, err := refreshToken.SignedString([]byte(conf.JwtSecret))
	if err != nil {
		return nil, err
	}
	tokens["refresh_token"] = rt

	return tokens, nil
}

// RefreshTokens - Refresh a new pairs token from a refresh token string and an Authenticable interface
func RefreshTokens(a Authenticable, refreshToken string) (map[string]string, error)
