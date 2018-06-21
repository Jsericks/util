package util

type BytesToBytes func([]byte) []byte

func ForSelectBytesChan(in, out chan []byte, cc chan struct{}, f BytesToBytes) {
	for {
		select {
		case m := <-in:
			out <- f(m)
		case <-cc:
			return
		}
	}
}
