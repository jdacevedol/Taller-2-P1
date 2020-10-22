package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/jdacevedol/taller/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
