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
	// APIAvInfo the no of api `https://api.bilibili.com/x/web-interface/view?aid=%d`
	APIAvInfo = iota
	// APIBvInfo the no of api `https://api.bilibili.com/x/web-interface/view?bvid=%s`
	APIBvInfo
	// APIPageList the no of api `https://api.bilibili.com/x/player/pagelist?bvid=%s`
	APIPageList
	// APICInfo the no of api `https://api.bilibili.com/x/player/playurl?bvid=%s&cid=%d&otype=json`
	APICInfo
)

const (
	// ResCodeOK res code: success
	ResCodeOK = 0
)

// Api地址
var apis = map[int]string{
	APIAvInfo:   "https://api.bilibili.com/x/web-interface/view?aid=%d",
	APIBvInfo:   "https://api.bilibili.com/x/web-interface/view?bvid=%s",
	APIPageList: "https://api.bilibili.com/x/player/pagelist?bvid=%s",
	APICInfo:    "https://api.bilibili.com/x/player/playurl?bvid=%s&cid=%d&otype=json",
}

// SDK bilibili api sdk
type SDK struct{}

// NewSDK get the new bilibili sdk
func NewSDK() *SDK {
	return &SDK{}
}

// GetAvInfo get av info by avid
func (bl *SDK) GetAvInfo(avid int) (*AvInfo, error) {
	var avInfo AvInfo
	err := bl.http(fmt.Sprintf(apis[APIAvInfo], avid), nil, &avInfo)
	if err != nil {
		return nil, err
	}
	return &avInfo, nil
}

// GetBvInfo get bvinfo by bvid
func (bl *SDK) GetBvInfo(bvid string) (*BvInfo, error) {
	var bvInfo BvInfo
	err := bl.http(fmt.Sprintf(apis[APIBvInfo], bvid), nil, &bvInfo)
	if err != nil {
		return nil, err
	}
	return &bvInfo, nil
}

// GetPagesByBVID get pages by bvid
func (bl *SDK) GetPagesByBVID(bvid string) ([]Page, error) {
	var pages []Page
	err := bl.http(fmt.Sprintf(apis[APIPageList], bvid), nil, &pages)
	if err != nil {
		return nil, err
	}
	return pages, nil
}

// GetCInfo get cinfo bv bvid and cid
func (bl *SDK) GetCInfo(bvid string, cid int) (*CInfo, error) {
	var cInfo CInfo
	err := bl.http(fmt.Sprintf(apis[APICInfo], bvid, cid), nil, &cInfo)
	if err != nil {
		return nil, err
	}
	return &cInfo, nil
}

// DownloadVideo down video by url and bvid
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

// http send http request
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
	var biliRes Res
	_ = json.Unmarshal(b, &biliRes)
	if biliRes.Code != ResCodeOK {
		return errors.New(biliRes.Message)
	}
	biliData := biliRes.Data
	b, _ = json.Marshal(biliData)
	_ = json.Unmarshal(b, data)
	return nil
}
