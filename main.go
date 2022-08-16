package main

import (
	"go-docker-cli/internal/docker"
)

/**
基于docker的自动化构建工具
*/
func main() {

	//检查并安装docker
	// docker.InitDocker()
	//启动docker 服务

	//创建接受gitlab的http服务

	//模拟接受到打包信息
	containerName := "node14"
	projectName := "juzi.sdxxtop.com"
	cliEnvType := "build"
	docker.RunDockerContainer(containerName, projectName, cliEnvType)
	/***


	RunCommand("docker-compose up -d " + containerName)

	cd := "cd /opt/" + containerName + "/" + name
	makefile := "make " + ci

	RunCommand("docker-compose exec -it " + containerName + " bash -c \" " + cd + " && " + makefile + " \"")
	*/
	// enterIdOutput1 := RunCommand("node -v")
	// fmt.Println(enterIdOutput1)

}
