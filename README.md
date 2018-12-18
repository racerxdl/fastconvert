[![Build Status](https://api.travis-ci.org/racerxdl/fastconvert.svg?branch=master)](https://travis-ci.org/racerxdl/fastconvert) [![Apache License](https://img.shields.io/badge/license-Apache-blue.svg)](https://tldrlegal.com/license/apache-license-2.0-(apache-2.0)) [![Go Report](https://goreportcard.com/badge/github.com/racerxdl/limedrv)](https://goreportcard.com/report/github.com/racerxdl/fastconvert)

# FastConvert

FastConvert is a library to be able to convert byte arrays to typed arrays in a faster way than `binary.Read`.

The thing about `binary.Read` is that it uses reflection to read data from an array and that makes it slow. How much slow? Well, check this benchmark:

```
┌─[lucas@nblucas] - [~/go/src/github.com/racerxdl/fastconvert] - [Tue Dec 18, 02:55]
└─[$] <git:(master)> go test -test.bench="." ./...
goos: linux
goarch: amd64
pkg: github.com/racerxdl/fastconvert
BenchmarkBinaryReadComplex64-8     	    2000	    857901 ns/op
BenchmarkFastConvertComplex64-8    	   20000	     83264 ns/op
BenchmarkBinaryReadComplex128-8    	    2000	    870663 ns/op
BenchmarkFastConvertComplex128-8   	   20000	     82742 ns/op
BenchmarkBinaryReadInt16-8         	    2000	    616818 ns/op
BenchmarkFastConvertInt16-8        	   30000	     41748 ns/op
BenchmarkBinaryReadInt32-8         	    2000	    649259 ns/op
BenchmarkFastConvertInt32-8        	   30000	     41402 ns/op
BenchmarkBinaryReadInt64-8         	    2000	    659684 ns/op
BenchmarkFastConvertInt64-8        	   30000	     41575 ns/op
PASS
ok  	github.com/racerxdl/fastconvert	17.834s
```

You can see that for `complex` types it is 10x faster than binary.Read, and for other primitives are 15x times faster.
That makes a LOT of difference when you have a loop that does this conversion (for example with my SDR Drivers which constantly receives data from the hardware).

The usage is pretty straightforward, check godoc for more info:

[https://godoc.org/github.com/racerxdl/fastconvert](https://godoc.org/github.com/racerxdl/fastconvert)

