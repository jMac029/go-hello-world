/*
Go / Golang
- Strong-typed
- Compiled to machine code
- Not Object Oriented (OO)
- Struct Oriented/Based
- Block Scoped
- Interface
- Composition
*/

var i int
// declaring i as an int will make it forever be an int.
// it cannot be re-assigned, hence strong-typed

i = "fooo" // compile err

var s string
// s will always be a string and cannot be re-assigned


// Go uses minimal punctuation
// no need to wrap a for loop in () as you would in JS
for i := 0; i < 10; i++ {
	fmt.Println(i)
	//fmt.Println is how you console.log in golang
}

// Go assignment
var foo string //without type derivation
foo = "bar"

foo := "bar" //with type derivation
// := means create the variable then assign it's type

bar = "blah"
// this will fail, you have to assign a type first


// Go export module 
package apackagename 
//define package name so that it is available to other files

// function (func) decleration
func Bar (s string) string {
	//Bar will be exported
	//structs/funcs need to be Captilized in order to be exported via package
}

func bar (s string) string {

}

bool // boolean declartion
int // integar declartion


// Go import
import (
    "fmt" // part of Go’s standard library
    "github.com/foo/foo" // imported directly from repository
)
foo.Bar()


// Go - return multiple values
func bar() (int, error) {
    return 1, nil
}
a, b := bar()


//Go error handling
foo, err := bar()
if err != nil {
    // handle error with defer, panic, recover, or log.fatal, etc...
}


// Variadic function
func foo (args ...int) {
	fmt.Println(len(args))
}

func main() {
	foo() // 0
	foo(1,2,3) // 3
}