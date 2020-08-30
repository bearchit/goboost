package structs

import (
	"errors"
	"reflect"

	"github.com/fatih/structtag"
)

type Field struct {
	Name  string
	Value interface{}
}

func FieldByTag(v interface{}, matcher func(tag structtag.Tag) bool) (*Field, error) {
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		vf := val.Field(i)
		tf := val.Type().Field(i)

		tags, err := structtag.Parse(string(tf.Tag))
		if err != nil {
			return nil, err
		}

		for _, tag := range tags.Tags() {
			if matcher(*tag) {
				return &Field{
					Name:  tf.Name,
					Value: vf.Interface(),
				}, nil
			}
		}
	}

	return nil, errors.New("no such field")
}

func FieldByTagKeyNamePairs(v interface{}, pairs []string) (*Field, error) {
	if len(pairs)%2 != 0 {
		return nil, errors.New("invalid key, name pairs")
	}

	return FieldByTag(v, func(tag structtag.Tag) bool {
		for i := 0; i < len(pairs); i += 2 {
			if tag.Key == pairs[i] && tag.Name == pairs[i+1] {
				return true
			}
		}

		return false
	})
}
