# Middware Demo

We work at a super secret spy agency.
Our super secret spies need to be able to access top secret information about where they need to drop the packagefor their super secret spy friends.
Let's provide that information to them in a not-so-secure way by serving that info at 127.0.0.1:8080/rendezvous

The only way they can see it is if they know the super secret passcode.

To do this we will use http middleware!

## For Demo

### no middleware:
```console
$ go run no_middleware.go &
$ curl -H "Password: 007" localhost:8080/rendezvous
```

### with middleware:
```console
$ go run middleware.go &
$ curl -v localhost:8080/rendezvous
$ curl -H "Password: 007" localhost:8080/rendezvous
$ curl -H "Password: 007" localhost:8080/enemy-spies
$ curl  -H "Password: 007" localhost:8080/secure-line
```