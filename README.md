
goviz
=====

a visualization tool for golang project dependency

This tool is for helping source code reading. 
The dependency of the whole code can be visualized quickly. 

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/github_com_hirokidaichi_goviz-ignore-test.png)



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
goviz -i github.com/hashicorp/serf -ignore-test -ignore-libs | dot -Tpng -o hoge.png
```

The result is [here](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/github_com_hashicorp_serf-ignore-libs_-ignore-test.png).

### Option

```
 -i :input path from $GOPATH/src
 -o :output filename (default STDOUT )
 -ignore-test :  Ignore "test" or "example"
 -ignore-libs :  Ignore pkgs
 -level  : dependency depth (default math.MaxInt8 )
```

## Samples

### [go-xslate](https://github.com/lestrrat/go-xslate)


```
go get github.com/lestrrat/go-xslate
goviz -i github.com/lestrrat/go-xslate -ignore-test -ignore-libs | dot -Tpng -o xslate.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/github_com_lestrrat_go-xslate-ignore-libs_-ignore-test_.png)

### [anko](https://github.com/mattn/anko)


```
go get github.com/mattn/anko
goviz -i github.com/mattn/anko -ignore-test -ignore-libs | dot -Tpng -o anko.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/github_com_mattn_anko-ignore-libs_-ignore-test.png)

### [packer](https://github.com/mitchellh/packer)


```
go get github.com/mitchellh/packer
goviz -i github.com/mitchellh/packer -ignore-test -ignore-libs | dot -Tpng -o packer.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/github_com_mitchellh_packer-ignore-libs_ignore-test.png)

### [serf](https://github.com/hashicorp/serf)

```
go get github.com/hashicorp/serf
goviz -i github.com/hashicorp/serf -ignore-test -ignore-libs | dot -Tpng -o serf.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/github_com_hashicorp_serf-ignore-libs_-ignore-test.png)


### [vegeta](https://github.com/tsenart/vegeta)

```
# not ignore-libs
go get github.com/tsenart/vegeta
goviz -i github.com/tsenart/vegeta -ignore-test | dot -Tpng -o vegeta.png
```
![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/github_com_tsenart_vegeta-ignore-test.png)

### [docker](https://github.com/dotcloud/docker)

```
# use level option 
go get github.com/dotcloud/docker
goviz -i github.com/dotcloud/docker/docker -ignore-test -ignore-libs -level 1| dot -Tpng -o docker-level1.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/github_com_dotcloud_docker_docker-ignore-test_-ignore-libs_-level_1.png)

```
# use level option 
go get github.com/dotcloud/docker
goviz -i github.com/dotcloud/docker/docker -ignore-test -ignore-libs -level 2| dot -Tpng -o docker-level2.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/github_com_dotcloud_docker_docker-ignore-test_-ignore-libs_-level_2.png)

```
# use level option 
go get github.com/dotcloud/docker
goviz -i github.com/dotcloud/docker/docker -ignore-test -ignore-libs -level 3| dot -Tpng -o docker-level3.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/github_com_dotcloud_docker_docker-ignore-test_-ignore-libs_-level_3.png)

```
# see only execdriver
go get github.com/dotcloud/docker
goviz -i github.com/dotcloud/docker/runtime/execdriver/execdrivers -ignore-test -ignore-libs | dot -Tpng -o docker-executiondrivers.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/github_com_dotcloud_docker_runtime_execdriver_execdrivers_-ignore-libs_-ignore-test.png)
## Ricense

MIT

## Author

hirokidaichi [at] gmail.com



