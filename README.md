# SMS2

SMS2 - stands for a simple in memory storage written in go.

###### There are two cache providers:
**Fixed** - a realisation based on a priority heap with a fixed size, when the heap overflows the item with smallest ttl overwrites. 
```sh
$ ./sms2 fixed 100 60 #First argument is capacaty, the second is ttl(sec) 
```
**Agile** - a ttl realisation based on goroutines that start right after the `Set` operation and delete the expired item when ttl is over. For reference see `agile.go:Set` 
```sh
$ ./sms2 agile
```

### Comand examples

```sh
$ telnet localhost 5555
```

| Command | Arguments
| ------ | ------ |
| set | myname sergei |
| get | myname |
| keys |  |
| remove | myname |
| capacity |  |
| lset | listname [1,2,3,4,5] |
| ladd | listname 6 |
| lget | listname [3:5] |
| lremove | listname [3:5] |

#### agile cache type (per key ttl for all set operations):
| Command | Arguments
| ------ | ------ |
| set | myname sergei 60 |


The http server is implemented as well. Send Post to `/` with body like: 
`operation=set&key=first9&value=42&ttl=60`

### Tech

SMS2 uses a number of open source libraries:

* [go-telnet] - a telnet server api in a style similar to the "net/http"
* [log15] - a simple toolkit for best-practice logging
------------------------------


### Implemented:
 - Two types of cache: a)heap based. per heap ttl. b) goroutine based. per key ttl. You can choose the cache type by passing arguments into the main function. 
 - string and list support
 - http and telnet servers
 - a simple performance test
 - dockerization 

### TODO:
 - add dictionary support
 - add full functionality support to the http server and http/telnet clients
 - create a more realistic performance test suit




   [go-telnet]: <https://github.com/reiver/go-telnet>
   [log15]: <https://github.com/inconshreveable/log15>
