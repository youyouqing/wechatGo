package service

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	. "wechatGin.cthai.cn/model"
	. "wechatGin.cthai.cn/service"
)

func HandleTaskAdd(ctx *gin.Context) {
	query := TaskAdd{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(500, err.Error())
	} else {
		result, errMsg := addTask(query)
		if !result {
			ctx.JSON(500, gin.H{
				"data": false,
				"msg":  errMsg,
				"code": 500,
			})
		}
		ctx.JSON(200, gin.H{
			"data": true,
			"msg":  "",
			"code": 200,
		})
	}
}

func addTask(task TaskAdd) (bool, string) {
	if task.RepeatFlag == 1 {
		// cron
		c := cron.New(cron.WithSeconds())
		entryId, err := c.AddFunc(task.RepeatCrontStr, func() {

		})
		if err != nil {
			return false, err.Error()
		}
		job := NewTaskJob(task.Id,task.UserId,TASKTYPE_CRON,entryId)
		NewJobQueue().AddJob(job)

	}
	return false, ""
}

func HandleTaskAddAndStart(ctx *gin.Context) {

}

func HandleTaskStart(ctx *gin.Context) {

}

func HandleTaskStop(ctx *gin.Context) {

}
