package middlewares

import (
	"strings"

	"github.com/eron97/bff-golang.git/cmd/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		zap.L().Info("Starting JWT protection", zap.String("ip", ip))
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			zap.L().Warn("Missing or malformed token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or malformed token"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				zap.L().Warn("Unexpected subscription method")
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return &config.PrivateKey.PublicKey, nil

		})

		if err != nil || !token.Valid {
			zap.L().Warn("Invalid token", zap.Error(err))
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		zap.L().Info("Valid JWT Token", zap.String("ip", ip))
		c.Locals("user", token.Claims)
		return c.Next()
	}
}

func JWTClaimsRequired(claimKey string, claimValue string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		zap.L().Info("Verifying JWT Claims", zap.String("ip", ip), zap.String("claimKey", claimKey), zap.String("claimValue", claimValue))
		userClaims := c.Locals("user").(jwt.MapClaims)
		if userClaims[claimKey] != claimValue {
			zap.L().Warn("Permission denied", zap.String("ip", ip), zap.String("claimKey", claimKey), zap.String("claimValue", claimValue))
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Permission denied"})
		}
		zap.L().Info("Valid JWT declarations", zap.String("ip", ip), zap.String("claimKey", claimKey), zap.String("claimValue", claimValue))
		return c.Next()
	}
}
