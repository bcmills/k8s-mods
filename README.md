# k8s-mods

This was a transitional module to specify a consistent set of requirements for
various `k8s.io` modules (and their transitive dependencies).

As of `v1.15.0`, the upstream Kubernetes modules now have `go.mod` files, so
this bootstrapping step should no longer be needed (see
[kubernetes/kubernetes#68577](https://github.com/kubernetes/kubernetes/issues/68577)).
**Please use those modules directly instead of bootstrapping through this one.**

Note that only the published Kubernetes subcomponents are supported as module
dependencies. See
[kubernetes/kubernetes#79384](https://github.com/kubernetes/kubernetes/issues/79384).
