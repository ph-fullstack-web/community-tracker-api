package middleware

import (
	"strings"

	"github.com/VncntDzn/community-tracker-api/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/futurenda/google-auth-id-token-verifier"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": fiber.StatusUnauthorized, "message": "Unathorized Request."})
	}

	authFields := strings.Fields(authHeader)
	if authFields[0] != "Bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": fiber.StatusUnauthorized, "message": "Unathorized Request."})
	}

	token := authFields[1]
	_, tErr := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("JWT_SECRET")), nil
	})

	if tErr != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": fiber.StatusUnauthorized, "message": "Unathorized Request."})
	}

	// validate role on add, edit and delete
	if ctx.Method() != "GET" {
		// decode token
		claimSet, decodeErr := googleAuthIDTokenVerifier.Decode(token)

		if decodeErr != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": fiber.StatusUnauthorized, "message": "Unathorized Request."})
		}

		if claimSet.Iss == "member" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": fiber.StatusUnauthorized, "message": "Unathorized Request."})
		}
	}

	return ctx.Next()
}
