# AcFun Comments Observer Backend
A站评论观察者后端，尽可能完整收录A站文章区下所有评论，无外部访问功能，需搭配独立WEB端进行查询数据。

WEB端Repo请访问这里：[LINK](https://github.com/WangHongshuo/acfun_comments_observer_web)

## 在线预览

http://47.100.72.81:5000 （丐中丐配置，勿压测）

## 依赖

- 数据库
- [proxy_pool](https://github.com/jhao104/proxy_pool)（可选）

## 使用组件

- gorm
- protoactor-go
- zap
- vipper

## Actor结构

```
Main
├── HttpServer(ToDo)
├── obctrl (ObserverController)
    ├── articleslistob (ArticlesListObserver)
        ├── commentsob (CommentsObserver)
```

## 业务流程

### 启动流程

```mermaid
sequenceDiagram
note over Main: 初始全局化配置
Main ->> obctrl: spawn（创建新actor）
obctrl ->> articleslistob: spawn（创建新actor，可以创建多个）
note over obctrl: 初始化全局db
obctrl ->> articleslistob: ResourceReadyMsg
articleslistob ->> commentsob: spwan（创建新actor，可以创建多个）
articleslistob ->> commentsob: ResourceReadyMsg
note over commentsob: 初始化db资源
commentsob ->> articleslistob: CommentsObReadyMsg
note over articleslistob: 等待收到所有commentsob的CommentsObReadyMsg
articleslistob ->> obctrl: ArticlesListObReadyMsg
note over obctrl,commentsob: 根据配置开始观测
```

### 运行流程

```mermaid
sequenceDiagram
obctrl ->> articleslistob_x: ObserveArticlesListTaskMsg（需要观测的文章区列表）
articleslistob_x -->> proxy pool: 获取https代理
proxy pool -->> articleslistob_x: 返回https代理
articleslistob_x ->> Internet: 获取文章区列表的aid（Article ID）
Internet ->> articleslistob_x: 获取文章区列表的aid（Article ID）
articleslistob_x ->> commentsob_x: ObserveCommentsTaskMsg（Article ID List）
commentsob_x -->> proxy pool: 获取https代理
proxy pool -->> commentsob_x: 返回https代理
loop Article ID List不为空
	commentsob_x ->> Internet: 获取comments
	Internet ->> commentsob_x: comments
	note over commentsob_x: 保存db
end
commentsob_x ->> articleslistob_x: ObserveCommentsTaskFinishedMsg消息
note over articleslistob_x: 等待收到所有commentsob的ObserveCommentsTaskFinishedMsg消息
note over articleslistob_x, Internet: 开始下一轮观测

```

### 异常流程

```mermaid
sequenceDiagram
obctrl ->> articleslistob_x: ObserveArticlesListTaskMsg（需要观测的文章区列表）
articleslistob_x -->> proxy pool: 获取https代理
proxy pool -->> articleslistob_x: 返回https代理
articleslistob_x ->> Internet: 获取文章区列表的aid（Article ID）
Internet -x articleslistob_x: 获取文章区列表的aid（Article ID）
articleslistob_x -->> proxy pool: 获取https代理（尝试更换代理）
proxy pool -->> articleslistob_x: 返回https代理
articleslistob_x ->> Internet: 获取文章区列表的aid（Article ID）
Internet ->> articleslistob_x: 获取文章区列表的aid（Article ID）
```

