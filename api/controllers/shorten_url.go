package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/asaskevich/govalidator"
	"time"
	"os"
	"github.com/jesseinvent/go-fiber-redis-url-shortener/database"
	"strconv"
	"github.com/jesseinvent/go-fiber-redis-url-shortener/helpers"
	"github.com/jesseinvent/go-fiber-redis-url-shortener/models"

	"github.com/google/uuid"
)

type RequestBody struct {
	URL			string			`json:"url"`
	Alias		string			`json:"alias"`
	Expiry		time.Duration	`json:"expiry"` //hours
}

func ShortenUrl(c *fiber.Ctx) error {

	fmt.Print(c.IP());

	body := new(models.URL);

	err := c.BodyParser(&body);

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid fields supplied or field is missing"});
	}

	// Implement rate limiting
	
	// Get IP key
	value, _ := database.Get(c.IP());

	if value == "" {
		// Set IP in redis with API QUOTA (No of requests) as value
		// 10 requests per 30minutes
		_, _ = database.Set(c.IP(), os.Getenv("API_QUOTA"), 30 * 60 * time.Second);

	} else {
		value, _ := strconv.ParseInt(value, 10, 0);

		if value <= 0 {
			limit, _ := database.TTL(c.IP());
			 
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "rate limit exceeded", "rate_limit": limit / time.Nanosecond / time.Minute });
		}
	}

	// Check if input is an actual url
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":  "invalid URL."});
	}

	// Check for domain error
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Invalid Domain"});
	}

	// Enforce https, SSL
	body.URL = helpers.EnforceHTTPS(body.URL);

	var id string;

	// Generate URL string
	if body.Alias == "" {
		id = uuid.New().String()[:6];
	} else {
		id = body.Alias;
	}

	value, _ = database.Get(id);

	if value != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "URL custom short is already in use."})
	}

	if body.Expiry == 0 {
		body.Expiry = 24;
	}

	// Set URL to DB
	result, err := database.Set(id, body.URL, body.Expiry * 3600 * time.Second);

	if !result && err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "unable to connect to server"});
	}

	remainingLimit, _ := database.DecrementKeyValue(c.IP());

	response := &models.URL {
		URL:			 body.URL,
		Alias: 	 		 "",
		Expiry: 		 body.Expiry,
		XRateRemaining:  remainingLimit,
		XRateLimitReset: 30,
	}

	ttl, _  := database.TTL(c.IP());

	response.XRateLimitReset = ttl / time.Nanosecond / time.Minute;

	response.Alias = os.Getenv("DOMAIN") + "/" + id;

	return c.Status(fiber.StatusOK).JSON(response); 
}