### mitum-storage

*mitum-storage* is a storage contract model based on the second version of mitum(aka [mitum2](https://github.com/ProtoconNet/mitum2)).

#### Installation

Before you build `mitum-storage`, make sure to run `docker run` for digest api.

```sh
$ git clone https://github.com/ProtoconNet/mitum-storage

$ cd mitum-storage

$ go build -o ./mitum-storage
```

#### Run

```sh
$ ./mitum-storage init --design=<config file> <genesis file>

$ ./mitum-storage run <config file> --dev.allow-consensus
```

[standalong.yml](standalone.yml) is a sample of `config file`.

[genesis-design.yml](genesis-design.yml) is a sample of `genesis design file`.
