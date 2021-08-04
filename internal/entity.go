package internal

// Task download task
type Task struct {
	ID           string `json:"id" dsc:"unique task id"`
	Date         string `json:"date" desc:"download date"`
	Video        *Video `json:"video" desc:"video info"`
	Status       int    `json:"status" desc:"status of task, 0 for fail and 1 for success"`
	Error        string `json:"error" desc:"task error"`
	StartTime    int64  `json:"start_time" desc:"start time of task"`
	StartTimeFmt string `json:"start_time_fmt" desc:"formatted start time of task"`
	EndTime      int64  `json:"end_ime" desc:"format end time of task"`
	EndTimeFmt   string `json:"end_ime_fmt" desc:"formatted end time of task"`
}

// Video video info
type Video struct {
	ID       interface{} `json:"id" desc:"unique video id from video site, not page id"`
	Title    string      `json:"title" desc:"video title"`
	Duration int         `json:"duration" desc:"total duration of task videos"`
	Pages    []Page      `json:"pages" desc:"all pages of task"`
}

// Page the specific page info
type Page struct {
	ID       interface{} `json:"id" desc:"id of video"`
	Title    string      `json:"title" desc:"title of video"`
	Duration int         `json:"duration" desc:"duration of video"`
	Width    int         `json:"width" desc:"width of video"`
	Height   int         `json:"height" desc:"height of video"`
	Rotate   int         `json:"rotate" desc:"rotate of video"`
}

const (
	// TaskStatusFail status: fail
	TaskStatusFail = iota
	// TaskStatusSuccess status: success
	TaskStatusSuccess
)
