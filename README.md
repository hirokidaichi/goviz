iz has a function which outputs the metrics (instability) of go project. goviz
=====

a visualization tool for golang project dependency

This tool is for helping source code reading. 
The dependency of the whole code can be visualized quickly. 

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/own.png)



## Installation

```
go install github.com/hirokidaichi/goviz
```

and if not being installed [graphviz](http://www.graphviz.org), install it :

```
brew install graphviz
```

## Usage

```
goviz -i github.com/hashicorp/serf | dot -Tpng -o hoge.png
```

### Option

```
Usage:
  goviz [OPTIONS]

Application Options:
  -i, --input=   intput ploject name
  -o, --output=  output file (STDOUT)
  -d, --depth=   max plot depth of the dependency tree (128)
  -f, --focus=   focus on the specific module
  -s, --search=  top directory of searching
  -l, --leaf     whether leaf nodes are plotted (false)
  -m, --metrics  display module metrics (false)

Help Options:
  -h, --help     Show this help message

Usage:
  goviz [OPTIONS]

Application Options:
  -i, --input=   intput ploject name
  -o, --output=  output file (STDOUT)
  -d, --depth=   max plot depth of the dependency tree (128)
  -f, --focus=   focus on the specific module
  -s, --search=  top directory of searching
  -l, --leaf     whether leaf nodes are plotted (false)
  -m, --metrics  display module metrics (false)

Help Options:
  -h, --help     Show this help message

exit status 1

```

## Samples

### [anko](https://github.com/mattn/anko)


```
goviz -i github.com/mattn/anko | dot -Tpng -o anko.png
```
![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/anko.png)


### [serf](https://github.com/hashicorp/serf)


```
goviz -i github.com/hashicorp/serf | dot -Tpng -o serf.png
```
![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/serf.png)


### [go-xslate](https://github.com/lestrrat/go-xslate)


```
goviz -i github.com/lestrrat/go-xslate | dot -Tpng -o xslate.png
```
![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/xslate.png)


### [vegeta](https://github.com/tsenart/vegeta)


```
goviz -i github.com/tsenart/vegeta -l| dot -Tpng -o vegeta.png
```
![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/vegeta.png)


### [packer](https://github.com/mitchellh/packer)


```
goviz -i github.com/mitchellh/packer --search SELF -l| dot -Tpng -o packer.png
```
![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/packer.png)


### [docker plot depth 1](https://github.com/dotcloud/docker/docker)


```
goviz -i github.com/dotcloud/docker/docker -s github.com/dotcloud/docker -d 1| dot -Tpng -o docker1.png
```
![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/docker1.png)


### [docker plot depth 2](https://github.com/dotcloud/docker/docker)


```
goviz -i github.com/dotcloud/docker/docker -s github.com/dotcloud/docker -d 2| dot -Tpng -o docker2.png
```
![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/docker2.png)


### [docker plot depth 3](https://github.com/dotcloud/docker/docker)


```
goviz -i github.com/dotcloud/docker/docker -s github.com/dotcloud/docker -d 3| dot -Tpng -o docker3.png
```
![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/docker3.png)


### [docker&#39;s execdrivers](https://github.com/dotcloud/docker/runtime/execdriver/execdrivers/)


```
goviz -i github.com/dotcloud/docker/runtime/execdriver/execdrivers/ -s github.com/dotcloud/docker| dot -Tpng -o docker-execdrivers.png
```
![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/docker-execdrivers.png)


### docker's metrics
goviz has a function which outputs the metrics (instability) of go project. 

```
goviz -i github.com/dotcloud/docker/docker -m 
```
Instability is a value of 0 to 1. 
It suggests that it is such an unstable module that this value is high. 
It becomes easy to distinguish whether it is a module nearer to  application layer, and whether it is a module near a common library. 


```
Inst:1.000 Ca(  0) Ce(  9)	github.com/dotcloud/docker/docker
Inst:0.960 Ca(  1) Ce( 24)	github.com/dotcloud/docker/pkg/libcontainer/nsinit
Inst:0.956 Ca(  2) Ce( 43)	github.com/dotcloud/docker/runtime
Inst:0.950 Ca(  1) Ce( 19)	github.com/dotcloud/docker/api/client
Inst:0.950 Ca(  1) Ce( 19)	github.com/dotcloud/docker/server
Inst:0.909 Ca(  1) Ce( 10)	github.com/dotcloud/docker/api/server
Inst:0.867 Ca(  2) Ce( 13)	github.com/dotcloud/docker/runtime/execdriver/native
Inst:0.857 Ca(  1) Ce(  6)	github.com/dotcloud/docker/runtime/graphdriver/devmapper
Inst:0.833 Ca(  1) Ce(  5)	github.com/dotcloud/docker/runtime/graphdriver/aufs
Inst:0.800 Ca(  1) Ce(  4)	github.com/dotcloud/docker/builtins
Inst:0.800 Ca(  2) Ce(  8)	github.com/dotcloud/docker/runtime/networkdriver/bridge
Inst:0.800 Ca(  1) Ce(  4)	github.com/dotcloud/docker/runtime/execdriver/execdrivers
Inst:0.778 Ca(  2) Ce(  7)	github.com/dotcloud/docker/pkg/libcontainer/network
Inst:0.750 Ca(  1) Ce(  3)	github.com/dotcloud/docker/sysinit
Inst:0.750 Ca(  3) Ce(  9)	github.com/dotcloud/docker/runtime/execdriver/lxc
Inst:0.750 Ca(  1) Ce(  3)	github.com/dotcloud/docker/runtime/execdriver/native/template
Inst:0.727 Ca(  3) Ce(  8)	github.com/dotcloud/docker/graph
Inst:0.667 Ca(  1) Ce(  2)	github.com/dotcloud/docker/runtime/execdriver/native/configuration
Inst:0.667 Ca(  1) Ce(  2)	github.com/dotcloud/docker/runtime/networkdriver/portmapper
Inst:0.667 Ca(  1) Ce(  2)	github.com/dotcloud/docker/runtime/networkdriver/ipallocator
Inst:0.667 Ca(  1) Ce(  2)	github.com/dotcloud/docker/links
Inst:0.571 Ca(  9) Ce( 12)	github.com/dotcloud/docker/runconfig
Inst:0.500 Ca(  2) Ce(  2)	github.com/dotcloud/docker/pkg/selinux
Inst:0.500 Ca(  5) Ce(  5)	github.com/dotcloud/docker/api
Inst:0.500 Ca(  2) Ce(  2)	github.com/dotcloud/docker/daemonconfig
Inst:0.500 Ca(  5) Ce(  5)	github.com/dotcloud/docker/image
Inst:0.500 Ca(  1) Ce(  1)	github.com/dotcloud/docker/pkg/libcontainer/capabilities
Inst:0.500 Ca(  1) Ce(  1)	github.com/gorilla/mux
Inst:0.500 Ca(  1) Ce(  1)	github.com/dotcloud/docker/runtime/graphdriver/btrfs
Inst:0.500 Ca(  1) Ce(  1)	github.com/dotcloud/docker/runtime/graphdriver/vfs
Inst:0.444 Ca( 10) Ce(  8)	github.com/dotcloud/docker/archive
Inst:0.333 Ca(  2) Ce(  1)	github.com/dotcloud/docker/opts
Inst:0.333 Ca(  2) Ce(  1)	github.com/dotcloud/docker/runtime/networkdriver/portallocator
Inst:0.250 Ca(  6) Ce(  2)	github.com/dotcloud/docker/registry
Inst:0.250 Ca(  6) Ce(  2)	github.com/dotcloud/docker/pkg/cgroups
Inst:0.250 Ca(  3) Ce(  1)	github.com/dotcloud/docker/pkg/sysinfo
Inst:0.250 Ca(  3) Ce(  1)	github.com/dotcloud/docker/runtime/networkdriver
Inst:0.154 Ca( 11) Ce(  2)	github.com/dotcloud/docker/runtime/graphdriver
Inst:0.125 Ca(  7) Ce(  1)	github.com/dotcloud/docker/pkg/label
Inst:0.091 Ca( 10) Ce(  1)	github.com/dotcloud/docker/nat
Inst:0.083 Ca( 11) Ce(  1)	github.com/dotcloud/docker/runtime/execdriver
Inst:0.077 Ca( 36) Ce(  3)	github.com/dotcloud/docker/utils
Inst:0.067 Ca( 14) Ce(  1)	github.com/dotcloud/docker/engine
Inst:0.056 Ca( 17) Ce(  1)	github.com/dotcloud/docker/pkg/libcontainer
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/collections
Inst:0.000 Ca(  1) Ce(  0)	github.com/gorilla/context
Inst:0.000 Ca(  1) Ce(  0)	code.google.com/p/go.net/websocket
Inst:0.000 Ca(  7) Ce(  0)	github.com/dotcloud/docker/dockerversion
Inst:0.000 Ca(  3) Ce(  0)	github.com/dotcloud/docker/pkg/mflag
Inst:0.000 Ca(  4) Ce(  0)	github.com/dotcloud/docker/pkg/mount
Inst:0.000 Ca(  1) Ce(  0)	github.com/dotcloud/docker/pkg/namesgenerator
Inst:0.000 Ca(  4) Ce(  0)	github.com/dotcloud/docker/pkg/netlink
Inst:0.000 Ca(  1) Ce(  0)	github.com/dotcloud/docker/pkg/proxy
Inst:0.000 Ca(  1) Ce(  0)	github.com/dotcloud/docker/pkg/listenbuffer
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/signal
Inst:0.000 Ca( 10) Ce(  0)	github.com/dotcloud/docker/pkg/system
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/systemd
Inst:0.000 Ca(  6) Ce(  0)	github.com/dotcloud/docker/pkg/term
Inst:0.000 Ca(  3) Ce(  0)	github.com/dotcloud/docker/pkg/user
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/version
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/libcontainer/utils
Inst:0.000 Ca(  4) Ce(  0)	github.com/dotcloud/docker/pkg/libcontainer/apparmor
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/iptables
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/graphdb
Inst:0.000 Ca(  5) Ce(  0)	github.com/dotcloud/docker/vendor/src/code.google.com/p/go/src/pkg/archive/tar

```
## Ricense

MIT

## Author

hirokidaichi [at] gmail.com



