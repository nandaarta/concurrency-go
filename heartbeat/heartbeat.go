package heartbeat

import (
	"fmt"
	"time"
)

type Beat string

const SomeBeat Beat = "jedag-jedug"

func DoSomething(done <-chan any, in []string, interval time.Duration) ([]string, []Beat) {
	var (
		res   []string
		beats []Beat
	)

	pulseResultCh, heartbeatCh := doSendPulse(done, GeneratePulse(done, in), interval)

	for {
		select {
		case <-done:
			return res, beats
		case beat, ok := <-heartbeatCh:
			if !ok {
				return res, beats
			}

			beats = append(beats, beat)
		case val, ok := <-pulseResultCh:
			if !ok {
				return res, beats
			}

			res = append(res, val)
		}
	}

	// return res, beats
}

func GeneratePulse(done <-chan any, in []string) <-chan string {
	chStream := make(chan string)

	go func() {
		defer close(chStream)

		for _, val := range in {
			select {
			case <-done:
				return
			case chStream <- val:
				// do something if you want
			}
		}
	}()

	return chStream
}

func doSendPulse(done <-chan any, dataStream <-chan string, interval time.Duration) (<-chan string, <-chan Beat) {
	heartBeat := make(chan Beat)
	result := make(chan string)

	go func() {
		defer close(heartBeat)
		defer close(result)

		pulse := time.NewTicker(interval)

		// define a closure to send the pulse
		sendPulse := func(done <-chan any) {
			select {
			case <-done:
				return
			case heartBeat <- SomeBeat:
			default:
			}
		}

		sendPulseResult := func(s string) {
			for {
				select {
				case <-done:
					return
				case <-pulse.C:
					fmt.Println("Send Pulse here")
					sendPulse(done)
				case result <- s:
					fmt.Println("Send string to result channel")
					return
				}
			}
		}

		for {
			select {
			case <-done:
				return
			case <-pulse.C:
				sendPulse(done)
			case val, ok := <-dataStream:
				if !ok {
					return
				}

				time.Sleep(2 * time.Millisecond)
				sendPulseResult(val)
			}
		}

	}()

	return result, heartBeat
}
