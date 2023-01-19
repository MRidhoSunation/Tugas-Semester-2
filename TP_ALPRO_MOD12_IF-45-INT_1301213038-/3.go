package main

func CariSekuensial(tab [9999]int, v int) int {
	//Returns the location where the value v is in the tab array; or -1 if v is not found
	for idx, val := range tab {
		if val == v {
			return idx
		}
	}
	return -1
}
