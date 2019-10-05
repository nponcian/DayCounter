package main

import (
  "time"
)

type DateInfo struct {
    timeData time.Time
    year int
    month time.Month
    day int
    weekDay string
}

func (dateInfo *DateInfo) addDate(years int, months int, days int) {
    dateInfo.timeData = dateInfo.timeData.AddDate(years, months, 1)
    dateInfo.year, dateInfo.month, dateInfo.day = dateInfo.timeData.Date()
    dateInfo.weekDay = dateInfo.timeData.Format("Monday")
}

func (dateInfo DateInfo) isEarlierThan(other DateInfo) bool {
    return dateInfo.year < other.year ||
            (dateInfo.year == other.year &&
            (dateInfo.month < other.month ||
                (dateInfo.month == other.month && dateInfo.day < other.day)))
}

func (dateInfo DateInfo) isWeekEnd() bool {
    return dateInfo.weekDay == "Saturday" || dateInfo.weekDay == "Sunday"
}

func (dateInfo DateInfo) isWeekStart() bool {
    return dateInfo.weekDay == "Monday"
}

func (dateInfo DateInfo) isHoliday(holidays []Holiday) bool {
    for _, currentHoliday := range holidays {
        if dateInfo.month.String() == currentHoliday.month && dateInfo.day == currentHoliday.day {
            return true
        }
    }
    return false
}

func (dateInfo DateInfo) isLeaveDay(leaveDays []LeaveDay) bool {
    for _, currentLeaveDay := range leaveDays {
        if dateInfo.month.String() == currentLeaveDay.month && dateInfo.day == currentLeaveDay.day {
            return true
        }
    }
    return false
}

func getStartDateInfo() DateInfo {
    startTime := time.Now()

    // <MODIFY THIS IF START TIME IS DIFFERENT>
    startTime = time.Date(2019, 10, 2, 0, 0, 0, 0, time.UTC)
    // </MODIFY THIS IF START TIME IS DIFFERENT>

    year, month, day := startTime.Date()
    weekDay := startTime.Format("Monday")
    return DateInfo{startTime, year, month, day, weekDay}
}

func getTargetDateInfo() DateInfo {
    // <MODIFY THIS IF TARGET TIME IS DIFFERENT>
    targetTime := time.Date(2019, 12, 22, 0, 0, 0, 0, time.UTC)
    // </MODIFY THIS IF TARGET TIME IS DIFFERENT>

    year, month, day := targetTime.Date()
    weekDay := targetTime.Format("Monday")
    return DateInfo{targetTime, year, month, day, weekDay}
}
