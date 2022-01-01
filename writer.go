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

package pbm

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"image/color"
	"io"

	"github.com/icza/bitio"
)

var errUnsupportedColorMode = errors.New("pbm: color mode not supported, please use 'color.RGBAModel'")

// Encode writes the Image img to Writer w in PBM format.
func Encode(w io.Writer, img image.Image) error {
	bw := bufio.NewWriter(w)

	switch img.ColorModel() {
	case color.RGBAModel:
		rec := img.Bounds()

		// write header
		fmt.Fprintf(bw, "P4\n%d %d\n", rec.Dx(), rec.Dy())
		bw.Flush()
		w := bitio.NewWriter(bw)
		defer func() {
			err := w.Close()
			if err != nil {
				panic(err)
			}

		}()

		// write pixels
		for y := rec.Min.Y; y < rec.Max.Y; y++ {
			for x := rec.Min.X; x < rec.Max.X; x++ {
				p := getBooleanFromColor(img.At(x, y))
				w.WriteBool(p)
			}
		}

	default:
		return errUnsupportedColorMode
	}
	return nil
}

func getBooleanFromColor(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	lum := (r + g + b) / 3
	if lum > uint32(0xffff)/2 {
		return false
	} else {
		return true
	}
}
