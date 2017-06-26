package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	//if mode is not device(file) then need to read from stdin and get md5sum
	if fi.Mode()&os.ModeDevice == 0 {
		fmt.Println("Calculate the md5sum of contents of file contents from stdin ")
		fmt.Printf("%x\n", MD5sum())
		//if mode is device then need to read data from device(file)
	} else if len(os.Args) > 1 {
		fmt.Println("Calculate the md5sum of contents of given file name ")
		fmt.Printf("%x\n", MD5SumFile(os.Args[1]))
	} else {
		//if no file name supplied then use own file name to return md5sum of it's own
		fmt.Println("Calculate the md5sum of contents of own file name ")
		fmt.Printf("%x\n", MD5SumFile("md5sum"))
	}

}

//Function to read from the Stdin and compute md5sum and return the same
func MD5sum() []byte {
	reader := bufio.NewReader(os.Stdin)
	data := make([]byte, 1024)
	h := md5.New()
	for {
		n, err := reader.Read(data)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		data = data[:n]
		if _, err := io.Copy(h, bytes.NewReader(data)); err != nil {
			log.Fatal(err)
		}
	}
	return h.Sum(nil)
}

//Function to compute the md5sum of given file name   and return the same
func MD5SumFile(file string) []byte {

	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return h.Sum(nil)
}
