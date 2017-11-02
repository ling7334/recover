# recover
[![Build Status](https://travis-ci.org/ling7334/recover.svg?branch=master)](https://travis-ci.org/ling7334/recover) [![codecov](https://codecov.io/gh/ling7334/recover/branch/master/graph/badge.svg)](https://codecov.io/gh/ling7334/recover)

recover form panic, convert Panicking information to error, and pass to next level. Prevent panic crash the program.

## Useage
```golang
  package main

  import "github.com/ling7334/recover"

  var revr = recover.Recover
  
  func main() {
    var err error
    defer revr(&err)
    panic("test panic")
  }
```
