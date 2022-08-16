package config

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

//读取外部配置，外部配置是项目构建配置等

// 获取项目的公共配置
func GetProjectCommonConfig(projectName string) string {
	ctx := gctx.New()
	value, err := g.Cfg("index").Get(ctx, "imageName")
	if err != nil {
		fmt.Println(err)
	}
	return value.String()
}
