# Go Radamsa

The Go Radamsa package provides a Go interface to the [Radamsa](https://gitlab.com/akihe/radamsa) mutational fuzzer using [cgo](https://golang.org/cmd/cgo/). This means that you can get mutational fuzzing in your Go programs without having to ever touch the `radamsa` binary.

This is development quality software. (Refer to _[On My Funny Ideas About What Beta Means](https://inessential.com/2019/09/02/on_my_funny_ideas_about_what_beta_means)_ for the definition of what that means.). Go Radamsa has some rough edges, including but not limited to: missing functionality, no tests, and no documentation. All of these issues will be addressed once `libradamsa` has been formally released (soonish). In the meantime, I may break things at any moment. You have been advised.

## Considerations

Radamsa is included as a submodule dependency. Be sure to use the `--recursive` flag when cloning this repository.

Because `libradamsa` has not yet been released, you have to work within Radamsa's `develop` branch. Currently, here is how to build `libradamsa` on macOS:

```bash
pushd radamsa
git checkout develop
make libradamsa-test
mkdir -p ../cache
cp c/radamsa.h lib/libradamsa.a ../cache
rm -fr bin/libradamsa-test c/libradamsa.c lib
popd
```

Now you are ready to build the test driver and give it a spin:

```bash
pushd cmd/goradamsa
make
./goradamsa -n 100 'Yay, fuzzing!'
```

The `goradamsa` command serves a dual purpose: it is my test driver and it illustrates how the `radamsa` package should be used. Type `goradamsa -h` to display the currently supported set of flags.

## Legal

Radamsa is © 2013 Aki Helin. The Go interface is © 2019 Ramsey Dow. Go Radamsa is released under the same license as Radamsa.

SPDX-License-Identifier: [MIT](https://spdx.org/licenses/MIT)
