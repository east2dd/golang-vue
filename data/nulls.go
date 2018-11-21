package data

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/guregu/null"
)

var parseLocation = time.UTC

func MakeNullString(val string) null.String {
	return null.NewString(val, val != "")
}

func NowPtr() *time.Time {
	t := time.Now()
	return &t
}

func NullStr(val string) null.String {
	return null.NewString(val, true)
}

func NullStringSlice(vals ...string) []null.String {
	nullVals := make([]null.String, len(vals))
	for i, x := range vals {
		nullVals[i] = null.NewString(x, true)
	}
	return nullVals
}

func ParseNullString(dest *null.String, value string) error {
	dest.String = value
	dest.Valid = value != ""
	return nil
}

func ParseNullInt(dest *null.Int, value string) error {
	if value == "" {
		dest.Int64 = 0
		dest.Valid = false
		return nil
	} else {
		value = strings.Replace(value, ",", "", -1)
		val, err := strconv.ParseFloat(value, 64)
		dest.Int64 = int64(val)
		dest.Valid = err == nil
		return err
	}
}

func ParseNullFloat(dest *null.Float, value string) error {
	if value == "" {
		dest.Float64 = 0
		dest.Valid = false
		return nil
	} else {
		value = strings.Replace(value, ",", "", -1)
		val, err := strconv.ParseFloat(value, 64)
		dest.Float64 = val
		dest.Valid = err == nil
		return err
	}
}

func ParseNullBool(dest *sql.NullBool, value string) error {
	if value == "" {
		dest.Bool = false
		dest.Valid = false
		return nil
	} else {
		val, err := strconv.ParseBool(value)
		dest.Bool = val
		dest.Valid = err == nil
		return err
	}
}

func ParseNullUSDate(dest *null.Time, value string) error {
	if value == "" {
		dest = &null.Time{}
		return nil
	} else {
		val, err := time.Parse("01/02/2006", value)
		dest.Time = val
		dest.Valid = err == nil
		return err
	}
}

func ParseFloat(dest *float64, value string) error {
	value = strings.Replace(value, ",", "", -1)
	val, err := strconv.ParseFloat(value, 64)
	*dest = val
	return err
}

func ParseInt(dest *int64, value string) error {
	value = strings.Replace(value, ",", "", -1)
	val, err := strconv.ParseFloat(value, 64)
	*dest = int64(val)
	return err
}

func SetNonZeroKey(store map[string]interface{}, column string, value interface{}) {
	switch value.(type) {
	case string:
		if value.(string) != "" {
			store[column] = value
		}
	case null.String:
		if value.(null.String).String != "" {
			store[column] = value
		}
	default:
		panic(errors.New("Type not handled."))
	}
}
