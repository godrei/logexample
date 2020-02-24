package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if err != nil && err != io.EOF {
			panic(err)
		}
	}
	log.Println("Bytes:", nBytes, "Chunks:", nChunks)
}
