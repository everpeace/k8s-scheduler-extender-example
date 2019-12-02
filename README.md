# k8s-scheduler-extender-example
This is a repo forked from an [k8s scheduler extender example](https://github.com/everpeace/k8s-scheduler-extender-example)
developed by @everpeace. 

We restructured the code to wrap all shared functions and variables into a package `k8s-scheduler-extender`.

## Usage
You can reference the functions and variables by import the package.
```golang
import "github.com/wangchen615/k8s-scheduler-extender-example/k8s-scheduler-extender"
```

The `main.go` is an example extender revised using the above reference.