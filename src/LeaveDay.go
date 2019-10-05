package main

type LeaveDay struct {
    month string
    day int
}

func getLeaveDays() []LeaveDay {
    // <MODIFY THIS LIST>
    return []LeaveDay {
        LeaveDay{"November", 22},
        LeaveDay{"December", 17},
        LeaveDay{"December", 18},
        LeaveDay{"December", 19},
        LeaveDay{"December", 20},
    }
    // </MODIFY THIS LIST>
}
