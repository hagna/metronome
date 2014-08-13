package main

import (
	"code.google.com/p/go.tools/godoc/vfs"
	"bitbucket.org/mjl/asset"
	"log"
_	"encoding/binary"
	"flag"
	"io/ioutil"
	"os"
_	"fmt"
_	"io"
	"net"
)


const (
	bits = 16
	rate = 44100
)

var tempo = flag.Int("t", 80, "tempo in beats per minute")

func handleConnection(conn net.Conn) {
	
}

func main() {
	flag.Parse()	
	fs := asset.Fs()
	if err := asset.Error(); err != nil {
		log.Println(err)
		fs = vfs.OS(".")
	}

	f, err := fs.Open("/sounds/click.raw")
	if err != nil {
		log.Fatal(err)
	}
	dat, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	fsb := float64((rate * 60) / float64(*tempo))
	log.Println("Samples per beat", fsb)
	spb := int(fsb)
	if len(dat) % 2 != 0 {
		log.Fatal("raw file contains an odd number of bytes")
	}
	pad := spb - len(dat) / 2
	log.Println("len dat", len(dat), "len beat", spb, "extra is", pad)

	for i := 0; i < pad + spb; i++ {
		dat = append(dat, []byte{0,0}...)
	} 

	fourbars := []byte{}
	fourbars = append(fourbars, dat...)
	fourbars = append(fourbars, dat...)
	fourbars = append(fourbars, dat...)
	fourbars = append(fourbars, dat...)


	go func() {
		ln, err := net.Listen("tcp", ":8080")
		if err != nil {
			// handle error
		}
		for {
			conn, err := ln.Accept()
			if err != nil {
				// handle error
				continue
			}
			go handleConnection(conn)
		}
	}()
	for {
	/*	for _, d := range dat {
			//fmt.Printf("%c", d) %c is unicode char not binary
			binary.Write(os.Stdout, binary.LittleEndian, byte(d))
		}
	*/
		_, _ = os.Stdout.Write(fourbars)
	 
	}
}
	
