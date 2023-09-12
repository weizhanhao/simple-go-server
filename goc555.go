package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main555() {
	// 设置要拉取的项目地址
	repoURL := "https://github.com/CarlJi/simple-go-server.git"

	// 获取URL的最后一个斜杠后面的部分
	lastSlashIndex := strings.LastIndex(repoURL, "/")
	projectNameWithGit := repoURL[lastSlashIndex+1:] // 包括.git后缀

	// 去掉.git后缀
	projectName := strings.TrimSuffix(projectNameWithGit, ".git")

	fmt.Println(projectName)

	// 设置要保存的本地文件夹路径
	localPath := "/Users/klook/GolandProjects/simple-go-server/test" + "/" + projectName

	err2 := os.Mkdir(localPath, 0755)
	//err := os.Mkdir("myfolder", 0755)
	if err2 != nil {
		fmt.Println("Failed to create folder:", err2)
		return
	}
	fmt.Println("Folder created successfully")
	// 设置要切换的commit id
	commitID := "2bc1a281"

	// 创建一个执行git命令的命令对象
	cmd := exec.Command("git", "clone", repoURL, localPath)

	// 设置命令的输出和错误输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to clone repository:", err)
		return
	}

	// 切换到指定的commit id
	cmd = exec.Command("git", "checkout", commitID)
	cmd.Dir = localPath

	// 设置命令的输出和错误输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行命令
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to switch to commit:", err)
		return
	}

	fmt.Println("Repository cloned and switched to commit successfully!")
}
