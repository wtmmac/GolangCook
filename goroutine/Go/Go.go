package main

import "log"

func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
			}
		}()

		f()
	}()
}
