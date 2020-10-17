# Wallpaper downloader

## Inspiration

This project is inspired by the wonderful people of [Reddit](https://www.reddit.com/) and the [reddit-wallpaper-downloader](https://github.com/mrsorensen/reddit-wallpaper-downloader) project. Though the above mentioned project downloads images from Reddit well, its written in Python which tends to make it slow. This project however is written in [Go lang](https://golang.org/) which is an open source programming language made by Google. This helps programs run at lightning speeds and makes multithreading a breeze.

## Advantages of this project

- Written in Go
- Lightning speed
- Multithreaded for even better speeds
- Fetch images based on Top of the Day, Month, Week, Year or even All time
- Choose subreddit to fetch images from

## Speed comparison

| Number of Photos |      Go      |    Python    |
| :--------------: | :----------: | :----------: |
|        10        |    10secs    |     47 secs  |
|       100        |    42secs    |  5min 28secs |

- Golang was set to run on 4 threads, which is the default
- The benchmarking was NOT scientific by any means. However you are welcome to try both projects and see how Go out performs Python.
- The benchmarks were done on same hardware and similar internet conditions

## Flags or Arguments

Usage:  
wallpaper-downloader    [-h|--help] [-t|--threads \<integer\>]  
                        [-r|--range (day|week|month|year|all)]  
                        [-s|--subreddit "\<value\>"]

Arguments:
```
  -h  --help       Print help information  
  -t  --threads    Number of Threads. Default: 4    
  -r  --range      Range for top posts. Default: all  
  -s  --subreddit  Name of Subreddit. Default: wallpaper  
```