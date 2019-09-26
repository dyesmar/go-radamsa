# Go Radamsa

The _Go Radamsa_ package provides a native interface to the [Radamsa](https://gitlab.com/akihe/radamsa) mutational fuzzer via [cgo](https://golang.org/cmd/cgo/). This means that you can get mutational fuzzing right in your Go process without having to ever touch the `radamsa` binary.

This is development quality software. (Refer to _[On My Funny Ideas About What Beta Means](https://inessential.com/2019/09/02/on_my_funny_ideas_about_what_beta_means)_ for the definition of what that means in this context.). _Go Radamsa_ has some some rough edges, including but not limited to: missing functionality, no tests, and no documentation. All of these issues will be addressed once libradamsa has been formally released (soonish). In the meantime, I may break things at any moment. You have been advised.

**SPDX-License-Identifier: [MIT](https://spdx.org/licenses/MIT)**
