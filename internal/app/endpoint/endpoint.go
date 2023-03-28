package endpoint

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {

	d := e.s.DaysLeft()

	s := fmt.Sprintf("Days left: %d\n", d)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}