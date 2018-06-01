package main

import (
	"github.com/nfnt/resize"
	// "image/jpeg"
	"log"
	"os"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"image/png"
)


func main() {
	searchDir := "/Users/azim/Documents/2DGAMEASSETS/LowPoly_Pixel_RPG_Assets_DevilsGarage_v01/2D/view-Ortho/images"
	// imgFiles, err := run(searchDir)
	// if err != nil {
	// 	log.Fatal("find error: ",err)
	// }
	files, err := ioutil.ReadDir(searchDir)
  if err != nil {
    log.Fatal(err)
  }

  for index, f := range files {
    fmt.Println(f.Name())
    tmpFileName := fmt.Sprintf("%s/%s",searchDir,f.Name())
    file, err := os.Open(tmpFileName)
		if err != nil {
			log.Fatal("open error: ",err)
		}
		img, err := png.Decode(file)
		if err != nil {
			log.Fatal("decode error: ",err)
		}
		file.Close()

		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(64, 64, img, resize.Lanczos3)

		out, err := os.Create(fmt.Sprintf("%s/%d.png",searchDir,index))
		if err != nil {
			log.Fatal("filecreate error: ", err)
		}
		defer out.Close()

		// write new image to file
		png.Encode(out, m)

  }
	
}