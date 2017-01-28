// +build darwin

package initlayer

// Setup for Darwin. Included here so that Visual Studio Code will
// stop getting into a twist trying to build an unbuildable class
// during source code editing on macOS.
func Setup(initLayer string, rootUID, rootGID int) error {
	panic("darwin not really supported for initlayer")
}
