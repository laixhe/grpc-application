package gormx

import "fmt"

// PageOffset 分页
func PageOffset(page, pageSize int) (size int, offset int) {
	return pageSize, (page - 1) * pageSize
}

// PageLimitSql 分页SQL
func PageLimitSql(offset, limit int) string {
	if offset == 0 {
		return fmt.Sprintf("LIMIT %d", limit)
	}
	return fmt.Sprintf("LIMIT %d OFFSET %d", offset, limit)
}
