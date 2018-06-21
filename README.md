# Watch'n'Go

**WORK IN PROGRESS**

 * Watch a single file
 * Watch files recursively in a directory, with an optional pattern
 * Run a command on modifications through `/bin/sh -c <command>`

## Install

```
go get -u github.com/Leryan/watchngo/cmd/watchngo
go install github.com/Leryan/watchngo/cmd/watchngo
```

## Usage

```
watchngo [-conf watchngo.ini] [-command <your command> -match <match> [-filter <filter>] [-debug]]
```

When using `-command -match -filter` options, configuration will be ignored. This makes it possible to use `watchngo` without writing a configuration file.

## Configuration

See [watchngo.sample.ini](watchngo.sample.ini) configuration example.

## TODO

 * [x] Recursive directory watching
 * [x] Match files using `path/filepath.Glob()`
 * [ ] Override the default command (`/bin/sh -c <command>`) that starts the actual command by configuration
 * [x] Command interpolation: `%match`, `%filter`, `%event.file`, `%event.op`
