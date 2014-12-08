#ricons [![Build Status](https://travis-ci.org/drbig/ricons.svg?branch=master)](https://travis-ci.org/drbig/ricons) [![GoDoc](https://godoc.org/github.com/drbig/ricons?status.svg)](http://godoc.org/github.com/drbig/ricons)

A Go library implementing a minimalistic random icon/avatar generator framework.

Features / Bugs:

- Simple and minimal code
- Base framework code has complete test coverage
- Includes four simple generators
- Benchmarks for generators
- Includes full-fledged HTTP icon server command
- Probably not so bad performance via HTTP (in terms of k req/s)
- Currently not optimised in any way (except for being simple)
- It's fun

## Showcase

Output of included generators, created with `showcase/generate.rb`:

Generator | Size | 1 | 2 | 3 | 4
--------- | ---- | - | - | - | -
grid | 16x16 | ![grid 1](https://raw.github.com/drbig/ricons/master/showcase/grid-16x16-1.png) | ![grid 2](https://raw.github.com/drbig/ricons/master/showcase/grid-16x16-2.png) | ![grid 3](https://raw.github.com/drbig/ricons/master/showcase/grid-16x16-3.png) | ![grid 4](https://raw.github.com/drbig/ricons/master/showcase/grid-16x16-4.png)
grid | 100x100 | ![grid 1](https://raw.github.com/drbig/ricons/master/showcase/grid-100x100-1.png) | ![grid 2](https://raw.github.com/drbig/ricons/master/showcase/grid-100x100-2.png) | ![grid 3](https://raw.github.com/drbig/ricons/master/showcase/grid-100x100-3.png) | ![grid 4](https://raw.github.com/drbig/ricons/master/showcase/grid-100x100-4.png)
symsquare | 16x16 | ![symsquare 1](https://raw.github.com/drbig/ricons/master/showcase/symsquare-16x16-1.png) | ![symsquare 2](https://raw.github.com/drbig/ricons/master/showcase/symsquare-16x16-2.png) | ![symsquare 3](https://raw.github.com/drbig/ricons/master/showcase/symsquare-16x16-3.png) | ![symsquare 4](https://raw.github.com/drbig/ricons/master/showcase/symsquare-16x16-4.png)
symsquare | 100x100 | ![symsquare 1](https://raw.github.com/drbig/ricons/master/showcase/symsquare-100x100-1.png) | ![symsquare 2](https://raw.github.com/drbig/ricons/master/showcase/symsquare-100x100-2.png) | ![symsquare 3](https://raw.github.com/drbig/ricons/master/showcase/symsquare-100x100-3.png) | ![symsquare 4](https://raw.github.com/drbig/ricons/master/showcase/symsquare-100x100-4.png)
uniform | 16x16 | ![uniform 1](https://raw.github.com/drbig/ricons/master/showcase/uniform-16x16-1.png) | ![uniform 2](https://raw.github.com/drbig/ricons/master/showcase/uniform-16x16-2.png) | ![uniform 3](https://raw.github.com/drbig/ricons/master/showcase/uniform-16x16-3.png) | ![uniform 4](https://raw.github.com/drbig/ricons/master/showcase/uniform-16x16-4.png)
uniform | 100x100 | ![uniform 1](https://raw.github.com/drbig/ricons/master/showcase/uniform-100x100-1.png) | ![uniform 2](https://raw.github.com/drbig/ricons/master/showcase/uniform-100x100-2.png) | ![uniform 3](https://raw.github.com/drbig/ricons/master/showcase/uniform-100x100-3.png) | ![uniform 4](https://raw.github.com/drbig/ricons/master/showcase/uniform-100x100-4.png)
vgrad | 16x16 | ![vgrad 1](https://raw.github.com/drbig/ricons/master/showcase/vgrad-16x16-1.png) | ![vgrad 2](https://raw.github.com/drbig/ricons/master/showcase/vgrad-16x16-2.png) | ![vgrad 3](https://raw.github.com/drbig/ricons/master/showcase/vgrad-16x16-3.png) | ![vgrad 4](https://raw.github.com/drbig/ricons/master/showcase/vgrad-16x16-4.png)
vgrad | 100x100 | ![vgrad 1](https://raw.github.com/drbig/ricons/master/showcase/vgrad-100x100-1.png) | ![vgrad 2](https://raw.github.com/drbig/ricons/master/showcase/vgrad-100x100-2.png) | ![vgrad 3](https://raw.github.com/drbig/ricons/master/showcase/vgrad-100x100-3.png) | ![vgrad 4](https://raw.github.com/drbig/ricons/master/showcase/vgrad-100x100-4.png)

You may notice the generators have been inspired by what I've already seen online. The uniform generator though is mostly for testing (to have full coverage of base framework code).

### Raw generator benchmarks

On Linux x64, Intel i7-2620M @ 2.70GHz:

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

My non-scientific tests with `httperf` (two machines, 1Gbps LAN) indicate around 2k ~ 5k req/s for 32x32 icons depending on generator, though please treat it as a guesstimate.

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
