package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dogayaglicioglu/go-oauth2/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GoogleLogin(c *fiber.Ctx) error {

	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")
	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}
func GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")
	googlecon := config.GoogleConfig()

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}

	usereInfo, err := fetchUserInfo(token.AccessToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("User data fetch failed")
	}

	jwtToken, err := createJWTToken(usereInfo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Token Creation Failed")
	}

	return c.JSON(fiber.Map{
		"token": jwtToken,
	})

}

func fetchUserInfo(accessToken string) (map[string]interface{}, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	fmt.Println(userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

var privateKey = []byte("myprivatekey")

func createJWTToken(userInfo map[string]interface{}) (string, error) {
	id, ok := userInfo["id"].(string)
	if !ok {
		return "", fmt.Errorf("user ID is not a string")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		fmt.Println("JWT Signing Error:", err)
		return "", err
	}

	return tokenString, nil
}
