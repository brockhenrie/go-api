package middleware

import (
	"api/util"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	//parses jwt token on cookie to check for any auth errors
	cookie := c.Cookies("jwt")

	if _,err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "User not authenticated",
		})
	}
	return c.Next()
}