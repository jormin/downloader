package bilibili

import (
	"errors"
	"testing"
)

func TestFailedPages_String(t *testing.T) {
	type fields struct {
		Pages Pages
		Err   error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "01",
			fields: fields{
				Pages: Pages{},
				Err:   nil,
			},
			want: "pages: {Cid:0 Page:0 From: Part: Duration:0 Vid: Weblink: Dimension:{Width:0 Height:0 Rotate:0}}, failed error: <nil>",
		},
		{
			name: "02",
			fields: fields{
				Pages: Pages{
					Cid:      187262840,
					Page:     1,
					From:     "vupload",
					Part:     "佩奇又开始玩mc",
					Duration: 331,
					Vid:      "",
					Weblink:  "",
					Dimension: Dimension{
						Width:  1280,
						Height: 720,
						Rotate: 0,
					},
				},
				Err: errors.New("test error"),
			},
			want: "pages: {Cid:187262840 Page:1 From:vupload Part:佩奇又开始玩mc Duration:331 Vid: Weblink: Dimension:{Width:1280 Height:720 Rotate:0}}, failed error: test error",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				f := &FailedPages{
					Pages: tt.fields.Pages,
					Err:   tt.fields.Err,
				}
				if got := f.String(); got != tt.want {
					t.Errorf("String() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
