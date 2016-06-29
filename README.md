# Config Store

## Synopsis

This service is responsible for allowing other service to access configuration data for shared resources like databases and message queues.

## Build status

* Master: [![CircleCI](https://circleci.com/gh/ErnestIO/config-store/tree/master.svg?style=svg)](https://circleci.com/gh/ErnestIO/config-store/tree/master)
* Develop: [![CircleCI](https://circleci.com/gh/ErnestIO/config-store/tree/develop.svg?style=svg)](https://circleci.com/gh/ErnestIO/config-store/tree/develop)


## Usage

To make use of this to get configuration over nats, you can use `config.get.SERVICE`, where `SERVICE` is the name of the service you want to get configuration for.

This can be accomplished with go using:

```
n, _ := nats.Connect(natsURI)
resp, _ := n.Request("service.get.redis", nil, time.Second)
```

## Installing

```
$ make deps
$ make install
```

## Running

```
$ NATS_URI='nats://localhost:4222' config-store -config=/etc/ernest/config.json
```
Please note, if no conf dir is specified, the local directory will be used.


## Tests

Running the tests:
```
make test
```

## Contributing

Please read through our
[contributing guidelines](CONTRIBUTING.md).
Included are directions for opening issues, coding standards, and notes on
development.

Moreover, if your pull request contains patches or features, you must include
relevant unit tests.

## Versioning

For transparency into our release cycle and in striving to maintain backward
compatibility, this project is maintained under [the Semantic Versioning guidelines](http://semver.org/).

## Copyright and License

Code and documentation copyright since 2015 r3labs.io authors.

Code released under
[the Mozilla Public License Version 2.0](LICENSE).
