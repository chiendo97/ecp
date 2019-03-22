package ecp

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// parseSlice support slice of string, int, int8, int16, int32, int64
// float32, float64, uint, uint8, uint16, uint32, uint64, bool
func parseSlice(v string, field reflect.Value) error {

	if !field.CanAddr() {
		return fmt.Errorf("field is not addressable")
	}
	if field.Kind() != reflect.Slice {
		return fmt.Errorf("field is not slice")
	}

	// either space nor commas is perfect, but I think space is better
	// since it's more natural: fmt.Println([]int{1, 2, 3}) = [1 2 3]
	stringSlice := strings.Split(v, " ") // split by space
	if v == "" {
		stringSlice = nil
	}

	field.Set(reflect.MakeSlice(field.Type(), len(stringSlice), cap(stringSlice)))

	typ := field.Type().String()[2:]
	switch typ {
	case reflect.String.String():
		field.Set(reflect.ValueOf(stringSlice))

	case reflect.Int.String():
		slice := []int{}
		for _, s := range stringSlice {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return err
			}
			slice = append(slice, int(i))
		}
		field.Set(reflect.ValueOf(slice))

	case reflect.Int8.String():
		slice := []int8{}
		for _, s := range stringSlice {
			i, err := strconv.ParseInt(s, 10, 8)
			if err != nil {
				return err
			}
			slice = append(slice, int8(i))
		}
		field.Set(reflect.ValueOf(slice))

	case reflect.Int16.String():
		slice := []int16{}
		for _, s := range stringSlice {
			i, err := strconv.ParseInt(s, 10, 16)
			if err != nil {
				return err
			}
			slice = append(slice, int16(i))
		}
		field.Set(reflect.ValueOf(slice))

	case reflect.Int32.String():
		slice := []int32{}
		for _, s := range stringSlice {
			i, err := strconv.ParseInt(s, 10, 32)
			if err != nil {
				return err
			}
			slice = append(slice, int32(i))
		}
		field.Set(reflect.ValueOf(slice))

	case reflect.Int64.String():
		slice := []int64{}
		for _, s := range stringSlice {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return err
			}
			slice = append(slice, i)
		}
		field.Set(reflect.ValueOf(slice))

	case reflect.Float32.String():
		slice := []float32{}
		for _, s := range stringSlice {
			i, err := strconv.ParseFloat(s, 32)
			if err != nil {
				return err
			}
			slice = append(slice, float32(i))
		}
		field.Set(reflect.ValueOf(slice))

	case reflect.Float64.String():
		slice := []float64{}
		for _, s := range stringSlice {
			i, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return err
			}
			slice = append(slice, float64(i))
		}
		field.Set(reflect.ValueOf(slice))

	case reflect.Uint.String():
		slice := []uint{}
		for _, s := range stringSlice {
			i, err := strconv.ParseUint(s, 10, 64)
			if err != nil {
				return err
			}
			slice = append(slice, uint(i))
		}
		field.Set(reflect.ValueOf(slice))

	case reflect.Uint8.String():
		slice := []uint8{}
		for _, s := range stringSlice {
			i, err := strconv.ParseUint(s, 10, 8)
			if err != nil {
				return err
			}
			slice = append(slice, uint8(i))
		}
		field.Set(reflect.ValueOf(slice))

	case reflect.Uint16.String():
		slice := []uint16{}
		for _, s := range stringSlice {
			i, err := strconv.ParseUint(s, 10, 16)
			if err != nil {
				return err
			}
			slice = append(slice, uint16(i))
		}
		field.Set(reflect.ValueOf(slice))

	case reflect.Uint32.String():
		slice := []uint32{}
		for _, s := range stringSlice {
			i, err := strconv.ParseUint(s, 10, 32)
			if err != nil {
				return err
			}
			slice = append(slice, uint32(i))
		}
		field.Set(reflect.ValueOf(slice))

	case reflect.Uint64.String():
		slice := []uint64{}
		for _, s := range stringSlice {
			i, err := strconv.ParseUint(s, 10, 64)
			if err != nil {
				return err
			}
			slice = append(slice, i)
		}
		field.Set(reflect.ValueOf(slice))

	case reflect.Bool.String():
		slice := []bool{}
		for _, s := range stringSlice {
			i, err := strconv.ParseBool(strings.ToLower(s))
			if err != nil {
				return err
			}
			slice = append(slice, i)
		}
		field.Set(reflect.ValueOf(slice))

	default:
		return fmt.Errorf("doesn't support type %s", typ)

	}

	return nil
}
