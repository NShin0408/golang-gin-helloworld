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

	// ランダムな国名を取得するエンドポイント
	r.GET("/", func(c *gin.Context) {
		var countries []models.Country

		// 国名を全て取得
		if err := dbConn.Find(&countries).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve countries"})
			return
		}

		// ランダムな国名を選択
		if len(countries) == 0 {
			c.String(http.StatusNotFound, "No countries found")
			return
		}

		rand.Seed(time.Now().UnixNano())
		randomCountry := countries[rand.Intn(len(countries))]

		// 国名をレスポンス
		c.String(http.StatusOK, randomCountry.Name)
	})

	// サーバーを起動
	r.Run() // listen and serve on 0.0.0.0:8080
}
