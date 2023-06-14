# Middware Demo

Let's create a server with route 127.0.0.1:8080/rendezvous

Our super secret spies need to be able to access top secret information about where they need to meet their super secret spy friends.
Let's provide that information to them in a not-so-secure way by serving that info at 127.0.0.1:8080/rendezvous

The only way they can see it is if they know the super secret passcode. 

We also want to know how long it took them to get the response.

To do this we will make two middleware functions (BadSecurity & RequestTimer) and chain them to the http Handler.

## For Demo:

### no middleware:
```console
$ go run no_middleware.go &
$ curl -H "Password: 007" localhost:8080/rendezvous
```

### with middleware:
```console
$ go run middleware.go &
$ curl localhost:8080/rendezvous
$ curl localhost:8080/enemy-spies
$ curl localhost:8080/secure-line
```