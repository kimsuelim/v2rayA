# V2rayA & Panda

v2rayA is a V2Ray client supporting global transparent proxy. Panda is a global network acceleration service.
We are committed to providing the simplest operation and meet most needs.

## Usage

Client mainly supporting on Windows and macOS.
Binary file and installation package from GitHub releases.

See [Releases](https://github.com/kimsuelim/v2rayA/releases)

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

## Credits

[v2raya/v2raya](https://github.com/v2rayA/v2rayA)

## License

[![License: AGPL v3-only](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
