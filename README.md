goviz
=====

a visualization tool for golang project dependency

This tool is for helping source code reading. 
The dependency of the whole code can be visualized quickly. 

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/hirokidaichi-goviz-l.png)



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

The result is [here](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/hashicorp-serf.png).

### Option

```
 -i : input path from $GOPATH/src
 -o : output filename (default STDOUT )
 -l : display libraries
 -seek-in : ( SELF | github.com/hoge/fuga )
 -d  : dependency depth (default math.MaxInt8 )
```

## Samples

### [go-xslate](https://github.com/lestrrat/go-xslate)


```
go get github.com/lestrrat/go-xslate
goviz -i github.com/lestrrat/go-xslate  | dot -Tpng -o xslate.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/lestrrat-go-xslate.png)

### [anko](https://github.com/mattn/anko)


```
go get github.com/mattn/anko
goviz -i github.com/mattn/anko   | dot -Tpng -o anko.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/mattn-anko.png)

### [packer](https://github.com/mitchellh/packer)


```
go get github.com/mitchellh/packer
goviz -i github.com/mitchellh/packer   | dot -Tpng -o packer.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/mitchellh-packer-seek-in-SELF--l.png)

### [serf](https://github.com/hashicorp/serf)

```
go get github.com/hashicorp/serf
goviz -i github.com/hashicorp/serf   | dot -Tpng -o serf.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/hashicorp-serf.png)


### [vegeta](https://github.com/tsenart/vegeta)

```
# not ignore-libs
go get github.com/tsenart/vegeta
goviz -i github.com/tsenart/vegeta -l  | dot -Tpng -o vegeta.png
```
![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/tsenart-vegeta-l.png)

### [docker](https://github.com/dotcloud/docker)

```
# use level option 
go get github.com/dotcloud/docker
goviz -i github.com/dotcloud/docker/docker -d 1 -seek-in github.com/dotcloud/docker| dot -Tpng -o docker-d1.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/dotcloud-docker-docker-seek-in-dotcloud-docker--d-1.png)

```
# use level option 
go get github.com/dotcloud/dockerå
goviz -i github.com/dotcloud/docker/docker -d 2 -seek-in github.com/dotcloud/docker| dot -Tpng -o docker-d2.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/dotcloud-docker-docker-seek-in-dotcloud-docker--d-2.png)

```
# use level option 
go get github.com/dotcloud/docker
goviz -i github.com/dotcloud/docker/docker -d 3 -seek-in github.com/dotcloud/docker| dot -Tpng -o docker-d3.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/dotcloud-docker-docker-seek-in-dotcloud-docker--d-3.png)

```
# see only execdriver
go get github.com/dotcloud/docker
goviz -i github.com/dotcloud/docker/runtime/execdriver/execdrivers -seek-in github.com/dotcloud/docker | dot -Tpng -o docker-executiondrivers.png
```

![](https://raw.githubusercontent.com/hirokidaichi/goviz/master/images/dotcloud-docker-runtime-execdriver-execdrivers--seek-in-dotcloud-docker.png)
## Ricense

MIT

## Author

hirokidaichi [at] gmail.com



