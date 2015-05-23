// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package identicon

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"math"
)

const (
	minSize = 16
)

type Identicon struct {
	fore, back color.Color
	size       int
}

// size九宫格中每个格子的像素
func New(back, fore color.Color, size int) (*Identicon, error) {
	if size < minSize {
		return nil, fmt.Errorf("New:产生的图片尺寸(%v)不能小于%v", size, minSize)
	}

	return &Identicon{
		fore: fore,
		back: back,
		size: size,
	}, nil
}

func (i *Identicon) Make(data []byte) image.Image {
	return makeImage(i.back, i.fore, i.size, data)
}

func Make(back, fore color.Color, size int, data []byte) (image.Image, error) {
	if size < minSize {
		return nil, fmt.Errorf("New:产生的图片尺寸(%v)不能小于%v", size, minSize)
	}

	return makeImage(back, fore, size, data), nil
}

func makeImage(back, fore color.Color, size int, data []byte) image.Image {
	h := md5.New()
	h.Write(data)
	sum := h.Sum(nil)

	// 第一个方块
	index := int(math.Abs(float64(sum[0]+sum[1]+sum[2]+sum[3]))) % len(blocks)
	b1 := blocks[index]

	// 第二个方块
	index = int(math.Abs(float64(sum[4]+sum[5]+sum[6]+sum[7]))) % len(blocks)
	b2 := blocks[index]

	// 中间方块
	index = int(math.Abs(float64(sum[8]+sum[9]+sum[10]+sum[11]))) % len(centerBlocks)
	c := centerBlocks[index]

	// 旋转角度
	angle := int8(math.Abs(float64(sum[12]+sum[13]+sum[14]+sum[15]))) % 4

	p := image.NewPaletted(image.Rect(0, 0, size, size), []color.Color{back, fore})
	draw(p, size, c, b1, b2, angle)
	return p
}

// 将完整的头像画到p上。
func draw(p *image.Paletted, size int, c, b1, b2 blockFunc, angle int8) {
	blockSize := float64(size / 3) // 每个格子的长宽

	incr := func() { // 增加angle的值，但不会大于3
		if angle > 2 {
			angle = 0
		} else {
			angle++
		}
	}

	c(p, blockSize, blockSize, blockSize, 0)

	b1(p, 0, 0, blockSize, angle)
	b2(p, blockSize, 0, blockSize, angle)

	incr()
	b1(p, 2*blockSize, 0, blockSize, angle)
	b2(p, 2*blockSize, blockSize, blockSize, angle)

	incr()
	b1(p, 2*blockSize, 2*blockSize, blockSize, angle)
	b2(p, blockSize, 2*blockSize, blockSize, angle)

	incr()
	b1(p, 0, 2*blockSize, blockSize, angle)
	b2(p, 0, blockSize, blockSize, angle)
}
