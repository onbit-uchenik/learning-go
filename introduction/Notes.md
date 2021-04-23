# Notes
```
Random notes learned while reading book An Introduction To Programming in Go and for references The Go Programming Language
```
* Interface is an abstract type it doesn't exposes the representation or internal structure of its values, or the set of basic operations they support; it reveals only some of their methods. When we ahve an value of interface type we know nothing about what it is, we only know what it can do, what I can expect from it, what will happen if I use it, but never how it do it. For example error is an interface type.
* To know the type of any variable we can use package reflect.