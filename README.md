# discountPhotoShop

When you're on a tight budget but need a drawing tool

### Installation

discountPhotoShop requires latest [Golang](https://golang.org/doc/install) to run.

### Run the app
Go to project folder and type:

```sh
$ cd discountPhotoShop
$ go run main.go
```

Create a canvas:
```sh
C 22 4
```

Create a line:
```sh
L 1 3 6 3
```

Create another line:
```sh
L 7 3 7 4
``` 

Create a rectangle:
```sh
R 14 1 18 3
``` 

Create a fill:
```sh
B 10 3 o
``` 

Quit the app:
```sh
Q
``` 

You can place the commands on one line
```sh
C 22 4 L 1 3 6 3 L 7 3 7 4 R 14 1 18 3 B 10 3 o
```

Or run the app and create automatically the drawing

```sh
$ go run main.go 
$ C 20 4 L 1 2 6 2 L 6 3 6 4 R 14 1 18 3 B 10 3 o
```

Run the app and create automatically the drawing then quit

```sh
$ go run main.go 
$ C 20 4 L 1 2 6 2 L 6 3 6 4 R 14 1 18 3 B 10 3 o Q
```

Or just draw a star
```sh
$ go run main.go 
$ C 27 23 L 1 23 13 1 L 13 1 25 23 L 25 23 1 8 L 1 8 25 8 L 26 8 1 23 Q
```

### Tests
Run the tests:
```sh
$ go test -cover ./... -v
```
### Documentation
Run the docs:
```sh
$ godoc -http=":6060"
```
Then visit: [http://localhost:6060/pkg/github.com/clybs/](http://localhost:6060/pkg/github.com/clybs/)