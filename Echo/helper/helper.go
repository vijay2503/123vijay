package helper

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func TokenValidation(c echo.Context) error {
	reqToken := c.Request().Header.Get("Authorization")
	//basicAuth:=c.Request().Header.Get("")
	parseReqToken := strings.SplitAfter(reqToken, "Bearer ")
	if len(parseReqToken) == 1 {
		return errors.New("BEARER TOKEN HEADER IS EMPTY")
	}
	parseStringToken := parseReqToken[1]
	token, err := jwt.Parse(parseStringToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return err
	}
	CheckToken := token.Valid
	if CheckToken != true {
		return err
	}
	return nil
}
