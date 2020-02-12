package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"strconv"
	. "wechatGin.cthai.cn/model"
)

func HandleTaskAdd(ctx *gin.Context) {
	query := TaskAdd{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(500, gin.H{
			"data": false,
			"msg":  err.Error(),
			"code": 500,
		})
		ctx.Abort()
	} else {
		errMsgChan := make(chan string)
		go addTask(query, errMsgChan)
		var errMsg string
		for {
			errMsg = <-errMsgChan
			if len(errMsg) > 0 {
				ctx.JSON(500, gin.H{
					"data": false,
					"msg":  errMsg,
					"code": 500,
				})
				ctx.Abort()
				break
			} else {
				ctx.JSON(200, gin.H{
					"data": query,
					"msg":  "",
					"code": 200,
				})
				ctx.Abort()
				break
			}
		}
	}
}

func addTask(task TaskAdd, errMsgChan chan string) {
	if task.RepeatFlag == 1 {
		// cron
		queue := ShareInstanceQueue()
		fmt.Println(queue.Jobs)

		job := NewTaskJob(task.Id, task.UserId, TASKTYPE_CRON, 0,false,task.RepeatCrontStr)
		_, err := queue.AddJob(job)
		if err != nil {
			errMsgChan <- err.Error()
			return
		}
		errMsgChan <- ""
	}else {
		// 延时一次性任务	TODO
	}
}

func addTaskAndStart(task TaskAdd, errMsgChan chan string) {
	if task.RepeatFlag == 1 {
		// cron
		queue := ShareInstanceQueue()
		fmt.Println(queue.Jobs)
		c := cron.New(cron.WithSeconds())
		entryId, err := c.AddFunc(task.RepeatCrontStr, func() {
			//fmt.Println("正在执行任务: taskId=" + string(task.Id) + "" + "时间表达式为:" + task.RepeatCrontStr + "当前时间:" + time.Now().Format("2006-01-02 15:04:05"))
		})
		c.Start()
		if err != nil {
			errMsgChan <- err.Error()
			return
		}
		job := NewTaskJob(task.Id, task.UserId, TASKTYPE_CRON, entryId,true,task.RepeatCrontStr)
		_, err = queue.AddJob(job)
		if err != nil {
			errMsgChan <- err.Error()
			return
		}
		errMsgChan <- ""
	}else {
		// 延时一次性任务	TODO
	}
}

func HandleTaskAddAndStart(ctx *gin.Context) {
	query := TaskAdd{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(500, gin.H{
			"data": false,
			"msg":  err.Error(),
			"code": 500,
		})
		ctx.Abort()
	} else {
		errMsgChan := make(chan string)
		go addTaskAndStart(query, errMsgChan)
		var errMsg string
		for {
			errMsg = <-errMsgChan
			if len(errMsg) > 0 {
				ctx.JSON(500, gin.H{
					"data": false,
					"msg":  errMsg,
					"code": 500,
				})
				ctx.Abort()
				break
			} else {
				ctx.JSON(200, gin.H{
					"data": query,
					"msg":  "",
					"code": 200,
				})
				ctx.Abort()
				break
			}
		}
	}
}

func HandleTaskStart(ctx *gin.Context) {
	taskIdStr := ctx.Param("id")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		ctx.JSON(500, gin.H{
			"data": false,
			"msg":  err.Error(),
			"code": 500,
		})
		ctx.Abort()
	}
	if j, ok := ShareInstanceQueue().Jobs[taskId]; ok {
		j.IsRunning = true
		c := cron.New(cron.WithSeconds())
		entryId,err := c.AddFunc(j.RepeatCrontStr, func() {
			fmt.Println("开始任务taskid="+string(j.TaskId))
		})
		if err != nil {
			// TODO
		}
		j.EntryId = entryId
		c.Start()

	} else {
		ctx.JSON(500, gin.H{
			"data": false,
			"msg":  "不存在任务，id="+taskIdStr,
			"code": 500,
		})
		ctx.Abort()
	}
}

func HandleTaskStop(ctx *gin.Context) {

}
