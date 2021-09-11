# Get year, month, day from time

yourbasic.org/golang

The [`Date`](https://golang.org/pkg/time/#Time.Date) function returns the year, month and day of a [`time.Time`](https://golang.org/pkg/time/#Time).

```
func (t Time) Date() (year int, month Month, day int)
```

In use:

```
year, month, day := time.Now().Date()
fmt.Println(year, month, day)      // For example 2009 November 10
fmt.Println(year, int(month), day) // For example 2009 11 10
```

You can also extract the information with seperate calls:

```
t := time.Now()
year := t.Year()   // type int
month := t.Month() // type time.Month
day := t.Day()     // type int
```

The [`time.Month`](https://golang.org/pkg/time/#Month) type specifies a month of the year (January = 1, …).

```
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
```


# How to get current timestamp

Use [`time.Now`](https://golang.org/pkg/time/#Now) and one of [`time.Unix`](https://golang.org/pkg/time/#Time.Unix) or [`time.UnixNano`](https://golang.org/pkg/time/#Time.UnixNano) to get a timestamp.

```
now := time.Now()      // current local time
sec := now.Unix()      // number of seconds since January 1, 1970 UTC
nsec := now.UnixNano() // number of nanoseconds since January 1, 1970 UTC

fmt.Println(now)  // time.Time
fmt.Println(sec)  // int64
fmt.Println(nsec) // int64
2009-11-10 23:00:00 +0000 UTC m=+0.000000000
1257894000
1257894000000000000
```



# How to find the day of week

The [`Weekday`](https://golang.org/pkg/time/#Time.Weekday) function returns returns the day of the week of a [`time.Time`](https://golang.org/pkg/time/#Time).

```
func (t Time) Weekday() Weekday
```

In use:

```
weekday := time.Now().Weekday()
fmt.Println(weekday)      // "Tuesday"
fmt.Println(int(weekday)) // "2"
```

## Type Weekday

The [`time.Weekday`](https://golang.org/pkg/time/#Weekday) type specifies a day of the week (Sunday = 0, …).

```
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
```



# Days between two dates

```
func main() {
    // The leap year 2016 had 366 days.
    t1 := Date(2016, 1, 1)
    t2 := Date(2017, 1, 1)
    days := t2.Sub(t1).Hours() / 24
    fmt.Println(days) // 366
}

func Date(year, month, day int) time.Time {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
```
