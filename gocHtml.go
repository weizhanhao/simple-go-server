package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main5() {
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// 定义输入和输出文件名
	inFile := filepath.Join(wd, "coveragenew.out")
	outFile := filepath.Join(wd, "coveragenew.html")

	// 执行转换命令
	cmd := exec.Command("go", "tool", "cover", "-html="+inFile, "-o="+outFile)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("转换完成！已生成文件：%s\n", outFile)
}
