# My-ls

My-ls is a custom command-line tool programmed in Go that replicates the functionality of the 'ls' command with specific arguments. It provides an alternative way to list files and directories in a directory.

## Features

- Supports the following options:
  - `-a`: Show all files, including hidden files.
  - `-r`: Reverse the order of listing.
  - `-t`: Sort files by modification time.
  - `-l`: Display files in long format, including permissions, size, and modification time.
  - `-R`: Recursively list files and directories.

## Requirements

- Go (version X.X.X)

## Installation
 Clone the repository:

```shell
git clone < link > 
cd my-ls 
```
## USAGE
``` shell
go run . < path > < arg >
```
### exemple
``` shell
go run . ./ -a -l -rt
```
the *go run .*  represente the ls command
## License & Copyright
**Alpapie**