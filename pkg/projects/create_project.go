package projects

import (
	"errors"
	"strings"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	requests "github.com/VncntDzn/community-tracker-api/pkg/projects/requests"
	projectValidation "github.com/VncntDzn/community-tracker-api/pkg/validations/projects"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h handler) CreateProject(c *fiber.Ctx) error {
	projectRequest := new(requests.CreateProjectRequest)

	c.BodyParser(projectRequest)

	validateErr := projectValidation.ValidateCreateCommunity(*projectRequest)
	if validateErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "error": validateErr})
	}

	//@Todo add validation for project name

	projectData := models.Project{
		ProjectName: projectRequest.ProjectName,
		ProjectCode: projectRequest.ProjectCode,
		IsActive:    projectRequest.IsActive,
	}

	project := &models.Project{}
	uppercaseProjectName := strings.ToUpper(projectRequest.ProjectName)
	projectSearchErr := h.DB.Where("UPPER(projectdesc) = ?", uppercaseProjectName).First(&project).Error
	if projectSearchErr != nil && !errors.Is(projectSearchErr, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "error": projectSearchErr.Error()})
	}

	if project.ProjectId != 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "error": "Project already exists."})
	}

	projectResult := h.DB.Create(&projectData)
	if projectResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "error": projectResult.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success!", "data": &projectRequest})
}
