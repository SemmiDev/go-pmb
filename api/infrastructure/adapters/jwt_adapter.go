package adapters

import (
	"fmt"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/environments"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/twinj/uuid"
	"strings"
	"time"
)

type IJwtAdapter interface {
	CreateTokenJWT(userid, email string) (*TokenDetails, error)
	ExtractToken(r *fiber.Ctx) (token string)
	TokenValid(token string) error
	VerifyToken(token string) (*jwt.Token, error)
	VerifyRefresh(token string) (*jwt.Token, error)
	ExtractTokenMetadata(tokenString string) (*AccessDetails, error)
}

type JwtAdapter struct{}

type AccessDetails struct {
	TokenUuid string
	UserId    string
	Email     string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUuid    string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func NewJwtAdapter() IJwtAdapter {
	return &JwtAdapter{}
}

func (t *JwtAdapter) CreateTokenJWT(userid, email string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.TokenUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + userid

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.TokenUuid
	atClaims["user_id"] = userid
	atClaims["email"] = email
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(environments.AccessSecret))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["email"] = email
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(environments.RefreshSecret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (t *JwtAdapter) ExtractTokenMetadata(tokenString string) (*AccessDetails, error) {
	token, err := t.VerifyToken(tokenString)
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId := claims["user_id"].(string)
		email := claims["email"].(string)

		return &AccessDetails{
			TokenUuid: accessUuid,
			UserId:    userId,
			Email:     email,
		}, nil
	}
	return nil, err
}

func (t *JwtAdapter) TokenValid(tokenString string) error {
	token, err := t.VerifyToken(tokenString)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func (t *JwtAdapter) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(environments.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (t *JwtAdapter) VerifyRefresh(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(environments.RefreshSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken get the token from the request body
func (t *JwtAdapter) ExtractToken(r *fiber.Ctx) string {
	bearToken := r.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 && strArr[0] == "Bearer" {
		return strArr[1]
	}
	return ""
}
