# Freq - Frequency Analysis API

Perform frequency analysis on the letters in a piece of text.  Works in the console or via API.

## Command Line Usage

First compile the binary:

```
$ go build
```

You can provide the text to be analyzed as an argument to the command:

```
$ ./freq --text "You have power over your mind - not outside events. Realize this, and you will find strength. - Marcus Aurelius"
```

Alternatively, you can specify a file to be analyzed:

```
$ ./freq --file _data/marcus.txt
```

## API Usage

First compile the binary:

```
$ go build
```

Launch the server:

```
$ ./freq --serve --host 0.0.0.0 --port 8080
```

If not otherwise specified, the default host will be `0.0.0.0` and the default port will be `80`.

You can test the API response with curl:

```
$ curl -X POST -id "text=You have power over your mind - not outside events. Realize this, and you will find strength. - Marcus Aurelius" localhost:8080/freq
```

# Changes from Kurtis

## Split code across files

The namespacing (package) in go is awesome. If a file is in the same folder and has the same package name, then there's no need to import it, it'll be auto included by the compiler. I've split all your various functions out into their own files so that it's easier to see all the different functionality while looking at the code files.

I pulled all the code related to LetterSet out into its own file, so now we can see the type and methods that this struct has.

## Change exported vs. unexported

If code is in the same package, and it doesn't need to work with any standard library packages, then there is no reason to have it be exported (start with a cap letter).
