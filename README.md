# Package pbm [![PkgGoDev](https://pkg.go.dev/badge/github.com/slashformotion/pbm)](https://pkg.go.dev/github.com/slashformotion/pbm) [![Go Report Card](https://goreportcard.com/badge/github.com/slashformotion/pbm)](https://goreportcard.com/report/github.com/slashformotion/pbm) [![Tests](https://github.com/slashformotion/pbm/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/slashformotion/pbm/actions/workflows/test.yml)


```
import "github.com/slashformotion/pbm"
```
Package pbm implements a Portable Bit Map (PBM) image decoder and encoder. The supported image color model is [color.RGBAModel](https://pkg.go.dev/image/color#RGBAModel).

The PBM specification is at http://netpbm.sourceforge.net/doc/pbm.html.


## func [Decode](reader.go#L28)
<pre>
func Decode(r <a href="https://pkg.go.dev/io">io</a>.<a href="https://pkg.go.dev/io#Reader">Reader</a>) (<a href="https://pkg.go.dev/image">image</a>.<a href="https://pkg.go.dev/image#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)
</pre>
Decode reads a PBM image from Reader r and returns it as an image.Image.


## func [DecodeConfig](reader.go#L39)
<pre>
func DecodeConfig(r <a href="https://pkg.go.dev/io">io</a>.<a href="https://pkg.go.dev/io#Reader">Reader</a>) (<a href="https://pkg.go.dev/image">image</a>.<a href="https://pkg.go.dev/image#Config">Config</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)
</pre>
DecodeConfig returns the color model and dimensions of a PBM image without decoding the entire image.


## func [Encode](writer.go#L15)
<pre>
func Encode(w <a href="https://pkg.go.dev/io">io</a>.<a href="https://pkg.go.dev/io#Writer">Writer</a>, img <a href="https://pkg.go.dev/image">image</a>.<a href="https://pkg.go.dev/image#Image">Image</a>) <a href="https://pkg.go.dev/builtin#error">error</a>
</pre>
Encode writes the Image img to Writer w in PBM format.
