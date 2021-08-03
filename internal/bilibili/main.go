package bilibili

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	bvid := "BV1Kr4y1A7G7"
	sdk := NewSDK()
	// 查询page列表
	bvInfo, err := sdk.GetBvInfo(bvid)
	if err != nil {
		panic(err)
	}
	var failedPages []FailedPages
	var wg sync.WaitGroup
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
				return
			}
			// 下载视频
			var file string
			if len(bvInfo.Pages) == 1 {
				file = fmt.Sprintf("%s.mp4", strings.Replace(bvInfo.Title, " ", "_", -1))
			} else {
				file = fmt.Sprintf("%s.mp4", strings.Replace(p.Part, " ", "_", -1))
			}
			err = sdk.DownloadVideo(cInfo.Durl[0].URL, file, bvid)
			if err != nil {
				failedPages = append(
					failedPages, FailedPages{
						Pages: p,
						Err:   err,
					},
				)
				return
			}
		}(p)
	}
	wg.Wait()
	for _, v := range failedPages {
		fmt.Println(v)
	}
}
