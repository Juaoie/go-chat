package docker

import (
	"fmt"
	"go-chat/internal/config"
	"os/exec"
	"runtime"
)

//检查服务器是否安装了docker，没安装则安装
func InitDocker() {
	cmd := exec.Command("docker", "-v")
	out, err := cmd.CombinedOutput()
	if err != nil {
		//安装docker
		fmt.Println("docker未安装，准备安装docker")
		runCommand(`curl -sSL ` + config.DockerUrl + ` | sh`)
		fmt.Println("安装docker成功")

		//重启docker解决没有daemon.json的问题
		runCommand("service docker start")
	} else {
		//输出docker版本号
		fmt.Println(string(out))
	}

}

func runCommand(cmd string) string {
	if runtime.GOOS == "windows" {
		return runInWindows(cmd)
	} else {
		return runInLinux(cmd)
	}
}

//执行linux 命令
func runInLinux(cmd string) string {
	fmt.Println("Running Linux cmd ======> :", cmd)
	result := exec.Command("/bin/sh", "-c", cmd)
	printCommand(result)
	d, _ := result.Output()
	return string(d)
}

func runInWindows(cmd string) string {
	//TODO:

	return "待完成"

}

//日志持续输出
func printCommand(cmd *exec.Cmd) error {
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	// 从管道中实时获取输出并打印到终端
	fmt.Println(">>>>>>开始实时打印到终端")
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}
	fmt.Println("结束<<<<<<")

	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
