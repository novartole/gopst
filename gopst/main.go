package main

import "github.com/ebitengine/purego"

func main() {
	// Register library
	typstlib, err := purego.Dlopen("liblaunch_typst.dylib", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		// Handle error
	}
	defer purego.Dlclose(typstlib)

	// Register the function
	var launch func(string) uint8
	purego.RegisterLibFunc(&launch, typstlib, "launch")

	// Calling
	if launch(getArgs()) == 0 {
		println("DONE")
	} else {
		println("FAILED")
	}
}

// Getting or building args
func getArgs() string {
	return "typst compile ../example/text.typ ../example/text.pdf"
}
