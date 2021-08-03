package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jormin/download/errors"
	"github.com/jormin/download/internal"
	"github.com/mitchellh/go-homedir"
	"github.com/rs/xid"
	"github.com/urfave/cli/v2"
)

// unique task
var task internal.Task

// unique video id from video site, not page id
var vid interface{}

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

// Before
func BeforeFunc(ctx *cli.Context) error {
	// check video id
	if ctx.Args().Len() == 0 {
		return errors.MissingRequiredArgumentErr
	}
	vid = ctx.Args().Get(0)
	if vid == "" {
		return errors.ArgumentVdValidateErr
	}
	// check log file
	path := getLogFilePath()
	_, err := os.Stat(path)
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
		Date:         time.Unix(curTime, 0).Format("2006-01-02"),
		StartTime:    curTime,
		StartTimeFmt: time.Unix(curTime, 0).Format("2006-01-02 15:04:05"),
	}
	return nil
}

// After
func AfterFunc(ctx *cli.Context) error {
	// complete task info
	curTime := time.Now().Unix()
	task.EndTime = curTime
	task.EndTimeFmt = time.Unix(curTime, 0).Format("2006-01-02 15:04:05")
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
