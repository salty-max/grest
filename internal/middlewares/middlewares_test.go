package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TestTimeoutMiddleware(t *testing.T) {
	// Create a new Fiber app
	app := fiber.New()

	// Add the TimeoutMiddleware to the app
	app.Use(TimeoutMiddleware(1 * time.Millisecond))

	// Add a route to the app
	app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(5 * time.Millisecond)
		return c.SendString("Hello, World!")
	})

	// Perform a request to the route and assert the response
	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusGatewayTimeout {
		t.Errorf("Expected status code %d, but got %d", http.StatusGatewayTimeout, resp.StatusCode)
	}
}
