package docker

import (
	"context"
	"fmt"
	"go-docker-cli/internal/config"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var ctx = context.Background()
var cli, _ = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

// 运行docker，打包构建单个项目

// 一个项目的一个环境对应一个容器

// containerName 需要使用到的docker容器名称；
// projectName		进行打包构建的项目名称；
// cliEnvType		执行cli工具的环境类型；
func RunDockerContainer(containerName string, projectName string, cliEnvType string) bool {
	//获取项目配置
	//判断容器是否正在执行
	containerList := GetRunDockerContainers()
	// var container types.Container
	container := types.Container{}
	//容器命名规范
	name := containerName + `/` + "projectName" + "/" + cliEnvType
	//docker名称系统会默认设置一个不重复的名称，例如：/nice_hugle
	//docker名称为数组字符串

	for _, containerItem := range containerList {
		if containerItem.Names[0] == name {
			container = containerItem
			break
		}
	}
	if container.ID != "" {
		//容器正在运行，可以添加日志
		return false
	}

	//拉取镜像
	fmt.Println(config.GetProjectCommonConfig(projectName))
	reader, err := cli.ImagePull(ctx, containerName, types.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
	}
	defer reader.Close()
	return true
}

//创建docker容器
func CreateDockerContainer(containerName string, projectName string, cliEnvType string) {

}

//获取所有正在运行的容器，包括前台执行和后台执行的容器
func GetRunDockerContainers() []types.Container {
	containerList, _ := cli.ContainerList(ctx, types.ContainerListOptions{})

	return containerList
}
