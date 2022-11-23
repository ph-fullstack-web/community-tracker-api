package community_members

import (
	"github.com/gofiber/fiber/v2"
	"github.com/VncntDzn/community-tracker-api/pkg/middleware"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	communityMembersRoute := app.Group("/api/community-members", middleware.AuthMiddleware)
	communityMembersRoute.Get("/:communityId", h.GetCommunityMembers)
	communityMembersRoute.Put("/:people_id", h.SoftDeleteMember)
}
