package sap_api_time_value_converter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

func ConvertToSAPTimeFormat(t time.Time) string {
	return fmt.Sprintf(`\/Date(%d)\/`, t.UnixMilli())
}

func ConvertToTimeFormat(sapTime string) time.Time {
	err := validatePrefix(sapTime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err)
		return time.Time{}
	}
	err = validateSuffix(sapTime)
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

func validatePrefix(sapTime string) error {
	if !strings.HasPrefix(sapTime, `\/Date(`) {
		return xerrors.Errorf(
			"%s is not type of SAP timestamp", sapTime,
		)
	}
	return nil
}

func validateSuffix(sapTime string) error {
	if !strings.HasSuffix(sapTime, `)\/`) {
		return xerrors.Errorf(
			"%s is not type of SAP timestamp", sapTime,
		)
	}
	return nil
}

func getUnixmilli(sapTime string) (int64, error) {
	num := sapTime[len(`\/Date(`) : len(sapTime)-len(`)\/`)]
	milli, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return -1, xerrors.Errorf("given word '%s' can not be converted to number: %w", sapTime, err)
	}
	return milli, nil
}
