# dicks [![Build Status](https://travis-ci.org/jimeh/dicks.png)](https://travis-ci.org/jimeh/dicks)

Print a number of random length ASCII dicks (`8====D`) to stdout. Inspired
by the [dicks](https://rubygems.org/gems/dicks) RubyGem.

```
$ dicks
8======D
8===D
8====D
8======D
8=========D
$ dicks 2
8======D
8=========D
$ dicks --balls=":" --shaft="-" --head=">" 3
:->
:------>
:--->
$ dicks --min-length=10 --max-length=20 3
8==============D
8===================D
8===============D
```

## Why?

I started playing with Go a few days ago, and wanted a stupidly simple
throw-away project to experiment with a few things like `go test`,
[goxc](https://github.com/laher/goxc) for cross-platform binary building,
Travis-CI integration, and more.

Recreating the [dicks](https://rubygems.org/gems/dicks) gem seemed like a
stupid enough pebble-sized mountain to tackle.


## Installation

### Binaries

Binaries are available for download on the
[Releases](https://github.com/jimeh/dicks/releases) page for:

- Mac OS X / Darwin (x86, amd64)
- FreeBSD (x86, amd64, arm)
- Linux (x86, amd64, arm)
- Windows (x86, amd64)
- NetBSD (x86, amd64, arm)
- OpenBSD (x86, amd64)
- Plan9 (x86)

### From Source

```bash
go get github.com/jimeh/dicks
```

## License

```
        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                    Version 2, December 2004

 Copyright (C) 2014 Jim Myhrberg

 Everyone is permitted to copy and distribute verbatim or modified
 copies of this license document, and changing it is allowed as long
 as the name is changed.

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

  0. You just DO WHAT THE FUCK YOU WANT TO.
```
