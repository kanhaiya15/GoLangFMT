### mould
Golang boilerplate for LambdaTest projects

[![Build Status](https://drone.lambdatest.io/api/badges/LambdaTest/mould/status.svg)](https://drone.lambdatest.io/LambdaTest/mould)

## Intro
This boilerplate is based on [go modules](https://blog.golang.org/using-go-modules), so you don't need regular golang file structure. This
project can live in any directory on your local system.


### Using this boilerplate
- Visit https://github.com/LambdaTest/mould
- Click `Use this template` to create a new repository
- Git clone your new repo
- After cloning, add this repo to git remote of your project to regularly receive improvements.
  > Eg. `git remote add template https://github.com/LambdaTest/mould`
- Then to update your repo with this app, fire `git pull template master`

## Building
- This can be built just like any other bo project `go build -o mould main.go`
- This project is docker ready. Dockerfile in the root directory can be used to create service docker
  - Sample docker usage
  ```
  docker build . --tag mould
  docker run -p 12223:12223  --env MLD_PORT=12223 mould
  ```
- A sample shell file is also included to automate the build steps. `build.sh` in the root directory can be used for special adding build steps.

## CI/CD
This repo is initialized with drone based CI/CD. The pipeline for CI/CD is defined in `.drone.yml` placed in root directory. This sample CI/CD script does the following tasks
- Clones the repo
- Make builds for windows, linux, mac for 32bit and 64 bit
- Pushes the builds to S3
- Invalidates the CloudFront distribution attached to the bucket
- Finally sends slack notification on build status
- The jobs defined in the `.drone.yml` uses some environment variables which can be viewed on https://drone.lambdatest.io/LambdaTest/mould/settings

There are couple of examples of Drone recipes which are currently being used in production based projects.
- [Burrow](https://github.com/LambdaTest/burrow). CI/CD of this repo is inspired from Burrow.
- [Tunnel server](https://github.com/LambdaTest/tunnel-server) This project has CI/CD which builds the server images and pushes to predefined servers using ansible
- [UnderPass](https://github.com/LambdaTest/underpass). This creates electron based apps which are signed, notarized and then published to S3


## Development

### Configuration management
This are many ways to configure the app in accordance to [12 factor app methodology](https://12factor.net/). This app uses [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper) to provide idiomatic way of creating command line based configurable apps.
- The main config model is defined in `config/model.go` Here you define the main config object of the application. LogConfig should be default in all the apps
- `config/default.go` contains the defaults for configs.
- This app can take configurations from following mediums
  - Environment variables prefixed with `MLD_`. This prefix is customized in `config/loader.go`
  - `.env` files in the current directory can have list of environment variables that can be used to provide config values. These should correspond to values defined in `config/model.go`. The environment variables are lowercased to find the match. Eg `MLD_PORT` will match to `Port` config defined in config file.
  - Config can be specified and overriden using command line arguments. `cmd/flags.go` can be used to create command line flags. The flags are automatically bounded to config names defined in `config/model.go` after lowercasing them. For eg. `Port` config defined in model is automatically bounded to `port` cmd flag. If you need different names for command line flags and config name, you need to define json tags for config names in `config/model.go`.  For eg. `SomeWeirdConfig` is configured to `some-weird-config` in command line flags
  - Config can also be defined in json, yaml, toml, ini files having naming convention like this `mould.json`, `.mould.toml` etc. See [Viper home page](https://github.com/spf13/viper). The base name can be overriden in `config/loader.go`
- Global constants are defined in `pkg/global/constants.go`

### Logging
- This app has a logging instance wrapper. A wrapper instance is configured with either [Uber Zap](https://github.com/uber-go/zap) or [Logrus](https://github.com/sirupsen/logrus). By default we use zap. This is configured in `cmd/bin.go` like `lumber.NewLogger(cfg.LogConfig, cfg.Verbose, lumber.InstanceZapLogger)`
- If you want to use any other logging library you can do without changing the source code by implementing Logger interface defined in `lumber/setup.go` and then initializing with your lib like in previous step
- Logging configs are defined in `lumber/setup.go`.
- Currently, logs are streamed on stdout as well as written in `mould.log`. Logs on stdout are non verbose whereas the logs in file are in verbose mode and include the timestamp. This behaviour can be customized using logging config. The name of default log file is set in `config/default.go`

### Cmd configuration
The main cmd line application defintion is defined in `cmd/bin.go` This is also the main starting file. main.go just triggers the function defined in this file
- Since project is using [Cobra](https://github.com/spf13/cobra). `--help` is already available on the built binary which displays help. This again can be configured in `cmd/bin.go`
- Main version of the project is defined in `pkg/global/version.go`. This version is configured to be displayed when app is called with `--version` flag.
- Command line flags are defined in `cmd/flags.go`. See [Cobra](https://github.com/spf13/cobra) on various options available in command line customization

### Hot reloading
This project is configured to support [fresh](https://github.com/gravityblast/fresh) runner which reloads the application actively whenever any golang file (or any other file configured for hot reloading) changes. This is very useful while actively developing as it removes the need to recompile and run the application again and again. `runner.conf` in the root directory is used to configure the fresh runner. More information can be viewed on [their github project](https://github.com/gravityblast/fresh)

### Design pattern
This app is built using assumption that multiple services (http, ws, crons, queues) can be part of the app. These all are separate subofferings but share common configuration and logger.
- These sub offerings run concurrently using go routines as defined in `cmd/bin.go`
- A central golang context is created in `bin/cmd.go` which is passed to all the subprocesses. It's the responsibility of subofferings to listen to the context and gracefully shutdown on cancellation of the context.
- Tha main process waits for all these subofferings to quit gracefully and then quits the application.
- These subsofferings are passed an instance of waitGroup which needs to signalled when these subprocesses are completed using `wg.Done()`
- If these subofferings are not shutdown gracefully within specified time, then main app quits forcefully.
- Sigint, (`ctrl-c`) can be used to gracefully kill the app
- All the subofferings are passed context, config, logger and waitgroup instance from `cmd/bin.go`
- subofferings can be defined as separate pkg in `pkg` subdirectory

An alternative plugin based design pattern where all the offerings can be customized using plugins is implemented in [Burrow](https://github.com/LambdaTest/burrow)

### Examples in this app
 By default the http server runs of 9876 as defined in `config/default.go`
- This app contains an HTTP server with global and versioned routes which can be accessed at `http://localhost/health`, `http://localhost/signal`, `http://localhost:9876/v1.0/hello` (curl -i -X POST -H "Authorization: Basic bWF5YW5rOmJob2xh" http://localhost:9876/v1.0/hello)
- Websocket echo endpoint is available at `ws://localhost:12223/ws`. You can use [websocat](https://github.com/vi/websocat) to test this using this command `websocat ws://localhost:12223/ws -H 'Authorization:Basic bWF5YW5rYjpLMEVVZ2NIdzNKR3hENjRDbGFrQVZvZDVpVjFaM2J5ZmhmQzdtMWt0amJ2VERzemwwaAo='`
- A cron setup which prints on screen every minute. Defined in `pkg/cron/`


## Builds
Prod Builds for this boilerplate for experimentation can be downloaded from following endpoints:
- Windows: https://downloads.lambdatest.com/mould/alpha/windows/64bit/mould
- Mac https://downloads.lambdatest.com/mould/alpha/mac/64bit/mould
- Linux: https://downloads.lambdatest.com/mould/alpha/linux/64bit/mould


## TODO
* [ ] Add basic test cases


