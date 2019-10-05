DayCounter

## PURPOSE
Count the number of days and working days between two dates. It is possible to enter holidays and work leave days to consider for the count. The application of this program is wide and endless, such as counting the exact days before the target delivery of products to customers, counting remaining days before deadline of submission of a school project, counting days spent in the hospital, or even counting remaining days upon resignation at work!

## REQUIREMENTS
1. GNU/Linux environment
    * Either:
        1. Full operating system
        2. Virtual machine (as a guest OS)
        3. Cygwin
        4. etc.
2. Git (optional if you just prefer to do a download and extract from GitHub)
    ~~~
    sudo apt install git
    ~~~
3. Go
    ~~~
    sudo snap install go --classic
    ~~~

## BUILDING
~~~
git clone https://github.com/nponcian/DayCounter.git
cd DayCounter/src
go build
~~~

## RUNNING
~~~
cd DayCounter/src
./src
~~~
Example
~~~
nponcian@nponcian-VirtualBox-Ubuntu:~/Documents/Programs$ cd DayCounter/src/
nponcian@nponcian-VirtualBox-Ubuntu:~/Documents/Programs/DayCounter/src$ go build
nponcian@nponcian-VirtualBox-Ubuntu:~/Documents/Programs/DayCounter/src$ ./src 
StartDateInfo: {{0 63705571200 <nil>} 2019 10 2 Wednesday}
TargetDateInfo: {{0 63712569600 <nil>} 2019 12 22 Sunday}

Day: 1 Thursday WorkingDay: 1 Week: 1 Date: October 3 2019
Day: 2 Friday WorkingDay: 2 Week: 1 Date: October 4 2019
Day: 3 WeekEnd
Day: 4 WeekEnd
Day: 5 Monday WorkingDay: 3 Week: 2 Date: October 7 2019
Day: 6 Tuesday WorkingDay: 4 Week: 2 Date: October 8 2019
Day: 7 Wednesday WorkingDay: 5 Week: 2 Date: October 9 2019
Day: 8 Thursday WorkingDay: 6 Week: 2 Date: October 10 2019
Day: 9 Friday WorkingDay: 7 Week: 2 Date: October 11 2019
Day: 10 WeekEnd
Day: 11 WeekEnd
Day: 12 Monday WorkingDay: 8 Week: 3 Date: October 14 2019
Day: 13 Tuesday WorkingDay: 9 Week: 3 Date: October 15 2019
Day: 14 Wednesday WorkingDay: 10 Week: 3 Date: October 16 2019
Day: 15 Thursday WorkingDay: 11 Week: 3 Date: October 17 2019
Day: 16 Friday WorkingDay: 12 Week: 3 Date: October 18 2019
Day: 17 WeekEnd
Day: 18 WeekEnd
Day: 19 Monday WorkingDay: 13 Week: 4 Date: October 21 2019
Day: 20 Tuesday WorkingDay: 14 Week: 4 Date: October 22 2019
Day: 21 Wednesday WorkingDay: 15 Week: 4 Date: October 23 2019
Day: 22 Thursday WorkingDay: 16 Week: 4 Date: October 24 2019
Day: 23 Friday WorkingDay: 17 Week: 4 Date: October 25 2019
Day: 24 WeekEnd
Day: 25 WeekEnd
Day: 26 Monday WorkingDay: 18 Week: 5 Date: October 28 2019
Day: 27 Tuesday WorkingDay: 19 Week: 5 Date: October 29 2019
Day: 28 Wednesday WorkingDay: 20 Week: 5 Date: October 30 2019
Day: 29 Thursday WorkingDay: 21 Week: 5 Date: October 31 2019
Day: 30 Holiday
Day: 31 WeekEnd
Day: 32 WeekEnd
Day: 33 Monday WorkingDay: 22 Week: 6 Date: November 4 2019
Day: 34 Tuesday WorkingDay: 23 Week: 6 Date: November 5 2019
Day: 35 Wednesday WorkingDay: 24 Week: 6 Date: November 6 2019
Day: 36 Thursday WorkingDay: 25 Week: 6 Date: November 7 2019
Day: 37 Friday WorkingDay: 26 Week: 6 Date: November 8 2019
Day: 38 WeekEnd
Day: 39 WeekEnd
Day: 40 Monday WorkingDay: 27 Week: 7 Date: November 11 2019
Day: 41 Tuesday WorkingDay: 28 Week: 7 Date: November 12 2019
Day: 42 Wednesday WorkingDay: 29 Week: 7 Date: November 13 2019
Day: 43 Thursday WorkingDay: 30 Week: 7 Date: November 14 2019
Day: 44 Friday WorkingDay: 31 Week: 7 Date: November 15 2019
Day: 45 WeekEnd
Day: 46 WeekEnd
Day: 47 Monday WorkingDay: 32 Week: 8 Date: November 18 2019
Day: 48 Tuesday WorkingDay: 33 Week: 8 Date: November 19 2019
Day: 49 Wednesday WorkingDay: 34 Week: 8 Date: November 20 2019
Day: 50 Thursday WorkingDay: 35 Week: 8 Date: November 21 2019
Day: 51 LeaveDay
Day: 52 WeekEnd
Day: 53 WeekEnd
Day: 54 Monday WorkingDay: 36 Week: 9 Date: November 25 2019
Day: 55 Tuesday WorkingDay: 37 Week: 9 Date: November 26 2019
Day: 56 Wednesday WorkingDay: 38 Week: 9 Date: November 27 2019
Day: 57 Thursday WorkingDay: 39 Week: 9 Date: November 28 2019
Day: 58 Friday WorkingDay: 40 Week: 9 Date: November 29 2019
Day: 59 WeekEnd
Day: 60 WeekEnd
Day: 61 Monday WorkingDay: 41 Week: 10 Date: December 2 2019
Day: 62 Tuesday WorkingDay: 42 Week: 10 Date: December 3 2019
Day: 63 Wednesday WorkingDay: 43 Week: 10 Date: December 4 2019
Day: 64 Thursday WorkingDay: 44 Week: 10 Date: December 5 2019
Day: 65 Friday WorkingDay: 45 Week: 10 Date: December 6 2019
Day: 66 WeekEnd
Day: 67 WeekEnd
Day: 68 Monday WorkingDay: 46 Week: 11 Date: December 9 2019
Day: 69 Tuesday WorkingDay: 47 Week: 11 Date: December 10 2019
Day: 70 Wednesday WorkingDay: 48 Week: 11 Date: December 11 2019
Day: 71 Thursday WorkingDay: 49 Week: 11 Date: December 12 2019
Day: 72 Friday WorkingDay: 50 Week: 11 Date: December 13 2019
Day: 73 WeekEnd
Day: 74 WeekEnd
Day: 75 Monday WorkingDay: 51 Week: 12 Date: December 16 2019
Day: 76 LeaveDay
Day: 77 LeaveDay
Day: 78 LeaveDay
Day: 79 LeaveDay
Day: 80 WeekEnd
Day: 81 WeekEnd
nponcian@nponcian-VirtualBox-Ubuntu:~/Documents/Programs/DayCounter/src$ 
~~~
