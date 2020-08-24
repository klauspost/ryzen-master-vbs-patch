$patterns = @(
    @{
        Desc = "Ryzen Master v.1.5 -> v2.2";
        Search = [Byte[]](0x44, 0x39, 0x6D, 0xA8, 0x0F, 0x84, 0xF7);
        Replace = [Byte[]](0x44, 0x39, 0x6D, 0xA8, 0x90, 0xe9, 0xf7)
    },
    @{
        Desc = "Ryzen Master v2.3 -> ?";
        Search = [Byte[]](0x44, 0x39, 0xad, 0xf8, 0, 0, 0, 0x0f, 0x84);
        Replace = [Byte[]](0x44, 0x39, 0xad, 0xf8, 0, 0, 0, 0x90, 0xe9)
    },
    @{
        Desc = "Ryzen Master Threadripper";
        Search = [Byte[]](0x00, 0x39, 0x7D, 0x90, 0x0F, 0x84, 0xE8, 0x00);
        Replace = [Byte[]](0x00, 0x39, 0x7D, 0x90, 0x90, 0xE9, 0xE8, 0x00)
    }
)

$exe = 'C:\Program Files\AMD\RyzenMaster\bin\AMD Ryzen Master.exe'

if (!(Test-Path $exe))
{
    Write-Host "'$exe' not found. Press any key to continue..."
    [void][System.Console]::ReadKey($true)
    
    return
}

Write-Host "Reading $exe"
$input = [System.IO.File]::ReadAllBytes($exe)

$output = $input

$i = 1

Foreach ($pattern in $patterns)
{
    $desc = $pattern.Desc
    $original = $pattern.Search
    $substitute = $pattern.Replace

    Write-Host "$i. Checking patch for $desc"

    $output = [Byte[]]("$output" -Replace "\b$original\b", "$substitute" -Split '\s+')

    $notChanged = [System.Linq.Enumerable]::SequenceEqual($input, $output)

    if ($notChanged)
    {
        $i++

        continue;
    }

    Write-Host "Writing to $exe"
    [System.IO.File]::WriteAllBytes($exe, $output)

    Write-Host "Done. Press any key to continue..."
    [void][System.Console]::ReadKey($true)

    return
}

Write-Host "FAILED: no byte sequence found. Press any key to continue..."
[void][System.Console]::ReadKey($true)
