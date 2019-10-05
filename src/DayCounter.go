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

func main() {
  currentTime := time.Now()
      currentTime = time.Date(2019, 9, 11, 0, 0, 0, 0, time.UTC) // hack due to incorrect time in this website
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

  fmt.Println("Current date:", month, day, year)
  fmt.Println("End date:", endMonth, endDay, endYear)
  fmt.Println()

  for daysCtr, workingDaysCtr := 1, 1; month.String() != endMonth || day != endDay || year != endYear; daysCtr++ {
    currentTime = currentTime.AddDate(0, 0, 1)
    month = currentTime.Month()
    day = currentTime.Day()
    year = currentTime.Year()
    weekDay = currentTime.Format("Monday")
    
    if weekDay == "Saturday" || weekDay == "Sunday" {
        fmt.Println("Day:", daysCtr, weekDay)
        continue
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

    fmt.Println("Day:", daysCtr, weekDay, "WorkingDay:", workingDaysCtr , ":", month, day, year)
    workingDaysCtr++
  }
}
