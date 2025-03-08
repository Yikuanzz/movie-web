package log

import (
	"io"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	/*
		将数据持久化写入到日志文件中
	*/
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

func Run(destination string) {
	/*
		设置日志的输出格式并启动日志
	*/
	log = stlog.New(fileLog(destination), "[GO] ", stlog.LstdFlags)
	log.Println("Starting the application...")
}

func RegisterHandler() {
	/*
		注册 log 服务到 http 服务中

		POST /log 读取上传的内容写入到日志中
	*/
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := io.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func write(msg string) {
	/*
		处理 http 请求中的日志逻辑
	*/
	log.Printf("%v\n", msg)
}
