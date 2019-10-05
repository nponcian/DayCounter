package main

type Holiday struct {
    month string
    day int
}

func getHolidays() []Holiday {
    // <MODIFY THIS LIST>
    return []Holiday {
        Holiday{"November", 1},
        Holiday{"November", 2},
        Holiday{"November", 30},
        Holiday{"December", 8},
    }
    // </MODIFY THIS LIST>
}
