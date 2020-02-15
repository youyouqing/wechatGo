package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"time"
	. "wechatGin.cthai.cn/model"
	"wechatGin.cthai.cn/utils"
)

// 添加任务
func HandleTaskAdd(ctx *gin.Context) {
	query := TaskAdd{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(500, utils.NewResultError500(err.Error()))
		ctx.Abort()
	} else {
		addTask(ctx, query)
	}
}

// 添加并开启任务
func HandleTaskAddAndStart(ctx *gin.Context) {
	query := TaskAdd{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(500, utils.NewResultError500(err.Error()))
		ctx.Abort()
	} else {
		addTaskAndStart(ctx, query)
	}
}

// 开启任务
func HandleTaskStart(ctx *gin.Context) {
	taskIdStr := ctx.Param("id")
	taskId, err := utils.StringToInt(taskIdStr)
	if err != nil {
		ctx.JSON(500, utils.NewResultError500(err.Error()))
		ctx.Abort()
	}
	queue := ShareInstanceQueue()
	taskJob, err := queue.GetJob(taskId)
	if err != nil {
		ctx.JSON(500, utils.NewResultError500("不存在任务，id="+taskIdStr))
		ctx.Abort()
		return
	}
	if taskJob.IsRunning == true {
		ctx.JSON(500, utils.NewResultError500("任务已经开启，请勿重复开启"))
		ctx.Abort()
		return
	}
	if taskJob.TaskType == TASKTYPE_CRON {
		c := cron.New(cron.WithSeconds())
		entryId, err := c.AddFunc(taskJob.RepeatCrontStr, func() {
			fmt.Println("正在执行任务: taskId=" + utils.IntToString(taskJob.TaskId) + "" + "时间表达式为:" + taskJob.RepeatCrontStr + "当前时间:" + time.Now().Format("2006-01-02 15:04:05"))
		})
		if err != nil {
			queue.DelJob(taskId)
			ctx.JSON(500, utils.NewResultError500(err.Error()))
			ctx.Abort()
		}
		taskJob.Entry = *NewEntry(entryId, c)
		c.Start()
		taskJob.IsRunning = true
		ctx.JSON(200, utils.NewResultSuccess200(true))
		ctx.Abort()
	} else {
		taskJob.IsRunning = true
		distanceSec := utils.TimeStringToTimestamp(taskJob.KnockTime) - utils.CurrentTimestamp()
		afterTime := time.Duration(distanceSec) * time.Second
		if afterTime < 0 {
			ctx.JSON(500, utils.NewResultError500("一次性任务时间不能比当前时间早"))
			ctx.Abort()
			return
		}
		afterTimeInstance := time.AfterFunc(afterTime, func() {
			fmt.Println("正在执行一次性任务: taskId=" + utils.IntToString(taskJob.TaskId) + "任务设定时间:" + taskJob.KnockTime + "当前时间:" + time.Now().Format("2006-01-02 15:04:05"))
		})
		taskJob.AfterTimeInstance = afterTimeInstance
	}
}

// 停止任务
func HandleTaskStop(ctx *gin.Context) {
	taskIdStr := ctx.Param("id")
	taskId, err := utils.StringToInt(taskIdStr)
	if err != nil {
		ctx.JSON(500, utils.NewResultError500(err.Error()))
		ctx.Abort()
	}
	queue := ShareInstanceQueue()
	taskJob, err := queue.GetJob(taskId)
	if err != nil {
		ctx.JSON(500, utils.NewResultError500(err.Error()))
		ctx.Abort()
		return
	}
	if taskJob.TaskType == TASKTYPE_CRON {
		if taskJob.Entry.EntryId == 0 {
			ctx.JSON(500, utils.NewResultError500("任务尚未开启，请勿停止"))
			ctx.Abort()
			return
		}
		if taskJob.IsRunning == false {
			ctx.JSON(500, utils.NewResultError500("任务已经停止，请勿重复停止"))
			ctx.Abort()
			return
		}
		taskJob.Entry.Cron.Stop()
		taskJob.IsRunning = false
		ctx.JSON(200, utils.NewResultSuccess200(true))
	}else {
		stopRes := taskJob.AfterTimeInstance.Stop()
		ctx.JSON(200, utils.NewResultSuccess200(stopRes))
	}

}

func addTask(ctx *gin.Context, task TaskAdd) {
	queue := ShareInstanceQueue()
	job := NewTaskJob(task.Id, task.UserId, TASKTYPE_CRON, task.KnockTime, task.NotifyTitle, task.NotifyContent,nil, *NewEntry(0, nil), false, task.RepeatCrontStr)
	_, err := queue.AddJob(job)
	if err != nil {
		ctx.JSON(500, utils.NewResultError500(err.Error()))
		return
	}
	if task.RepeatFlag == 1 {
		// cron
		ctx.JSON(200, utils.NewResultSuccess200(job))
	} else {
		// 延时一次性任务
		job.TaskType = TASKTYPE_ONCE
		ctx.JSON(200, utils.NewResultSuccess200(job))
	}
}
func addTaskAndStart(ctx *gin.Context, task TaskAdd) {
	queue := ShareInstanceQueue()
	entry := NewEntry(0, nil)
	job := NewTaskJob(task.Id, task.UserId, TASKTYPE_CRON, task.KnockTime, task.NotifyTitle, task.NotifyContent,nil, *entry, true, task.RepeatCrontStr)
	if task.RepeatFlag == 1 {
		// cron
		_, err := queue.AddJob(job)
		if err != nil {
			ctx.JSON(500, utils.NewResultError500(err.Error()))
			ctx.Abort()
			return
		}
		c := cron.New(cron.WithSeconds())
		entryId, err := c.AddFunc(task.RepeatCrontStr, func() {
			fmt.Println("正在执行任务: taskId=" + utils.IntToString(task.Id) + "" + "时间表达式为:" + task.RepeatCrontStr + "当前时间:" + time.Now().Format("2006-01-02 15:04:05"))
		})
		if err != nil {
			queue.DelJob(task.Id)
			ctx.JSON(500, utils.NewResultError500(err.Error()))
			ctx.Abort()
			return
		}
		entry.EntryId = entryId
		entry.Cron = c
		queue.UpdateEntry(task.Id, *entry)
		c.Start()
		ctx.JSON(200, utils.NewResultSuccess200(task))
		ctx.Abort()
		PrintQueueJob()
		return

	} else {
		// 延时一次性任务
		job.TaskType = TASKTYPE_ONCE
		distanceSec := utils.TimeStringToTimestamp(job.KnockTime) - utils.CurrentTimestamp()
		afterTime := time.Duration(distanceSec) * time.Second
		if afterTime < 0 {
			ctx.JSON(500, utils.NewResultError500("一次性任务时间不能比当前时间早"))
			ctx.Abort()
			return
		}
		time.AfterFunc(afterTime, func() {
			fmt.Println("正在执行一次性任务: taskId=" + utils.IntToString(task.Id) + "任务设定时间:" + task.KnockTime + "当前时间:" + time.Now().Format("2006-01-02 15:04:05"))
		})
		ctx.JSON(200, utils.NewResultSuccess200(task))
	}
	PrintQueueJob()
}

// debug print queue
func PrintQueueJob() {
	jobs := ShareInstanceQueue().Jobs
	for key := range jobs {
		fmt.Println(key)
		fmt.Println("TaskId==>", jobs[key].TaskId)
		fmt.Println("TaskType==>", jobs[key].TaskType)
		fmt.Println("RepeatCrontStr==>", jobs[key].RepeatCrontStr)
		fmt.Println("UserId==>", jobs[key].UserId)
		fmt.Println("IsRunning==>", jobs[key].IsRunning)
		fmt.Println("EntryId==>", jobs[key].Entry.EntryId)
		fmt.Println("cron==>", jobs[key].Entry.Cron)
	}
}
