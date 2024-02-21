package main

import (
	"log"
	"myfiber/routes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFiberRoute(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        "/TEST",
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 200",
			route:        "/TEST1",
			expectedCode: 200,
		},
		// Second test case
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/undef_url",
			expectedCode: 404,
		},
	}

	// Define Fiber app.
	App := Setup("./views")
	App.Get("/TEST1", RouteTest1)

	// App.Get("/TEST", RouteJson1) -  !!! Падает с ошибкой !!!

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := App.Test(req, 1)
		// if resp != nil && err != nil {
		log.Println("response code:", resp.StatusCode)
		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		// }
	}
}

func RouteTest1(c *fiber.Ctx) error {

	jj := routes.MyJson{
		Par1: "Param-11",
		Par2: "P2",
		Par3: "P3",
	}
	return c.JSON(jj)
}
