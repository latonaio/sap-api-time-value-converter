package sap_api_time_value_converter

import (
	"strings"

	"golang.org/x/xerrors"
)

func isSAPDateFormat(s string) bool {
	if err := validateSAPDateFormat(s); err != nil {
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
		return xerrors.Errorf(
			"%s is not type of SAP timestamp", sapTime,
		)
	}
	return nil
}

func validateSuffix(sapTime string) error {
	if !(strings.HasSuffix(sapTime, `)\/`) || strings.HasSuffix(sapTime, `)/`)) {
		return xerrors.Errorf(
			"%s is not type of SAP timestamp", sapTime,
		)
	}
	return nil
}
