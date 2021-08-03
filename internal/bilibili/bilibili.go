package bilibili

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/jormin/downloader/internal"
	"github.com/rs/xid"
)

// BiliBili
type BiliBili struct {
	bvInfo   *BvInfo
	savePath string
}

func (b *BiliBili) GetSiteName() string {
	return "bilibili"
}

func (b *BiliBili) GetSiteUrl() string {
	return "https://www.bilibili.com/"
}

func (b *BiliBili) GetTaskID() string {
	return xid.New().String()
}

func (b *BiliBili) GetVideoID() interface{} {
	return b.bvInfo.Bvid
}

func (b *BiliBili) GetVideoTitle() interface{} {
	return b.bvInfo.Title
}

func (b *BiliBili) GetSavePath() string {
	return b.savePath
}

func (b *BiliBili) GetVideoInfo() (*internal.Video, error) {
	if b.bvInfo == nil {
		return nil, errors.New("video info is empty")
	}
	video := internal.Video{
		ID:    b.bvInfo.Bvid,
		Title: b.bvInfo.Title,
	}
	duration := 0
	var pages []internal.Page
	for _, p := range b.bvInfo.Pages {
		duration += p.Duration
		pages = append(
			pages, internal.Page{
				ID:       p.Cid,
				Title:    p.Part,
				Duration: p.Duration,
				Width:    p.Dimension.Width,
				Height:   p.Dimension.Height,
				Rotate:   p.Dimension.Rotate,
			},
		)
	}
	video.Duration = duration
	video.Pages = pages
	return &video, nil
}

func (b *BiliBili) Download(path string, id interface{}) (success int, fail int, err error) {
	bvid := id.(string)
	sdk := NewSDK()
	// 查询page列表
	bvInfo, err := sdk.GetBvInfo(bvid)
	if err != nil {
		return 0, 0, errors.New(fmt.Sprintf("get video info from site bilibili error, please check video id is right"))
	}
	b.bvInfo = bvInfo
	directory := strings.Replace(bvInfo.Title, " ", "_", -1)
	b.savePath = fmt.Sprintf("%s/%s", path, directory)
	_, err = os.Stat(b.savePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return 0, 0, err
		}
		err = os.Mkdir(b.savePath, os.ModePerm)
		if err != nil {
			return 0, 0, err
		}
	}
	var failedPages []FailedPages
	var wg sync.WaitGroup
	var locker sync.Mutex
	// 最多同时处理5个请求，防止被拦截
	dch := make(chan int, 5)
	for _, p := range bvInfo.Pages {
		wg.Add(1)
		go func(p Pages) {
			dch <- 1
			defer func() {
				wg.Done()
				<-dch
			}()
			// 查询C信息
			cInfo, err := sdk.GetCInfo(bvid, p.Cid)
			if err != nil {
				failedPages = append(
					failedPages, FailedPages{
						Pages: p,
						Err:   err,
					},
				)
				locker.Lock()
				fail += 1
				locker.Unlock()
				return
			}
			// 下载视频
			var file, fileName string
			if len(bvInfo.Pages) == 1 {
				fileName = directory
			} else {
				fileName = strings.Replace(p.Part, " ", "_", -1)
				file = fmt.Sprintf("%s/%s.mp4")
			}
			file = fmt.Sprintf("%s/%s.mp4", b.savePath, fileName)
			err = sdk.DownloadVideo(cInfo.Durl[0].URL, file, bvid)
			if err != nil {
				failedPages = append(
					failedPages, FailedPages{
						Pages: p,
						Err:   err,
					},
				)
				locker.Lock()
				fail += 1
				locker.Unlock()
				return
			}
			locker.Lock()
			success += 1
			locker.Unlock()
		}(p)
	}
	wg.Wait()
	if len(failedPages) > 0 {
		b, _ := json.Marshal(failedPages)
		log.Printf("failed pages: %s", b)
	}
	return success, fail, nil
}

// NewBiliBili
func NewBiliBili() *BiliBili {
	return &BiliBili{}
}
