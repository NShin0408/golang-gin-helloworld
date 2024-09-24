package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_gin_helloworld/db"
	"golang_gin_helloworld/models"
	"net/http"
)

func main() {
	// データベースの初期化
	dbConn, err := db.InitDB()
	if err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		return
	}

	// Ginのルーター設定
	r := gin.Default()

	// 静的ファイルの提供
	r.Static("/static", "./static")

	// テンプレートファイルの読み込み
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		var profiles []models.Profile

		if err := dbConn.Find(&profiles).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve profiles"})
			return
		}

		if len(profiles) == 0 {
			c.String(http.StatusNotFound, "No profiles found")
			return
		}

		c.HTML(http.StatusOK, "profiles.html", gin.H{
			"profiles": profiles,
		})
	})

	// サーバーを起動
	r.Run() // listen and serve on 0.0.0.0:8080
}
