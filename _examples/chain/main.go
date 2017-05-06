package main

import (
	"github.com/esemplastic/unis"
)

func NewPathNormalizer() unis.Processor {
	slash := "/"
	replacer := unis.NewReplacer(map[string]string{
		"\\": slash,
		"//": slash,
	})

	suffixRemover := unis.NewSuffixRemover(slash)

	slashPrepender := unis.NewPrependerIfNotExists(0, slash[0])

	return unis.NewChain(
		slashPrepender,
		replacer,
		suffixRemover,
	)
}

var pathNormalizer = NewPathNormalizer()

func NormalizePath(path string) string {
	if path == "" {
		return path
	}
	return pathNormalizer.Process(path)
}

func main() {
	original := "home\\/users//Downloads"
	result := NormalizePath(original) // /home/users/Downloads
	print(original)
	print(" |> ")
	println(result)
}
