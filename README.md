# shellwrap

shell wrapper for Windows

## Usage

```
C:\>shellwrap echo $OS
Windows_NT
```

```
C:\>shellwrap notepad $(ver).txt
```

## Installation

```
$ go get github.com/mattn/shellwrap
```

If you want to run the process with shell (i.e. cmd.exe), use `-s` flag.

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
