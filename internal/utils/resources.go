package utils

type Semaphore chan struct{}

func NewSemaphore(cap int) Semaphore {
	return make(Semaphore, cap)
}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) Release() {
	<-s
}
