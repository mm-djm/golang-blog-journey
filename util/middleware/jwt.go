package middleware

import (
	"golang-blog-journey/util/log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	jwtKey = "abcdefg"
)

var JWTMap = make(map[string]int)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(jwtKey),
	}
}

type MyClaims struct {
	Email string `json:"user"`
	jwt.StandardClaims
}

func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

func (j *JWT) ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, nil
	}

	return nil, nil
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader, err := c.Cookie("user")
		if err != nil {
			log.Errorw("get cookie", "err", err)
			c.String(http.StatusBadGateway, "Cookie error")
			c.Abort()
			return
		}
		if tokenHeader == "" {
			c.String(http.StatusBadGateway, "Cookie not set,pls login first.")
			c.Abort()
			return
		}

		if len(tokenHeader) == 0 {
			c.String(http.StatusBadGateway, "Cookie not set")
			c.Abort()
			return
		}

		j := NewJWT()
		claims, err := j.ParseToken(tokenHeader)
		if err != nil {
			log.Errorw("ParserToken", "err", err)
			c.String(http.StatusBadGateway, "Cookie parse error")
			c.Abort()
			return
		}

		if role, ok := JWTMap[claims.Email]; !ok {
			c.String(http.StatusBadGateway, "pls login first")
			c.Abort()
			return
		} else if role == 1 {
			c.String(http.StatusBadGateway, "guest cannot access")
			c.Abort()
			return
		}

		c.Next()
	}
}

func JwtAPIToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader, err := c.Cookie("user")
		if err != nil {
			log.Errorw("get cookie", "err", err)
			c.String(http.StatusBadGateway, "Cookie error")
			c.Abort()
			return
		}
		if tokenHeader == "" {
			c.String(http.StatusBadGateway, "Cookie not set,pls login first.")
			c.Abort()
			return
		}

		if len(tokenHeader) == 0 {
			c.String(http.StatusBadGateway, "Cookie not set,pls login first.")
			c.Abort()
			return
		}

		j := NewJWT()
		claims, err := j.ParseToken(tokenHeader)
		if err != nil {
			log.Errorw("ParserToken", "err", err)
			c.String(http.StatusBadGateway, "Cookie parse error")
			c.Abort()
			return
		}

		if _, ok := JWTMap[claims.Email]; !ok {
			c.String(http.StatusBadGateway, "pls login first")
			c.Abort()
			return
		}

		c.Next()
	}
}

//Guest=1
//Host=0
func RegisterMap(email, role string) {
	if role == "host" {
		JWTMap[email] = 0
	} else if role == "guest" {
		JWTMap[email] = 1
	}
}
