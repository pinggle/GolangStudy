package main

/**
	go build -o server main.go server.go
	./server

测试:
nc 127.0.0.1 8888

*/
func main() {
	server := NewServer("0.0.0.0", 8888)
	server.start()
}
