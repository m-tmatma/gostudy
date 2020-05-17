package main

import "fmt"
import "hash/crc32"
import "io"
import "os"

func getHash(filename string) (uint32, error){
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	h := crc32.NewIEEE()
	_, err = io.Copy(h, f)
	if err != nil {
		return 0, err
	}

	return h.Sum32(), nil
}

func main(){
	h := crc32.NewIEEE()
	h.Write([]byte("test"))


	v := h.Sum32()
	fmt.Println(v)

	fmt.Println(getHash("crc32.go"))
}