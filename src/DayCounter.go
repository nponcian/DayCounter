// You can edit this code!
// Click here and start typing.
package main

import (
  "fmt"
  "time"
)

// ----------------
// DATE INFORMATION
// ----------------

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
    return dateInfo.year <= other.year &&
            (dateInfo.month < other.month ||
            (dateInfo.month == other.month && dateInfo.day < other.day))
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
    startTime = time.Date(2019, 9, 14, 0, 0, 0, 0, time.UTC)
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

// -------------------
// HOLIDAY INFORMATION
// -------------------

type Holiday struct {
    month string
    day int
}

func getHolidays() []Holiday {

    // <MODIFY THIS FOR LIST OF HOLIDAYS>
    return []Holiday {
        Holiday{"November", 1},
        Holiday{"November", 2},
        Holiday{"November", 30},
        Holiday{"December", 8},
    }
    // </MODIFY THIS FOR LIST OF HOLIDAYS>

}

// -------------------
// LEAVES INFORMATION
// -------------------

type LeaveDay struct {
    month string
    day int
}

func getLeaveDays() []LeaveDay {

    // <MODIFY THIS FOR LIST OF LEAVEDAYS>
    return []LeaveDay {
        LeaveDay{"November", 22},
        LeaveDay{"December", 17},
        LeaveDay{"December", 18},
        LeaveDay{"December", 19},
        LeaveDay{"December", 20},
    }
    // </MODIFY THIS FOR LIST OF LEAVEDAYS>

}

// -------------------
// MAIN
// -------------------

func main() {  
    dateInfo := getStartDateInfo()
    targetDateInfo := getTargetDateInfo()

    holidays := getHolidays()
    leaveDays := getLeaveDays()

    fmt.Println("StartDateInfo:", dateInfo)
    fmt.Println("TargetDateInfo:", targetDateInfo)
    fmt.Println()

    for daysCtr, workingDaysCtr, weekCtr := 1, 1, 1; dateInfo.isEarlierThan(targetDateInfo); daysCtr++ {
        dateInfo.addDate(0, 0, 1)

        if dateInfo.isWeekEnd() {
            fmt.Println("Day:", daysCtr, "WeekEnd")
            continue
        } else if dateInfo.isHoliday(holidays) {
            fmt.Println("Day:", daysCtr, "Holiday")
            continue
        } else if dateInfo.isLeaveDay(leaveDays) {
            fmt.Println("Day:", daysCtr, "LeaveDay")
            continue
        }

        if dateInfo.isWeekStart() {
            weekCtr++
        }
        fmt.Println("Day:", daysCtr, dateInfo.weekDay,
                    "WorkingDay:", workingDaysCtr,
                    "Week:", weekCtr,
                    "Date:", dateInfo.month, dateInfo.day, dateInfo.year)
        workingDaysCtr++
    }
}
