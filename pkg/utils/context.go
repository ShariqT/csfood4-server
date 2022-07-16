package utils

import (
	"github.com/ShariqT/csfood4/pkg/core"
	"github.com/labstack/echo"
)

type CommonstockContext struct {
	echo.Context
}

func (c CommonstockContext) DB() (core.DB, error) {
	db, err := NewDB()
	return db, err
}
