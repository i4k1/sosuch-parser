# sosuch-parser
Simple utility for parsing files by keywords on the 2ch.hk (sosuch) imageboard.

Inspired by [ValdikSS/endless-sosuch](https://github.com/ValdikSS/endless-sosuch).

## Build
The program uses only the standard golang library, so it is enough to simply [install the golang compiler](https://go.dev/dl) and build it into a binary with the command: `go build main.go`.

## Usage
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

start_mpv() {
    sleep 10 # we wait until the program downloads at least some files
    mpv --shuffle webm/*.webm
}

# the program will search for webm-threads, and download webm files
./sosuch-parser --keywords="webm" --fileformats="webm" --path="webm" & start_mpv
```
Now you can watch WebM-threads *endlessly*!
