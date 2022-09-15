package sap_api_time_value_converter

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func ConvertToSAPTimeFormat(t time.Time) string {
	return fmt.Sprintf(`/Date(%d)/`, t.UnixMilli())
}

func ChangeFormatToReadableDateTime(sapTime string) string {
	if sapTime == "" {
		return ""
	}
	t := ConvertToTimeFormat(sapTime)
	if t.Year() <= 1 {
		return ""
	}

	if t.UTC().Hour() == 0 && t.UTC().Minute() == 0 && t.UTC().Second() == 0 && t.UTC().Nanosecond() == 0 {
		return t.Format("2006-01-02")
	}

	return t.Format(time.RFC3339)
}

func ChangeFormatToReadableTime(sapTime string) string {
	if sapTime == "" {
		return ""
	}

	t, err := time.Parse("PT15H04M05S", sapTime)
	if err != nil {
		return sapTime
	}

	return t.Format("15:04:05")
}

func ChangeFormatToSAPFormat(readableTime string) string {
	if readableTime == "" {
		return ""
	}
	t, err := time.Parse(time.RFC3339, readableTime)
	if err != nil {
		return ""
	}
	return ConvertToSAPTimeFormat(t)
}

func ChangeTimeFormatToReadableForStruct(str interface{}) {
	rv := reflect.ValueOf(str)
	pickStringToReadable(rv)
}

func ChangeTimeFormatToSAPFormatStruct(str interface{}) {
	rv := reflect.ValueOf(str)
	pickStringToSAPFormat(rv)
}

func pickStringToSAPFormat(rv reflect.Value) {
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface:
		pickStringToSAPFormat(rv.Elem())
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			pickStringToSAPFormat(rv.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			pickStringToSAPFormat(rv.Field(i))
		}

	}
	if rv.Kind() == reflect.String {
		changeValueToSAPFormat(rv)
	}
}

func changeValueToSAPFormat(rv reflect.Value) {
	if rv.Kind() != reflect.String {
		return
	}
	if !rv.CanSet() {
		return
	}

	strValue := rv.String()
	if isReadableTimeFormat(strValue) {
		rv.SetString(ChangeFormatToSAPFormat(strValue))
	}
}

func pickStringToReadable(rv reflect.Value) {
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface:
		pickStringToReadable(rv.Elem())
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			pickStringToReadable(rv.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			pickStringToReadable(rv.Field(i))
		}

	}
	if rv.Kind() == reflect.String {
		changeValueToReadable(rv)
	}
}

func changeValueToReadable(rv reflect.Value) {
	if rv.Kind() != reflect.String {
		return
	}
	if !rv.CanSet() {
		return
	}

	strValue := rv.String()
	if isSAPDateFormat(strValue) {
		rv.SetString(ChangeFormatToReadableDateTime(strValue))
	}

	if isSAPDurationFormat(strValue) {
		rv.SetString(ChangeFormatToReadableTime(strValue))
	}

}

func ConvertToTimeFormat(sapTime string) time.Time {
	err := validateSAPDateFormat(sapTime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err)
		return time.Time{}
	}

	milli, err := getUnixmilli(sapTime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err)
		return time.Time{}
	}
	return time.UnixMilli(milli)
}

// getUnixmilli unixミリ秒を返す
func getUnixmilli(sapTime string) (int64, error) {
	fixedString := strings.Join(strings.Split(sapTime, `\`), "")
	num := fixedString[len(`/Date(`) : len(fixedString)-len(`)/`)]
	num = strings.Split(num, "+")[0]
	milli, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("given word '%s' can not be converted to number: %w", sapTime, err)
	}
	return milli, nil
}

