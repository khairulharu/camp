package auth

import (
	"campsite/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type adminAuth struct {
	adminService domain.AdminService
}

func NewAdmin(app *fiber.App, adminService domain.AdminService, authMid fiber.Handler) {
	handler := adminAuth{
		adminService: adminService,
	}

}
