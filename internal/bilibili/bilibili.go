package bilibili

import (
	"errors"
	"fmt"
	"os"

	"github.com/jormin/downloader/helper"
	"github.com/jormin/downloader/internal"
)

// BiliBili
type BiliBili struct {
	bvInfo   *BvInfo
	savePath string
	sdk      SDK
}

// GetSiteName the name of site to download video, such as `bilibili`.
func (b *BiliBili) GetSiteName() string {
	return "bilibili"
}

// GetSiteUrl the url of site to download video, such as `https://www.bilibili.com/`.
func (b *BiliBili) GetSiteUrl() string {
	return "https://www.bilibili.com/"
}

// GetVideoID the id of video that will be downloaded
func (b *BiliBili) GetVideoID() interface{} {
	return b.bvInfo.Bvid
}

// GetVideoTitle the title of video that will be downloaded
func (b *BiliBili) GetVideoTitle() interface{} {
	return b.bvInfo.Title
}

// GetSavePath the path to save video
func (b *BiliBili) GetSavePath() string {
	return b.savePath
}

// GetVideoInfo get video info
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

// Download download video by id
func (b *BiliBili) Download(path string, id interface{}) (success int, fail int, err error) {
	bvid := id.(string)
	// 查询page列表
	bvInfo, err := b.sdk.GetBvInfo(bvid)
	if err != nil {
		return 0, 0, errors.New(fmt.Sprintf("get video info from site bilibili error, please check video id is right"))
	}
	b.bvInfo = bvInfo
	directory := helper.RemoveIllegalCharacters(bvInfo.Title)
	b.savePath = fmt.Sprintf("%s/%s", path, directory)
	_, err = os.Stat(b.savePath)
	if err == nil {
		return 0, 0, errors.New("save path exist")
	}
	if os.IsNotExist(err) {
		err = os.Mkdir(b.savePath, os.ModePerm)
		if err != nil {
			return 0, 0, err
		}
	}
	success, fail = b.DownloadPages(bvid, bvInfo.Pages, directory)
	return success, fail, nil
}

// DownloadPage download page
func (b *BiliBili) DownloadPages(bvid string, pages []Pages, directory string) (success int, fail int) {
	// considering that external services have frequency restrictions, do not use goroutine
	for _, p := range pages {
		// download page
		var fileName string
		if len(pages) == 1 {
			fileName = directory
		} else {
			fileName = helper.RemoveIllegalCharacters(p.Part)
		}
		// get page info
		cInfo, err := b.sdk.GetCInfo(bvid, p.Cid)
		if err != nil {
			fmt.Printf("download video %d【%s】 fail, can not get video info\n", p.Cid, p.Part)
			continue
		}
		// download video
		file := fmt.Sprintf("%s/%s.flv", b.savePath, fileName)
		err = b.sdk.DownloadVideo(bvid, cInfo.Durl[0].URL, file)
		if err != nil {
			fmt.Printf("download video %d【%s】 fail, %v\n", p.Cid, p.Part, err)
			fail += 1
			continue
		}
		fmt.Printf("download video %d【%s】 success\n", p.Cid, p.Part)
		success += 1
	}
	return success, fail
}

// NewBiliBili
func NewBiliBili() *BiliBili {
	return &BiliBili{
		sdk: SDK{},
	}
}
