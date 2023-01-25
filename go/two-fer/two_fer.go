// Package twofer is a simple test exercise
// Given a name, return a string with the message:
// "One for name, one for me."
// Where "name" is a variable that will be replaces by what you pass to the function.
package twofer

// ShareWith returns the two-fer string with you name or default "you" taking tame arg 
func ShareWith(name string) string {
    if name == "" {
        name = "you"
    }

	return "One for "+ name +", one for me."
}
