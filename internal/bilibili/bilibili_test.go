package bilibili

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/jormin/downloader/internal"
)

func TestBiliBili_Download(t *testing.T) {
	type fields struct {
		bvInfo   *BvInfo
		savePath string
	}
	type args struct {
		path string
		id   interface{}
	}
	tests := []struct {
		name        string
		bili        *BiliBili
		args        args
		wantSuccess int
		wantFail    int
		wantErr     bool
	}{
		{
			name: "01",
			bili: &BiliBili{
				bvInfo:   nil,
				savePath: "",
			},
			args: args{
				path: "./test",
				id:   "BV1QX4y1G7fG",
			},
			wantSuccess: 1,
			wantFail:    0,
			wantErr:     false,
		},
		{
			name: "02",
			bili: &BiliBili{
				bvInfo:   nil,
				savePath: "",
			},
			args: args{
				path: "./test",
				id:   "BV1QX4y1G7fG",
			},
			wantSuccess: 0,
			wantFail:    0,
			wantErr:     true,
		},
		{
			name: "03",
			bili: &BiliBili{
				bvInfo:   nil,
				savePath: "",
			},
			args: args{
				path: "./test",
				id:   "BV1NU4y1H71S",
			},
			wantSuccess: 2,
			wantFail:    0,
			wantErr:     false,
		},
		{
			name: "04",
			bili: &BiliBili{
				bvInfo:   nil,
				savePath: "",
			},
			args: args{
				path: "",
				id:   "BV1QX4y1G7fGxxx",
			},
			wantSuccess: 0,
			wantFail:    0,
			wantErr:     true,
		},
		{
			name: "05",
			bili: &BiliBili{
				bvInfo:   nil,
				savePath: "",
			},
			args: args{
				path: "abcdefg",
				id:   "BV1QX4y1G7fG",
			},
			wantSuccess: 0,
			wantFail:    0,
			wantErr:     true,
		},
		{
			name: "06",
			bili: &BiliBili{
				bvInfo:   nil,
				savePath: "/abcdefg",
			},
			args: args{
				path: "abcdefg",
				id:   "BV1QX4y1G7fG",
			},
			wantSuccess: 0,
			wantFail:    0,
			wantErr:     true,
		},
	}
	// create category `test` to save download files before test
	_ = os.Mkdir("./test", os.ModePerm)
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				gotSuccess, gotFail, err := tt.bili.Download(tt.args.path, tt.args.id)
				fmt.Println(err)
				if (err != nil) != tt.wantErr {
					t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotSuccess != tt.wantSuccess {
					t.Errorf("Download() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
				}
				if gotFail != tt.wantFail {
					t.Errorf("Download() gotFail = %v, want %v", gotFail, tt.wantFail)
				}
			},
		)
	}
	// delete download files after test
	_ = os.RemoveAll("./test")
}

func TestBiliBili_DownloadPages(t *testing.T) {
	savePath := "./test_page"
	type args struct {
		bvid      string
		pages     []Pages
		directory string
	}
	tests := []struct {
		name        string
		bili        *BiliBili
		args        args
		wantSuccess int
		wantFail    int
	}{

		{
			name: "01",
			bili: &BiliBili{
				sdk:      SDK{},
				savePath: savePath,
			},
			args: args{
				bvid: "BV1nJ411R7gk",
				pages: []Pages{
					{
						Cid:  133079147,
						Part: "test_1",
					},
				},
				directory: "test",
			},
			wantSuccess: 1,
			wantFail:    0,
		},
		{
			name: "02",
			bili: &BiliBili{
				sdk:      SDK{},
				savePath: savePath,
			},
			args: args{
				bvid: "BV1nJ411R7gkxxx",
				pages: []Pages{
					{
						Cid:  133079147,
						Part: "test_2",
					},
				},
				directory: "test",
			},
			wantSuccess: 0,
			wantFail:    0,
		},
		{
			name: "03",
			bili: &BiliBili{
				sdk:      SDK{},
				savePath: savePath,
			},
			args: args{
				bvid: "BV1nJ411R7gk",
				pages: []Pages{
					{
						Cid:  133079147,
						Part: "test_3",
					},
					{
						Cid:  133079147,
						Part: "test_4",
					},
				},
				directory: "test",
			},
			wantSuccess: 2,
			wantFail:    0,
		},
		{
			name: "04",
			bili: &BiliBili{
				sdk:      SDK{},
				savePath: "/",
			},
			args: args{
				bvid: "BV1nJ411R7gk",
				pages: []Pages{
					{
						Cid:  133079147,
						Part: "test_3",
					},
					{
						Cid:  133079147,
						Part: "test_4",
					},
				},
				directory: "test",
			},
			wantSuccess: 0,
			wantFail:    2,
		},
	}
	// create category `test` to save download files before test
	_ = os.Mkdir(savePath, os.ModePerm)
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tt.bili.savePath = savePath
				gotSuccess, gotFail := tt.bili.DownloadPages(tt.args.bvid, tt.args.pages, tt.args.directory)
				if gotSuccess != tt.wantSuccess {
					t.Errorf("DownloadPages() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
				}
				if gotFail != tt.wantFail {
					t.Errorf("DownloadPages() gotFail = %v, want %v", gotFail, tt.wantFail)
				}
			},
		)
	}
	// delete download files after test
	_ = os.RemoveAll(savePath)
}

func TestBiliBili_GetSavePath(t *testing.T) {
	type fields struct {
		bvInfo   *BvInfo
		savePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "01",
			fields: fields{
				savePath: "/test",
			},
			want: "/test",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				b := &BiliBili{
					bvInfo:   tt.fields.bvInfo,
					savePath: tt.fields.savePath,
				}
				if got := b.GetSavePath(); got != tt.want {
					t.Errorf("GetSavePath() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestBiliBili_GetSiteName(t *testing.T) {
	type fields struct {
		bvInfo   *BvInfo
		savePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "01",
			fields: fields{},
			want:   "bilibili",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				b := &BiliBili{
					bvInfo:   tt.fields.bvInfo,
					savePath: tt.fields.savePath,
				}
				if got := b.GetSiteName(); got != tt.want {
					t.Errorf("GetSiteName() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestBiliBili_GetSiteUrl(t *testing.T) {
	type fields struct {
		bvInfo   *BvInfo
		savePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "01",
			fields: fields{
				bvInfo:   nil,
				savePath: "",
			},
			want: "https://www.bilibili.com/",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				b := &BiliBili{
					bvInfo:   tt.fields.bvInfo,
					savePath: tt.fields.savePath,
				}
				if got := b.GetSiteURL(); got != tt.want {
					t.Errorf("GetSiteURL() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestBiliBili_GetVideoID(t *testing.T) {
	type fields struct {
		bvInfo   *BvInfo
		savePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "01",
			fields: fields{
				bvInfo: &BvInfo{
					Bvid: "BV1nJ411R7gk",
				},
				savePath: "",
			},
			want: "BV1nJ411R7gk",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				b := &BiliBili{
					bvInfo:   tt.fields.bvInfo,
					savePath: tt.fields.savePath,
				}
				if got := b.GetVideoID(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetVideoID() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestBiliBili_GetVideoInfo(t *testing.T) {
	type fields struct {
		bvInfoStr string
		savePath  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *internal.Video
		wantErr bool
	}{
		{
			name: "01",
			fields: fields{
				bvInfoStr: `{"bvid":"BV1QX4y1G7fG","aid":586724070,"videos":1,"tid":27,"tname":"综合","copyright":1,"pic":"http://i1.hdslb.com/bfs/archive/cf435a1f9d28f2504c8631c8833524c266b25a08.png","title":"超级宝贝JOJO：一起去爬山","pubdate":1613483877,"ctime":1613483877,"desc":"-","desc_v2":[{"raw_text":"-","type":1,"biz_id":0}],"state":0,"duration":68,"mission_id":16360,"rights":{"bp":0,"elec":0,"download":1,"movie":0,"pay":0,"hd5":1,"no_reprint":1,"autoplay":1,"ugc_pay":0,"is_cooperation":0,"ugc_pay_preview":0,"no_background":0,"clean_mode":0,"is_stein_gate":0},"owner":{"mid":631032199,"name":"宝宝巴士动画","face":"http://i1.hdslb.com/bfs/face/dea719ffd2f000506c4ae48348b3035165431dc7.jpg"},"stat":{"aid":586724070,"view":52072,"danmaku":2,"reply":4,"favorite":14,"coin":13,"share":17,"now_rank":0,"his_rank":0,"like":64,"dislike":0,"evaluation":"","argue_msg":""},"dynamic":"","cid":298528945,"dimension":{"width":1920,"height":1080,"rotate":0},"no_cache":false,"pages":[{"cid":298528945,"page":1,"from":"vupload","part":"1613483514414.mp4","duration":68,"vid":"","weblink":"","dimension":{"width":1920,"height":1080,"rotate":0}}],"subtitle":{"allow_submit":false,"list":[]},"user_garb":{"url_image_ani_cut":""}}`,
			},
			want: &internal.Video{
				ID:       "BV1QX4y1G7fG",
				Title:    "超级宝贝JOJO：一起去爬山",
				Duration: 68,
				Pages: []internal.Page{
					{
						ID:       298528945,
						Title:    "1613483514414.mp4",
						Duration: 68,
						Width:    1920,
						Height:   1080,
						Rotate:   0,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "02",
			fields: fields{
				bvInfoStr: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				b := &BiliBili{
					savePath: tt.fields.savePath,
				}
				if tt.fields.bvInfoStr == "" {
					b.bvInfo = nil
				} else {
					var bvInfo BvInfo
					if tt.fields.bvInfoStr != "" {
						_ = json.Unmarshal([]byte(tt.fields.bvInfoStr), &bvInfo)
					}
					b.bvInfo = &bvInfo
				}
				got, err := b.GetVideoInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("GetVideoInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetVideoInfo() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestBiliBili_GetVideoTitle(t *testing.T) {
	type fields struct {
		bvInfo   *BvInfo
		savePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "01",
			fields: fields{
				bvInfo: &BvInfo{
					Title: "超级宝贝 JOJO 儿歌 合集 (+156P)",
				},
			},
			want: "超级宝贝 JOJO 儿歌 合集 (+156P)",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				b := &BiliBili{
					bvInfo:   tt.fields.bvInfo,
					savePath: tt.fields.savePath,
				}
				if got := b.GetVideoTitle(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetVideoTitle() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestNewBiliBili(t *testing.T) {
	tests := []struct {
		name string
		want *BiliBili
	}{
		{
			name: "01",
			want: NewBiliBili(),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NewBiliBili(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewBiliBili() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
