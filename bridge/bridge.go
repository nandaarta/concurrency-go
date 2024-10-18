package bridge

func Bridge(done <-chan any, chanStream <-chan (<-chan any)) <-chan any {
	bridgeCh := make(chan any)

	go func() {
		defer close(bridgeCh)

		for {
			var stream <-chan any

			select {
			case <-done:
				return
			case ch, ok := <-chanStream:
				if !ok {
					return
				}

				stream = ch
			}

			// read values off stream and add them to bridge stream
			// once the stream is closed we continue with the other
			// channels
			for value := range stream {
				select {
				case <-done:
				case bridgeCh <- value:
				}
			}
		}
	}()

	return bridgeCh
}

func Generate(values []any) <-chan (<-chan any) {
	channelStream := make(chan (<-chan any))

	go func() {
		defer close(channelStream)

		for _, v := range values {
			valueStream := make(chan any, 1)
			valueStream <- v
			close(valueStream)

			channelStream <- valueStream
		}
	}()

	return channelStream
}
