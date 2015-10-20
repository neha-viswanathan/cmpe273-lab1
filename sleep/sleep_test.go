package sleep

import (
  "time"
  "testing"
)

const (
    RFC822 = "02 Jan 06 15:04 MST"
)
var time1 int64
var time2 int64

func sleepTime(i int) {
	now1 := time.Now().Format(RFC822)
	start1, _ := time.Parse(RFC822, now1)
	time.Sleep(time.Duration(i)* time.Second)
	time1 = int64(time.Since(start1)/time.Second)

}

func newSleep(j int )  {
	now2 := time.Now().Format(RFC822)
	start2, _ := time.Parse(RFC822, now2)
	Sleep(j)
	time2 = int64(time.Since(start2)/time.Second)
	//return duration2
}

func TestSleep(t *testing.T) {

	for k:=1; k<3; k++{
		go sleepTime(k)
		go newSleep(k)

		if time2 != time1{
			t.Error("Error in new sleep function ", time1, time2)
		}
	}

}
