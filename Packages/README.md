## Packages

This is one of important thing in go and this helps us to reduce our development time, as we can reuse the things.

### Important

* Only variables starting with a capital letter gets exported to other packages. They have to import the package in their program.

* We can have any number of sub directories and go programs inside them. In the main package or other package using these just have to import the package location properly.

* Packages must be inside `src` folder and we should set its root folder in GOPATH or GOROOT. Otherwise you will get an error similar to this

> main.go:3:8: cannot find package "greeting" in any of:
	/Github/Go_Tutorial/src/greeting (from $GOROOT)
