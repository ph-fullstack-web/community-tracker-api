package login

import (
	"time"
	"log"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	login "github.com/VncntDzn/community-tracker-api/pkg/login/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"github.com/futurenda/google-auth-id-token-verifier"
	"github.com/dgrijalva/jwt-go"
	"github.com/VncntDzn/community-tracker-api/config"
)

func (h handler) GoogleLogin(ctx *fiber.Ctx) error {
	ISSUER := config.GetEnv("ISSUER")
	CLIENT_ID := config.GetEnv("CLIENT_ID")
	JWT_SECRET := config.GetEnv("JWT_SECRET")

	var loginRequest login.GoogleLoginRequest
	ctx.BodyParser(&loginRequest)

	var validate = validator.New()
	validateErr := validate.Struct(loginRequest)

	if validateErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "Please complete inputs."})
	}

	// validate token
	token := loginRequest.Token
	v := googleAuthIDTokenVerifier.Verifier{}
	aud :=  CLIENT_ID
	tokenErr := v.VerifyIDToken(token, []string{
		aud,
	})

	if tokenErr != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": fiber.StatusUnauthorized, "message": "Unathorized Request."})
	}

	// validate decoded token
	claimSet, decodeErr := googleAuthIDTokenVerifier.Decode(token)

	if decodeErr != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": fiber.StatusUnauthorized, "message": "Unathorized Request."})
	}

	if (claimSet.Iss != ISSUER ||
		claimSet.Aud != CLIENT_ID || 
		claimSet.Exp * 1000 < time.Now().UnixMilli()) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": fiber.StatusUnauthorized, "message": "Unathorized Request."})
	}

	// get employee details by email
	csvemail := claimSet.Email

	if csvemail == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "Not Found"})
	}

	// check if employee is admin manager
	employeeRole := "";
	adminManager := &models.AdminManager{}
	if qErr := h.DB.Where(&models.AdminManager{Email: csvemail}).First(&adminManager).Error; qErr != nil {
		employeeRole = "member";
		(*adminManager).AdminName = claimSet.Name
		(*adminManager).Email = csvemail
		log.Println(qErr.Error())
	} else {
		employeeRole = adminManager.RoleType;
	}
	
	// create claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: employeeRole,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	newtoken, newTokenErr := claims.SignedString([]byte(JWT_SECRET))
	if newTokenErr != nil {
		log.Println(newTokenErr.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": "Unable to login"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "access_token": newtoken, "data": adminManager})

}
