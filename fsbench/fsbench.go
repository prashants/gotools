package main

import (
	"fmt"
	"os"
)

import "./random"

var err error

func main() {
	dname := random.GenDirName()
	err = os.Mkdir(dname, 0766)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory : %s\n", dname)
		return
	}

	createFiles(1000, 100, 1, dname)
}

func createFiles(small int, medium int, large int, dirname string) {
	var fname string
	var fpath string
	var fsize int
	var datablock [524288000]byte
	var data []byte
	var c int

	randFile, err := os.Open("/dev/urandom")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening /dev/urandom file\n")
		return
	}
	defer randFile.Close()

	/* creating small files */
	fmt.Println("Creating small files...")
	for c = 1; c <= small; c++ {
		fname = random.GenFileName()
		fpath = dirname + "/" + fname
		fsize = random.GetSmallSize()
		outputFile, err := os.OpenFile(fpath, os.O_RDWR | os.O_CREATE, 0766)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file : %s (%s)\n", fpath, err)
			continue
		}
		fmt.Fprintf(os.Stdout, "Created file : %s (%d)\n", fpath, fsize)
		data = datablock[0:fsize]
		randFile.Read(data)
		outputFile.Write(data)
		outputFile.Close()
	}

	/* creating medium files */
	fmt.Println("Creating medium files...")
	for c = 1; c <= medium; c++ {
		fname = random.GenFileName()
		fpath = dirname + "/" + fname
		fsize = random.GetMediumSize()
		outputFile, err := os.OpenFile(fpath, os.O_RDWR | os.O_CREATE, 0766)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file : %s (%s)\n", fpath, err)
			continue
		}
		fmt.Fprintf(os.Stdout, "Created file : %s (%d)\n", fpath, fsize)
		data = datablock[0:fsize]
		randFile.Read(data)
		outputFile.Write(data)
		outputFile.Close()
	}

	/* creating large files */
	fmt.Println("Creating large files...")
	for c = 1; c <= large; c++ {
		fname = random.GenFileName()
		fpath = dirname + "/" + fname
		fsize = random.GetMediumSize()
		outputFile, err := os.OpenFile(fpath, os.O_RDWR | os.O_CREATE, 0766)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file : %s (%s)\n", fpath, err)
			continue
		}
		fmt.Fprintf(os.Stdout, "Created file : %s (%d)\n", fpath, fsize)
		data = datablock[0:fsize]
		randFile.Read(data)
		outputFile.Write(data)
		outputFile.Close()
	}
}
