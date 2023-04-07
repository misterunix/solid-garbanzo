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

	// Start needs to be assigned when the process starts
	start := time.Now()                                                   // get current time
	stu := int64(time.Nanosecond) * start.UnixNano() / int64(time.Second) // convert to unix time in seconds

	// end needs to be assigned when time limit is assigned.
	end := start.Add(time.Minute * 5)                                   // add end time in minutes to the current time
	etu := int64(time.Nanosecond) * end.UnixNano() / int64(time.Second) // convert to unix time in seconds

	ez := float64(etu - stu) // end time - start time for the timer

	// test TimeStr - works
	/*
		left := time.Since(end)
		s := TimeStr(int(left.Seconds()))
		fmt.Println(left, s)
		fmt.Println("------------")
	*/

	//tu := int64(time.Nanosecond) * timeT.UnixNano() / int64(time.Millisecond)

	for i := 0; i < 300; i++ {

		// needs to be set when checking the time
		n := time.Now()                                                   // get current time
		ntu := int64(time.Nanosecond) * n.UnixNano() / int64(time.Second) // convert to unix time in seconds

		nz := float64(ntu - stu) // current time - start time for the timer
		tt := nz / ez
		rt := tt * 100.0 // ratio of current time to end time
		// now I have a percent completion

		//fmt.Println("ez:", ez, "nz:", nz, "rt:", rt)

		fmt.Printf("ratio: %5.2f %6.3f\n", rt, tt)

		time.Sleep(1 * time.Second)
	}

}
