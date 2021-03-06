#ricons [![Build Status](https://travis-ci.org/drbig/ricons.svg?branch=master)](https://travis-ci.org/drbig/ricons) [![GoDoc](https://godoc.org/github.com/drbig/ricons?status.svg)](http://godoc.org/github.com/drbig/ricons)

A Go library implementing a minimalistic random icon/avatar generator framework.

Features / Bugs:

- Simple and minimal code
- Base framework code has complete test coverage
- Includes four simple generators
- Benchmarks for generators
- Includes full-fledged HTTP icon server command
- Probably not so bad performance via HTTP (in terms of k req/s)
- *Proper* image pool for better performance
- It's fun

### Current word from the author

Sorry for the practically broken code with the previous version. This version should be better. And it's still fun.

Also I believe there might be a bug in `math.Rand` (test under high concurrent load), but maybe it's just me.

## Showcase

Output of included generators, created with `showcase/generate.rb`:

| Generator | Size    | 1   | 2   | 3   | 4   |
| --------- | ------- | --- | --- | --- | --- |
| grid      | 16x16   | ![grid 1](https://raw.github.com/drbig/ricons/master/showcase/grid-16x16-1.png) | ![grid 2](https://raw.github.com/drbig/ricons/master/showcase/grid-16x16-2.png) | ![grid 3](https://raw.github.com/drbig/ricons/master/showcase/grid-16x16-3.png) | ![grid 4](https://raw.github.com/drbig/ricons/master/showcase/grid-16x16-4.png) |
| grid      | 100x100 | ![grid 1](https://raw.github.com/drbig/ricons/master/showcase/grid-100x100-1.png) | ![grid 2](https://raw.github.com/drbig/ricons/master/showcase/grid-100x100-2.png) | ![grid 3](https://raw.github.com/drbig/ricons/master/showcase/grid-100x100-3.png) | ![grid 4](https://raw.github.com/drbig/ricons/master/showcase/grid-100x100-4.png) |
| symsquare | 16x16   | ![symsquare 1](https://raw.github.com/drbig/ricons/master/showcase/symsquare-16x16-1.png) | ![symsquare 2](https://raw.github.com/drbig/ricons/master/showcase/symsquare-16x16-2.png) | ![symsquare 3](https://raw.github.com/drbig/ricons/master/showcase/symsquare-16x16-3.png) | ![symsquare 4](https://raw.github.com/drbig/ricons/master/showcase/symsquare-16x16-4.png) |
| symsquare | 100x100 | ![symsquare 1](https://raw.github.com/drbig/ricons/master/showcase/symsquare-100x100-1.png) | ![symsquare 2](https://raw.github.com/drbig/ricons/master/showcase/symsquare-100x100-2.png) | ![symsquare 3](https://raw.github.com/drbig/ricons/master/showcase/symsquare-100x100-3.png) | ![symsquare 4](https://raw.github.com/drbig/ricons/master/showcase/symsquare-100x100-4.png) |
| uniform   | 16x16   | ![uniform 1](https://raw.github.com/drbig/ricons/master/showcase/uniform-16x16-1.png) | ![uniform 2](https://raw.github.com/drbig/ricons/master/showcase/uniform-16x16-2.png) | ![uniform 3](https://raw.github.com/drbig/ricons/master/showcase/uniform-16x16-3.png) | ![uniform 4](https://raw.github.com/drbig/ricons/master/showcase/uniform-16x16-4.png) |
| uniform   | 100x100 | ![uniform 1](https://raw.github.com/drbig/ricons/master/showcase/uniform-100x100-1.png) | ![uniform 2](https://raw.github.com/drbig/ricons/master/showcase/uniform-100x100-2.png) | ![uniform 3](https://raw.github.com/drbig/ricons/master/showcase/uniform-100x100-3.png) | ![uniform 4](https://raw.github.com/drbig/ricons/master/showcase/uniform-100x100-4.png) |
| vgrad     | 16x16   | ![vgrad 1](https://raw.github.com/drbig/ricons/master/showcase/vgrad-16x16-1.png) | ![vgrad 2](https://raw.github.com/drbig/ricons/master/showcase/vgrad-16x16-2.png) | ![vgrad 3](https://raw.github.com/drbig/ricons/master/showcase/vgrad-16x16-3.png) | ![vgrad 4](https://raw.github.com/drbig/ricons/master/showcase/vgrad-16x16-4.png) |
| vgrad     | 100x100 | ![vgrad 1](https://raw.github.com/drbig/ricons/master/showcase/vgrad-100x100-1.png) | ![vgrad 2](https://raw.github.com/drbig/ricons/master/showcase/vgrad-100x100-2.png) | ![vgrad 3](https://raw.github.com/drbig/ricons/master/showcase/vgrad-100x100-3.png) | ![vgrad 4](https://raw.github.com/drbig/ricons/master/showcase/vgrad-100x100-4.png) |

You may notice the generators have been inspired by what I've already seen online. The uniform generator though is mostly for testing (to have full coverage of base framework code).

### Raw generator benchmarks

v0.0.3 benchmark, done on Linux x64, Intel i7-2620M @ 2.70GHz:

    PASS
    BenchmarkGrid16x16        300000              3634 ns/op
    BenchmarkGrid32x32        300000              4261 ns/op
    BenchmarkSymsquare16x16   200000              6077 ns/op
    BenchmarkSymsquare32x32   200000              7381 ns/op
    BenchmarkUniform16x16    3000000               473 ns/op
    BenchmarkUniform32x32    2000000               725 ns/op
    BenchmarkVgrad16x16      1000000              1555 ns/op
    BenchmarkVgrad32x32       200000              6391 ns/op
    ok      github.com/drbig/ricons 12.332s

Original benchmark, same hardware as above:

    $ go test -bench .
    PASS
    BenchmarkGrid16x16        500000              5114 ns/op
    BenchmarkGrid32x32        500000              7443 ns/op
    BenchmarkSymsquare16x16   500000              7175 ns/op
    BenchmarkSymsquare32x32   200000             10495 ns/op
    BenchmarkUniform16x16    1000000              1570 ns/op
    BenchmarkUniform32x32     500000              3502 ns/op
    BenchmarkVgrad16x16      1000000              2820 ns/op
    BenchmarkVgrad32x32       200000              9409 ns/op
    ok      github.com/drbig/ricons 20.477s

### riconsd HTTP server

Basic command usage:

    $ ./riconsd -h
    Usage: ./riconsd [options]
      -a=":3232": server bind address
      -b=256: image bound (e.g. max 256x256)
      -l=false: show generators and exit
      -q=false: disable logging
    $ ./riconsd -l
    grid: grid-based icons
    symsquare: symmetric square-based icons
    uniform: single uniform color
    vgrad: simple vertical gradient

The HTTP API is as follows:

- `/info.json` - returns description of what the instance offers
- `/<generator>/<format>/<width>/<height>` - get a random icon

You can discover generators and formats via the first URI:

    $ curl -s http://127.0.0.1:3232/info.json | python2 -mjson.tool
    {
        "formats": [
            "gif",
            "jpeg",
            "png"
        ],
        "generators": {
            "grid": "grid: grid-based icons",
            "symsquare": "symsquare: symmetric square-based icons",
            "uniform": "uniform: single uniform color",
            "vgrad": "vgrad: simple vertical gradient"
        },
        "versions": {
            "ricons": "0.0.1",
            "riconsd": "0.0.1"
        }
    }

The icon request URI should be self-explanatory:

    $ curl -o grid.png http://127.0.0.1:3232/grid/png/200/200

As usual, the server exposes some statistics via `expvar` at `/debug/vars`:

- `formats` - map counting successful encodings for each image format
- `generators` - map counting successful icon creations for each generator
- `hits` - total number of requests recived

### riconsd httperf tests

v0.0.3 results below have the following setup:

- `riconsd` runs on the same i7 laptop, with `export GOMAXPROCS=4`
- `httperf` runs on an old Core 2 Duo laptop, with `ulimit -n 4096`
- The LAN is 1 Gbps
- Both laptops don't do much else during the tests (i.e. no other significant load)
- All generators are tested with the same arguments to `httperf`
- All generators are outputting 32px x 32px PNG files
- Still, these are general ballpark test

Full output:

    $ httperf --server 192.168.0.11 --port 3232 --hog --num-conns 256 --num-calls 32 --rate 256 --uri=/uniform/png/32/32
    httperf --hog --client=0/1 --server=192.168.0.11 --port=3232 --uri=/uniform/png/32/32 --rate=256 --send-buffer=4096 --recv-buffer=16384 --num-conns=256 --num-calls=32
    httperf: warning: open file limit > FD_SETSIZE; limiting max. # of open files to FD_SETSIZE
    Maximum connect burst length: 1
    
    Total: connections 256 requests 8192 replies 8192 test-duration 3.750 s
    
    Connection rate: 68.3 conn/s (14.6 ms/conn, <=250 concurrent connections)
    Connection time [ms]: min 266.7 avg 2586.8 max 3452.7 median 2751.5 stddev 577.0
    Connection time [ms]: connect 0.3
    Connection length [replies/conn]: 32.000
    
    Request rate: 2184.7 req/s (0.5 ms/req)
    Request size [B]: 82.0
    
    Reply rate [replies/s]: min 0.0 avg 0.0 max 0.0 stddev 0.0 (0 samples)
    Reply time [ms]: response 80.8 transfer 0.0
    Reply size [B]: header 102.0 content 106.0 footer 0.0 (total 208.0)
    Reply status: 1xx=0 2xx=8192 3xx=0 4xx=0 5xx=0
    
    CPU time [s]: user 0.31 system 3.43 (user 8.4% system 91.6% total 99.9%)
    Net I/O: 620.4 KB/s (5.1*10^6 bps)
    
    Errors: total 0 client-timo 0 socket-timo 0 connrefused 0 connreset 0
    Errors: fd-unavail 0 addrunavail 0 ftab-full 0 other 0

Tabulated:

| Generator | req/s  |
| --------- | ------ |
| uniform   | 2184.7 |
| vgrad     | 2101.3 |
| symsquare | 2116.1 |
| grid      | 2123.1 |

## Contributing

Follow the usual GitHub development model:

1. Clone the repository
2. Make your changes on a separate branch
3. Make sure you at least run `gofmt` before committing
4. Make a pull request

See licensing for legalese.

## Licensing

Standard two-clause BSD license, see LICENSE.txt for details.

Copyright (c) 2014 Piotr S. Staszewski
