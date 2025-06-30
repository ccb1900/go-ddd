package main

import (
	"goddd/internal/domain"
	"goddd/internal/infra/database"
	"goddd/pkg/config"
	"goddd/query"

	"gorm.io/gen"
)

type Test struct {
	ID int `gorm:"type:int;column:PKEY"`
}

func (t Test) TableName() string {
	return "T_LASER_MEASURED_BASE"
}

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE PKEY = @name
	FilterWithNameAndRole(name int) ([]gen.T, error)
	// SELECT * FROM @@table WHERE PKEY = @name
	FilterWithNameAndRole2(name int) (domain.Test, error)
}

func main() {
	config.Init("config")
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(database.NewDB(config.ProvideConfig())) // reuse your gorm db

	g.ApplyBasic(domain.Book{})
	g.ApplyBasic(domain.Test{})

	g.ApplyInterface(func(Querier) {}, domain.Book{})

	// Generate the code
	g.Execute()

	query.Book.FilterWithNameAndRole(1)
}
