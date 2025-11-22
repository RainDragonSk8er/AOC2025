package scaffold

import "fmt"

// GetTemplate returns the funny "Hello World" content for a given extension and day.
func GetTemplate(ext string, day int) string {
	switch ext {
	case ".lua":
		return fmt.Sprintf("print(\"8 scripts a day keeps the doctor away\")\n-- Solution for Day %d", day)
	case ".rs":
		return fmt.Sprintf("fn main() {\n    println!(\"blazingly fast from prod to customer (laughably slow from dev to prod)\");\n    // Solution for Day %d\n}", day)
	case ".js":
		return fmt.Sprintf("console.log(\"wow look how i built the web with slow code\");\n// Solution for Day %d", day)
	case ".ts":
		return fmt.Sprintf("let s: string = \"wow look how i build the web with slow code\";\nconsole.log(s);\n// Solution for Day %d", day)
	case ".py":
		return fmt.Sprintf("print(\"I'm not slow, I'm just abstracting away your patience.\")\n# Solution for Day %d", day)
	case ".go":
		return fmt.Sprintf("package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"if err != nil { panic('I have no generics for my feelings') }\")\n    // Solution for Day %d\n}", day)
	case ".hs":
		return fmt.Sprintf("main = putStrLn \"Monads are just monoids in the category of endofunctors.\"\n-- Solution for Day %d", day)
	case ".c":
		return fmt.Sprintf("#include <stdio.h>\n\nint main() {\n    printf(\"Segmentation fault (core dumped) - just kidding, Hello World!\\n\");\n    // Solution for Day %d\n    return 0;\n}", day)
	case ".cpp":
		return fmt.Sprintf("#include <iostream>\n\nint main() {\n    std::cout << \"I hope you like 500 line error messages\" << std::endl;\n    // Solution for Day %d\n    return 0;\n}", day)
	case ".cs":
		return fmt.Sprintf("using System;\n\nclass Program {\n    static void Main() {\n        Console.WriteLine(\"I'm just Java with a cool hat.\");\n        // Solution for Day %d\n    }\n}", day)
	case ".sh":
		return fmt.Sprintf("#!/bin/bash\necho 'rm -rf / --no-preserve-root # Just kidding'\n# Solution for Day %d", day)
	case ".jl":
		return fmt.Sprintf("println(\"I'm like Python but I actually lift.\")\n# Solution for Day %d", day)
	case ".md":
		return fmt.Sprintf("# Solution for Day %d\n\nTODO: Write solution", day)
	default:
		return fmt.Sprintf("// Solution for Day %d", day)
	}
}
