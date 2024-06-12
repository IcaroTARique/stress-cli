# stress-cli

A simple CLI tool to stress test your server.

## Use
Use with docker building the image
```bash
 docker build -t stress-cli .
```

### Flags
SpitFire uses the following flags:
- **-u or --url**: The URL to stress test
- **-j or --job**: The number of jobs to run
- **-w or --workers**: The number of workers to run
- **-v or --verbose**: To see the internal work

### Verbose 
If you would like to see the internal work
set --verbose or -v.
```bash
go run main.go SpitFire -u="http://google.com" -j=100 -w=50 -v
```
### NotVerbose
Interested only in the results? Do the following
```bash
go run main.go SpitFire -u="http://google.com" -j=100 -w=50
```
