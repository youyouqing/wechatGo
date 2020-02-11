package service

import (
	"github.com/robfig/cron/v3"
	"wechatGin.cthai.cn/err"
)

type JobQueue struct {
	Jobs map[int]*TaskJob
}

type TaskJob struct {
	TaskId    int
	UserId    int
	TaskType  int
	EntryId   cron.EntryID
	IsRunning bool
}

const TASKTYPE_CRON = 1 // cron表达式任务
const TASKTYPE_ONCE = 2 // 一次性任务

func NewJobQueue() *JobQueue {
	return &JobQueue{}
}

func NewTaskJob(taskId int, userId int, taskType int, entryId cron.EntryID) *TaskJob {
	return &TaskJob{
		TaskId:    taskId,
		UserId:    userId,
		TaskType:  taskType,
		EntryId:   entryId,
		IsRunning: false,
	}
}

func (this *JobQueue) AddJob(job *TaskJob) (*TaskJob, error) {

	if j, ok := this.Jobs[job.TaskId]; ok {
		return j, &err.JobError{
			JobId: job.TaskId,
			Msg:   "有正在执行的任务，任务重复创建",
		}
	} else {
		this.Jobs[job.TaskId] = job
	}
	return job, nil
}
