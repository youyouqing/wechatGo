package service

import (
	"github.com/robfig/cron/v3"
	"sync"
	"wechatGin.cthai.cn/err"
)

var queueInstance *JobQueue
var once sync.Once

// 队列实体
type JobQueue struct {
	Jobs map[int]*TaskJob
}

// 任务实体
type TaskJob struct {
	TaskId         int
	UserId         int
	TaskType       int
	Entry          Entry
	IsRunning      bool
	RepeatCrontStr string
}

// cron 定时器实体
type Entry struct {
	EntryId cron.EntryID
	Cron    *cron.Cron
}

func NewEntry(entryId cron.EntryID, cron *cron.Cron) *Entry {
	return &Entry{
		EntryId: entryId,
		Cron:    cron,
	}
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
	return &JobQueue{make(map[int]*TaskJob)}
}

func NewTaskJob(taskId int, userId int, taskType int, entry Entry, isRunning bool, repeatCrontStr string) *TaskJob {
	return &TaskJob{
		TaskId:         taskId,
		UserId:         userId,
		TaskType:       taskType,
		Entry:          entry,
		IsRunning:      isRunning,
		RepeatCrontStr: repeatCrontStr,
	}
}

func (this *JobQueue) GetJob(jobId int) (*TaskJob, error) {
	if j, ok := this.Jobs[jobId]; ok {
		return j,nil
	}else {
		return nil,&err.JobError{JobId: jobId,Msg: "队列不存在该任务"}
	}
}

func (this *JobQueue) AddJob(job *TaskJob) (*TaskJob, error) {

	if j, ok := this.Jobs[job.TaskId]; ok {
		er := new(err.JobError)
		er.JobId = job.TaskId
		if !j.IsRunning {
			er.Msg = "已存在相同的任务为停止状态，请开启"
			return j, er
		}
		er.Msg = "有正在执行的任务，任务重复创建"
		return j, er
	} else {
		this.Jobs[job.TaskId] = job
	}
	return job, nil
}

func (this *JobQueue) DelJob(jobId int) *JobQueue {
	if _, ok := this.Jobs[jobId]; ok {
		delete(this.Jobs, jobId)
	}
	return this
}

func (this *JobQueue) UpdateEntry(jobId int, entry Entry) *JobQueue {
	if j, ok := this.Jobs[jobId]; ok {
		j.Entry = entry
	}
	return this
}
