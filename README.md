# Chest

![Stable version](https://img.shields.io/badge/Stable-v0.0.1-green)
![GitHub All Releases](https://img.shields.io/github/downloads/artemkakun/chest/total?color=green&label=Downloads)
![CircleCI](https://img.shields.io/circleci/build/github/ArtemkaKun/Chest)

"Chest" is a program, created in Golang, that allows administrators to back up their Minecraft servers automatically.

## What Chest can do?
✅ Automatically stop the Minecraft server (only when you use *screen* tool for run your server)

✅ Pack server files to .7z archive (actually it pack it to .tar archive and then to .7z)

✅ Upload a backup archive to Google Drive

✅ Automatically start the Minecraft server (only when you use *screen* tool for run your server)

✅ Delete old backup archives from Google Drive

## How it does this?
The main program was written in Go, but for work it uses Bash scripts (**Chest can create these scripts automatically** when you start the tool with flags or you can create scripts manually). 

For uploading backups to cloud storage it uses APIs (for example Google Drive API), but you also need to generate some additional files to perform this action (please, check the wiki for instructions).

## What servers Chest tool support?
- [x] Linux servers
- [ ] Windows servers

## What cloud storages Chest tool suppport?
- [x] Google Drive
- [ ] Yandex Disk
- [ ] Mega
- [ ] OneDrive
- [ ] Dropbox
- [ ] FTP clouds

## How to start using Chest tool?
[First start](https://github.com/ArtemkaKun/Chest/wiki/First-start)

[How to build](https://github.com/ArtemkaKun/Chest/wiki/How-to-build)
