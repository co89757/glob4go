# glob4go
[![Build Status](https://travis-ci.org/co89757/glob4go.svg?branch=master)](https://travis-ci.org/co89757/glob4go)

A lightweight glob wildcard pattern matching utility for Go, when regex is an overkill 
## Installation 
```
go get github.com/co89757/glob4go 
```
## Supported Glob feature
The glob syntax is largely similar to UNIX globbing execept for a few unsupported syntax 
 - `*` :matches any number of characters 
 - `?` :matches a single character
 - `[...]` :matches a range/group of characters. For example, `[0-9]` matches any single-digit number, `[axz]` matches any character in 'a','x','z'
 - `[^...]` :inverse range/group match. The inverse of the above range match case. It matches any character not in the range 
 - all keywords above can be escaped by `\`, i.e. `\?` matches a literal `?` 

## Example
Currently, the API of the pacakge is really simple:
```go
glob4go.Glob(pattern, str []byte, ignoreCase bool) 
```
Here is an example usage 

```go 
import (
  "github.com/co89757/glob4go"
)
func globExample(){
  var match bool 
pattern := []byte("abcd[0-9]?")
s1 := []byte("abcd8e")
//case sensitive match
match = glob4go.Glob(pattern,s1,false) // true, it is a match 

s2 := []byte("AbCd8e")
match = glob4go.Glob(pattern, s2, true) // true, it is a match when ignore case 
}

```

## TO-DO
* Performance and benchmarking 
* Add unicode support 