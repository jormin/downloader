package entity

// Data
type Data struct {
	Todos map[string]Todo `json:"todos" remark:"存储数据"`
}

type Todo struct {
	ID         string `json:"id" remark:"唯一ID"`
	Content    string `json:"content" remark:"内容"`
	Date       string `json:"date" remark:"日期"`
	Level      int    `json:"level" remark:"等级，1-3，1最小，3最大"`
	Status     int    `json:"status" remark:"状态 0:未完成 1:已完成"`
	CreateTime int64  `json:"create_time" remark:"创建时间"`
	UpdateTime int64  `json:"update_time" remark:"更新时间"`
}

const (
	TodoLevelLow          = 1 // level: low
	TodoLevelMiddle       = 2 // level: middle
	TodoLevelHigh         = 3 // level: high
	TodoStatusUncompleted = 0 // status: uncompleted
	TodoStatusCompleted   = 1 // status: completed
	TodoStatusDiscarded   = 2 // status: discarded
)
