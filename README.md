# SMS2

SMS2 - stands for simple in memory storage written in go.

### Tech

SMS2 uses a number of open source libraries:

* [go-telnet] - telnet server api in a style similar to the "net/http"
* [log15] - simple toolkit for best-practice logging

### Run
There are two cache providers.

**Fixed** - realisation based on fixed size priority heap, when overflows the item with least ttl overwrites. Items in heap sorted in ttl order.
Running fixedCache provider.
```sh
$ ./sms2 fixed 100 60
```
Capasity is 100. Ttl is 60 seconds.


**Agile** -  realisation based on groutines that are starting right after `Set` operation is done and waiting until ttl time is over - then deleting the item.
```sh
$ ./sms2 agile
```

### Comand exmples
#### fixed cache type:
| Command | Arguments
| ------ | ------ |
| set | myname sergei |
| get | myname |
| keys |  |
| remove | myname |
| capacity |  |

#### agile cache type:
| Command | Arguments
| ------ | ------ |
| set | myname sergei 60 |

For agile cache type we need ttl for each `set` command.

Also you can use http Server and send post request with body:
operation=set&key=first9&value=42&ttl=60

### Todos

 - Write unit Tests
 - File backup functionality
 - Performance tests

License
----

MIT


   [go-telnet]: <https://github.com/reiver/go-telnet>
   [log15]: <https://github.com/inconshreveable/log15>

