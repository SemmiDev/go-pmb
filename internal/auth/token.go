package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/twinj/uuid"
	"go-clean/internal/app/model"
	"go-clean/internal/config"
	"os"
	"strings"
	"time"
)

type AccessDetails struct {
	AccessUuid string
	UserId     string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

var client = config.NewRedisClient()

func Logout(c *fiber.Ctx) error {
	metadata, err := ExtractTokenMetadata(c)
	if err != nil {
		return c.JSON(model.WebResponse{
			Code:   fiber.StatusUnauthorized,
			Status: c.Status(fiber.StatusUnauthorized).String(),
			Data:   nil,
		})
	}

	delErr := DeleteTokens(c, metadata)
	if delErr != nil {
		return c.JSON(model.WebResponse{
			Code:   fiber.StatusUnauthorized,
			Status: c.Status(fiber.StatusUnauthorized).String(),
			Data:   delErr,
		})
	}

	return c.JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: c.Status(fiber.StatusOK).String(),
		Data:   "successfully logged out",
	})
}

func Refresh(c *fiber.Ctx) error {
	mapToken := map[string]string{}
	if err := c.BodyParser(&mapToken); err != nil {
		return c.JSON(model.WebResponse{
			Code:   fiber.StatusUnprocessableEntity,
			Status: c.Status(fiber.StatusUnprocessableEntity).String(),
			Data:   err.Error(),
		})
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	//if there is an error, the token must have expired
	if err != nil {
		fmt.Println("the error: ", err)
		return c.JSON(model.WebResponse{
			Code:   fiber.StatusUnauthorized,
			Status: c.Status(fiber.StatusUnauthorized).String(),
			Data:   "refresh token expired",
		})
	}

	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return c.JSON(model.WebResponse{
			Code:   fiber.StatusUnauthorized,
			Status: c.Status(fiber.StatusUnauthorized).String(),
			Data:   "error occurred",
		})
	}

	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			return c.JSON(model.WebResponse{
				Code:   fiber.StatusUnprocessableEntity,
				Status: c.Status(fiber.StatusUnprocessableEntity).String(),
				Data:   "error occurred",
			})
		}
		userId := claims["user_id"].(string)

		deleted, delErr := DeleteAuth(c, refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			//Delete the previous Refresh Token
			return c.JSON(model.WebResponse{
				Code:   fiber.StatusUnauthorized,
				Status: c.Status(fiber.StatusUnauthorized).String(),
				Data:   "unauthorized",
			})
		}

		//Create new pairs of refresh and access tokens
		ts, createErr := CreateToken(userId)
		if createErr != nil {
			return c.JSON(model.WebResponse{
				Code:   fiber.StatusUnauthorized,
				Status: c.Status(fiber.StatusUnauthorized).String(),
				Data:   createErr,
			})
		}

		saveErr := CreateAuth(c, userId, ts)
		if saveErr != nil {
			//save the tokens metadata to redis
			return c.JSON(model.WebResponse{
				Code:   fiber.StatusForbidden,
				Status: c.Status(fiber.StatusForbidden).String(),
				Data:   saveErr,
			})
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}

		return c.JSON(model.WebResponse{
			Code:   fiber.StatusCreated,
			Status: c.Status(fiber.StatusCreated).String(),
			Data:   tokens,
		})

	} else {
		return c.JSON(model.WebResponse{
			Code:   fiber.StatusUnauthorized,
			Status: c.Status(fiber.StatusUnauthorized).String(),
			Data:   "refresh expired",
		})
	}
}

func CreateToken(userid string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.AccessUuid + "++" + userid

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func CreateAuth(c *fiber.Ctx, userid string, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := client.Set(c.Context(), td.AccessUuid, userid, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(c.Context(), td.RefreshUuid, userid, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func ExtractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(c *fiber.Ctx) error {
	token, err := VerifyToken(c)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(c *fiber.Ctx) (*AccessDetails, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId := claims["user_id"].(string)
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func FetchAuth(c *fiber.Ctx, authD *AccessDetails) (string, error) {
	userid, err := client.Get(c.Context(), authD.AccessUuid).Result()
	if err != nil {
		return "", err
	}
	if authD.UserId != userid {
		return "", errors.New("unauthorized")
	}
	return userid, nil
}

func DeleteAuth(c *fiber.Ctx, givenUuid string) (int64, error) {
	deleted, err := client.Del(c.Context(), givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func DeleteTokens(c *fiber.Ctx, authD *AccessDetails) error {
	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%s", authD.AccessUuid, authD.UserId)
	//delete access token
	deletedAt, err := client.Del(c.Context(), authD.AccessUuid).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := client.Del(c.Context(), refreshUuid).Result()
	if err != nil {
		return err
	}
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}
