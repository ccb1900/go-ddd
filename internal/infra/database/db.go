package database

// import (

// 	"gorm.io/gorm"
// )

// func NewDB(cfg config.AppConfig) *gorm.DB {
// 	return &gorm.DB{}
// }

import (
	"goddd/pkg/config"

	"fmt"
	"sync"
	"time"

	oracle "github.com/godoes/gorm-oracle"
	// oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbOnce sync.Once
var db *gorm.DB

func NewDB(cfg config.AppConfig) *gorm.DB {
	dbOnce.Do(func() {
		options := map[string]string{
			"CONNECTION TIMEOUT": "90",
			"LANGUAGE":           "SIMPLIFIED CHINESE",
			"TERRITORY":          "CHINA",
			"SSL":                "false",
		}
		// oracle://user:password@127.0.0.1:1521/service
		c := cfg.DB
		dbConfig, ok := c.Databases[c.Default]

		if !ok {
			panic(fmt.Sprintf("default db(%s) not found", c.Default))
		}
		url := oracle.BuildUrl(dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.Username, dbConfig.Password, options)
		dialector := oracle.New(oracle.Config{
			DSN:                     url,
			IgnoreCase:              false, // query conditions are not case-sensitive
			NamingCaseSensitive:     true,  // whether naming is case-sensitive
			VarcharSizeIsCharLength: true,  // whether VARCHAR type size is character length, defaulting to byte length

			// RowNumberAliasForOracle11 is the alias for ROW_NUMBER() in Oracle 11g, defaulting to ROW_NUM
			RowNumberAliasForOracle11: "ROW_NUM",
		})
		// dialector := postgres.Open("user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable")
		var err error
		db, err = gorm.Open(dialector, &gorm.Config{
			// Logger: slogGorm.New(
			// 	slogGorm.WithHandler(Log().Handler()),                         // since v1.3.0
			// 	slogGorm.WithTraceAll(),                                       // trace all messages
			// 	slogGorm.SetLogLevel(slogGorm.DefaultLogType, slog.Level(32)), // Define the default logging level
			// ),
			SkipDefaultTransaction:                   true, // 是否禁用默认在事务中执行单次创建、更新、删除操作
			DisableForeignKeyConstraintWhenMigrating: true, // 是否禁止在自动迁移或创建表时自动创建外键约束
			// 自定义命名策略
			NamingStrategy: schema.NamingStrategy{
				NoLowerCase:         true, // 是否不自动转换小写表名
				IdentifierMaxLength: 30,   // Oracle: 30, PostgreSQL:63, MySQL: 64, SQL Server、SQLite、DM: 128
			},
			PrepareStmt:     false, // 创建并缓存预编译语句，启用后可能会报 ORA-01002 错误
			CreateBatchSize: 50,    // 插入数据默认批处理大小
		})
		if err != nil {
			// panic error or log error info
			panic(err)
		}

		// "TIME_ZONE":               "+08:00",                       // ALTER SESSION SET TIME_ZONE = '+08:00';
		// "NLS_DATE_FORMAT":         "DD-MON-RR",                    // ALTER SESSION SET NLS_DATE_FORMAT = 'YYYY-MM-DD';
		// "NLS_TIME_FORMAT":         "HH.MI.SSXFF AM",               // ALTER SESSION SET NLS_TIME_FORMAT = 'HH24:MI:SS.FF3';
		// "NLS_TIMESTAMP_FORMAT":    "DD-MON-RR HH.MI.SSXFF AM",     // ALTER SESSION SET NLS_TIMESTAMP_FORMAT = 'YYYY-MM-DD HH24:MI:SS.FF3';
		// "NLS_TIME_TZ_FORMAT":      "HH.MI.SSXFF AM TZR",           // ALTER SESSION SET NLS_TIME_TZ_FORMAT = 'HH24:MI:SS.FF3 TZR';
		// "NLS_TIMESTAMP_TZ_FORMAT": "DD-MON-RR HH.MI.SSXFF AM TZR", // ALTER SESSION SET NLS_TIMESTAMP_TZ_FORMAT = 'YYYY-MM-DD HH24:MI:SS.FF3 TZR';
		// set session parameters
		// if sqlDB, err := db.DB(); err == nil {
		// 	_, _ = oracle.AddSessionParams(sqlDB, map[string]string{
		// 		"TIME_ZONE":               "+08:00",                       // ALTER SESSION SET TIME_ZONE = '+08:00';
		// 		"NLS_DATE_FORMAT":         "YYYY-MM-DD",                   // ALTER SESSION SET NLS_DATE_FORMAT = 'YYYY-MM-DD';
		// 		"NLS_TIME_FORMAT":         "HH24:MI:SSXFF",                // ALTER SESSION SET NLS_TIME_FORMAT = 'HH24:MI:SS.FF3';
		// 		"NLS_TIMESTAMP_FORMAT":    "YYYY-MM-DD HH24:MI:SSXFF",     // ALTER SESSION SET NLS_TIMESTAMP_FORMAT = 'YYYY-MM-DD HH24:MI:SS.FF3';
		// 		"NLS_TIME_TZ_FORMAT":      "HH24:MI:SS.FF TZR",            // ALTER SESSION SET NLS_TIME_TZ_FORMAT = 'HH24:MI:SS.FF3 TZR';
		// 		"NLS_TIMESTAMP_TZ_FORMAT": "YYYY-MM-DD HH24:MI:SSXFF TZR", // ALTER SESSION SET NLS_TIMESTAMP_TZ_FORMAT = 'YYYY-MM-DD HH24:MI:SS.FF3 TZR';
		// 	})
		// }
		sdb, err := db.DB()
		if err != nil {
			panic(err)
		}
		_, _ = oracle.AddSessionParams(sdb, map[string]string{
			"TIME_ZONE":               "+08:00",                       // ALTER SESSION SET TIME_ZONE = '+08:00';
			"NLS_DATE_FORMAT":         "YYYY-MM-DD",                   // ALTER SESSION SET NLS_DATE_FORMAT = 'YYYY-MM-DD';
			"NLS_TIME_FORMAT":         "HH24:MI:SSXFF",                // ALTER SESSION SET NLS_TIME_FORMAT = 'HH24:MI:SS.FF3';
			"NLS_TIMESTAMP_FORMAT":    "YYYY-MM-DD HH24:MI:SSXFF",     // ALTER SESSION SET NLS_TIMESTAMP_FORMAT = 'YYYY-MM-DD HH24:MI:SS.FF3';
			"NLS_TIME_TZ_FORMAT":      "HH24:MI:SS.FF TZR",            // ALTER SESSION SET NLS_TIME_TZ_FORMAT = 'HH24:MI:SS.FF3 TZR';
			"NLS_TIMESTAMP_TZ_FORMAT": "YYYY-MM-DD HH24:MI:SSXFF TZR", // ALTER SESSION SET NLS_TIMESTAMP_TZ_FORMAT = 'YYYY-MM-DD HH24:MI:SS.FF3 TZR';
		})
		sdb.SetConnMaxIdleTime(time.Duration(c.DbPoolConfig.MaxIdleTime) * time.Hour)
		sdb.SetConnMaxLifetime(time.Duration(c.DbPoolConfig.MaxLifeTime) * time.Hour)
		sdb.SetMaxIdleConns(c.DbPoolConfig.MaxIdle)
		sdb.SetMaxOpenConns(c.DbPoolConfig.MaxOpen)
		if cfg.Debug {
			db = db.Debug()
		}
	})

	return db
}
