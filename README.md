# ryzen-master-vbs-patch

AMD Ryzen Master VBS patcher allows AMD Ryzen Master to run while Virtualization is enabled.


Thanks to [@TOM_RUS](https://twitter.com/TOM_RUS/status/1204867886197755904) for finding this.

This patch disables the VBS check on startup. It does NOT fix what might be the cause of this being disabled.

This MAY break functionality and it MAY stop working on future versions.

This has been confirmed to work on v2.1.1.1472 (2020) and matches the change described above for v2.1.0.1424 (2019) so versions inbetween are likely fine as well. 
The exact version does not have to match, but if AMD changes the code for detection it will likely break.

Note that on Threadripper CPUs the actually installed version might be a v1.5.x for which there is no patch currently. 
The version in the "About" page in the software should reflect the actual version. 

# Running

Various parts of this will require administrator access.

* Download and unzip binary from [releases](https://github.com/klauspost/ryzen-master-vbs-patch/releases).
* Copy `AMD Ryzen Master.exe` from `c:\Program Files\AMD\RyzenMaster\bin\` or where your software is installed.
* In Explorer Drag & Drop `AMD Ryzen Master.exe` in your new location on to `ryzen-master-vbs-patch.exe`
* If successful `patched-AMD Ryzen Master.exe` should now be created.
* Copy this back to the `RyzenMaster` folder.
* Rename your existing `AMD Ryzen Master.exe` to `AMD Ryzen Master BACKUP.exe` or similar.
* Rename `patched-AMD Ryzen Master.exe` to `AMD Ryzen Master.exe`.

From the commandline, the syntax is:

```
 Usage: ryzen-master-vbs-patch [-p=patched] "AMD Ryzen Master.exe"
   -p string
         Specify prefix for output file. Set to "", overwrite input. (default "patched-")
```

# Building 

Requires Go SDK  installed.
 
`go get github.com/klauspost/ryzen-master-vbs-patch`
`go install github.com/klauspost/ryzen-master-vbs-patch`

# License

```
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
```
