package transcode

import (
	"fmt"
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
	"log"
)

func Transcode() {
	filename := "test.mov"

	// Register all formats and codecs
	avformat.AvRegisterAll()

	ctx := avformat.AvformatAllocContext()
	//outctx := avformat.AvformatAllocContext()

	// Open video file
	if avformat.AvformatOpenInput(&ctx, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctx.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")

		// Close input file and free context
		ctx.AvformatCloseInput()
		return
	}
	for i:=0; i<10; i++ {
		fmt.Println()
	}
	//avformat.AvformatAllocOutputContext2(&outctx, ctx.Oformat(), "", "")
	a := avcodec.AvcodecFindEncoderByName("libx265")
	fmt.Println(a)
}
