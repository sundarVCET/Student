package validation

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

var val *validator.Validate

func Init() {
	val = validator.New()
}

func Validate(obj interface{}) error {
	fmt.Println("Entering into validate")
	err := val.Struct(obj)
	if err == nil {
		return nil
	}
	structName := getType(obj)
	errs := err.(validator.ValidationErrors)
	message := strings.ReplaceAll(errs[0].Namespace(), structName+".", "") + " is invalid or missing"
	fmt.Println("MESSAGE", message)
	return fmt.Errorf(fmt.Sprintf("code:%s;message:%s", "BAD_INPUT", message))
}

func Valid(v interface{}) error {
	err := val.Struct(v)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func PartialValid(v interface{}, field string) error {
	err := val.StructPartial(v, field)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ExceptValid(v interface{}, field string) error {
	err := val.StructExcept(v, field)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
