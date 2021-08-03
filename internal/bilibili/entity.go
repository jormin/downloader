package bilibili

import "fmt"

// BiliBiliRes
type BiliBiliRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	TTL     int         `json:"ttl"`
	Data    interface{} `json:"data"`
}

// Page
type Page struct {
	Cid       int    `json:"cid"`
	Page      int    `json:"page"`
	From      string `json:"from"`
	Part      string `json:"part"`
	Duration  int    `json:"duration"`
	Vid       string `json:"vid"`
	Weblink   string `json:"weblink"`
	Dimension struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		Rotate int `json:"rotate"`
	} `json:"dimension"`
}

// AVInfo
type AvInfo struct {
	Bvid      string    `json:"bvid"`
	Aid       int       `json:"aid"`
	Videos    int       `json:"videos"`
	Tid       int       `json:"tid"`
	Tname     string    `json:"tname"`
	Copyright int       `json:"copyright"`
	Pic       string    `json:"pic"`
	Title     string    `json:"title"`
	Pubdate   int       `json:"pubdate"`
	Ctime     int       `json:"ctime"`
	Desc      string    `json:"desc"`
	DescV2    []DescV2  `json:"desc_v2"`
	State     int       `json:"state"`
	Duration  int       `json:"duration"`
	Rights    Rights    `json:"rights"`
	Owner     Owner     `json:"owner"`
	Stat      Stat      `json:"stat"`
	Dynamic   string    `json:"dynamic"`
	Cid       int       `json:"cid"`
	Dimension Dimension `json:"dimension"`
	NoCache   bool      `json:"no_cache"`
	Pages     []Pages   `json:"pages"`
	Subtitle  Subtitle  `json:"subtitle"`
	UserGarb  UserGarb  `json:"user_garb"`
}
type DescV2 struct {
	RawText string `json:"raw_text"`
	Type    int    `json:"type"`
	BizID   int    `json:"biz_id"`
}
type Rights struct {
	Bp            int `json:"bp"`
	Elec          int `json:"elec"`
	Download      int `json:"download"`
	Movie         int `json:"movie"`
	Pay           int `json:"pay"`
	Hd5           int `json:"hd5"`
	NoReprint     int `json:"no_reprint"`
	Autoplay      int `json:"autoplay"`
	UgcPay        int `json:"ugc_pay"`
	IsCooperation int `json:"is_cooperation"`
	UgcPayPreview int `json:"ugc_pay_preview"`
	NoBackground  int `json:"no_background"`
	CleanMode     int `json:"clean_mode"`
	IsSteinGate   int `json:"is_stein_gate"`
}
type Owner struct {
	Mid  int    `json:"mid"`
	Name string `json:"name"`
	Face string `json:"face"`
}
type Stat struct {
	Aid        int    `json:"aid"`
	View       int    `json:"view"`
	Danmaku    int    `json:"danmaku"`
	Reply      int    `json:"reply"`
	Favorite   int    `json:"favorite"`
	Coin       int    `json:"coin"`
	Share      int    `json:"share"`
	NowRank    int    `json:"now_rank"`
	HisRank    int    `json:"his_rank"`
	Like       int    `json:"like"`
	Dislike    int    `json:"dislike"`
	Evaluation string `json:"evaluation"`
	ArgueMsg   string `json:"argue_msg"`
}

type Dimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Rotate int `json:"rotate"`
}
type Pages struct {
	Cid       int       `json:"cid"`
	Page      int       `json:"page"`
	From      string    `json:"from"`
	Part      string    `json:"part"`
	Duration  int       `json:"duration"`
	Vid       string    `json:"vid"`
	Weblink   string    `json:"weblink"`
	Dimension Dimension `json:"dimension"`
}
type Subtitle struct {
	AllowSubmit bool          `json:"allow_submit"`
	List        []interface{} `json:"list"`
}
type UserGarb struct {
	URLImageAniCut string `json:"url_image_ani_cut"`
}

// BvInfo
type BvInfo struct {
	Bvid      string    `json:"bvid"`
	Aid       int       `json:"aid"`
	Videos    int       `json:"videos"`
	Tid       int       `json:"tid"`
	Tname     string    `json:"tname"`
	Copyright int       `json:"copyright"`
	Pic       string    `json:"pic"`
	Title     string    `json:"title"`
	Pubdate   int       `json:"pubdate"`
	Ctime     int       `json:"ctime"`
	Desc      string    `json:"desc"`
	DescV2    []DescV2  `json:"desc_v2"`
	State     int       `json:"state"`
	Duration  int       `json:"duration"`
	Rights    Rights    `json:"rights"`
	Owner     Owner     `json:"owner"`
	Stat      Stat      `json:"stat"`
	Dynamic   string    `json:"dynamic"`
	Cid       int       `json:"cid"`
	Dimension Dimension `json:"dimension"`
	NoCache   bool      `json:"no_cache"`
	Pages     []Pages   `json:"pages"`
	Subtitle  Subtitle  `json:"subtitle"`
	UserGarb  UserGarb  `json:"user_garb"`
}

// CInfo
type CInfo struct {
	From              string           `json:"from"`
	Result            string           `json:"result"`
	Message           string           `json:"message"`
	Quality           int              `json:"quality"`
	Format            string           `json:"format"`
	Timelength        int              `json:"timelength"`
	AcceptFormat      string           `json:"accept_format"`
	AcceptDescription []string         `json:"accept_description"`
	AcceptQuality     []int            `json:"accept_quality"`
	VideoCodecid      int              `json:"video_codecid"`
	SeekParam         string           `json:"seek_param"`
	SeekType          string           `json:"seek_type"`
	Durl              []Durl           `json:"durl"`
	SupportFormats    []SupportFormats `json:"support_formats"`
	HighFormat        interface{}      `json:"high_format"`
}
type Durl struct {
	Order     int      `json:"order"`
	Length    int      `json:"length"`
	Size      int      `json:"size"`
	Ahead     string   `json:"ahead"`
	Vhead     string   `json:"vhead"`
	URL       string   `json:"url"`
	BackupURL []string `json:"backup_url"`
}
type SupportFormats struct {
	Quality        int    `json:"quality"`
	Format         string `json:"format"`
	NewDescription string `json:"new_description"`
	DisplayDesc    string `json:"display_desc"`
	Superscript    string `json:"superscript"`
}

// 失败信息
type FailedPages struct {
	Pages Pages
	Err   error
}

// String
func (f *FailedPages) String() string {
	return fmt.Sprintf("pages: %+v, failed error: %v", f.Pages, f.Err)
}
