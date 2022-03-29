package filters

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/zsmartex/pkg/repository"
)

func WithSchema(schema, table schema.Tabler) repository.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Table(fmt.Sprintf("%s.%s", schema, table))
	}
}

func WithPreload(tableName string, cond ...interface{}) repository.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Preload(tableName, cond...)
	}
}

func WithPreloadAll() repository.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Preload(clause.Associations)
	}
}

func WithJoin(modelName string, args ...interface{}) repository.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Joins(modelName, args...)
	}
}

func WithSelect(query string, args ...interface{}) repository.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Select(query, args...)
	}
}

func WithOmit(fields ...string) repository.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Omit(fields...)
	}
}