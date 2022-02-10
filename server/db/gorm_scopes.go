package db

import (
	"fmt"
	"strings"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"gorm.io/gorm"
)

func Search(v *pb.Search) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if v == nil {
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
	for _, s := range columns {
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
		db = db.Or(fmt.Sprintf("%s %s ?", s, expresion), value)
	}
	if len(columns) > 0 {
		db.Where(db)
	}
	return db
}
