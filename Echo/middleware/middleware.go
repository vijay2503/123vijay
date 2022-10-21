package middleware

import (
	jw "echo-task/helper"
	vm "echo-task/model"
	ss "echo-task/studentservice"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func GetAllStudents(c echo.Context) error {
	err := jw.TokenValidation(c)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bearer Token In valid")
	}
	err, s := ss.GetAllStudentDetails()
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request from GetAllStudents ")
	} else {
		json, err := json.Marshal(&s)
		if err != nil {
			return c.String(http.StatusBadRequest, " RECORD NOT FOUND in json")
		} else {
			jsonResponse := fmt.Sprintf("successfully Get Details%s", string(json))
			return c.String(http.StatusOK, jsonResponse)
		}
	}
}

func DeleteStudentDetails(c echo.Context) error {
	roll_no := c.Param("roll_no")
	err, req := ss.DeleteStudentDetails(roll_no)
	if err != nil {
		if req != nil {
			return c.String(http.StatusBadRequest, "sending Bad request  Row Doesn't Exist")
		}
		return c.String(http.StatusBadRequest, "sending Bad request func : DeleteStudentDetails")
	} else {
		return c.String(http.StatusOK, "Deleted sucessfully")
	}
}

func Login(c echo.Context) error {
	var signin vm.Login
	err := c.Bind(&signin)
	if err != nil {
		return echo.ErrUnauthorized
	}
	err, _ = ss.GetStudentDetails(signin.RollNo, signin.Name)
	if err != nil {
		c.String(401, "Registered Student Only Allowed to Login !!!!\n\n")
		return c.JSON(http.StatusOK, echo.Map{
			"Status": "Login Failed ",
		})
	}
	// Set custom claims
	claims := &JwtCustomClaims{
		"vijay",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token":  t,
		"Status": " Successfully login",
	})
}

func SignIN(c echo.Context) error {
	var newStudentManagement vm.StudentDetails
	err := c.Bind(&newStudentManagement)
	if err != nil {
		return err

	}
	err = ss.Registration(newStudentManagement)
	if err != nil {
		return err
	}
	jsonResponse, err := json.Marshal(newStudentManagement)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"Status":       "Successfully Registered",
		"Student Data": string(jsonResponse),
	})
}
