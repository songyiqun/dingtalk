### **请求示例（go）**：

#### 一、获取access_token    

```go
//获取access_token
token := OapiGettoken()
```

#### 二、待办任务

#####     1、**发起待办** 

```go
//发起待办
OapiWorkrecordAdd(token)
```

#####      2、**更新待办** 

```go
//更新待办
OapiWorkrecordUpdate(token)
```

#####     3、获取用户待办事项

```go
//获取用户待办事项
OapiWorkrecordGetbyuseridUpdate(token)
```



#### 二、消息通知

##### 		1、工作通知消息

###### 				1.1发送工作通知消息

```go
//发送工作消息
v2 := Test.OapiMessageCorpconversationAsyncsendV2(token)
```

###### 				1.2、查询工作通知消息的发送进度

```go
//工作消息发送进度
Test.OapiMessageCorpconversationGetsendprogress(token, v2.TaskId)
```

###### 				1.3、查询工作通知消息的发送结果

```go
//工作消息发送结果
Test.OapiMessageCorpconversationGetsendresult(token, v2.TaskId)
```

###### 				1.4、工作通知消息撤回

```go
//工作消息撤回
Test.OapiMessageCorpconversationRecall(token, v2.TaskId)
```