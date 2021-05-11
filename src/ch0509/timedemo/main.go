package main

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

const (
	AUTO_TIME_ZONE  = true
	DATE_LAYOUT     = "2006-01-02"
	TIME_LAYOUT     = "15:04:05"
	DATETIME_LAYOUT = "2006-01-02 15:04:05"
)

func main() {
	dateSeri()
	mothSeri()
	yearSeri()
}

func dateSeri() {
	minDateTime, _ := dateFrom("2019-01-01")
	maxDateTime, _ := dateFrom("2021-05-04")
	diff := maxDateTime.Sub(minDateTime)
	dateStrList := []string{dateTo(minDateTime)}
	for nextDate := minDateTime; ; {
		nextDate = nextDate.AddDate(0, 0, 1)
		if maxDateTime.Sub(nextDate) < 0 {
			break
		}
		dateStrList = append(dateStrList, dateTo(nextDate))
	}
	fmt.Println(diff)
	fmt.Println(dateStrList)
	fmt.Println(reflect.TypeOf(diff))
}

func mothSeri() {
	//minMothTime, _ := timeParse("2019-01", "2006-01")
	minMothTime1, _ := time.Parse("2006-01", "2019-01")
	maxMothTime2, _ := time.Parse("2006-01", "2021-05")
	minMothTime, _ := timeParse("2019-01", "2006-01")
	maxMothTime, _ := timeParse("2021-05", "2006-01")
	fmt.Println("-------------------moth info-----------------")
	fmt.Println(minMothTime1.Local())
	fmt.Println(maxMothTime2.Local())
	fmt.Println("+++++++++++++++++++++++++")

	fmt.Println(minMothTime)
	fmt.Println(maxMothTime)
	diff := maxMothTime.Sub(minMothTime)
	mothStrList := []string{dateTo(minMothTime)}
	for nextDate := minMothTime; ; {
		nextDate = nextDate.AddDate(0, 1, 0)
		if maxMothTime.Sub(nextDate) < 0 {
			break
		}
		mothStrList = append(mothStrList, dateTo(nextDate))
	}
	fmt.Println(diff)
	fmt.Println(mothStrList)
	fmt.Println(reflect.TypeOf(diff))
}

func yearSeri() {
	minYearTime, _ := timeParse("2019", "2006")
	maxYearTime, _ := timeParse("2021", "2006")
	minYear, _ := time.Parse("2006", "2019")
	maxYear, _ := time.Parse("2006", "2021")
	fmt.Println("-------------------year info-----------------")
	fmt.Println(minYear)
	fmt.Println(maxYear)
	fmt.Println("------------------------------------------")
	fmt.Println(minYearTime)
	fmt.Println(maxYearTime)
	diff := maxYearTime.Sub(minYearTime)
	YearStrList := []string{dateTo(minYearTime)}
	for nextDate := minYearTime; ; {
		nextDate = nextDate.AddDate(1, 0, 0)
		if maxYearTime.Sub(nextDate) < 0 {
			break
		}
		YearStrList = append(YearStrList, dateTo(nextDate))
	}
	//linq.Form(YearStrList)
	fmt.Println(diff)
	fmt.Println(YearStrList)
	fmt.Println(reflect.TypeOf(diff))
}

// dateFrom
func dateFrom(val string) (time.Time, error) {
	return timeParse(val, DATE_LAYOUT)
}

// timeFrom
func timeFrom(val string) (time.Time, error) {
	return timeParse(val, TIME_LAYOUT)
}

// datetimeFrom
func datetimeFrom(val string) (time.Time, error) {
	return timeParse(val, DATETIME_LAYOUT)
}

// dateTo
func dateTo(val time.Time) string {
	return timeFormat(val, DATE_LAYOUT)
}

// timeTo
func timeTo(val time.Time) string {
	return timeFormat(val, TIME_LAYOUT)
}

// DatetimeTo
func DatetimeTo(val time.Time) string {
	return timeFormat(val, DATETIME_LAYOUT)
}

// timeParse
func timeParse(val string, layout string) (time.Time, error) {
	if s, err := time.Parse(layout, val); err == nil {
		if AUTO_TIME_ZONE {
			var duration time.Duration
			_, offset := time.Now().Zone()
			if offset > 0 {
				tz := math.Ceil(float64(offset / 3600))
				if tz > 0 {
					duration = time.Duration(-tz)
				} else {
					duration = time.Duration(tz)
				}
			}

			return s.Add(time.Hour * duration).Local(), nil
		}

		return s, nil
	} else {
		return time.Now(), err
	}
}

// timeFormat
func timeFormat(val time.Time, layout string) string {
	if val.IsZero() {
		return ""
	}
	return val.Format(layout)
}
