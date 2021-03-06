package internal

// Downloader the tool to download video from third-party sites
type Downloader interface {
	// GetSiteName the name of site to download video, such as `bilibili`.
	GetSiteName() string
	// GetSiteURL the url of site to download video, such as `https://www.bilibili.com/`.
	GetSiteURL() string
	// GetVideoID the id of video that will be downloaded
	GetVideoID() interface{}
	// GetVideoTitle the title of video that will be downloaded
	GetVideoTitle() interface{}
	// GetSavePath the path to save video
	GetSavePath() string
	// GetVideoInfo get video info
	GetVideoInfo() (*Video, error)
	// Download download video by id
	Download(path string, id interface{}) (success int, fail int, err error)
}
