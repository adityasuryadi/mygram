package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func NewValidation(db *gorm.DB) *validator.Validate {
	validate := validator.New()
	
	validate.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		// fmt.Println(fl.StructFieldName())
		// // get parameter dari tag struct validate
		table:=fl.Param()
		field:= strings.ToLower(fl.StructFieldName())
		// // get nama tagsturct fl.Field()
		// // get value tagsturct fl.Field()
		var total int64
		err := db.Table(table).Where(""+field+" = ?",fl.Field()).Count(&total).Error
		if err != nil {
			fmt.Println(err)
		}
		// // Return true if the count is zero (i.e., the value is unique)
		return total == 0
	})
	return validate
}

