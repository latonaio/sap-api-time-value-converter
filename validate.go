package sap_api_time_value_converter

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func isSAPDateFormat(s string) bool {
	if err := validateSAPDateFormat(s); err != nil {
		return false
	}
	return true
}
func isSAPDurationFormat(sapTime string) bool {
	ok, _ := regexp.MatchString(`PT[0-2]\dH[0-6]\dM[0-6]\dS`, sapTime)
	return ok
}
func isSAPDateTimeFormat(sapTime string) bool {
	_, err := time.Parse("20060102150405", sapTime)
	return err == nil
}

func isReadableTimeFormat(s string) bool {
	if _, err := time.Parse(time.RFC3339, s); err != nil {
		return false
	}
	return true
}

func isReadableDateFormat(s string) bool {
	if _, err := time.Parse("2006-01-02", s); err != nil {
		return false
	}
	return true
}

func validateSAPDateFormat(sapTime string) error {
	err := validatePrefix(sapTime)
	if err != nil {
		return err
	}
	err = validateSuffix(sapTime)
	if err != nil {
		return err
	}
	return nil
}

func validatePrefix(sapTime string) error {
	if !(strings.HasPrefix(sapTime, `\/Date(`) || strings.HasPrefix(sapTime, `/Date(`)) {
		return fmt.Errorf(
			"%s is not type of SAP timestamp", sapTime,
		)
	}
	return nil
}

func validateSuffix(sapTime string) error {
	if !(strings.HasSuffix(sapTime, `)\/`) || strings.HasSuffix(sapTime, `)/`)) {
		return fmt.Errorf(
			"%s is not type of SAP timestamp", sapTime,
		)
	}
	return nil
}
