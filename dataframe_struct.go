package pandas

import (
	"fmt"
	"reflect"
	"strings"

	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/num"
)

var (
	DefaultTagName = api.DefaultTagName
)

// LoadStructs creates a new DataFrame from arbitrary struct slices.
//
// LoadStructs will ignore unexported fields inside an struct. Note also that
// unless otherwise specified the column names will correspond with the name of
// the field.
//
// You can configure each field with the `dataframe:"name[,type]"` struct
// tag. If the name on the tag is the empty string `""` the field name will be
// used instead. If the name is `"-"` the field will be ignored.
//
// Examples:
//
//	// field will be ignored
//	field int
//
//	// Field will be ignored
//	Field int `dataframe:"-"`
//
//	// Field will be parsed with column name Field and type int
//	Field int
//
//	// Field will be parsed with column name `field_column` and type int.
//	Field int `dataframe:"field_column"`
//
//	// Field will be parsed with column name `field` and type string.
//	Field int `dataframe:"field,string"`
//
//	// Field will be parsed with column name `Field` and type string.
//	Field int `dataframe:",string"`
//
// If the struct tags and the given LoadOptions contradict each other, the later
// will have preference over the former.
func LoadStructs(i any, options ...LoadOption) DataFrame {
	if i == nil {
		return DataFrame{Err: fmt.Errorf("load: can't create DataFrame from <nil> value")}
	}

	// Set the default load options
	cfg := loadOptions{
		defaultType: SERIES_TYPE_STRING,
		detectTypes: true,
		hasHeader:   true,
		nanValues:   num.PossibleNaOfString,
	}

	// Set any custom load options
	for _, option := range options {
		option(&cfg)
	}

	tpy, val := reflect.TypeOf(i), reflect.ValueOf(i)
	switch tpy.Kind() {
	case reflect.Slice:
		if tpy.Elem().Kind() != reflect.Struct {
			return DataFrame{Err: fmt.Errorf(
				"load: type %s (%s %s) is not supported, must be []struct", tpy.Name(), tpy.Elem().Kind(), tpy.Kind())}
		}
		if val.Len() == 0 {
			return DataFrame{Err: fmt.Errorf("load: can't create DataFrame from empty slice")}
		}

		numFields := val.Index(0).Type().NumField()
		var columns []Series
		for j := 0; j < numFields; j++ {
			// Extract field metadata
			if !val.Index(0).Field(j).CanInterface() {
				continue
			}
			field := val.Index(0).Type().Field(j)
			fieldName := field.Name
			fieldType := field.Type.String()

			// Process struct tags
			fieldTags := field.Tag.Get(DefaultTagName)
			if fieldTags == "-" {
				continue
			}
			tagOpts := strings.Split(fieldTags, ",")
			if len(tagOpts) > 2 {
				return DataFrame{Err: fmt.Errorf("malformed struct tag on field %s: %s", fieldName, fieldTags)}
			}
			if len(tagOpts) > 0 {
				if name := strings.TrimSpace(tagOpts[0]); name != "" {
					fieldName = name
				}
				if len(tagOpts) == 2 {
					if tagType := strings.TrimSpace(tagOpts[1]); tagType != "" {
						fieldType = tagType
					}
				}
			}

			// Handle `types` option
			var t Type
			if cfgType, ok := cfg.types[fieldName]; ok {
				t = cfgType
			} else {
				// Handle `detectTypes` option
				if cfg.detectTypes {
					// Parse field type
					parsedType, err := parseType(fieldType)
					if err != nil {
						return DataFrame{Err: err}
					}
					t = parsedType
				} else {
					t = cfg.defaultType
				}
			}

			// Create Series for this field
			elements := make([]any, val.Len())
			for i := 0; i < val.Len(); i++ {
				fieldValue := val.Index(i).Field(j)
				elements[i] = fieldValue.Interface()

				// Handle `nanValues` option
				if findInStringSlice(fmt.Sprint(elements[i]), cfg.nanValues) != -1 {
					elements[i] = nil
				}
			}

			// Handle `hasHeader` option
			if !cfg.hasHeader {
				tmp := make([]any, 1)
				tmp[0] = fieldName
				elements = append(tmp, elements...)
				fieldName = ""
			}
			if t == SERIES_TYPE_BOOL {
				columns = append(columns, NewSeriesWithType(SERIES_TYPE_BOOL, fieldName, elements...))
			} else if t == SERIES_TYPE_INT32 {
				columns = append(columns, NewSeriesWithType(SERIES_TYPE_INT32, fieldName, elements...))
			} else if t == SERIES_TYPE_INT64 {
				columns = append(columns, NewSeriesWithType(SERIES_TYPE_INT64, fieldName, elements...))
			} else if t == SERIES_TYPE_FLOAT32 {
				columns = append(columns, NewSeriesWithType(SERIES_TYPE_FLOAT32, fieldName, elements...))
			} else if t == SERIES_TYPE_FLOAT64 {
				columns = append(columns, NewSeriesWithType(SERIES_TYPE_FLOAT64, fieldName, elements...))
			} else {
				// 默认string
				columns = append(columns, NewSeriesWithType(SERIES_TYPE_STRING, fieldName, elements...))
			}
		}
		return NewDataFrame(columns...)
	}
	return DataFrame{Err: fmt.Errorf(
		"load: type %s (%s) is not supported, must be []struct", tpy.Name(), tpy.Kind())}
}
