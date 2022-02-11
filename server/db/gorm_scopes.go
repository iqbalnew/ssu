package db

import (
	"fmt"
	"math"
	"strings"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"gorm.io/gorm"
)

func Paginate(value interface{}, v *pb.PaginationResponse, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if v.Limit > 0 || v.Page > 0 {
		var totalRows int64
		db.Model(value).Count(&totalRows)

		v.TotalRows = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(v.Limit)))
		v.TotalPages = int32(totalPages)
	}

	return func(db *gorm.DB) *gorm.DB {
		if v.Limit < 1 || v.Page < 1 {
			return db
		}

		offset := (v.Page - 1) * v.Limit
		if v != nil {
			return db.Limit(int(v.Limit)).Offset(int(offset))
		}
		return db
	}
}

func Sort(v *pb.Sort) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if v == nil || v.Column == "" {
			return db
		}
		v.Column = columnNameBuilder(v.Column)
		if v != nil {
			return db.Order(v.Column + " " + v.Direction)
		}
		return db
	}
}

func Search(v *pb.Search) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if v == nil || v.Value == "" || v.Columns == "" {
			return db
		}

		columns := strings.Split(v.Columns, ",")
		expression := ""
		value := v.Value

		if len(v.Value) > 2 {
			if string(v.Value[0:1]) == "%" {
				expression = string(v.Value[0:2])
			}
		}

		switch expression {
		case "%%":
			value := "%" + string(v.Value[2:len(v.Value)]) + "%"
			db = searchColumnsLoop(db, columns, "LIKE", value)

		case "%!":
			value := "%" + string(v.Value[2:len(v.Value)]) + "%"
			db = searchColumnsLoop(db, columns, "ILIKE", value)

		case "":
			db = searchColumnsLoop(db, columns, "=", value)
		}

		return db
	}
}

func searchColumnsLoop(db *gorm.DB, columns []string, expresion string, value string) *gorm.DB {
	for j, s := range columns {
		s = columnNameBuilder(s)
		if j == 0 {
			db = db.Where(fmt.Sprintf("%s %s ?", s, expresion), value)
		} else {
			db = db.Or(fmt.Sprintf("%s %s ?", s, expresion), value)
		}
	}
	return db
}

func columnNameBuilder(s string) string {
	if strings.Contains(s, "->") {
		nested := strings.Split(s, "->")
		s = ""
		for i, t := range nested {
			if i == 0 {
				s = fmt.Sprintf("\"%s\"", t)
			} else if i == len(nested)-1 {
				s = s + fmt.Sprintf("->>'%s'", t)
			} else {
				s = s + fmt.Sprintf("->'%s'", t)
			}
		}
	} else {
		s = fmt.Sprintf("\"%s\"", s)
	}

	return s
}
