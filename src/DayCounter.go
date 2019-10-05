// You can edit this code!
// Click here and start typing.
package main

import (
  "fmt"
  "time"
)

type Holiday struct {
    month string
    day int
}

type Leaveday struct {
    month string
    day int
}

func main() {
  
  currentTime := time.Now()
      currentTime = time.Date(2019, 9, 12, 0, 0, 0, 0, time.UTC) // hack due to incorrect time in this website
  month := currentTime.Month()
  day := currentTime.Day()
  year := currentTime.Year()
  weekDay := currentTime.Format("Monday") // currentTime.Weekday()

  endMonth := "December"
  endDay := 22
  endYear := 2019
  // endTime := time.Date(2019, 12, 22, 0, 0, 0, 0, time.UTC)
  
  holidays := []Holiday{
      Holiday{"November", 1},
      Holiday{"November", 2},
      Holiday{"November", 30},
      Holiday{"December", 8},
  }
  leavedays := []Leaveday{
      Leaveday{"November", 22},
      Leaveday{"December", 17},
      Leaveday{"December", 18},
      Leaveday{"December", 19},
      Leaveday{"December", 20},
  }

  fmt.Println("Current date:", month, day, year)
  fmt.Println("End date:", endMonth, endDay, endYear)
  fmt.Println()

  for daysCtr, workingDaysCtr, weekCtr := 1, 1, 1; month.String() != endMonth || day != endDay || year != endYear; daysCtr++ {
    currentTime = currentTime.AddDate(0, 0, 1)
    month = currentTime.Month()
    day = currentTime.Day()
    year = currentTime.Year()
    weekDay = currentTime.Format("Monday")
    
    if weekDay == "Saturday" || weekDay == "Sunday" {
        fmt.Println("Day:", daysCtr, weekDay)
        continue
    } else if weekDay == "Monday" {
        weekCtr++
    }

    isHoliday := false
    for _, currentHoliday := range holidays {
        if month.String() == currentHoliday.month && day == currentHoliday.day {
            isHoliday = true
            break
        }
    }
    if isHoliday {
        fmt.Println("Day:", daysCtr, "Holiday")
        continue
    }


    isLeaveday := false
    for _, currentLeaveday := range leavedays {
        if month.String() == currentLeaveday.month && day == currentLeaveday.day {
            isLeaveday = true
            break
        }
    }
    if isLeaveday {
        fmt.Println("Day:", daysCtr, "Leaveday")
        continue
    }

    fmt.Println("Day:", daysCtr, weekDay, "WorkingDay:", workingDaysCtr , "Week:", weekCtr, ":", month, day, year)
    workingDaysCtr++
  }
}
