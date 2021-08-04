package bilibili

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestNewSDK(t *testing.T) {
	tests := []struct {
		name string
		want *SDK
	}{
		{
			name: "normal",
			want: &SDK{},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NewSDK(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewSDK() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestSDK_DownloadVideo(t *testing.T) {
	type args struct {
		url  string
		file string
		bvid string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "01",
			args: args{
				url:  fmt.Sprintf("http://download.lerzen.com/test.txt?%d", time.Now().Unix()),
				file: "./test_sdk/a.txt",
				bvid: "BV1nJ411R7gk",
			},
			wantErr: false,
		},
		{
			name: "02",
			args: args{
				url:  fmt.Sprintf("http://download.lerzen.com/test.txt?%d", time.Now().Unix()),
				file: "./test_sdk/a.txt",
				bvid: "BV1nJ411R7gk",
			},
			wantErr: true,
		},
		{
			name: "03",
			args: args{
				url:  fmt.Sprintf("http://download.lerzen.com/test111.txt?%d", time.Now().Unix()),
				file: "./test_sdk/b.txt",
				bvid: "BV1Zi4y1x7Q2",
			},
			wantErr: true,
		},
		{
			name: "04",
			args: args{
				url:  fmt.Sprintf("http://download.lerzen.com/test.txt?%d", time.Now().Unix()),
				file: "/zxcvasdf/a.txt",
				bvid: "BV1Zi4y1x7Q2",
			},
			wantErr: true,
		},
		{
			name: "05",
			args: args{
				url:  "123",
				file: "/a.txt",
				bvid: "BV1Zi4y1x7Q2",
			},
			wantErr: true,
		},
	}
	// create category `test` to save download files before test
	_ = os.Mkdir("./test_sdk", os.ModePerm)
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				bl := &SDK{}
				if err := bl.DownloadVideo(tt.args.bvid, tt.args.url, tt.args.file); (err != nil) != tt.wantErr {
					t.Errorf("DownloadVideo() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
	// delete download files after test
	_ = os.RemoveAll("./test_sdk")
}

func TestSDK_GetAvInfo(t *testing.T) {
	type args struct {
		avid int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "01",
			args: args{
				avid: 540584194,
			},
			wantErr: false,
		},
		{
			name: "01",
			args: args{
				avid: 540584194121212,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				bl := &SDK{}
				_, err := bl.GetAvInfo(tt.args.avid)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetAvInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
}

func TestSDK_GetBvInfo(t *testing.T) {
	type args struct {
		bvid string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "01",
			args: args{
				bvid: "BV1Zi4y1x7Q2",
			},
			wantErr: false,
		},
		{
			name: "02",
			args: args{
				bvid: "BV1Zi4y1x7Q2xxxxx",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				bl := &SDK{}
				_, err := bl.GetBvInfo(tt.args.bvid)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetBvInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
}

func TestSDK_GetCInfo(t *testing.T) {
	type args struct {
		bvid string
		cid  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "01",
			args: args{
				bvid: "BV1Zi4y1x7Q2",
				cid:  187262840,
			},
			wantErr: false,
		},
		{
			name: "02",
			args: args{
				bvid: "BV1Zi4y1x7Q2xxx",
				cid:  187262840123,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				bl := &SDK{}
				_, err := bl.GetCInfo(tt.args.bvid, tt.args.cid)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetCInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
}

func TestSDK_GetPagesByBVID(t *testing.T) {
	type args struct {
		bvid string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "01",
			args: args{
				bvid: "BV1Zi4y1x7Q2",
			},
			wantErr: false,
		},
		{
			name: "02",
			args: args{
				bvid: "BV1Zi4y1x7Q2xxxxx",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				bl := &SDK{}
				_, err := bl.GetPagesByBVID(tt.args.bvid)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetPagesByBVID() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
}

func TestSDK_http(t *testing.T) {
	type args struct {
		url     string
		headers map[string]string
		data    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "01",
			args: args{
				url: "https://api.bilibili.com/x/player/pagelist?bvid=BV1Zi4y1x7Q2",
				headers: map[string]string{
					"a": "1",
					"b": "2",
				},
				data: &[]Page{},
			},
			wantErr: false,
		},
		{
			name: "02",
			args: args{
				url: "https://api.bilibili.com/x/player/pagelist?bvid=BV1Zi4y1x7Q2xxxxx",
				headers: map[string]string{
					"a": "1",
					"b": "2",
				},
				data: &[]Page{},
			},
			wantErr: true,
		},
		{
			name: "03",
			args: args{
				url:     "https://upos-sz-mirrorcoso1.bilivideo.com/upgcxcode/47/91/133079147/133079147_nb2-1-32.flv?e=ig8euxZM2rNcNbRVhwdVhwdlhWdVhwdVhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&uipk=5&nbs=1&deadline=1628052349&gen=playurlv2&os=coso1bv&oi=22403703&trid=6d8dcb37298648759460e682ea5df52fu&platform=pc&upsig=b162da52eb4bb0e21120b8c3e702ca21&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,platform&mid=0&bvc=vod&nettype=0&orderid=0,3&agrr=1&logo=80000000",
				headers: nil,
				data:    &[]Page{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				bl := &SDK{}
				if err := bl.http(tt.args.url, tt.args.headers, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("http() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}
