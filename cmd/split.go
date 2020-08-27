package cmd

/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"image"
	_ "image/jpeg" // to read
	_ "image/png"  // to read
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:   "split [image path]",
	Short: "Split a big image in many small images",
	Args:  cobra.MinimumNArgs(1),
	Long: `Split a large JPEG or PNG image that should be 600 pixels
high by multiples of 600 pixels wide into many 600 x 600 pixel
images automatically in the same path adding the order number
at the end.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, path := range args {
			split(path)
		}
	},
}

// basePixels is 600 x 600 pixels
const basePixels = 600

func init() {
	rootCmd.AddCommand(splitCmd)
}

func split(path string) {
	validFormat(path)

	img := readImage(path)
	width, height := getImageDimension(img)
	validSize(width, height)

	frames := width / basePixels
	log.Print("Total images: ", frames)

	for frame := 0; frame < frames; frame++ {
		cropAndSave(path, img, frame)
	}
	log.Print("Done!")
}

func cropAndSave(imagePath string, img image.Image, frame int) {
	cropImg := imaging.Crop(img, image.Rect(frame*basePixels, 0, (frame+1)*basePixels, basePixels))

	// save cropped image
	outputPath := buildPath(imagePath, frame)
	log.Print("Saving to: ", outputPath)

	err := imaging.Save(cropImg, outputPath)
	fatalIf(err)
}

func readImage(imagePath string) image.Image {
	file, err := os.Open(imagePath)
	fatalIf(err, imagePath)

	image, _, err := image.Decode(file)
	fatalIf(err)

	return image
}

func buildPath(imagePath string, frame int) string {
	ext := path.Ext(imagePath)
	outfile := imagePath[0:len(imagePath)-len(ext)] + "_" + strconv.Itoa(frame) + ext
	return outfile
}

func validFormat(path string) {
	ext := filepath.Ext(path)
	if ext != ".jpg" && ext != ".png" {
		log.Fatal("Unsupported format!", path)
	}
}

func validSize(width int, height int) {
	if height != basePixels || width%basePixels != 0 {
		log.Fatal("Format not valid: ", width, height)
	}
}

func getImageDimension(img image.Image) (int, int) {
	size := img.Bounds().Size()
	return size.X, size.Y
}
