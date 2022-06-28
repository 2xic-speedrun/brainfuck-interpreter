package brainfuck

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	output := Run("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
	if output != "Hello World!" {
		panic("Error in interpreter")
	}
}
