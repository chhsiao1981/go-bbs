package types

import (
	"time"
)

type Time4 int32

func NowTS() Time4 {
	// We don't need to worry about time-zone when using unix-timestamp.
	return Time4(time.Now().Unix())
}

//ToLocal
//
//Instead of using Local, we specify TIME_LOCATION
//to avoid the confusion. (also good for tests)
func (t Time4) ToLocal() time.Time {
	return time.Unix(int64(t), 0).In(TIMEZONE)
}

//Cdate
//
//Print date-time in string.
//23+1 bytes, "12/31/2007 00:00:00 Mon\0"
func (t Time4) Cdate() string {
	return t.ToLocal().Format("01/02/2006 15:04:05 Mon")
}

//Cdatelite
//
//Light-print date-time in string.
//19+1 bytes, "12/31/2007 00:00:00\0"
func (t Time4) Cdatelite() string {
	return t.ToLocal().Format("01/02/2006 15:04:05")
}

//Cdatedate
//
//Print date in string.
//10+1 bytes, "12/31/2007\0"
func (t Time4) Cdatedate() string {
	return t.ToLocal().Format("01/02/2006")
}

//CdateMd
//
//Print month/day in string
//5+1 bytes, "12/31\0"
func (t Time4) CdateMd() string {
	return t.ToLocal().Format("01/02")
}

//CdateMdHM
//
//Print month/day hour:minute in string
//11+1 bytes, "12/31 10:01\0"
func (t Time4) CdateMdHM() string {
	return t.ToLocal().Format("01/02 15:04")
}

//CdateMdHMS
//
//Print month/day hour:minute:second in string
//13+1 bytes, "12/31 10:01:01\0"
func (t Time4) CdateMdHMS() string {
	return t.ToLocal().Format("01/02 15:04:05")
}
