package panics

import "log"

//HandlePanic will recover from detected panics
func HandlePanic(loc string) {
	if err := recover(); err != nil {
		log.Printf("%s recovered from panic, %v", loc, err)
	}
}

//ConcurrentHandlePanic will be used in go routine to recover from detected panics
func ConcurrentHandlePanic(loc string, handlerFn func()) {
	defer HandlePanic(loc)
	handlerFn()
}
