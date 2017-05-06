package main

import (
	"path"
	"strings"

	"github.com/esemplastic/unis"
)

func NewPathNormalizer() unis.Processor {
	slash := "/"
	replacer := unis.NewReplacer(map[string]string{
		`\`:  slash,
		`//`: slash,
	})

	suffixRemover := unis.NewSuffixRemover(slash)
	slashPrepender := unis.NewTargetedJoiner(0, slash[0])

	toLower := unis.ProcessorFunc(strings.ToLower) // convert standard functions to UNIS and add to the chain.
	cleanPath := unis.ProcessorFunc(path.Clean)    // convert standard functions to UNIS and add to the chain.
	return unis.NewChain(
		cleanPath,
		slashPrepender,
		replacer,
		suffixRemover,
		toLower,
	)
}

var defaultPathNormalizer = NewPathNormalizer()

func NormalizePath(path string) string {
	if path == "" {
		return path
	}
	return defaultPathNormalizer.Process(path)
}

func main() {
	original := "api\\////users/"
	result := NormalizePath(original) // /api/users
	print(original)
	print(" |> ")
	println(result)
}
