package db

import (
	"fmt"
	"strings"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"gorm.io/gorm"
)

func Paginate(pagination *pb.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pagination != nil {
			return db.Limit(int(pagination.Limit)).Offset(int(pagination.Offset))
		}
		return db
	}
}

func Sort(sort *pb.Sort) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if sort != nil {
			return db.Order(sort.Column + " " + sort.Direction)
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
		if j == 0 {
			db = db.Where(fmt.Sprintf("%s %s ?", s, expresion), value)
		} else {
			db = db.Or(fmt.Sprintf("%s %s ?", s, expresion), value)
		}
	}
	return db
}
