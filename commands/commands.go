package commands

import (
	"github.com/urfave/cli/v2"
)

// Before
func BeforeFunc(ctx *cli.Context) error {
	return nil
}

// After
func AfterFunc(ctx *cli.Context) error {
	return nil
}
