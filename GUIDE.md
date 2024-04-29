## Setting Up a Development Environment
### UI
Just run `yarn` in gui directory and web app will be served.
`yarn` will automatically recompile code when it detects a change, then live-reloading web applications.

    $ cd gui
    $ yarn serve

App running at: http://localhost:8081

### Go
Generates a Web GUI application

    $ ./build.sh

then build the Docker container, developing in docker environment

    $ docker-compose -f docker-compose.dev.yml up --build

`gin` will automatically recompile code when it detects a change, then `docker` will be restarted.

App running at http://localhost:2017

## Tool
* [go tool link](https://pkg.go.dev/cmd/link#hdr-Command_Line)
* [sentry-go](https://github.com/getsentry/sentry-go)

## Library
* [go-sysinfo](https://github.com/elastic/go-sysinfo) : go-sysinfo is a library for collecting system information.
* [GoDotEnv](https://github.com/joho/godotenv) : A Go port of Ruby's dotenv library (Loads environment variables from .env files)