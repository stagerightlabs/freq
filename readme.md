# Freq - Frequency Analysis API

Generate a character frequency analysis for a given piece of text.  This tool can be used on the command line or via an API.

## Command Line Usage

First compile the binary:

```
$ go build
```

You can provide the text to be analyzed as an argument to the command:

```
$ ./freq --text "You have power over your mind - not outside events. Realize this, and you will find strength."
```

Alternatively, you can specify a file to be analyzed:

```
$ ./freq --file _data/monte_cristo.txt
```
