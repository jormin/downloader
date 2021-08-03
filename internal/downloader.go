package internal

// Downloader
type Downloader interface {
	// GetSiteName the name of site to download video, such as `bilibili`.
	GetSiteName() string
	// GetSiteUrl the url of site to download video, such as `https://www.bilibili.com/`.
	GetSiteUrl() string
	// GetTaskID the unique id of download task
	GetTaskID() string
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
