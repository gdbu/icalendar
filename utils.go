package icalendar

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func getHeader(company, product string) (out []byte) {
	return []byte(fmt.Sprintf(header, company, product))
}

func appendString(bs []byte, prefix, val, suffix string) []byte {
	if len(val) == 0 {
		return bs
	}

	val = string(bytes.Replace([]byte(val), []byte{'\n'}, []byte("<br>"), -1))
	return append(bs, splitTo(prefix+val, 75)+suffix...)
}

func appendInt64(bs []byte, prefix string, val int64, suffix string) []byte {
	if val == 0 {
		return bs
	}

	return appendString(bs, prefix, strconv.FormatInt(val, 10), suffix)

}

func appendTime(bs []byte, prefix string, val time.Time, suffix string) []byte {
	if val.IsZero() {
		return bs
	}

	return appendString(bs, prefix, val.Format(dateFmt), suffix)
}

func appendStringSlice(bs []byte, prefix string, val []string, suffix string) []byte {
	if len(val) == 0 {
		return bs
	}

	return appendString(bs, prefix, strings.Join(val, ","), suffix)
}

func appendStringer(bs []byte, prefix string, val Stringer, suffix string) []byte {
	if reflect.ValueOf(val).IsNil() {
		return bs
	}

	return appendString(bs, prefix, val.String(), suffix)
}

func splitTo(in string, limit int) string {
	if len(in) <= limit {
		return in
	}

	var bs []byte
	for i := 0; i < len(in); {
		var (
			delta int
			part  string
		)

		if delta = len(in) - i; delta > limit {
			delta = limit
			part = in[i:i+delta] + "\r\n "
		} else {
			part = in[i : i+delta]
		}

		bs = append(bs, part...)
		i += delta
	}

	return string(bs)
}

// Stringer represents data structures which can stringify themselves
type Stringer interface {
	String() string
}
