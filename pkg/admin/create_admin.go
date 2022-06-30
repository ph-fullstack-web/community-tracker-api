package admin

import (
	admin "github.com/VncntDzn/community-tracker-api/pkg/admin/requests"
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/VncntDzn/community-tracker-api/pkg/utils/hash"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (h handler) CreateAdmin(ctx *fiber.Ctx) error {
	var request admin.CreateAdminRequest
	ctx.BodyParser(&request)

	// validation
	var validate = validator.New()
	validateErr := validate.Struct(request)
	if validateErr != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "message": validateErr})
	}

	hashedPassword := hash.Make(request.CognizantId)
	adminData := &models.AdminManager{
		AdminName:   request.AdminName,
		CognizantID: request.CognizantId,
		Email:       request.Email,
		Password:    hashedPassword,
		RoleType:    "admin",
		IsActive:    true,
	}

	if insertErr := h.DB.Create(&adminData).Error; insertErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": insertErr.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"success": fiber.StatusCreated, "data": &adminData})
}