package main

import (
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type FileServer struct {
	Root string
}

func (f *FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := "coveragenew1.out"

	// 检查文件是否存在
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 设置响应头
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))

	// 将文件内容写入响应
	http.ServeContent(w, r, "", time.Now(), file)
}

func main() {
	fs := &FileServer{
		Root: "/", // 指定文件存储路径
	}

	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}
