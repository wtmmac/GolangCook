package main

import (
	"fmt"
	"time"
)

func main() {
	for attempt := 0; true; attempt++ {
		//log.Trace().Msgf("grpc client %v dial attempt %v time",g.name, attempt)
		if attempt != 0 {
			fmt.Println(backoff(attempt, time.Millisecond*10, time.Second*10))
		}

		if attempt > 1000000000000000 {
			break
		}

	}
}

func backoff(attempt int, min time.Duration, max time.Duration) time.Duration {
	d := time.Duration(attempt*attempt) * min
	if d > max {
		d = max
	}
	return d
}
