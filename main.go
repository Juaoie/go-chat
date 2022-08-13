package main

import (
	"fmt"
	"go-chat/internal/docker"
	"os/exec"
)

/**
重启docker
service docker start
*/

func main() {

	//检查并安装docker
	docker.InitDocker()
	//启动docker 服务

	//创建接受gitlab的http服务

	//模拟接受到打包信息
	containerName := "node14"
	name := "juzi.sdxxtop.com"
	ci := "build"
	/***


	RunCommand("docker-compose up -d " + containerName)

	cd := "cd /opt/" + containerName + "/" + name
	makefile := "make " + ci

	RunCommand("docker-compose exec -it " + containerName + " bash -c \" " + cd + " && " + makefile + " \"")
	*/
	// enterIdOutput1 := RunCommand("node -v")
	// fmt.Println(enterIdOutput1)

}

//安装docker-compose
func initDockerCompose() {
	cmd := exec.Command("docker-compose", "-v")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println("docker未安装，准备安装docker-compose")
		//安装docker
		RunCommand("curl -L https://get.daocloud.io/docker/compose/releases/download/v2.9.0/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose")
		fmt.Println("安装docker-compose成功")

		//给 docker-compose 执行权限
		RunCommand("sudo chmod +x /usr/local/bin/docker-compose")
	} else {
		fmt.Println(string(out))
	}
}
