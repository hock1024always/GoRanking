package main

import "Ranking/router"

// 入口文件
func main() {
	r := router.Router()
	r.Run(":9999")
}
