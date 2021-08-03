package bilibili

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/rs/xid"
)

// BiliBili
type BiliBili struct {
	BvInfo   *BvInfo `json:"bvinfo"`
	SavePath string  `json:"save_path"`
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
	return b.BvInfo.Bvid
}

func (b *BiliBili) GetVideoTitle() interface{} {
	return b.BvInfo.Title
}

func (b *BiliBili) GetSavePath() interface{} {
	return b.SavePath
}

func (b *BiliBili) GetVideoInfo() (interface{}, error) {
	return b.BvInfo, nil
}

func (b *BiliBili) Download(path string, id interface{}) (success int, fail int, err error) {
	bvid := id.(string)
	sdk := NewSDK()
	// 查询page列表
	bvInfo, err := sdk.GetBvInfo(bvid)
	if err != nil {
		return 0, 0, err
	}
	b.BvInfo = bvInfo
	directory := strings.Replace(bvInfo.Title, " ", "_", -1)
	b.SavePath = fmt.Sprintf("%s/%s", path, directory)
	_, err = os.Stat(b.SavePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return 0, 0, err
		}
		err = os.Mkdir(b.SavePath, os.ModePerm)
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
			file = fmt.Sprintf("%s/%s.mp4", b.SavePath, fileName)
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
	log.Printf("success: %d, fail: %d", success, fail)
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
