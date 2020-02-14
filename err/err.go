package err

import (
	"wechatGin.cthai.cn/utils"
)

type JobError struct {
	JobId int
	Msg   string
}

func (e *JobError) Error() string {
	if e == nil {
		return "<nil>"
	}
	var s string
	if e.Msg != "" {
		s = "任务错误原因: " + e.Msg + " 任务id:" + utils.IntToString(e.JobId)
	}
	return s
}

func (e *JobError) String() string {
	return "任务id=" + utils.IntToString(e.JobId) + "错误信息:" + e.Msg
}
