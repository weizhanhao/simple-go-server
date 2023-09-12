package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main222() {
	//解析覆盖率文件
	coverCmd := exec.Command("go", "tool", "cover", "-func=coveragenew1.out")
	coverOutput, err := coverCmd.Output()
	if err != nil {
		fmt.Println("解析覆盖率文件时出错:", err)
		os.Exit(1)
	}
	//输出覆盖率参数
	fmt.Println(string(coverOutput))

	// 编译并运行测试代码
	//cmd := exec.Command("go", "test", "-coverprofile=coveragenew.out")
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("运行测试代码时出错:", err)
	//	os.Exit(1)
	//}
	//
	//// 打开覆盖率文件
	//file, err := os.Open("coveragenew.out")
	//if err != nil {
	//	fmt.Println("打开覆盖率文件时出错:", err)
	//	os.Exit(1)
	//}
	//defer file.Close()
	//
	//// 统计代码行数
	//coveredLineCount := 0
	//totalLineCount := 0
	//coveredLines := make(map[string]bool)
	//for {
	//	var fileName string
	//	var lineNum int
	//	_, err := fmt.Fscanf(file, "%s %d\n", &fileName, &lineNum)
	//	if err != nil {
	//		break
	//	}
	//
	//	key := fmt.Sprintf("%s:%d", fileName, lineNum)
	//	if !coveredLines[key] {
	//		coveredLines[key] = true
	//		coveredLineCount++
	//	}
	//	totalLineCount++
	//}
	//
	//fmt.Println("覆盖的代码行数:", coveredLineCount)
	//fmt.Println("总代码行数:", totalLineCount)

}
