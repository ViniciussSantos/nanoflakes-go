package nanoflakes

import (
	"strconv"
	"time"
)

type Nanoflake struct {
	epoch int64
	value int64
}

func (n Nanoflake) ToString() string {
	return strconv.FormatInt(n.value, 10)
}

func (n Nanoflake) CreationTimeMillis() int64 {
	return n.epoch + TimeStampValue(n.value)
}

func (n Nanoflake) CreationTime() time.Time {
	return time.UnixMilli(n.CreationTimeMillis())
}
