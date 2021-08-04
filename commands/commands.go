package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jormin/downloader/errors"
	"github.com/jormin/downloader/helper"
	"github.com/jormin/downloader/internal"
	"github.com/mitchellh/go-homedir"
	"github.com/rs/xid"
	"github.com/urfave/cli/v2"
)

// unique task
var task internal.Task

// unique video id from video site, not page id
var vid interface{}

// the directory that videos saved in, default is command run directory
var dir string

// global error
var err error

// video info
var video *internal.Video

// video saved path
var savePath string

// success and fail num of task videos
var success, fail int

// Get log file path
func getLogFilePath() string {
	home, _ := homedir.Dir()
	return fmt.Sprintf("%s/downloader.log", home)
}

// BeforeFunc do something before execute command
func BeforeFunc(ctx *cli.Context) error {
	// check video id
	if ctx.Args().Len() == 0 {
		return errors.ErrMissingRequiredArgument
	}
	vid = ctx.Args().Get(0)
	if vid == "" {
		return errors.ErrArgumentVdValidate
	}
	// check save dir
	dir, _ = os.Getwd()
	flags := ctx.FlagNames()
	for _, v := range flags {
		switch v {
		case "d":
			dir = ctx.String("d")
			if dir == "" {
				return errors.ErrFlagDirValidate
			}
		}
	}
	// check log file
	path := getLogFilePath()
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		_, err = os.Create(path)
		if err != nil {
			return err
		}
	}
	// create unique task
	curTime := time.Now().Unix()
	task = internal.Task{
		ID:           xid.New().String(),
		Date:         helper.FormatDate(curTime),
		StartTime:    curTime,
		StartTimeFmt: helper.FormatTime(curTime),
	}
	return nil
}

// AfterFunc  do something before after command
func AfterFunc(ctx *cli.Context) error {
	// complete task info
	curTime := time.Now().Unix()
	task.EndTime = curTime
	task.EndTimeFmt = helper.FormatTime(curTime)
	task.Video = video
	if err == nil {
		task.Status = internal.TaskStatusSuccess
		fmt.Printf(
			"task %s of video %v download success %d and fail %d, video saved in %s.\n", task.ID, vid, success, fail,
			savePath,
		)
	} else {
		task.Error = err.Error()
		task.Status = internal.TaskStatusFail
		fmt.Printf("task %s of video %v download err: %v.\n", task.ID, vid, err)
	}
	// record task log
	path := getLogFilePath()
	b, _ := json.Marshal(task)
	log := fmt.Sprintf("%s\n", b)
	_ = os.WriteFile(path, []byte(log), os.ModeAppend)
	return nil
}
