package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Endpoint struct {
  Url string `json:"url"`
}

func main() {
  e := echo.New()
  e.POST("/v1/open", func(c echo.Context) error {
    endpoint := new(Endpoint)
    if err := c.Bind(endpoint); err != nil {
      return c.String(http.StatusInternalServerError, err.Error())
    }

    // try request to url
    resp, err := http.Get(endpoint.Url)
    if err != nil {
      return c.String(http.StatusInternalServerError, err.Error())
    }

    return c.String(http.StatusOK, resp.Status)
  })
  e.Logger.Fatal(e.Start(":1323"))
}

