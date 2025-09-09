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

## Contributing

If you find bugs issues are always welcome. If you can also make a PR, those
are even more welcome! A few things to keep in mind though:

- This is a side project of a side project. It might take me a while to get
  back to you. This project is designed to be easily fork-able so you are not
  dependent on me.
- If you can make a PR, it makes my life easier. Even if it isn't perfect or
  fully tested, that's still helpful to me.
- I might reject issues or PRs for any reason. I'm not trying to be an asshole,
  I promise! I'm just one person with limited time.

For new features, I might be a bit more selective but generally open to it if
it generally useful. My guideline guideline is:

> "if it is a pretty trivial implementation that should probably just be in the
> standard library, then it should be a gstb package."

Feel free to create new issues with ideas. If you already have an
implementation, feel free to create a PR too! Here are a few rules:

- The above contributing points apply. Please keep those in mind.
- Packages should be in a single `.go` file with any tests being in a separate
  `_test.go` file.
- Except for things in the Go standard library, no external libraries. Assume
  the latest version of Go.
- Avoid inter-dependencies between packages in this library, but consider how
  they should fit together. Since things are meant to be self-contained
  single-files, having hard dependencies makes integration more difficult.
- Tests are not required, but heavily favored.
