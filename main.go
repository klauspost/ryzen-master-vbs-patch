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

package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var patterns = []struct {
	desc            string
	search, replace []byte
}{
	{
		desc:    "Ryzen Master v.1.5 -> v2.2",
		search:  []byte{0x44, 0x39, 0x6D, 0xA8, 0x0F, 0x84, 0xF7},
		replace: []byte{0x44, 0x39, 0x6D, 0xA8, 0x90, 0xe9, 0xf7},
	},
	{
		desc:    "Ryzen Master v2.3 -> ?",
		search:  []byte{0x44, 0x39, 0xad, 0xf8, 0, 0, 0, 0x0f, 0x84},
		replace: []byte{0x44, 0x39, 0xad, 0xf8, 0, 0, 0, 0x90, 0xe9},
	},
	{
		desc:    "Ryzen Master Threadripper",
		search:  []byte{0x00, 0x39, 0x7D, 0x90, 0x0F, 0x84, 0xE8, 0x00},
		replace: []byte{0x00, 0x39, 0x7D, 0x90, 0x90, 0xE9, 0xE8, 0x00},
	},
}

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
		args = []string{"AMD Ryzen Master.exe"}
	}

	for _, f := range args {
		dir, file := filepath.Split(f)
		err := patch(filepath.Join(dir, file), filepath.Join(dir, *prefix+file))
		switch err {
		case nil:
		default:
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

var errCannotPatch = errors.New("no byte sequence not found")

func patch(in string, out string) error {
	fmt.Printf("Reading %q\n", in)
	b, err := ioutil.ReadFile(in)
	if err != nil {
		return err
	}
	for i, pattern := range patterns {
		c := bytes.Count(b, pattern.search)
		fmt.Printf("%d. Checking patch for %q. ", i, pattern.desc)
		fmt.Println("Matching byte sequences:", c, "(should be 1)")
		if c != 1 {
			continue
		}
		b = bytes.Replace(b, pattern.search, pattern.replace, -1)
		fmt.Printf("Writing to %q\n", out)
		return ioutil.WriteFile(out, b, os.ModePerm)
	}
	return errCannotPatch
}
