package err

type JobError struct {
	JobId  int
	Msg string
}

func (e *JobError) Error() string {
	if e == nil {
		return "<nil>"
	}
	s := string(e.JobId)
	if e.Msg != "" {
		s = "任务错误原因: " + e.Msg + " 任务id: " + s
	}
	return s
}