package util

type ForSelectCoordinator struct {
	In      chan interface{}
	Out     chan interface{}
	Process func(interface{}) interface{}
	Cancel  chan struct{}
}

func (f *ForSelectCoordinator) Start() {
	for {
		select {
		case i := <-f.In:
			o := f.Process(i)
			f.Out <- o
		case <-f.Cancel:
			return
		}
	}
}

// func ForSelectBytesChan(in, out chan []byte, cc chan struct{}, f BytesToBytes) {
// 	for {
// 		select {
// 		case m := <-in:
// 			fmt.Println("M: ", string(m))
// 			out <- f(m)
// 		case <-cc:
// 			return
// 		}
// 	}
// }
