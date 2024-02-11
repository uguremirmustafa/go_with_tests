package hello

import "fmt"

const (
	english = "English"
	french  = "French"
	turkish = "Turkish"

	engHelloPrefix = "Hello, "
	frHelloPrefix  = "Bonjour, "
	trHelloPrefix  = "Merhaba, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// prefix is our named return value in this function
// this will create a variable called `prefix` in the function
// it will be assigned the "zero" value. it depends on the type,
// for example int:0 string:""
// you can return whatever it's set to by just calling `return` rather than `return prefix`
func greetingPrefix(language string) (prefix string) {
	switch language {
	case turkish:
		prefix = trHelloPrefix
	case french:
		prefix = frHelloPrefix
	default:
		prefix = engHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("ugur", turkish))
}
