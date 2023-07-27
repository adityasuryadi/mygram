package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// contract
type Validation interface {
	ValidateRequest(request interface{}) interface{}
}

// func NewValidation(db *gorm.DB) *validator.Validate {
// 	validate := validator.New()

// 	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
// 		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
// 		// skip if tag key says it should be ignored
// 		if name == "-" {
// 			return ""
// 		}
// 		return name
// 	})

// 	validate.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
// 		// fmt.Println(fl.StructFieldName())
// 		// // get parameter dari tag struct validate
// 		table:=fl.Param()
// 		field:= strings.ToLower(fl.StructFieldName())
// 		// // get nama tagsturct fl.Field()
// 		// // get value tagsturct fl.Field()
// 		var total int64
// 		err := db.Table(table).Where(""+field+" = ?",fl.Field()).Count(&total).Error
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		// // Return true if the count is zero (i.e., the value is unique)
// 		return total == 0
// 	})
// 	return validate
// }

func NewValidation(db *gorm.DB) Validation {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

	validate.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		// fmt.Println(fl.StructFieldName())
		// // get parameter dari tag struct validate
		table := fl.Param()
		field := strings.ToLower(fl.StructFieldName())
		// // get nama tagsturct fl.Field()
		// // get value tagsturct fl.Field()
		var total int64
		err := db.Table(table).Where(""+field+" = ?", fl.Field()).Count(&total).Error
		if err != nil {
			fmt.Println(err)
		}
		// // Return true if the count is zero (i.e., the value is unique)
		return total == 0
	})

	// custom validasi untuk cek image
	validate.RegisterValidation("image_validation", func(fl validator.FieldLevel) bool {
		mimeTypes := []string{
			"jpg",
			"jpeg",
			"gif",
			"png",
		}

		value := fl.Field()
		var isValid bool
		for _, v := range mimeTypes {
			isValid = strings.ToLower(value.String()) == strings.ToLower("."+v)
			if isValid {
				return isValid
			}
		}
		return isValid
	})

	return &ValidationImpl{
		Validate: validate,
	}
}

type ValidationImpl struct {
	Validate *validator.Validate
}

// ValidateRequest implements Validation
func (validateImpl *ValidationImpl) ValidateRequest(request interface{}) interface{} {
	err := validateImpl.Validate.Struct(request)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		out := make([]ErrorMessage, len(validationErrors))
		for i, fieldError := range validationErrors {
			out[i] = ErrorMessage{
				fieldError.Field(),
				GetErrorMsg(fieldError),
				getGroup(fieldError.Namespace()),
			}
		}
		return out
	}
	return nil
}

func getGroup(nameSpace string) string {
	split := strings.Split(nameSpace, ".")
	var group string
	lenGroup := len(split) - 1
	group = ""
	if lenGroup > 1 {
		group = split[lenGroup-1]
	}
	return group
}
