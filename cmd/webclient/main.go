package main

import (
	"github.com/dobyte/due/component/http/v2"
	"github.com/dobyte/due/v2"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	// 创建容器
	container := due.NewContainer()

	fiberComponent := http.NewServer(
		http.WithAddr(":8081"),
	)

	router := fiberComponent.Proxy().Router()
	// 设置静态文件目录
	router.Get("/*", static.New("./web_resource/dist"))

	// 设置路由
	router.Get("/hello", func(c fiber.Ctx) error {
		name := c.Query("name", "world")
		c.JSON(map[string]string{
			"message": "hello " + name,
		})
		return nil
	})

	// 添加组件
	container.Add(fiberComponent)

	// 启动容器
	container.Serve()
}
