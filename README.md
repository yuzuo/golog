
### 日志收集
使用redis生成消费方式
使用tail的方式，用logagent 发送日志到server，server消费存储到mongo

#### 使用方式
配置config.yaml
在需要收集的地方  go run logagent.go -log=./a.log
在服务位置启用 go run server.go 
