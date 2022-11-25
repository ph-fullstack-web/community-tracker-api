package login

import (
	"log"
	"time"

	"github.com/VncntDzn/community-tracker-api/config"
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	login "github.com/VncntDzn/community-tracker-api/pkg/login/requests"
	"github.com/VncntDzn/community-tracker-api/pkg/utils/hash"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

func (h handler) Login(ctx *fiber.Ctx) error {
	var loginRequest login.LoginRequest
	ctx.BodyParser(&loginRequest)

	var validate = validator.New()
	validateErr := validate.Struct(loginRequest)
	if validateErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "Please complete inputs."})
	}

	adminManager := &models.AdminManager{}

	if qErr := h.DB.Where(&models.AdminManager{CognizantID: loginRequest.CognizantId}).First(&adminManager).Error; qErr != nil {
		log.Println(qErr.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": "Unable to login"})
	}
	// check if record found
	if adminManager.ID == 0 {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": fiber.StatusForbidden, "message": "Invalid credentials."})
	}

	//validate password
	loginSuccess := hash.Check(loginRequest.Password, adminManager.Password)
	if !loginSuccess {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": fiber.StatusForbidden, "message": "Invalid credentials."})
	}

	// create claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id: adminManager.CognizantID,
		Issuer: adminManager.RoleType,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, tokenErr := claims.SignedString([]byte(config.GetEnv("JWT_SECRET"))) // secret key stored in .env
	if tokenErr != nil {
		log.Println(tokenErr.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": "Unable to login"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "access_token": token, "data": adminManager})
}
