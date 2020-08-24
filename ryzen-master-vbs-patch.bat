@echo off

PowerShell -NoProfile -ExecutionPolicy Bypass -Command "& {Start-Process PowerShell -ArgumentList '-NoProfile -ExecutionPolicy Bypass -File \"%~dp0ryzen-master-vbs-patch.ps1\"' -Verb RunAs}"
