go-tls-watch
============

ssl certificate validation service


Usage
-----

start the application:

```bash
me@t:~$ go run cmd/gotlswatch/main.go
```

and then start asking for certificate checks:

```bash
me@t:~$ curl -XPOST --data "https://google.de" http://localhost:9991/api/validate
{"website":"https://google.de","valid":true,"expires":"2019-04-23T14:58:00Z"}
me@t:~$
```

