// Copyright 2015 by caixw, All rights reserved
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package identicon

import (
	"image"
)

// 所有block函数的类型
type blockFunc func(img *image.Paletted, x, y, size float64, angle int8)

var (
	// 可以出现在中间的方块，必须是对称的
	centerBlocks = []blockFunc{b0, b1, b2, b3}

	// 所有方块
	blocks = []blockFunc{b0, b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11, b12, b13, b14, b15}
)

// 将多边形points旋转angle个角度，然后输出到img上，起点为x,y坐标
func draw(img *image.Paletted, x, y, size float64, angle int8, points [][]float64) {
	m := size / 2
	for index, point := range points {
		x1, y1 := rotate(point[0], point[1], x+m, y+m, angle)
		points[index] = []float64{x1, y1}
	}

	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			if pointInPolygon(i, j, points) {
				img.SetColorIndex(int(i), int(j), 1)
			}
		}
	}
}

// 全空白
//
//  --------
//  |      |
//  |      |
//  |      |
//  --------
func b0(img *image.Paletted, x, y, size float64, angle int8) {
}

// 全填充正方形
//
//  --------
//  |######|
//  |######|
//  |######|
//  --------
func b1(img *image.Paletted, x, y, size float64, angle int8) {
	isize := int(size)
	ix := int(x) + 1 // 留一白边
	iy := int(y) + 1
	for i := ix; i < ix+isize-1; i++ {
		for j := iy; j < iy+isize-1; j++ {
			img.SetColorIndex(i, j, 1)
		}
	}
}

// 中间小空白
//  ----------
//  |        |
//  |  ####  |
//  |  ####  |
//  |        |
//  ----------
func b2(img *image.Paletted, x, y, size float64, angle int8) {
	l := size / 4
	x = x + l
	y = y + l

	for i := x; i < x+2*l; i++ {
		for j := y; j < y+2*l; j++ {
			img.SetColorIndex(int(i), int(j), 1)
		}
	}
}

// b3
//
//  ---------
//  |   #   |
//  |  ###  |
//  | ##### |
//  |#######|
//  | ##### |
//  |  ###  |
//  |   #   |
//  ---------
func b3(img *image.Paletted, x, y, size float64, angle int8) {
	m := size / 2
	points := [][]float64{{x + m, y}, {x + size, y + m}, {x + m, y + size}, {x, y + m}}

	for i := x; i < x+size; i++ {
		for j := y; j < x+size; j++ {
			if pointInPolygon(i, j, points) {
				img.SetColorIndex(int(i), int(j), 1)
			}
		}
	}
}

// b4
//
//  -------
//  |#####|
//  |#### |
//  |###  |
//  |##   |
//  |#    |
//  |------
func b4(img *image.Paletted, x, y, size float64, angle int8) {
	points := [][]float64{{x, y}, {x + size, y}, {x, y + size}}
	draw(img, x, y, size, angle, points)
}

// b5
//
//  ---------
//  |   #   |
//  |  ###  |
//  | ##### |
//  |#######|
func b5(img *image.Paletted, x, y, size float64, angle int8) {
	m := size / 2
	points := [][]float64{{x + m, y}, {x + size, y + size}, {x, y + size}}
	draw(img, x, y, size, angle, points)
}

// b6 矩形
//
//  --------
//  |###   |
//  |###   |
//  |###   |
//  --------
func b6(img *image.Paletted, x, y, size float64, angle int8) {
	m := size / 2
	points := [][]float64{{x, y}, {x + m, y}, {x + m, y + size}, {x, y + size}}
	draw(img, x, y, size, angle, points)
}

// b7 斜放的锥形
//
//  ---------
//  | #     |
//  |  ##   |
//  |  #####|
//  |   ####|
//  |--------
func b7(img *image.Paletted, x, y, size float64, angle int8) {
	m := size / 2
	points := [][]float64{{x, y}, {x + size, y + m}, {x + size, y + size}, {x + m, y + size}}
	draw(img, x, y, size, angle, points)
}

// b8 三个堆叠的三角形
//
//  -----------
//  |    #    |
//  |   ###   |
//  |  #####  |
//  |  #   #  |
//  | ### ### |
//  |#########|
//  -----------
func b8(img *image.Paletted, x, y, size float64, angle int8) {
	m := size / 2
	mm := m / 2

	// 顶部三角形
	points := [][]float64{{x + m, y}, {x + 3*mm, y + m}, {x + mm, y + m}}
	draw(img, x, y, size, angle, points)

	// 底下左边
	points = [][]float64{{x + mm, y + m}, {x + m, y + size}, {x, y + size}}
	draw(img, x, y, size, angle, points)

	// 底下右边
	points = [][]float64{{x + 3*mm, y + m}, {x + size, y + size}, {x + m, y + size}}
	draw(img, x, y, size, angle, points)
}

// b9 斜靠的三角形
//
//  ---------
//  |#      |
//  | ####  |
//  |  #####|
//  |  #### |
//  |   #   |
//  ---------
func b9(img *image.Paletted, x, y, size float64, angle int8) {
	m := size / 2
	points := [][]float64{{x, y}, {x + size, y + m}, {x + m, y + size}}
	draw(img, x, y, size, angle, points)
}

// b10
//
//  ----------
//  |    ####|
//  |    ### |
//  |    ##  |
//  |    #   |
//  |####    |
//  |###     |
//  |##      |
//  |#       |
//  ----------
func b10(img *image.Paletted, x, y, size float64, angle int8) {
	m := size / 2
	points := [][]float64{{x + m, y}, {x + size, y}, {x + m, y + m}}
	draw(img, x, y, size, angle, points)

	points = [][]float64{{x, y + m}, {x + m, y + m}, {x, y + size}}
	draw(img, x, y, size, angle, points)
}

// b11 左上角1/4大小的方块
//
//  ----------
//  |####    |
//  |####    |
//  |####    |
//  |        |
//  |        |
//  ----------
func b11(img *image.Paletted, x, y, size float64, angle int8) {
	m := size / 2
	points := [][]float64{{x, y}, {x + m, y}, {x + m, y + m}, {x, y + m}}
	draw(img, x, y, size, angle, points)
}

// b12
//
//  -----------
//  |         |
//  |         |
//  |#########|
//  |  #####  |
//  |    #    |
//  -----------
func b12(img *image.Paletted, x, y, size float64, angle int8) {
	m := size / 2
	points := [][]float64{{x, y + m}, {x + size, y + m}, {x + m, y + size}}
	draw(img, x, y, size, angle, points)
}

// b13
//
//  -----------
//  |         |
//  |         |
//  |    #    |
//  |  #####  |
//  |#########|
//  -----------
func b13(img *image.Paletted, x, y, size float64, angle int8) {
	m := size / 2
	points := [][]float64{{x + m, y + m}, {x + size, y + size}, {x, y + size}}
	draw(img, x, y, size, angle, points)
}

// b14
//
//  ---------
//  |   #   |
//  | ###   |
//  |####   |
//  |       |
//  |       |
//  ---------
func b14(img *image.Paletted, x, y, size float64, angle int8) {
	m := size / 2
	points := [][]float64{{x + m, y}, {x + m, y + m}, {x, y + m}}
	draw(img, x, y, size, angle, points)
}

// b15
//
//  ----------
//  |#####   |
//  |###     |
//  |#       |
//  |        |
//  |        |
//  ----------
func b15(img *image.Paletted, x, y, size float64, angle int8) {
	m := size / 2
	points := [][]float64{{x, y}, {x + m, y}, {x, y + m}}
	draw(img, x, y, size, angle, points)
}