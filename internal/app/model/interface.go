package model

import "database/sql/driver"

type ItemInterface interface {
	Value() (driver.Value, error)
	Scan(value interface{}) error
}
