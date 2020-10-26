# Wallpaper downloader

## Table of Contents

  - [Inspiration](#inspiration)
  - [Advantages](#advantages)
  - [Speed comparison](#speed-comparison)
  - [Flags or Arguments](#flags-or-arguments)
  - [Installation](#installation)
    - [GNU/Linux, Unix or MacOS](#gnulinux-unix-or-macos)
    - [Windows](#windows)
  - [Customization](#customization)

## Inspiration

This project is inspired by the wonderful people of [Reddit](https://www.reddit.com/) and the [reddit-wallpaper-downloader](https://github.com/mrsorensen/reddit-wallpaper-downloader) project. Though the above mentioned project downloads images from Reddit well, its written in Python which tends to make it slow. This project however is written in [Go lang](https://golang.org/) which is an open source programming language made by Google. This helps programs run at lightning speeds and makes multithreading a breeze.

## Advantages

- Written in Go
- Lightning speed
- Multithreaded for even better speeds
- Fetch images based on Top of the Day, Month, Week, Year or even All time
- Choose subreddit to fetch images from

## Speed comparison

| Number of Photos |   Go   |   Python    |
| :--------------: | :----: | :---------: |
|        10        | 10secs |   47secs    |
|       100        | 42secs | 5min 28secs |

- Golang was set to run on 4 threads, which is the default
- The benchmarking was NOT scientific by any means. However you are welcome to try both projects and see how Go out performs Python.
- The benchmarks were done on same hardware, on Linux and similar internet conditions

## Flags or Arguments

Usage:  
wallpaper-downloader [-h|--help] [-o|--output "\<value>" [-o|--output \<value>] [-n|--number \<integer>]  
 [-t|--threads \<integer>] [-r|--range
(day|week|month|year|all)] [-s|--subreddit
"\<value>"] [-p|--portrait] [--width \<integer>] [--height \<integer>]

Arguments:

```
-h  --help        Print help information
-o  --output      Output directory path. Default: [Wallpapers/]
-n  --number      Maximum number of images to be fetched, rounded off to
                  smallest multiple of 10. Default: 50
-t  --threads     Number of Threads. Default: 4
-p  --popularity  Popularity of posts to fetch. Default: top
-r  --range       Range for top posts. Default: all
-s  --subreddit   Name of Subreddit. Default: wallpaper
-P  --portrait    Turn on to allow portrait images. Default: false
    --width       Minimum Width of images (in pixels). Default: 1920
    --height      Minimum Height of images (in pixels). Default: 1080
-v  --version     Check version of program. Default: false
```

## Installation

### GNU/Linux, Unix or MacOS

1. Download the latest release of the software from [here](https://github.com/allen505/wallpaper-downloader/releases/)
2. Open a terminal and navigate to the downloaded file
3. Make the file executable by running `chmod u+x wallpaper-downloader-linux-amd64`
4. Run `./wallpaper-downloader-linux-amd64` to run with the default settings
5. By default the wallpapers will be download to the directory where the file is present. You can use the `-o` or `--output` argument to specify an Output Directory
6. Run `./wallpaper-downloader-linux-amd64 -h` for the help menu

### Windows

1. Download the latest release of the software from [here](https://github.com/allen505/wallpaper-downloader/releases/)
2. Open a terminal and navigate to the downloaded file
3. Run `.\wallpaper-downloader-windows-amd64.exe` to run with the default settings
4. By default the wallpapers will be download to the directory where the file is present. You can use the `-o` or `--output` argument to specify an Output Directory
5. Run `.\wallpaper-downloader-windows-amd64.exe -h` for help menu  
