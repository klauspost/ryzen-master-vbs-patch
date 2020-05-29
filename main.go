package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func init() {
	flag.Usage = func() {
		fmt.Println("Usage: ryzen-master-vbs-patch [-p=patched] \"AMD Ryzen Master.exe\"")
		flag.PrintDefaults()
	}
}

var prefix = flag.String("p", "patched-", `Specify prefix for output file. Set to "", overwrite input.`)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}

	for _, f := range args {
		dir, file := filepath.Split(f)
		err := patch(filepath.Join(dir, file), filepath.Join(dir, *prefix+file))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

var search = []byte{0x44, 0x39, 0x6D, 0xA8, 0x0F, 0x84, 0xF7}
var replace = []byte{0x44, 0x39, 0x6D, 0xA8, 0x90, 0xe9, 0xf7}

func patch(in string, out string) error {
	fmt.Println("Patching", in)
	b, err := ioutil.ReadFile(in)
	if err != nil {
		return err
	}
	c := bytes.Count(b, search)
	fmt.Println("Matching byte sequences:", c, "(should be 1)")
	if c != 1 {
		fmt.Println("Skipping file...")
		return nil
	}
	b = bytes.Replace(b, search, replace, -1)
	fmt.Println("Writing to", out)
	err = ioutil.WriteFile(out, b, 0)
	if err != nil {
		return err
	}
	return nil
}

/*
The MIT License (MIT)

Copyright (c) 2020 Klaus Post

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
