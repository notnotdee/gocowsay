#### Instructions

Currently designed to work with any CLI tool (`fortune` is an easy one) or echoed text that can pipe its stdout into the stdin of the `gocowsay` program.

**To run the program within the project directory** (after cloning)
```
$ go build -o gocowsay
```

```
$ sudo apt install fortune
```

```
$ fortune | ./gocowsay
```

Or, 
```
$ echo 'your text here' | ./gocowsay
```


**To create a global executable**

_From within the project directory:_
```
$ go build -o gocowsay
```

```
$ chmod +x ./gocowsay
```

```
$ sudo cp ./gocowsay /usr/local/bin
```

_From any location:_
```
$ sudo apt install fortune
```

```
$ fortune | gocowsay
```

Or, 
```
$ echo 'your text here' | gocowsay
```


#### Next 
- [ ] Add flags logic to change animal avatar or animal eyes
- [ ] Publish package to pkg.go.dev such that `$ go get -u github.com/dl-watson/gocowsay` works as expected

#### Screenshots

<img src="img/screenshot.png">
