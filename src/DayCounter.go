package main

import (
  "fmt"
)

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
