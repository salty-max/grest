package middlewares

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func JSONMiddleWare(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json; charset=UTF-8")
	return c.Next()
}

func TimeoutMiddleware(duration time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), duration)
		defer cancel()

		// Store the context object in the request context using SetUserValue
		c.Context().SetUserValue("ctx", ctx)

		// Call the next middleware in the chain
		ch := make(chan error, 1)
		go func() {
			ch <- c.Next()
		}()

		select {
		case err := <-ch:
			if err != nil {
				return err
			}
			return nil
		case <-ctx.Done():
			return fiber.ErrGatewayTimeout
		}
	}
}
