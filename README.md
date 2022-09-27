# Exercism Go Track

[![Configlet Status](https://github.com/exercism/go/workflows/Configlet%20CI/badge.svg)](https://github.com/exercism/go/actions?query=workflow%3A%22Configlet+CI%22)
[![Exercise Test Status](https://github.com/exercism/go/workflows/Exercise%20tests/badge.svg)](https://github.com/exercism/go/actions?query=workflow%3A%22Exercise+tests%22)
[![Gitter chat](https://badges.gitter.im/exercism/go.svg)](https://gitter.im/exercism/go)

Exercism exercises in Go

## Issues

We welcome issues filed at https://github.com/exercism/go/issues for problems of any size.
Feel free to report typographical errors or poor wording.
We are most interested in improving the quality of the test suites.
You can greatly help us improve the quality of the exercises by filing reports of invalid solutions that pass tests or of valid solutions that fail tests.

## Development setup

Beyond filing issues, if you would like to contribute directly to the Go code in the Exercism Go track, you should follow some standard Go development practices.
You should have a [recent version of Go](http://golang.org/doc/install) installed, ideally either the current release, the previous release, or tip.

You will need a github account and you will need to fork exercism/go to your account.
See [GitHub Help](https://help.github.com/articles/fork-a-repo/) if you are unfamiliar with the process.
Clone your fork with the command: `git clone https://github.com/<you>/go`.
Test your clone by cding to the go directory and typing `bin/fetch-golangci-lint` and then
`bin/run-tests`. You should see tests pass for all exercises.

Note that unlike most other Go code, it is not necessary to clone this to your GOPATH.
This is because this repo only imports from the standard library and isn't expected to be imported by other packages.

Your Go code should be formatted using the [gofmt](https://golang.org/cmd/gofmt/) tool.

There is a [misspelling tool](https://github.com/client9/misspell). You can install and occasionally run it to
find low hanging typo problems. [#570](https://github.com/exercism/go/pull/570) It's not added into CI since it could give false positives.

## Contributing Guide

Please be familiar with the [contributing guide](https://github.com/exercism/legacy-docs/tree/main/contributing-to-language-tracks) in the docs repository.
This describes some great ways to get involved.
In particular, please read the [Pull Request Guidelines](https://github.com/exercism/legacy-docs/blob/main/contributing/pull-request-guidelines.md) before opening a pull request.

## Exercism Go style

Let's walk through an example, non-existent, exercise, which we'll call
`fizzbuzz` to see what could be included in its implementation.

### Exercise configuration

An exercise is [configured](https://github.com/exercism/docs/blob/master/language-tracks/configuration/exercises.md)
via an entry in the exercises array in [config.json file](/config.json). If `fizzbuzz` is an optional
exercise, it would have an entry below the core exercises that might look like:

```json
{
  "slug": "fizzbuzz",
  "uuid": "mumblety-peg-whatever",
  "core": false,
  "unlocked_by": "two-fer",
  "difficulty": 1,
  "topics": ["conditionals"]
}
```

See [Exercise Configuration](https://github.com/exercism/docs/blob/master/language-tracks/configuration/exercises.md)
for more info.

### Exercise files: Overview

For any exercise you may see a number of files present in a directory under `exercises/`:

```sh
~/go/exercises/fizzbuzz
$ tree -a
.
├── cases_test.go
├── example.go
├── fizzbuzz.go
├── fizzbuzz_test.go
├── .meta
│   └── description.md
│   └── gen.go
│   └── hints.md
│   └── metadata.yml
└── README.md
```

This list of files _can vary_ across exercises.
Not all exercises use all of these files.
Exercises originate their test data and README text from the Exercism [problem-specification repository](https://github.com/exercism/problem-specifications/tree/master/exercises).
This repository collects common information for all exercises across all tracks.
However, should track-specific documentation need to be included with the exercise, files in an exercise's `.meta/` directory can be used to override or augment the exercise's README.

Let's briefly describe each file:

- **cases_test.go** - Contains [generated test cases](#synchronizing-exercises-with-problem-specifications), using test data sourced from the problem-specifications repository.
  Only in some exercises.
  Automatically generated by `.meta/gen.go`.

- **example.go** - An [example solution](#example-solutions) for the exercise used to verify the test suite.
  Ignored by the `exercism fetch` command.
  See also [ignored files](#ignored-files).

- **fizzbuzz.go** - A _stub file_, in some early exercises to give users a starting point.

- **fizzbuzz_test.go** - The main test file for the exercise.

- **.meta/** - Contains files not to be included when a user fetches an
  exercise: See also [ignored files](#ignored-files).

- **.meta/description.md** - Use to generate a [track specific description](https://github.com/exercism/docs/blob/master/language-tracks/exercises/anatomy/readmes.md) of the exercise in the exercise's README.

- **.meta/gen.go** - Generates `cases_test.go` when present.
  See also [synchronizing exercises with problem specifications](#synchronizing-exercises-with-problem-specifications).

- **.meta/hints.md** - Use to add track specific information in addition to the generic exercise's problem-specification description in the README.

- **.meta/metadata.yml** - Track specific exercise metadata, overrides the exercise metadata from the problem-specifications repository.

In some exercises there can be extra files, for instance the [series](exercises/series/) exercise contains extra test files.

### Ignored files

When a user fetches an exercise, they do not need to get all the files within an exercise directory.
For instance; the _example.go_ files that contain an example solution, or the _gen.go_ files used to generate an exercise's test cases.
Therefore there are certain files and directories that are ignored when an exercise is fetched.
These are:

- The _.meta_ directory and anything within it.
- Any file that matches the `ignore_pattern` defined in [config.json file](/config.json).
  This currently matches any filename that contains the word `example`, _unless_ it is followed by the word `test`, with any number of characters in between.

### Example solutions

_example.go_ is a reference solution.
It is a valid solution that the CI (continuous integration) service can run tests against.
Files with _"example"_ in the file name are skipped by the `exercism fetch` command.
Because of this, there is less need for this code to be a model of style, expression and readability, or to use the best algorithm.
Examples can be plain, simple, concise, even naïve, as long as they are correct.

### Stub files

Stub files, such as _leap.go_, are a starting point for solutions.
Not all exercises need to have a stub file, only exercises early in the syllabus.
By convention, the stub file for an exercise with slug `exercise-slug`
must be named `exercise_slug.go`. This is because CI needs to delete stub files
to avoid conflicting definitions.

### Tests

The test file is fetched for the solver and deserves
attention for consistency and appearance.

The `leap` exercise makes use of data-driven tests.
Test cases are defined as data, then a test function iterates over the data.
In this exercise, as they are generated, the test cases are defined in the _cases_test.go_ file.
The test function that iterates over this data is defined in the _leap_test.go_ file.
The _cases_test.go_ file is generated by information found in [problem-specifications](https://github.com/exercism/problem-specifications) using [generators](#synchronizing-exercises-with-problem-specifications).
To add additional test cases (e.g. test cases that only make sense for Go) add the test cases to `<exercise>_test.go`.
An example of using additional test cases can be found in the exercise [two-bucket](exercises/practice/two-bucket/two_bucket_test.go).

Identifiers within the test function appear in actual-expected order as described
at [Useful Test Failures](https://github.com/golang/go/wiki/CodeReviewComments#useful-test-failures).
Here the identifier `observed` is used instead of actual. That's fine. More
common are words `got` and `want`. They are clear and short. Note [Useful Test
Failures](https://github.com/golang/go/wiki/CodeReviewComments#useful-test-failures)
is part of [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
Really we like most of the advice on that page.

In Go we generally have all tests enabled and do not ask the solver to edit the test program, to enable progressive tests for example.
`t.Fatalf()`, as seen in the _leap_test.go_ file, will stop tests at the first failure encountered, so the solver is not faced with too many failures at once.

### Testable examples

Some exercises can contain [Example tests](https://blog.golang.org/examples)
that document the exercise API. These examples are run alongside the standard
exercise tests and will verify that the exercise API is working as expected.
They are not required by all exercises and are not intended to replace the
data-driven tests. They are most useful for providing examples of how an
exercise's API is used. Have a look at the example tests in the [clock exercise](https://github.com/exercism/go/blob/master/exercises/clock/example_clock_test.go)
to see them in action.

### Errors

We like errors in Go. It's not idiomatic Go to ignore invalid data or have undefined
behavior. Sometimes our Go tests require an error return where other language
tracks don't.

### Benchmarks

In most test files there will also be benchmark tests, as can be seen at the end
of the _leap_test.go_ file. In Go, benchmarking is a first-class citizen of the
testing package. We throw in benchmarks because they're interesting, and because
it is idiomatic in Go to think about performance. There is no critical use for
these though. Usually they will just bench the combined time to run over all
the test data rather than attempt precise timings on single function calls. They
are useful if they let the solver try a change and see a performance effect.

## Synchronizing exercises with problem specifications

Some problems that are implemented in multiple tracks use the same inputs and
outputs to define the test suites.
Where the [problem-specifications](https://github.com/exercism/problem-specifications)
repository contains a _canonical-data.json_ file with these inputs and outputs,
we can generate the test cases programmatically.
The problem-specifications repo also defines the instructions for the exercises, which are also shared across tracks and must also be synchronized.

### Test structure

See the _gen.go_ file in the `leap` exercise for an example of how this
can be done.

Test case generators are named _gen.go_ and are kept in a special _.meta_
directory within each exercise that makes use of a test cases generator. This
_.meta_ directory will be ignored when a user fetches an exercise.

Whenever the shared JSON data changes, the test cases will need to be regenerated.
The generator will first look for a local copy of the **problem-specifications** repository.
If there isn't one it will attempt to get the relevant json data for the exercise from the **problem-specifications** repository on GitHub.
The generator uses the GitHub API to find some information about exercises when it cannot find a local copy of **problem-specifications**.
This can cause throttling issues if working with a large number of exercises.
_We therefore recommend using a local copy of the repository when possible_ (remember to keep it current :smile:).

To use a local copy of the **problem-specifications** repository, make sure that it has been cloned into the same parent-directory as the **go** repository.

```sh
$ tree -L 1 .
.
├── problem-specifications
└── go
```

#### Adding a test generator to an exercise

For some exercises, a test generator is used to generate the `cases_test.go` file with the test cases based on information from [problem-specifications](https://github.com/exercism/problem-specifications).
To add a new exercise generator to an exercise the following steps are needed:
1. Create the file `gen.go` in the directory `.meta` of the exercise
2. Add the following template to `gen.go`:
```go
package main

import (
    "log"
    "text/template"
  
    "../../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j = map[string]interface{}{
              "property_1":  &[]Property1Case{},
              "property_2":  &[]Property2Case{},
	}
	if err := gen.Gen("<exercise-name>", j, t); err != nil {
		log.Fatal(err)
	}
}
```
3. Insert the name of the exercise to the call of `gen.Gen`
4. Add all values for the field `property` in `canonical-data.json` to the map `j`. `canonical-data.json` can be found at [problem-specifications/exercises/\<exercise-name\>](https://github.com/exercism/problem-specifications)
5. Create the needed structs for storing the test cases from `canonical-data.json` (you can for example use [JSON-to-Go](https://mholt.github.io/json-to-go/) to convert the JSON to a struct)

**NOTE:** In some cases, the struct of the data in the `input`/`expected` fields is not the same for all test cases of one property. In those situations, an `interface{}` has to be used to represent the values for these fields.  These `interface{}` values then need to be handled by the test generator. A common way to handle these cases is to create methods on the test case structs that perform type assertions on the `interface{}` values and return something more meaningful. These methods can then be referenced/called in the `tmpl` template variable. Examples of this can be found in the exercises [forth](https://github.com/exercism/go/blob/main/exercises/practice/forth/.meta/gen.go) or [bowling](https://github.com/exercism/go/blob/main/exercises/practice/bowling/.meta/gen.go).

6. Add the variable `tmpl` to `gen.go`. This template will be used to create the `cases_test.go` file.

Example:
```go
var tmpl = `package <package of exercise>

{{.Header}}

var testCases = []struct {
	description    string
	input          int
	expected       int       
}{ {{range .J.<property>}}
{
	description: {{printf "%q"  .Description}},
	input: {{printf "%d"  .Score}},
	expected: {{printf "%d"  .Score}},
},{{end}}
}
`
```
7. Synchronize the test case using the exercise generator (as described in [Synchronizing tests and instructions](#synchronizing-tests-and-instructions))
8. Check the validity of `cases_test.go`
9. Use the generated test cases in the `<exercise>_test.go` file
10. Check if `.meta/example.go` passes all tests

### Synchronizing tests and instructions

To keep track of which tests are implemented by the exercise the file `.meta/tests.toml` is used by [configlet](https://github.com/exercism/configlet).

To synchronize the exercise with [problem-specifications](https://github.com/exercism/problem-specifications) and to regenerate the tests, navigate into the **go** directory and perform the following steps:

1. Synchronize your exercise with [`exercism/problem-specifications`](https://github.com/exercism/problem-specifications) using [configlet](https://github.com/exercism/configlet):

```console
$ configlet sync --update -e <exercise>
```

`configlet` synchronizes the following parts, if an updated is needed:

* docs: `.docs/instructions.md`, `.docs/introduction.md`
* metadata: `.meta/config.json`
* tests: `.meta/tests.toml`
* filepaths: `./meta/config.json`

For further instructions check out [configlet](https://github.com/exercism/configlet#configlet-sync).

2. Run the test case generator to update `<exercise>/cases_test.go`:

```console
$ GO111MODULE=off go run exercises/practice/<exercise>/.meta/gen.go
```

**NOTE**: If you see the error `json: cannot unmarshal object into Go value of type []gen.Commit` when running the generator you probably have been rate limited by GitHub.
Try providing a GitHub access token with the flag `-github_token="<Token>"`.
Using the token will result in a higher rate limit.
The token does not need any specific scopes as it is only used to fetch infos about commits.

You should see that some/all of the above files have changed. Commit the changes.

### Synchronizing all exercises with generators

```console
$ ./bin/run-generators <GitHub Access Token>
```

**NOTE**: If you see the error `json: cannot unmarshal object into Go value of type []gen.Commit` when running the generator you probably have been rate limited by GitHub.
Make sure you provided the GitHub access token as first argument to the script as shown above.
Using the token will result in a higher rate limit.
The token does not need any specific scopes as it is only used to fetch infos about commits.

## Managing the Go version

For an easy management of the Go version in the `go.mod` file in all exercises, we can use `gomod-sync`.
This is a tool made in Go that can be seen in the `gomod-sync/` folder.

To update all go.mod files according to the config file (`gomod-sync/config.json`) run:

```console
$ cd gomod-sync && go run main.go update
```

To check all exercise go.mod files specify the correct Go version, run:

```console
$ cd gomod-sync && go run main.go check
```

## Pull requests

Pull requests are welcome.
You forked, cloned, coded and tested and you have something good? Awesome!
Use git to add, commit, and push to your repository.
Checkout your repository on the web now.
You should see your commit and the invitation to submit a pull request!

<img src="img/mars1.png">

Click on that big green button.
You have a chance to add more explanation to your pull request here, then send it.
Looking at the exercism/go repository now instead of your own, you see this.

<img src="img/mars2.png">

That inconspicuous orange dot is important!
Hover over it (no, not on this image, on a real page) and you can see it's indicating that a CI build is in progress.
After a few minutes (usually) that dot will turn green indicating that tests passed.
If there's a problem, it comes up red:

<img src="img/mars3.png">

This means you've still got work to do.
Click on "details" to go to the CI build details. Look over the build log for clues.
Usually error messages will be helpful and you can correct the problem.

## Direction

Directions are unlimited.
This code is fresh and evolving.
Explore the existing code and you will see some new directions being tried.
Your fresh ideas and contributions are welcome. :sparkles:

### Go icon

The Go logo was designed by Renée French, and has been released under the Creative Commons 3.0 Attributions license.
