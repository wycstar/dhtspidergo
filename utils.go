package main

func getNeighbor(a []byte, b []byte) []byte {
	return append(a[:10], b[10:]...)
}