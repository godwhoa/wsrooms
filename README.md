Websocket rooms
===============

A rewrite of an old websocket server with rooms.

Testing it out
==============
```
# install
go get -v github.com/godwhoa/wsrooms
go get -v github.com/lafikl/telsocket

# run server
go run *.go

# connect the client
telsocket -url ws://localhost:8080/ws/test
```

TODO:
=====

+ ~~Clients talking to other clients~~
+ Fix the crash