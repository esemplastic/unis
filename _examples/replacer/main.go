package main

import (
	"github.com/esemplastic/unis"
)

const slash = "/"

// SlashFixer removes double (system) slashes and returns the new path.
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

	// if expected, got := result, SlashFixer(original); expected != got {
	// 	fmt.Printf("expected '%s' but got '%s'", expected, got)
	// 	return
	// }

}
