# UNIS

`UNIS` tries to share a common architecture and the necessary `interfaces` that will help you to refactor your project or application to a better place to work on. Choose one way to organise your `string utilities`, across your projects.

Developers can now, move forward and implement their own types of string utilities based on the UNIS architecture. 

**Apply good design patterns** from the beginning and you will be saved from a lot of work later on. Trust me.

## Documentation

UNIS contains some basic implementations that you can see and learn from!

First of all, let's take look at the `unis.Processor` interface. 
```go
// Processor is the most important interface of this package.
//
// It's being used to implement basic string processors.
// Users can use all these processors to build more on their packages.
//
// A Processor should change the "original" and returns its result based on that.
type Processor interface {
	// Process accepts an "original" string and returns a result based on that.
	Process(original string) (result string)
}
```
As you can see, it contains just a function which accepts a string and returns a string.
Everything implements that interface only -- **No, please don't close the browser yet!**


### Continue...

How many times are you using `strings.Replace` in your project? -- Correct, a lot.
Spawning `strings.Replace` many times is dangerous because you may forget a replacement somewhere after a change.

UNIS has a function which can help you structure all of your `strings.Replace` to one spot
based on the `replacements`. Replacements is just a map[$oldstring]$newstring.

```go
// NewReplacer accepts a map of old and new string values.
// The "old" will be replaced with the "new" one.
//
// Same as for loop with a strings.Replace("original", old, new, -1).
NewReplacer(replacements map[string]string) ProcessorFunc
```

> `ProcessorFunc` is just an ease-to-use "alias" for a `Processor`. It's a func that accepts the same arguments as `Processor.Process` and 
in the same time implements the `Processor` interface. Exactly like the `net/http.HandlerFunc` function.

Example:

```go
package main

import (
	"github.com/esemplastic/unis"
)

const slash = "/"

// SlashFixer removes double (system) slashes and returns the fixed path.
var SlashFixer = unis.NewReplacer(map[string]string{
	"\\": slash,
	"//": slash,
})

func main() {
	original := "\\home\\/users//Downloads"
	result := SlashFixer(original) // /home/users/Downloads
	print(original)
	print(" |> ")
	println(result)
}
```

> `SlashFixer` is an `unis.ProcessorFunc`, can be called as `SlashFixer.Process(string)` or simply `SlashFixer(string)`.

### Chain

```go
// Processors is a list of string Processor.
Processors []Processor

// NewChain returns a new chain of processors
// the result of the first goes to the second and so on.
NewChain(processors ...Processor) ProcessorFunc
```

### Conditional

```go
// NewConditional runs the 'p' processor, if the string didn't
// changed then it assumes that that processor has being failed
// and it returns a Chain of the 'alternative' processor(s).
NewConditional(p Processor, alternative ...Processor) ProcessorFunc
```


### Prefix

```go
// NewPrefixRemover accepts a "prefix" and returns a new processor
// which returns the result without that "prefix".
NewPrefixRemover(prefix string) ProcessorFunc

// NewPrepender accepts a "prefix" and returns a new processor
// which returns the result prepended with that "prefix".
NewPrepender(prefix string) ProcessorFunc

// NewPrependerIfNotExists accepts an "expectedIndex" as int
// and a "prefixChar" as byte and returns a new processor
// which returns the result prepended with that "prefixChar"
// if the "original" string[expectedIndex] != prefixChar.
NewPrependerIfNotExists(expectedIndex int, prefixChar byte) ProcessorFunc
```


### Suffix

```go
// NewSuffixRemover accepts a "suffix" and returns a new processor
// which returns the result without that "suffix".
NewSuffixRemover(suffix string) ProcessorFunc

// NewAppender accepts a "suffix" and returns a new processor
// which returns the result appended with that "suffix".
NewAppender(suffix string) ProcessorFunc 
```

### Range

```go
// NewRange accepts "begin" and "end" indexes.
// Returns a new processor which tries to
// return the "original[begin:end]".
NewRange(begin, end int) ProcessorFunc

// NewRangeBegin almost same as NewRange but it
// accepts only a "begin" index, that means that
// it assumes that the "end" index is the last of the "original" string.
//
// Returns the "original[begin:]".
NewRangeBegin(begin int) ProcessorFunc

// NewRangeEnd almost same as NewRange but it
// accepts only an "end" index, that means that
// it assumes that the "start" index is 0 of the "original".
//
// Returns the "original[0:end]".
NewRangeEnd(end int) ProcessorFunc
```


### Bonus: Divider

We saw that everything are `Processors` at the end.
UNIS has some other interfaces like `Divider` too, which should
split a string into two different string pieces, and `Joiner` which should joins two pieces into one. 


```go

// Divider should be implemented by all string dividers.
type Divider interface {
	// Divide takes a string "original" and splits it into two pieces.
	Devide(original string) (part1 string, part2 string)
}

// NewDivider returns a new divider which splits
// a string into two pieces, based on the "separator".
//
// On failure returns the original path as its first
// return value, and empty as it's second.
NewDivider(separator string) Divider

// NewInvertOnFailureDivider accepts a Divider "divider"
// and returns a new one.
//
// It calls the previous "divider" if succed then it returns
// the result as it is, otherwise it inverts the order of the result.
//
// Rembmer: the "divider" by its nature, returns the original string
// and empty as second parameter if the divide action has beeing a failure.
NewInvertOnFailureDivider(divider Divider) Divider

// Divide is an action which runs a new divider based on the "separator"
// and the "original" string.
Divide(original string, separator string) (string, string)
```

## Support

Help me share realistic design patterns by [starring](https://github.com/esemplastic/unis/stargazers) the project!
If you would like to add your implementation of a `UNIS Processor`, feel free to push a [PR](https://github.com/esemplastic/unis/pulls)!


## Philosophy

The UNIS philosophy is to provide small, robust tooling for common string actions, making it a great solution for an extensible project.

I would love to see UNIS as a common place of all Go's `extensible string utilities` that any Gopher will find and use with ease.
The goal is to make this repository authored by a `Community` which cares about code extensibility, stability and buety! 

## Versioning

Current: **0.0.1**  
Date: 6 May 2017

Read more about Semantic Versioning 2.0.0

 - http://semver.org/
 - https://en.wikipedia.org/wiki/Software_versioning
 - https://wiki.debian.org/UpstreamGuide#Releases_and_Versions

## Contributing

If you are interested in contributing to the UNIS project, please make a [PR](https://github.com/esemplastic/unis/pulls).

## People

A list of all contributors can be found [here](CONTRIBUTORS).

## TODO

- [ ] Tests, Examples & more complete Documentation on realistic usecases.

## License

Unless otherwise noted, the source files are distributed
under the 3-Clause BSD License found in the [LICENSE](LICENSE) file.