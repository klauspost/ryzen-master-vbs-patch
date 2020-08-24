# ryzen-master-vbs-patch Powershell version

This version does not require you to compile or run binary code and the code you are running can be easily inspected.

AMD Ryzen Master VBS patcher allows AMD Ryzen Master to run while Hyper-V Virtualization is enabled.


Thanks to [@TOM_RUS](https://twitter.com/TOM_RUS/status/1204867886197755904) for finding this.

This patch disables the VBS check on startup. It does NOT fix what might be the cause of this being disabled.

This MAY break functionality and it MAY stop working on future versions.

This has been confirmed to work on the versions listed below so versions inbetween are likely fine as well. 
The exact version does not have to match, but if AMD changes the code for detection it will likely break.

Confirmed versions:

 * v2.1.0.1424 (2019)
 * v2.1.1.1472 (2020)
 * v2.2.0.1543 (2020)
 * v2.3.0.1591 (2020)

Threadripper patch was supplied by [@neoKushan](https://github.com/neoKushan).

# Running

Various parts of this will require administrator access.

* Download `ryzen-master-vbs-patch.bat` and `ryzen-master-vbs-patch.ps1`.
* Run `ryzen-master-vbs-patch.bat`.
