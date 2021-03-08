package migrations

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/prometheus/common/log"

	// file
	_ "github.com/golang-migrate/migrate/database/sqlite3"
	// sqlite3
	_ "github.com/golang-migrate/migrate/v4/source/file"

	// 数据库
	_ "github.com/mattn/go-sqlite3"
)

// Sync 同步数据库
// file 文件路径
func Sync(file string) {
	m, err := migrate.New(
		file,
		"sqlite3://mirai/data.sqlite?query")
	if err != nil {
		log.Error(err)
		panic(err)
	}
	m.Steps(2)
}
