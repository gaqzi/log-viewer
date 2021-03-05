log-viewer (lv)
===============

A tool to pretty print and colorize a stream of JSON logs, optionally split by a prefix.

For example, to use with `docker-compose logs`:

```shell
$ docker-compose logs -f my-cool-app | lv -sep ' : '
```
