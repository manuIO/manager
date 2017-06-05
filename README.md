# Mainflux manager

[![license][badge:license]](LICENSE)
[![build][badge:ci]][www:ci]
[![go report card][badge:grc]][www:grc]

Mainflux manager provides an HTTP API for managing platform resources: users,
devices, applications and channels. Through this API clients are able to do
the following actions:

- register new accounts and obtain access tokens
- provision new devices
- provision new applications
- create new channels
- "plug" devices and applications into the channels

For in-depth explanation of the aforementioned scenarios, as well as thorough
understanding of Mainflux, please check out the [official documentation][doc].

[badge:license]: https://img.shields.io/badge/license-Apache%20v2.0-blue.svg
[badge:ci]: https://travis-ci.org/mainflux/manager.svg?branch=master
[badge:grc]: https://goreportcard.com/badge/github.com/mainflux/manager
[doc]: http://mainflux.io
[www:ci]: https://travis-ci.org/mainflux/manager
[www:grc]: https://goreportcard.com/report/github.com/mainflux/manager
