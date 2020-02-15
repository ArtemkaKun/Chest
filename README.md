# Chest

![Stable version](https://img.shields.io/badge/Stable-v0.0.1-green)
![GitHub All Releases](https://img.shields.io/github/downloads/artemkakun/chest/total?color=green&label=Downloads)

"Chest" is a program, created in Golang, that allows administrators to backup their Minecraft servers automatically.

## What the Chest can do?
:white_check_mark: Automatically stop the Minecraft server

:white_check_mark: Pack server files to .7z archive (actually it pack it to .tar archive and then to .7z)

:white_check_mark: Upload a backup archive to Google Drive

:white_check_mark: Automatically start the Minecraft server

:white_check_mark: Delete old backup archives from Google Drive

## How it do this?
Main program was writed in Go, but for work it use Bash scripts (**Chest can create these scripts automatically** when you start the tool with flags or you can create scripts manually). 

For uploading backups to cloud storage it use APIs (for example Google Drive API), but you also need to generate some additional files to perform this action (please, check the wiki for instructions).

## What servers the Chest tool suppport?
- [x] Linux servers
- [ ] Windows servers

## What cloud storages the Chest tool suppport?
- [x] Google Drive
- [ ] Yandex Disk
- [ ] Mega
- [ ] OneDrive
- [ ] Dropbox
- [ ] FTP clouds

## How to start use the Chest tool?
[First start](https://github.com/ArtemkaKun/Chest/wiki/First-start)

[How to build](https://github.com/ArtemkaKun/Chest/wiki/How-to-build)
