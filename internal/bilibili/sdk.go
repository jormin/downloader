package bilibili

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	ApiAvInfo = iota
	ApiBvInfo
	ApiPageList
	ApiCInfo
	ResCodeOK = 0
)

// Api地址
var apis = map[int]string{
	ApiAvInfo:   "https://api.bilibili.com/x/web-interface/view?aid=%d",
	ApiBvInfo:   "https://api.bilibili.com/x/web-interface/view?bvid=%s",
	ApiPageList: "https://api.bilibili.com/x/player/pagelist?bvid=%s",
	ApiCInfo:    "https://api.bilibili.com/x/player/playurl?bvid=%s&cid=%d&otype=json",
}

// SDK
type SDK struct{}

// NewSDK
func NewSDK() *SDK {
	return &SDK{}
}

// GetAvInfo 获取AV信息
func (bl *SDK) GetAvInfo(avid int) (*AvInfo, error) {
	var avInfo AvInfo
	err := bl.http(fmt.Sprintf(apis[ApiAvInfo], avid), nil, &avInfo)
	if err != nil {
		return nil, err
	}
	return &avInfo, nil
}

// GetBvInfo 获取BV信息
func (bl *SDK) GetBvInfo(bvid string) (*BvInfo, error) {
	var bvInfo BvInfo
	err := bl.http(fmt.Sprintf(apis[ApiBvInfo], bvid), nil, &bvInfo)
	if err != nil {
		return nil, err
	}
	return &bvInfo, nil
}

// GetPageListByBVID 根据BVID查询page列表
func (bl *SDK) GetPagesByBVID(bvid string) ([]Page, error) {
	var pages []Page
	err := bl.http(fmt.Sprintf(apis[ApiPageList], bvid), nil, &pages)
	if err != nil {
		return nil, err
	}
	return pages, nil
}

// GetCInfo 后去C信息
func (bl *SDK) GetCInfo(bvid string, cid int) (*CInfo, error) {
	var cInfo CInfo
	err := bl.http(fmt.Sprintf(apis[ApiCInfo], bvid, cid), nil, &cInfo)
	if err != nil {
		return nil, err
	}
	return &cInfo, nil
}

// DownloadVideo 下载视频
func (bl *SDK) DownloadVideo(bvid string, url string, file string) error {
	_, err := os.Stat(file)
	if err == nil || !os.IsNotExist(err) {
		return errors.New("file exist")
	}
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Referer", fmt.Sprintf("https://www.bilibili.com/video/%s", bvid))
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}
	var buf []byte = make([]byte, 1024)
	for {
		n, err := res.Body.Read(buf)
		if err == io.EOF {
			break
		}
		_, _ = out.WriteString(string(buf[:n]))
	}
	return nil
}

// http
func (bl *SDK) http(url string, headers map[string]string, data interface{}) error {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	var biliRes BiliBiliRes
	_ = json.Unmarshal(b, &biliRes)
	if biliRes.Code != ResCodeOK {
		return errors.New(biliRes.Message)
	}
	biliData := biliRes.Data
	b, _ = json.Marshal(biliData)
	_ = json.Unmarshal(b, data)
	return nil
}
