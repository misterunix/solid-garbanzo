package main

import (
	"fmt"
	"math"
	"time"
)

// Convert seconds to days, hours, minutes and seconds
func TimeStr(sec int) (res string) {
	ds, sec := sec/86400, sec%86400
	hrs, sec := sec/3600, sec%3600
	mins, sec := sec/60, sec%60
	if sec < 0 {
		res += "-" // add a negative if seconds is negative
	}
	ds = int(math.Abs(float64(ds)))     // make sure it's positive
	hrs = int(math.Abs(float64(hrs)))   // make sure it's positive
	mins = int(math.Abs(float64(mins))) // make sure it's positive
	sec = int(math.Abs(float64(sec)))   // make sure it's positive
	res += fmt.Sprintf("%02dd:", ds)    // days
	res += fmt.Sprintf("%02dh:", hrs)   // hours
	res += fmt.Sprintf("%02dm:", mins)  // minutes
	res += fmt.Sprintf("%02ds", sec)    // seconds
	return
}

func main() {

	start := time.Now()                                                   // get current time
	stu := int64(time.Nanosecond) * start.UnixNano() / int64(time.Second) // convert to unix time in seconds

	end := start.Add(time.Minute * 5)                                   // add end time in minutes to the current time
	etu := int64(time.Nanosecond) * end.UnixNano() / int64(time.Second) // convert to unix time in seconds

	// test TimeStr - works
	/*
		left := time.Since(end)
		s := TimeStr(int(left.Seconds()))
		fmt.Println(left, s)
		fmt.Println("------------")
	*/

	//tu := int64(time.Nanosecond) * timeT.UnixNano() / int64(time.Millisecond)

	for i := 0; i < 300; i++ {
		n := time.Now()                                                   // get current time
		ntu := int64(time.Nanosecond) * n.UnixNano() / int64(time.Second) // convert to unix time in seconds

		ez := float64(etu - stu) // end time - start time for the timer
		nz := float64(ntu - stu) // current time - start time for the timer
		rt := nz / ez            // ratio of current time to end time

		fmt.Println("ez:", ez, "nz:", nz, "rt:", rt)

		fmt.Println("ratio:", rt)

		time.Sleep(1 * time.Second)
	}

}
