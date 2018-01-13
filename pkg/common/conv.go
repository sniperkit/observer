package common

import "strconv"

func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func Uint32ToString(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}
