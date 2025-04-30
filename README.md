# gstb

Single-file public domain (or MIT licensed) libraries for Go, inspired by
[nothings/stb](https://github.com/nothings/stb).

*Note that this is _not_ a _binding_ to stb, just a project in the same spirit
of stb.*

---

## Usage

You can use this library in one of two ways:

### (recommended) Just copy it!

Each package is self-contained. Just copy the relevant `.go` file (and the
`_test.go` file if you want too) into your project.

### `go get`

Note that this library does _not_ use [SemVer](https://semver.org/). I only
develop against recent-ish Go versions and make no guarantees about backwards
or forwards compatibility.

```
go get -u github.com/sapslaj/gstb
```

## Background

`<yapping>`

Over the years I have found that there have been a few things that are missing
from Go's standard library or are shortcomings of it (for reasons both good and
bad). This collection of libraries are just things that I find myself copying
around a lot. I got tired of having to pull in these gigantic modules for tiny
pieces of functionality and resorted to just copying things around like the
good ol C days. And honestly, it has kinda worked out pretty well for me. I
don't think _every_ module should work this way, but for small stuff like this
I think it's totally fine. Turns out I'm not the only one who thinks this way,
hence the heavy [stb](https://github.com/nothings/stb) inspiration.

These packages do not generally follow things like SOLID, TDD, or any other
design principles. Everything is exposed and public; no private members.
Constructors are avoided; the zero value is meaningful. Only functions that
start with `Must...` are expected to ever panic. Generics are used (arguably
abused) heavily to provide more type safety at the cost of more "boilerplate"
and/or difficult to read code; I consider this a reasonable trade-off for the
most part.

`</yapping>`
