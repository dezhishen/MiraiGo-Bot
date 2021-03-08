package migrations

import (
	"github.com/golang-migrate/migrate/v4"

	// file
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
		panic(err)
	}
	m.Steps(2)
}
