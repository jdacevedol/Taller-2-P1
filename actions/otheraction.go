package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func OtherHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("other.html"))
}