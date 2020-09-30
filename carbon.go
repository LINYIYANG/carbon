package carbon

import (
	"strconv"
	"strings"
	"time"
)

type Carbon struct {
	Time time.Time
	loc  *time.Location
}

// New 初始化,设置默认时区
func New() *Carbon {
	loc, _ := time.LoadLocation(Local)
	return &Carbon{loc: loc}
}

// Timezone 设置时区，必须在New()之后，其他方法之前调用
func (c *Carbon) Timezone(name string) *Carbon {
	loc, _ := time.LoadLocation(name)
	c.loc = loc
	return c
}

// CreateFromTimestamp 将时间戳转成时间实例
func (c *Carbon) CreateFromTimestamp(timestamp int64) *Carbon {
	c.Time = time.Unix(timestamp, 0)
	return c
}

// CreateFromDateTime 将年月日时分秒转成时间实例
func (c *Carbon) CreateFromDateTime(year int, month time.Month, day int, hour int, minute int, second int) *Carbon {
	c.Time = time.Date(year, month, day, hour, minute, second, 0, c.loc)
	return c
}

// CreateFromDate 将年月日转成时间实例
func (c *Carbon) CreateFromDate(year int, month time.Month, day int) *Carbon {
	c.Time = time.Date(year, month, day, 0, 0, 0, 0, c.loc)
	return c
}

// CreateFromTime 将时分秒转成时间类型
func (c *Carbon) CreateFromTime(hour int, minute int, second int) *Carbon {
	c.Time = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, minute, second, 0, c.loc)
	return c
}

// CreateFromString 将字符串转成时间类型
func (c *Carbon) CreateFromString(duration string) *Carbon {
	d, _ := time.ParseDuration(duration)
	c.Time = c.Time.Add(d)
	return c
}

// Now 当前
func (c *Carbon) Now() *Carbon {
	c.Time = time.Now()
	return c
}

// AddYears N年后
func (c *Carbon) AddYears(years int) *Carbon {
	c.Time = c.Time.AddDate(years, 0, 0)
	return c
}

// AddYear 1年后
func (c *Carbon) AddYear() *Carbon {
	c.Time = c.Time.AddDate(1, 0, 0)
	return c
}

// SubYears N年前
func (c *Carbon) SubYears(years int) *Carbon {
	c.Time = c.Time.AddDate(-years, 0, 0)
	return c
}

// SubYear 1年前
func (c *Carbon) SubYear() *Carbon {
	c.Time = c.Time.AddDate(-1, 0, 0)
	return c
}

// AddMonths N月后
func (c *Carbon) AddMonths(months int) *Carbon {
	c.Time = c.Time.AddDate(0, months, 0)
	return c
}

// AddMonth 1月后
func (c *Carbon) AddMonth() *Carbon {
	c.Time = c.Time.AddDate(0, 1, 0)
	return c
}

// SubMonths N月前
func (c *Carbon) SubMonths(months int) *Carbon {
	c.Time = c.Time.AddDate(0, -months, 0)
	return c
}

// SubMonth 1月前
func (c *Carbon) SubMonth() *Carbon {
	c.Time = c.Time.AddDate(0, -1, 0)
	return c
}

// AddWeeks N周后
func (c *Carbon) AddWeeks(weeks int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, weeks*DaysPerWeek)
	return c
}

// AddWeek 1周后
func (c *Carbon) AddWeek() *Carbon {
	c.Time = c.Time.AddDate(0, 0, DaysPerWeek)
	return c
}

// SubWeeks N周前
func (c *Carbon) SubWeeks(weeks int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -weeks*DaysPerWeek)
	return c
}

// SubMonth 1月前
func (c *Carbon) SubWeek() *Carbon {
	c.Time = c.Time.AddDate(0, 0, -DaysPerWeek)
	return c
}

// AddDays N天后
func (c *Carbon) AddDays(days int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, days)
	return c
}

// AddDay 1天后
func (c *Carbon) AddDay() *Carbon {
	c.Time = c.Time.AddDate(0, 0, 1)
	return c
}

// SubDays N天前
func (c *Carbon) SubDays(days int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -days)
	return c
}

// SubDay 1天前
func (c *Carbon) SubDay() *Carbon {
	c.Time = c.Time.AddDate(0, 0, -1)
	return c
}

// AddHours N小时后
func (c *Carbon) AddHours(hours int) *Carbon {
	duration := strconv.Itoa(hours) + "h"
	return c.CreateFromString(duration)
}

// AddHour 1小时后
func (c *Carbon) AddHour() *Carbon {
	duration := "1h"
	return c.CreateFromString(duration)
}

// AddHours N小时前
func (c *Carbon) SubHours(hours int) *Carbon {
	duration := "-" + strconv.Itoa(hours) + "h"
	return c.CreateFromString(duration)
}

// AddHour 1小时前
func (c *Carbon) SubHour() *Carbon {
	duration := "-1h"
	return c.CreateFromString(duration)
}

// AddMinutes N分钟后
func (c *Carbon) AddMinutes(minutes int) *Carbon {
	duration := strconv.Itoa(minutes) + "m"
	return c.CreateFromString(duration)
}

// AddMinute 1分钟后
func (c *Carbon) AddMinute() *Carbon {
	duration := "1m"
	return c.CreateFromString(duration)
}

// SubMinutes N分钟前
func (c *Carbon) SubMinutes(minutes int) *Carbon {
	duration := "-" + strconv.Itoa(minutes) + "m"
	return c.CreateFromString(duration)
}

// SubMinute 1分钟前
func (c *Carbon) SubMinute() *Carbon {
	duration := "-1m"
	return c.CreateFromString(duration)
}

// AddSeconds N秒钟后
func (c *Carbon) AddSeconds(second int) *Carbon {
	duration := strconv.Itoa(second) + "s"
	return c.CreateFromString(duration)
}

// AddMinute 1秒钟后
func (c *Carbon) AddSecond() *Carbon {
	duration := "1s"
	return c.CreateFromString(duration)
}

// SubMinutes N秒钟前
func (c *Carbon) SubSeconds(second int) *Carbon {
	duration := "-" + strconv.Itoa(second) + "s"
	return c.CreateFromString(duration)
}

// SubMinute 1秒钟前
func (c *Carbon) SubSecond() *Carbon {
	duration := "-1s"
	return c.CreateFromString(duration)
}

// Parse 解析标准时间格式
func (c *Carbon) Parse(value string) *Carbon {
	value = strings.Trim(value, " ")
	layout := "2006-01-02 15:04:05"

	if len(value) == 10 && strings.Contains(value, "-") {
		layout = "2006-01-02"
	}

	if len(value) == 10 && strings.Contains(value, "/") {
		layout = "2006/01/02"
	}

	if len(value) == 19 && strings.Contains(value, "/") {
		layout = "2006/01/02 15:04:05"
	}

	if len(value) == 14 {
		layout = "20060102150405"
	}

	if len(value) == 8 {
		layout = "20060102"
	}

	t, _ := time.ParseInLocation(layout, value, c.loc)
	c.Time = t
	return c
}

// ParseByCustom 解析自定义时间格式
func (c *Carbon) ParseByCustom(value string, format string) *Carbon {
	value = strings.Trim(value, " ")
	layout := format2layout(format)
	t, _ := time.ParseInLocation(layout, value, c.loc)
	c.Time = t
	return c
}

// Format 格式化时间
func (c *Carbon) Format(format string) string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.loc).Format(format2layout(format))
}

// Today 今天
func (c *Carbon) Today() string {
	now := time.Now()
	return now.In(c.loc).Format("2006-01-02")
}

// StartOfToday 今天开始时间
func (c *Carbon) StartOfToday() string {
	return c.Today() + " 00:00:00"
}

// EndOfToday 今天结束时间
func (c *Carbon) EndOfToday() string {
	return c.Today() + " 23:59:59"
}

// Tomorrow 明天
func (c *Carbon) Tomorrow() string {
	tomorrow := time.Now().AddDate(0, 0, 1)
	return tomorrow.In(c.loc).Format("2006-01-02")
}

// StartOfTomorrow 明天开始时间
func (c *Carbon) StartOfTomorrow() string {
	return c.Tomorrow() + " 00:00:00"
}

// EndOfTomorrow 明天结束时间
func (c *Carbon) EndOfTomorrow() string {
	return c.Tomorrow() + " 23:59:59"
}

// Yesterday 昨天
func (c *Carbon) Yesterday() string {
	yesterday := time.Now().AddDate(0, 0, -1)
	return yesterday.In(c.loc).Format("2006-01-02")
}

// StartOfYesterday 昨天开始时间
func (c *Carbon) StartOfYesterday() string {
	return c.Yesterday() + " 00:00:00"
}

// EndOfYesterday 昨天结束时间
func (c *Carbon) EndOfYesterday() string {
	return c.Yesterday() + " 23:59:59"
}

// FirstDayInYear 当年第一天
func (c *Carbon) FirstOfYear() string {
	return c.CreateFromDate(c.Time.Year(), 01, 01).ToDateStartString()
}

// LastDayInYear 当年最后一天
func (c *Carbon) LastOfYear() string {
	return c.CreateFromDate(c.Time.Year(), 12, 31).ToDateStartString()
}

// FirstDayInMonth 当月第一天
func (c *Carbon) FirstOfMonth() string {
	return c.CreateFromDate(c.Time.Year(), c.Time.Month(), 01).ToDateStartString()
}

// LastDayInMonth 当月最后一天
func (c *Carbon) LastOfMonth() string {
	return c.CreateFromDate(c.Time.Year(), c.Time.Month(), c.getDays()).ToDateStartString()
}

// StartOfYear 当年开始时间
func (c *Carbon) StartOfYear() string {
	return c.CreateFromDate(c.Time.Year(), 01, 01).Format("2006-01-02 00:00:00")
}

// EndOfYear 当年结束时间
func (c *Carbon) EndOfYear() string {
	return c.CreateFromDateTime(c.Time.Year(), 12, 31, HoursPerDay-1, MinutesPerHour-1, SecondsPerMinute-1).Format("2006-01-02 15:04:05")
}

// StartOfMonth 当月开始时间
func (c *Carbon) StartOfMonth() string {
	return c.CreateFromDate(c.Time.Year(), c.Time.Month(), 01).ToDateStartString()
}

// EndOfMonth 当月结束时间
func (c *Carbon) EndOfMonth() string {
	return c.CreateFromDate(c.Time.Year(), c.Time.Month(), c.getDays()).ToDateEndString()
}

// FirstDayInMonth 当天开始时间
func (c *Carbon) StartOfDay() string {
	return c.CreateFromDate(c.Time.Year(), c.Time.Month(), c.Time.Day()).ToDateStartString()
}

// LastDayInMonth 当天结束时间
func (c *Carbon) EndOfDay() string {
	return c.CreateFromDate(c.Time.Year(), c.Time.Month(), c.Time.Day()).ToDateEndString()
}

// ToDateTimeString 转日期时间字符串
func (c *Carbon) ToDateTimeString() string {
	return c.Time.In(c.loc).Format("2006-01-02 15:04:05")
}

// ToDateString 转日期字符串
func (c *Carbon) ToDateString() string {
	return c.Time.In(c.loc).Format("2006-01-02")
}

// ToDateStartString 转日期开始时间字符串
func (c *Carbon) ToDateStartString() string {
	return c.Time.In(c.loc).Format("2006-01-02 00:00:00")
}

// ToDateEndString 转日期结束时间字符串
func (c *Carbon) ToDateEndString() string {
	return c.CreateFromDateTime(c.Time.Year(), c.Time.Month(), c.Time.Day(), HoursPerDay-1, MinutesPerHour-1, SecondsPerMinute-1).Format("2006-01-02 15:04:05")
}

// ToTimeString 转时间字符串
func (c *Carbon) ToTimeString() string {
	return c.Time.In(c.loc).Format("15:04:05")
}

// ToTimeStartString 转开始时间字符串
func (c *Carbon) ToTimeStartString() string {
	return c.Time.In(c.loc).Format("15:00:00")
}

// ToDateEndString 转结束时间字符串
func (c *Carbon) ToTimeEndString() string {
	return c.CreateFromDateTime(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), MinutesPerHour-1, SecondsPerMinute-1).Format("15:04:05")
}

// ToTimestamp 转时间戳
func (c *Carbon) ToTimestamp() int64 {
	return c.Time.Unix()
}

// IsLeapYear 是否是闰年
func (c *Carbon) IsLeapYear() bool {
	year := c.Time.Year()
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// IsJanuary 是否是一月
func (c *Carbon) IsJanuary() bool {
	return c.Time.In(c.loc).Month() == time.January
}

// IsMonday 是否是二月
func (c *Carbon) IsFebruary() bool {
	return c.Time.In(c.loc).Month() == time.February
}

// IsMarch 是否是三月
func (c *Carbon) IsMarch() bool {
	return c.Time.In(c.loc).Month() == time.March
}

// IsApril 是否是四月
func (c *Carbon) IsApril() bool {
	return c.Time.In(c.loc).Month() == time.April
}

// IsMay 是否是五月
func (c *Carbon) IsMay() bool {
	return c.Time.In(c.loc).Month() == time.May
}

// IsJune 是否是六月
func (c *Carbon) IsJune() bool {
	return c.Time.In(c.loc).Month() == time.June
}

// IsJuly 是否是七月
func (c *Carbon) IsJuly() bool {
	return c.Time.In(c.loc).Month() == time.July
}

// IsAugust 是否是八月
func (c *Carbon) IsAugust() bool {
	return c.Time.In(c.loc).Month() == time.August
}

// IsSeptember 是否是九月
func (c *Carbon) IsSeptember() bool {
	return c.Time.In(c.loc).Month() == time.September
}

// IsOctober 是否是十月
func (c *Carbon) IsOctober() bool {
	return c.Time.In(c.loc).Month() == time.October
}

// IsNovember 是否是十一月
func (c *Carbon) IsNovember() bool {
	return c.Time.In(c.loc).Month() == time.November
}

// IsDecember 是否是十二月
func (c *Carbon) IsDecember() bool {
	return c.Time.In(c.loc).Month() == time.December
}

// IsMonday 是否是周一
func (c *Carbon) IsMonday() bool {
	return c.Time.In(c.loc).Weekday() == time.Monday
}

// IsTuesday 是否是周二
func (c *Carbon) IsTuesday() bool {
	return c.Time.In(c.loc).Weekday() == time.Tuesday
}

// IsWednesday 是否是周三
func (c *Carbon) IsWednesday() bool {
	return c.Time.In(c.loc).Weekday() == time.Wednesday
}

// IsThursday 是否是周四
func (c *Carbon) IsThursday() bool {
	return c.Time.In(c.loc).Weekday() == time.Thursday
}

// IsFriday 是否是周五
func (c *Carbon) IsFriday() bool {
	return c.Time.In(c.loc).Weekday() == time.Friday
}

// IsSaturday 是否是周六
func (c *Carbon) IsSaturday() bool {
	return c.Time.In(c.loc).Weekday() == time.Saturday
}

// IsSunday 是否是周日
func (c *Carbon) IsSunday() bool {
	return c.Time.In(c.loc).Weekday() == time.Sunday
}

// IsToday 是否是今天
func (c *Carbon) IsToday() bool {
	return c.ToDateString() == c.Today()
}

// IsYesterday 是否是昨天
func (c *Carbon) IsYesterday() bool {
	return c.SubDay().ToDateString() == c.Yesterday()
}

// IsTomorrow 是否是明天
func (c *Carbon) IsTomorrow() bool {
	return c.AddDay().ToDateString() == c.Tomorrow()
}

// IsFirstOfYear 是否年初
func (c *Carbon) IsFirstOfYear() bool {
	_, month, day := c.Time.Date()
	if month == time.January && day == 1 {
		return true
	}
	return false
}

// IsLastOfYear 是否是年末
func (c *Carbon) IsLastOfYear() bool {
	_, month, day := c.Time.Date()
	if month == time.December && day == 31 {
		return true
	}
	return false
}

// IsFirstOfMonth 是否月初
func (c *Carbon) IsFirstOfMonth() bool {
	return c.ToDateTimeString() == c.FirstOfMonth()
}

// IsLastOfMonth 是否是月末
func (c *Carbon) IsLastOfMonth() bool {
	return c.ToDateTimeString() == c.LastOfMonth()
}

// IsStartOfYear 是否当年开始时间
func (c *Carbon) IsStartOfYear() bool {
	return c.ToDateTimeString() == c.StartOfYear()
}

// IsEndOfYear 是否当年结束时间
func (c *Carbon) IsEndOfYear() bool {
	return c.ToDateTimeString() == c.EndOfYear()
}

// IsStartOfMonth 是否当月开始时间
func (c *Carbon) IsStartOfMonth() bool {
	return c.ToDateTimeString() == c.StartOfMonth()
}

// IsEndOfMonth 是否当月结束时间
func (c *Carbon) IsEndOfMonth() bool {
	return c.ToDateTimeString() == c.EndOfMonth()
}

// IsStartOfDay 是否当天开始时间
func (c *Carbon) IsStartOfDay() bool {
	return c.ToDateTimeString() == c.StartOfDay()
}

// IsEndOfDay 是否当天结束时间
func (c *Carbon) IsEndOfDay() bool {
	return c.ToDateTimeString() == c.EndOfDay()
}

// IsStartOfToday 是否今天开始时间
func (c *Carbon) IsStartOfToday() bool {
	return c.ToDateTimeString() == c.StartOfToday()
}

// IsEndOfToday 是否今天结束时间
func (c *Carbon) IsEndOfToday() bool {
	return c.ToDateTimeString() == c.EndOfToday()
}

// IsStartOfTomorrow 是否明天开始时间
func (c *Carbon) IsStartOfTomorrow() bool {
	return c.ToDateTimeString() == c.StartOfTomorrow()
}

// IsEndOfTomorrow 是否明天结束时间
func (c *Carbon) IsEndOfTomorrow() bool {
	return c.ToDateTimeString() == c.EndOfTomorrow()
}

// IsStartOfYesterday 是否昨天开始时间
func (c *Carbon) IsStartOfYesterday() bool {
	return c.ToDateTimeString() == c.StartOfYesterday()
}

// IsEndOfYesterday 是否昨天结束时间
func (c *Carbon) IsEndOfYesterday() bool {
	return c.ToDateTimeString() == c.EndOfYesterday()
}
