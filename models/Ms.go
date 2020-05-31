package models

import (
	"log"
	"strings"
)

type Ms struct {
	Bdat      string `json:"bdat"`
	Table     string `json:"table"`
	RowId     int64  `json:"row_id" db:"row_id"`
	Style     int64  `json:"style" db:"style"`
	Name      string `json:"name" db:"name"`
	tablename string
}

func (ms *Ms) GetTableName() string {
	return ms.tablename
}

func (ms *Ms) SetTableName(tablename string) {
	ms.tablename = tablename
	s := strings.Split(tablename, ".")
	if len(s) != 2 {
		log.Fatalln(tablename + " 分割后不为2")
	}
	ms.Bdat = s[0]
	ms.Table = s[1]
}
