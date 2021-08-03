package commands

import (
	"github.com/jormin/downloader/config"
	"github.com/jormin/downloader/errors"
	"github.com/jormin/downloader/internal/bilibili"
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
	dir := ""
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
	success, fail, err = bili.Download(dir, vid)
	// do not return err here, unified external output is in function `AfterFunc`
	if err != nil {
		return nil
	}
	video, err = bili.GetVideoInfo()
	savePath = bili.GetSavePath()
	return nil
}
