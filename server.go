package main

import (
	"CypressModule/output"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func main() {
	e := echo.New()

	s := &http.Server{
		Addr:         "localhost:3399",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	e.GET("/", func(c echo.Context) error {


		select {
		default:
			fmt.Print("in default")

		case <-c.Request().Context().Done():
			fmt.Println("in chan")
			return c.Request().Context().Err()
		}

		e := c.JSON(http.StatusOK, output.Result(""))
		if e != nil {
			zap.L().Error(e.Error(), zap.Error(e))
			c.String(http.StatusInternalServerError, e.Error())
			return errors.Wrap(e, "json")
		}
		return c.String(http.StatusOK, "")
	})
	e.Logger.Info(e.StartServer(s))
}
