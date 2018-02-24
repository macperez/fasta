package io

import "log"

/*
Check is used to check errors
*/
func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
