package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

/**
重启docker
service docker start
*/

func main() {
	name := "test"
	containerName := "node14"
	initDocker()
	initDockerCompose()
	//启动docker 服务
	RunCommand("docker-compose up -d " + containerName)

	enterIdOutput := RunCommand("docker-compose exec " + containerName + " bash " + "sh -c \"cd /opt/" + containerName + "/" + name + "/ && ls\"")

	fmt.Println(enterIdOutput)
	enterIdOutput1 := RunCommand("node -v")
	fmt.Println(enterIdOutput1)

}

//安装docker-compose
func initDockerCompose() {
	cmd := exec.Command("docker-compose", "-v")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println("docker未安装，准备安装docker-compose")
		//安装docker
		res := RunCommand("curl -L https://get.daocloud.io/docker/compose/releases/download/v2.9.0/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose")
		fmt.Println("安装docker-compose成功")
		fmt.Println(res)
		RunCommand("sudo chmod +x /usr/local/bin/docker-compose")
		out, _ := cmd.CombinedOutput()
		fmt.Println(string(out))
	} else {
		fmt.Println(string(out))
	}
}

//安装docker
func initDocker() {
	cmd := exec.Command("docker", "-v")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println("docker未安装，准备安装docker")
		//安装docker
		res := RunCommand("curl -sSL https://get.daocloud.io/docker | sh")
		fmt.Println("安装docker成功")
		fmt.Println(res)
		out, _ := cmd.CombinedOutput()
		fmt.Println(string(out))
	} else {
		fmt.Println(string(out))
	}

}

func RunCommand(cmd string) string {
	if runtime.GOOS == "windows" {
		return runInWindows(cmd)
	} else {
		return runInLinux(cmd)
	}
}

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
	fmt.Println(">>>开始实时打印到终端>>>")
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}
	fmt.Println("<<<结束<<<")

	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
