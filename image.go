package gozbar

import (
	"image"
	"image/draw"
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <zbar.h>
import "C"

const y800 = 0x30303859

type Image struct {
	src       *image.Gray
	zbarImage *C.zbar_image_t
}

func NewImage(src image.Image) *Image {
	newImage := &Image{
		src:       image.NewGray(src.Bounds()),
		zbarImage: C.zbar_image_create(),
	}

	dims := newImage.src.Bounds().Size()
	C.zbar_image_set_size(newImage.zbarImage, C.uint(dims.X), C.uint(dims.Y))

	draw.Draw(
		newImage.src,
		newImage.src.Bounds(),
		src,
		image.ZP,
		draw.Over,
	)

	C.zbar_image_set_format(newImage.zbarImage, C.ulong(y800))
	C.zbar_image_set_data(
		newImage.zbarImage,
		unsafe.Pointer(&newImage.src.Pix[0]),
		C.ulong(len(newImage.src.Pix)),
		nil,
	)

	runtime.SetFinalizer(
		newImage,
		func(i *Image) {
			// The image data was allocated by the Go runtime, we
			// don't want zbar trying to free it
			C.zbar_image_set_data(newImage.zbarImage, nil, 0, nil)
			C.zbar_image_destroy(i.zbarImage)
		},
	)

	return newImage
}
