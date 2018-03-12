# SMS2

SMS2 - stands for simple in memory storage written in go.

### Run
There are two cache providers.

**Fixed** - realisation based on fixed size priority heap, when overflows the item with least ttl overwrites. Items in heap sorted in ttl order.
```sh
$ ./sms2 fixed 100 60
```
First argument is capacaty, second is ttl in seconds. 


**Agile** -  realisation based on groutines that are starting right after `Set` operation is done and waiting until ttl time is over - then deleting the item.
```sh
$ ./sms2 agile
```

### Comand exmples
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

#### agile cache type (per key ttl for all set operations):
| Command | Arguments
| ------ | ------ |
| set | myname sergei 60 |


Http server implemented as well. Send Post to `/` 
`operation=set&key=first9&value=42&ttl=60`

### Tech

SMS2 uses a number of open source libraries:

* [go-telnet] - telnet server api in a style similar to the "net/http"
* [log15] - simple toolkit for best-practice logging
------------------------------


### Implemented:
 - Two types of cache a)heap based. per heap ttl. b) goroutine based. per key ttl. You can choice cache type by passing arguments into main. 
 - string, list support.
 - http, telnet servers.
 - simple performance test
 - dockerization 

### TODO:
 - add dictionary support
 - add full functionality support to http server, and http/telnet clients
 - create more realistick performance test suit




   [go-telnet]: <https://github.com/reiver/go-telnet>
   [log15]: <https://github.com/inconshreveable/log15>