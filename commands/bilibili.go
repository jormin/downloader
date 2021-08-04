package commands

import (
	"github.com/jormin/downloader/config"
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
			Before:    BeforeFunc,
			After:     AfterFunc,
		},
	)
}

// BiliBili download video from bilibili (https://www.bilibili.com/)
func BiliBili(ctx *cli.Context) error {
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
