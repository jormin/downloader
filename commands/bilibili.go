package commands

import (
	"fmt"

	"github.com/jormin/download/config"
	"github.com/jormin/download/errors"
	"github.com/jormin/download/internal/bilibili"
	"github.com/urfave/cli/v2"
)

// init
func init() {
	config.RegisterCommand(
		"", &cli.Command{
			Name:      "bilibili",
			Usage:     "download video from bilibili (https://www.bilibili.com/)",
			Action:    BiliBili,
			ArgsUsage: "[bvid: the bvid of video started by 'BV' that can be found from url]",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "d",
					Usage:       "the directory to save video",
					Required:    false,
					DefaultText: "",
				},
			},
			Before: BeforeFunc,
			After:  AfterFunc,
		},
	)
}

// BiliBili download video from bilibili (https://www.bilibili.com/)
func BiliBili(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		return errors.MissingRequiredArgumentErr
	}
	id := ctx.Args().Get(0)
	dir := ""

	if id == "" {
		return errors.FlagContentValidateErr
	}
	flags := ctx.FlagNames()
	for _, v := range flags {
		switch v {
		case "d":
			dir = ctx.String("d")
			if dir == "" {
				return errors.FlagDirValidateErr
			}
		}
	}
	bili := bilibili.NewBiliBili()
	_, _, err := bili.Download(dir, id)
	fmt.Println(err)
	return nil
}
