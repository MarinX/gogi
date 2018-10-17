# GOGI
A car diagnostic tool written in go
> go + obd + grafana + influx

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/MarinX/gogi)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

# Demo
[![Video Showcase](http://img.youtube.com/vi/AVP4vLoic_g/0.jpg)](http://www.youtube.com/watch?v=AVP4vLoic_g)

# Usage

### Building
```sh
# normal build
go build

# building on rpi or arm device
GOOS=linux GOARCH=arm go build
```

### Configuration
!!! Lookup and copy enviroment file `.env.example` to `.env`

| Key          | Value       | Default | Description | 
| -------------|-------------|---------|-------------|
| APP_DEBUG | bool | false | Use debug logger |
| USE_FAKE | bool | false | Use fake OBD readings - good for testing|
| SERIAL_DEVICE | string | /dev/ttyUSB0 | Path to OBD serial device|
| DB_DRIVER | string | influx | Driver to use for storing OBD readings |
| DB_HOST | string | http://localhost | Host of the database|
| DB_PORT | int | 8086 | Port for the database |
| DB_USERNAME | string || Username for database|
| DB_PASSWORD | string || Password for database|
| DB_DATABASE | string | mb | Schema to use for database |

### Running
```sh
./gogi
```


# Roadmap
- [ ] Add tests
- [ ] Add MySQL driver
- [ ] Test on more cars

# License
MIT
 