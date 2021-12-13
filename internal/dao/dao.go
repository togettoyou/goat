package dao

import (
	"time"

	"goat-layout/internal/model"
	"goat-layout/pkg/conf"
	"goat-layout/pkg/log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// NewMock 模拟
func NewMock() (*model.Store, error) {
	return &model.Store{
		Book: newBook2(),
	}, nil
}

// NewSqlite 使用 Sqlite 数据库
func NewSqlite() (*model.Store, error) {
	db, err := gorm.Open(
		sqlite.Open("./data.db"), // 数据库文件存放路径
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "sys_", // 表名前缀，`User` 的表名应该是 `sys_users`
				SingularTable: true,   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `sys_user`
			},
			Logger: log.NewGormLogger(log.New("gorm").L()),
		})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		/*数据库实体模型*/
		&model.Book{},
	)
	if err != nil {
		return nil, err
	}

	return &model.Store{
		Book: newBook1(db),
	}, nil
}

// NewMysql 使用 MySQL 数据库
func NewMysql() (*model.Store, error) {
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN:                       conf.Mysql.Dsn,
			DefaultStringSize:         191,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "sys_", // 表名前缀，`User` 的表名应该是 `sys_users`
				SingularTable: true,   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `sys_user`
			},
			Logger: log.NewGormLogger(log.New("gorm").L()),
		})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(conf.Mysql.MaxIdle)
	sqlDB.SetMaxOpenConns(conf.Mysql.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(conf.Mysql.MaxLifetime) * time.Minute)
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		AutoMigrate(
			/*数据库实体模型*/
			&model.Book{},
		)
	if err != nil {
		return nil, err
	}

	return &model.Store{
		Book: newBook1(db),
	}, nil
}
