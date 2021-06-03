#### gocowsay

A [cowsay](https://en.wikipedia.org/wiki/Cowsay) implementation in Go. 

#### Instructions

Currently designed to work with any CLI program (like `fortune`) that can pipe its stdout into the stdin of the `gocowsay` executable. 

From within the project directory:
```
$ go build -o gocowsay
```

```
$ chmod +x ./gocowsay
```

```
$ sudo cp ./gocowsay /usr/local/bin
```

From any location: 
```
$ sudo apt install fortune
```

```
$ fortune | gocowsay
```

#### Next 
- [ ] Add flags logic to change animal avatar or animal eyes
- [ ] Publish package to pkg.go.dev such that `$ go get -u github.com/dl-watson/gocowsay` works as expected

#### Screenshots