// medoly - a clone of much better games
// Copyright (c) 2022 Michael D Henderson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// cmdPng runs the PNG command
var cmdPng = &cobra.Command{
	Use:   "png",
	Short: "png things",
	Run: func(cmd *cobra.Command, args []string) {
		src, err := os.ReadFile("bressea - 638bb263ac4546a8.png")
		if err != nil {
			log.Fatal(err)
		}
		img, err := png.Decode(bytes.NewReader(src))
		if err != nil {
			log.Fatal(err)
		} else if err = toText(img); err != nil {
			log.Fatal(err)
		} else if err = toPng(img); err != nil {
			log.Fatal(err)
		}
	},
}

func toPng(im image.Image) error {
	name := "medoly.png"
	const width, height = 256, 256

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		_ = f.Close()
		return err
	} else if err = f.Close(); err != nil {
		return err
	}
	log.Printf("created %s\n", name)
	return nil
}

func toText(img image.Image) error {
	levels := []string{" ", "░", "▒", "▓", "█"}
	buf := &bytes.Buffer{}
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		if y > 500 {
			break
		}
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if x > 500 {
				break
			}
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			buf.WriteString(levels[level])
		}
		buf.WriteByte('\n')
	}
	name := "bressea - 638bb263ac4546a8.txt"
	if err := os.WriteFile(name, buf.Bytes(), 0666); err != nil {
		return err
	}
	log.Printf("created %s\n", name)
	return nil
}

var argsPng struct {
}

func init() {
	cmdRoot.AddCommand(cmdPng)
}
