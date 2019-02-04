# k8s-mods

This module specifies a consistent set of requirements for various `k8s.io`
modules (and their transitive dependencies).

To the extent possible, the `k8s.io` modules are all taken from the
corresponding kubernetes-* version tags.

**DO NOT IMPORT FROM THIS MODULE.** Instead, use `go get -m` to impose its
minimum requirements within your module, then run `go mod tidy` to prune away
unnecessary dependencies:

```
$ go get -m github.com/bcmills/k8s-mods
$ go mod tidy
```
