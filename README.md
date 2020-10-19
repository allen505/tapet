# Wallpaper downloader

## Table of Contents

  - [Inspiration](#inspiration)
  - [Advantages of this project](#advantages-of-this-project)
  - [Speed comparison](#speed-comparison)
  - [Flags or Arguments](#flags-or-arguments)
  - [Installation](#installation)
    - [GNU/Linux, Unix or MacOS](#gnulinux-unix-or-macos)
    - [Windows](#windows)
  - [Customization](#customization)

## Inspiration

This project is inspired by the wonderful people of [Reddit](https://www.reddit.com/) and the [reddit-wallpaper-downloader](https://github.com/mrsorensen/reddit-wallpaper-downloader) project. Though the above mentioned project downloads images from Reddit well, its written in Python which tends to make it slow. This project however is written in [Go lang](https://golang.org/) which is an open source programming language made by Google. This helps programs run at lightning speeds and makes multithreading a breeze.

## Advantages of this project

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
wallpaper-downloader [-h|--help] [-t|--threads \<integer\>]  
 [-r|--range (day|week|month|year|all)]  
 [-s|--subreddit "\<value\>"]

Arguments:

```
  -h  --help       Print help information
  -t  --threads    Number of Threads. Default: 4
  -r  --range      Range for top posts. Default: all
  -s  --subreddit  Name of Subreddit. Default: wallpaper
```

## Installation

### GNU/Linux, Unix or MacOS

1. Download the latest release of the software from [here](https://github.com/allen505/wallpaper-downloader/releases/)
2. Open a terminal and navigate to the downloaded file
3. Run `./wallpaper-downloader` to run with the default settings
4. The wallpapers will be download to `~/Pictures/Wallpapers`
5. Run `./wallpaper-downloader -h` for help menu
6. (_Optional_) To regularly run _wallpaper-downloader_, use `crontab` while redirecting `stdout` and `stderr` to `/dev/null`. So your the commands would look like this:
   ```
   ./wallpaper-downloader any-arguments-here > /dev/null 2>&1
   ```
   You can learn more about crontab [here](https://www.geeksforgeeks.org/crontab-in-linux-with-examples/)

### Windows

1. Download the latest release of the software from [here](https://github.com/allen505/wallpaper-downloader/releases/)
2. Open a terminal and navigate to the downloaded file
3. Run `.\wallpaper-downloader.exe` to run with the default settings
4. The wallpapers will be download to `~\Pictures\Wallpapers`
5. Run `.\wallpaper-downloader.exe -h` for help menu

## Customization

Since the code is open sourced _wallpaper-downloader_ is highly customizable. The following parameters can be modified by editing the lines immediately after the `import()` statements in the `getWalls.go` file:

- Destination Folder
- Minimum Width
- Minimum Height
- Client Timeout Duration (in seconds)
- Posts per Request
- Cap limit of threads
- Cap limit of name size
