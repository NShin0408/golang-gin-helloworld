package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_gin_helloworld/db"
	"golang_gin_helloworld/models"
	"math/rand"
	"net/http"
	"time"
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
	r.Static("/static/css", "./static/css")
	r.Static("/static/images", "./static/images")

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

		rand.Seed(time.Now().UnixNano())
		randomProfile := profiles[rand.Intn(len(profiles))]

		imageURL := fmt.Sprintf("/static/images/%s", randomProfile.ImageFilename)

		c.HTML(http.StatusOK, "profile.html", gin.H{
			"name":   randomProfile.Name,
			"detail": randomProfile.Detail,
			"image":  imageURL,
		})
	})

	// サーバーを起動
	r.Run() // listen and serve on 0.0.0.0:8080
}
