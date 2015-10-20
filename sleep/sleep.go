package sleep

import "time"

func Sleep(sTime int) {
  select {
  case <-time.After(time.Duration(sTime) * time.Second):
  }
}
