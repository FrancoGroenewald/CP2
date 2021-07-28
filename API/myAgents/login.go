package myAgents

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/FrancoGroenewald/CP2/API/appConfig"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Claims struct {
	jwt.StandardClaims
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	result := make(map[string]interface{})
	var claimsReturn interface{}
	if err := appConfig.DB.Raw("exec [PractiseDatabase].[dbo].[spLogin] ?, ?", data["user"], data["password"]).Scan(&result); err != nil {
		log.Flags()
	}

	if userID, err := result["useID"].(int64); !err {
		fmt.Println(err)
	} else {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(userID)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		})

		token, err := claims.SignedString([]byte("secret"))

		if err != nil {
			log.Flags()
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		}

		c.Cookie(&cookie)

		cookieReturn := c.Cookies("jwt")

		tokenReturn, err := jwt.ParseWithClaims(cookieReturn, &Claims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil || !tokenReturn.Valid {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": "unauthenticated",
			})
		}
		claimsReturn = tokenReturn.Claims
	}

	return c.JSON(fiber.Map{
		"response": result,
		"claims":   claimsReturn,
		"message":  "success",
	})

}
