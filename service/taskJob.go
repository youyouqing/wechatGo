package service

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"sync"
	"wechatGin.cthai.cn/err"
)
var queueInstance *JobQueue
var once sync.Once

type JobQueue struct {
	Jobs map[int]*TaskJob
}

type TaskJob struct {
	TaskId    int
	UserId    int
	TaskType  int
	EntryId   cron.EntryID
	IsRunning bool
	RepeatCrontStr string
}

const TASKTYPE_CRON = 1 // cron表达式任务
const TASKTYPE_ONCE = 2 // 一次性任务

func ShareInstanceQueue() *JobQueue {
	once.Do(func() {
		queueInstance = newJobQueue()
	})
	return queueInstance
}

func newJobQueue() *JobQueue {
	return &JobQueue{make(map[int] *TaskJob)}
}

func NewTaskJob(taskId int, userId int, taskType int, entryId cron.EntryID,isRunning bool,repeatCrontStr string) *TaskJob {
	return &TaskJob{
		TaskId:    taskId,
		UserId:    userId,
		TaskType:  taskType,
		EntryId:   entryId,
		IsRunning: false,
		RepeatCrontStr: repeatCrontStr,
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
		fmt.Println(this.Jobs)
	}
	return job, nil
}

