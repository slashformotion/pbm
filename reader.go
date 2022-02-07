// Package pbm implements a Portable Bit Map. (PBM) image decoder and encoder. The supported image
// color model is color.RGBAModel.
//
// The PBM specification is at http://netpbm.sourceforge.net/doc/pbm.html.
package pbm

// Copyright [2022] slashformotion

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"bufio"
	"bytes"
	"errors"
	"image"
	"image/color"
	"io"
	"strconv"

	"github.com/icza/bitio"
)

func init() {
	image.RegisterFormat("pbm", "P4", Decode, DecodeConfig)
}

var (
	errBadHeader = errors.New("pbm: invalid header")
	errNotEnough = errors.New("pbm: not enough image data")
)

// Decode reads a PBM image from Reader r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
	var d decoder
	img, err := d.decode(r, false)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// DecodeConfig returns the dimensions of a PBM image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (image.Config, error) {
	var d decoder
	if _, err := d.decode(r, true); err != nil {
		return image.Config{}, err
	}
	return image.Config{
		ColorModel: color.RGBAModel,
		Width:      d.width,
		Height:     d.height,
	}, nil
}

type decoder struct {
	br *bufio.Reader

	// from header
	magicNumber string
	width       int
	height      int
}

func (d *decoder) decode(r io.Reader, configOnly bool) (image.Image, error) {
	d.br = bufio.NewReader(r)

	// decode header
	err := d.decodeHeader()
	if err != nil {
		return nil, err
	}
	if configOnly {
		return nil, nil
	}

	// decode image
	img := image.NewRGBA(image.Rect(0, 0, d.width, d.height))
	bitReader := bitio.NewReader(d.br)
	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {
			pixel, err := bitReader.ReadBits(1)
			// fmt.Print(pixel)
			pixel = pixel & 1
			if err != nil {
				return nil, errNotEnough
			}

			img.SetRGBA(x, y, getColorFromBoolean(pixel))
		}
	}
	return img, nil
}

func getColorFromBoolean(b uint64) color.RGBA {
	var pixelColor byte = 0xff
	if b != 0 {
		pixelColor = 0x0
	}
	return color.RGBA{pixelColor, pixelColor, pixelColor, 0xff}
}

func (d *decoder) decodeHeader() error {
	var err error
	var b byte
	header := make([]byte, 0)

	comment := false
	for fields := 0; fields < 3; {
		b, err = d.br.ReadByte()
		if err != nil {
			return errBadHeader
		}

		if b == '#' {
			comment = true
		} else if !comment {
			header = append(header, b)
		}
		if comment && b == '\n' {
			comment = false
		} else if !comment && (b == ' ' || b == '\n' || b == '\t') {
			fields++
		}
	}
	headerFields := bytes.Fields(header)

	d.magicNumber = string(headerFields[0])
	if d.magicNumber != "P4" {
		return errBadHeader
	}
	d.width, err = strconv.Atoi(string(headerFields[1]))
	if err != nil {
		return errBadHeader
	}
	d.height, err = strconv.Atoi(string(headerFields[2]))
	if err != nil {
		return errBadHeader
	}
	return nil
}
