# Config Store

master:  [![CircleCI](https://circleci.com/gh/ernestio/config-store/tree/master.svg?style=shield)](https://circleci.com/gh/ernestio/config-store/tree/master)  
develop: [![CircleCI](https://circleci.com/gh/ernestio/config-store/tree/develop.svg?style=shield)](https://circleci.com/gh/ernestio/config-store/tree/develop)

## Synopsis

This service is responsible for allowing other service to access and set configuration data for shared resources like databases and message queues.

## Usage

The getting and setting of Ernest configuration over NATS is dependant upon the structure of the NATS subject.

`config.get.<service name>` - gets the configuration for the specified service.
`config.set.<service name>` - sets the configuration for the specified service.



This can be accomplished in Go using:

```
// get configuration for a service
n, _ := nats.Connect(natsURI)
resp, _ := n.Request("service.get.redis", nil, time.Second)

// set configuration for a service
n, _ := nats.Connect(natsURI)
resp, _ := n.Request("service.set.redis", "{"hostname": "redis"}", time.Second)
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

Code and documentation copyright since 2015 ernest.io authors.

Code released under
[the Mozilla Public License Version 2.0](LICENSE).
