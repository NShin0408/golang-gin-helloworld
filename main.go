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
		var books []models.Book

		if err := dbConn.Find(&books).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
			return
		}

		if len(books) == 0 {
			c.String(http.StatusNotFound, "No books found")
			return
		}

		c.HTML(http.StatusOK, "books.html", gin.H{
			"books": books,
		})
	})

	// 推奨リンクを取得するハンドラー
	r.GET("/api/recommend/:id", func(c *gin.Context) {
		bookID := c.Param("id")
		var recommendations []models.Recommend

		if err := dbConn.Where("book_id = ?", bookID).Find(&recommendations).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve recommendations"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"recommendations": recommendations,
		})
	})

	// サーバーを起動
	r.Run() // listen and serve on 0.0.0.0:8080
}
