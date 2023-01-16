package cpu

import (
	"runtime"
	"time"
)

func RunCPULoad(coresCount int, timeSeconds int, percentage int) {
	runtime.GOMAXPROCS(coresCount)

	// second     ,s  * 1
	// millisecond,ms * 1000
	// microsecond,μs * 1000 * 1000
	// nanosecond ,ns * 1000 * 1000 * 1000

	// every loop : run + sleep = 1 unit

	// 1 unit = 1 ms
	unitHundresOfMicrosecond := 1000
	runMicrosecond := unitHundresOfMicrosecond * percentage
	sleepMicrosecond := unitHundresOfMicrosecond*100 - runMicrosecond
	for i := 0; i < coresCount; i++ {
		go func() {
			runtime.LockOSThread()
			// endless loop
			for {
				begin := time.Now()
				for {
					// run 100%

					if time.Since(begin).Milliseconds() > int64(runMicrosecond) {
						break
					}
				}
				// sleep
				time.Sleep(time.Duration(sleepMicrosecond) * time.Microsecond)
			}
		}()
	}
	// how long
	time.Sleep(time.Duration(timeSeconds) * time.Second)
}
