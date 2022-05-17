package communityHandler

import (
	"github.com/RegieCanadaRC/community-tracker-api/database"
	"github.com/RegieCanadaRC/community-tracker-api/internal/model"
	"github.com/gofiber/fiber/v2"
)

func GetCommunity(c *fiber.Ctx) error {
	db := database.DB
	var communities []model.Community

	// find all community in the database
	db.Find(&communities)

	// If no community is present return an error
	if len(communities) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No community present", "data": nil})
	}

	// Else return community
	return c.JSON(fiber.Map{"status": "success", "message": "Communities Found", "data": communities})

}

func AddCommunity(c *fiber.Ctx) error {
	db := database.DB
	community := new(model.Community)

	err := c.BodyParser(community)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = db.Create(&community).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create community", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created community", "data": community})
}

func GetCommunityById(c *fiber.Ctx) error {
	db := database.DB
	var community model.Community

	id := c.Params("communityid")

	if result := db.First(&community, "communityid = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Community Found", "data": community})
}
