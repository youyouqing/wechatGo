##独立的定时任务管理系统
对外提供http接口服务，使其业务耦合性尽量低


###字段解析
```
type TaskAdd struct {
    Id             int    `form:"id" json:"id"`  
	DeleteFlag     int    `form:"delete_flag" json:"delete_flag"`
	RawCreateTime  string `form:"raw_create_time" json:"raw_create_time"`
	KnockTime      string `form:"knock_time" json:"knock_time"`
	UserId         int    `form:"user_id" json:"user_id"`
	NotifyTitle    string `form:"notify_title" json:"notify_title"`
	NotifyContent  string `form:"notify_content" json:"notify_content"`
	NotifyLevel    int    `form:"notify_level" json:"notify_level"`
	RepeatFlag     int    `form:"repeat_flag" json:"repeat_flag"`
	RepeatCrontStr string `form:"repeat_cront_str" json:"repeat_cront_str"`
}

```
###HTTP API 案例
##### １、新增任务（定时任务）
```
POST http://localhost:8080/task/add
Content-Type: application/x-www-form-urlencoded

id=1&user_id=10&repeat_cront_str=* * * * * *&repeat_flag=1
```

##### ２、新增任务（一次性任务）
```
POST http://localhost:8080/task/add
Content-Type: application/x-www-form-urlencoded

id=1&user_id=10&knock_time=2020-02-15 22:52:00&repeat_flag=０
```

##### 3、开启任务
```
GET http://localhost:8080/task/start/{$taskId}
Accept: application/json

```

##### 4、暂停任务
```
GET http://localhost:8080/task/stop/{$taskId}
Accept: application/json

```


##参与公众号

```javascript
  
```