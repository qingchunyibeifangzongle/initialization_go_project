# initialization_go_project
initialization go project



### 热重启(热加载)

##### 1.安装`air`工具
```
go get -u github.com/cosmtrek/air
```
##### 2.在 $GOPATH/bin 目录下生成air命令
##### 3.在项目根目录创建一个名为 .air.conf 的配置文件。具体详细配置可参考 [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air) 下的 `air_example.toml` 
创建完毕之后，在文件中写入你应用运行的命令如
```markdown
go build main.go
```
##### 4.运行 `air` 
```markdown

      __    _   ___  
     / /\  | | | |_) 
    /_/--\ |_| |_| \_ , built with Go 
    
    watching .
    !exclude tmp
    building...
    running...
    hello world //main.go 里面 fmt.Println("hello world")

```
