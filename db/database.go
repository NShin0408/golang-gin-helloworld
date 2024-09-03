package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// .env ファイルをロード
	//err := godotenv.Load(".env")
	//if err != nil {
	//	return nil, fmt.Errorf("エラー: %v", err)
	//}

	// 接続文字列を組み立てる
	dataSourceName := GetDataSourceName()

	// PostgresSQLへの接続
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("データベース接続エラー: %v", err)
	}

	// データベース接続のPingを行う
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("データベース接続エラー: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("データベースPingエラー: %v", err)
	}

	// モデルの一括マイグレーションを実行
	if err := RunMigrations(db); err != nil {
		return nil, fmt.Errorf("マイグレーションエラー: %v", err)
	}

	return db, nil
}

func GetDataSourceName() string {
	//dbUser := os.Getenv("DB_USER")
	//dbPassword := os.Getenv("DB_PASSWORD")
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	//dbName := os.Getenv("DB_NAME")

	dbUser := "postgre_test_m91l_user"
	dbPassword := "HTHvdmPnv6PlhAEfMNVXLTuLhpDNaN79"
	dbHost := "dpg-crbe44bqf0us73dbi1hg-a"
	dbPort := 5432
	dbName := "postgre_test_m91l"

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	return dataSourceName
}
