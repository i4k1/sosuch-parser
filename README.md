# Sosuch parser
Simple utility for parsing files by keywords on the 2ch.hk (sosuch) imageboard.

Inspired by [ValdikSS/endless-sosuch](https://github.com/ValdikSS/endless-sosuch).

## Build
The program uses only the standard golang library, so it is enough to simply [install the golang compiler](https://go.dev/dl) and build it into a binary file using the following commands:
```sh
go mod tidy
go build
```

**Build Status:**

![Build on Ubuntu](https://github.com/i4k1/sosuch-parser/actions/workflows/build-ubuntu.yml/badge.svg) ![Build on Windows](https://github.com/i4k1/sosuch-parser/actions/workflows/build-windows.yml/badge.svg)

## Usage
>[!WARNING]
>I am not responsible for the content of the site. 2ch.hk is an imageboard where almost anything can be posted.

If you run the program with the `--help` (or `-h`) flag, the program will give you something like this:
```
Usage of sosuch-parser:
  -boards string
        Which board to parse (default "b")
  -fileformats string
        What file formats to download     
  -keywords string
        Use keywords for parsing
  -path string
        In which directory to save files (default "src")
```
If you run the program without any flags, the program will start parsing all threads from the /b/ board and downloading them to the `src` directory. You can parse threads using regular expressions, for this use the `--keywords` flag (works poorly, will be fixed in the future). You can also use the program in your scripts for automation, or anything else. For example, you can use a script that runs the [mpv player](https://mpv.io) in parallel to parsing, to immediately play downloaded videos:
```sh
#!/bin/sh

# the program will search for webm-threads, and download webm files
./sosuch-parser --keywords="webm" --fileformats="webm" --path="webm" && mpv --shuffle webm/*.webm
```
Now you can watch WebM-threads *endlessly*!

## License
This software is licensed under the [Creative Commons Zero v1.0 Universal](LICENSE) license, which means it is in the public domain, which roughly means you can do whatever you want with it!
