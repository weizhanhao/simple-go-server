package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main4() {
	// 获取当前分支和目标分支
	currentBranch, err := getCurrentBranch()
	if err != nil {
		fmt.Println("Failed to get current branch:", err)
		os.Exit(1)
	}
	targetBranch := "master" // 设置目标分支

	// 获取当前分支和目标分支之间的代码差异
	diffCmd := exec.Command("git", "diff", "--name-only", targetBranch)
	diffOutput, err := diffCmd.Output()
	fmt.Println(diffOutput)
	if err != nil {
		fmt.Println("Failed to get diff:", err)
		os.Exit(1)
	}

	// 运行当前分支的测试并生成测试覆盖率文件
	err = runTestsAndGenerateCoverage(currentBranch)
	if err != nil {
		fmt.Println("Failed to generate coverage for current branch:", err)
		os.Exit(1)
	}

	// 运行目标分支的测试并生成测试覆盖率文件
	err = runTestsAndGenerateCoverage(targetBranch)
	if err != nil {
		fmt.Println("Failed to generate coverage for target branch:", err)
		os.Exit(1)
	}

	// 对比两个覆盖率文件
	err = compareCoverageFiles(currentBranch, targetBranch)
	if err != nil {
		fmt.Println("Failed to compare coverage files:", err)
		os.Exit(1)
	}

	fmt.Println("Coverage comparison generated.")
}

func getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	branch := strings.TrimSpace(string(output))
	return branch, nil
}

func runTestsAndGenerateCoverage(branch string) error {
	// 运行测试并生成测试覆盖率文件
	cmd := exec.Command("go", "test", "-coverprofile=coverage_"+branch+".out")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func compareCoverageFiles(currentBranch, targetBranch string) error {
	// 读取当前分支和目标分支的覆盖率文件内容
	currentCoverage, err := ioutil.ReadFile("coverage_" + currentBranch + ".out")
	fmt.Println(currentCoverage)
	if err != nil {
		return err
	}
	targetCoverage, err := ioutil.ReadFile("coverage_" + targetBranch + ".out")
	fmt.Println(targetCoverage)
	if err != nil {
		return err
	}

	// 进行覆盖率对比
	// TODO: 在这里执行自定义的覆盖率对比逻辑

	return nil
}

func compareCoverage(coverage1, coverage2 []int) {
	// 创建一个 map 来存储覆盖到的点
	coveredPoints := make(map[int]bool)

	// 遍历 coverage1，将覆盖到的点添加到 map 中
	for _, point := range coverage1 {
		coveredPoints[point] = true
	}

	// 创建两个 map 来存储相同覆盖和不同覆盖的点
	sameCoveragePoints := make(map[int]bool)
	differentCoveragePoints := make(map[int]bool)

	// 遍历 coverage2，检查每个点是否在 coverage1 中也存在
	for _, point := range coverage2 {
		if _, exists := coveredPoints[point]; exists {
			sameCoveragePoints[point] = true
		} else {
			differentCoveragePoints[point] = true
		}
	}

	// 输出结果
	fmt.Println("相同覆盖的点：", sameCoveragePoints)
	fmt.Println("不同覆盖的点：", differentCoveragePoints)
}

func main3() {
	// 示例数据
	coverage1 := []int{1, 2, 3, 4, 5}
	coverage2 := []int{3, 4, 5, 6, 7}

	// 执行对比逻辑
	compareCoverage(coverage1, coverage2)
}

//sudo curl -s https://api.github.com/repos/qiniu/goc/releases/latest | grep "browser_download_url.*-darwin-amd64.tar.gz" | cut -d : -f 2,3 | tr -d \" | xargs -n 1 curl -L | tar -zx && chmod +x goc && mv goc /usr/local/bin
//
