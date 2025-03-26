// main.go
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

func main() {
	fmt.Println("GoCV测试程序 - Windows ARM64")
	fmt.Printf("GoCV版本: %s\n", gocv.Version())
	fmt.Printf("OpenCV版本: %s\n", gocv.OpenCVVersion())

	// 创建测试图像
	mat := gocv.NewMatWithSize(300, 400, gocv.MatTypeCV8UC3)
	defer mat.Close()

	// 填充白色背景
	mat.SetTo(gocv.NewScalar(255, 255, 255, 0))

	// 绘制一些形状
	blue := color.RGBA{0, 0, 255, 0}
	red := color.RGBA{255, 0, 0, 0}
	green := color.RGBA{0, 255, 0, 0}

	// 绘制圆形
	gocv.Circle(&mat, image.Pt(200, 150), 100, blue, 2)

	// 绘制矩形
	gocv.Rectangle(&mat, image.Rect(100, 50, 300, 250), green, 2)

	// 绘制线条
	gocv.Line(&mat, image.Pt(100, 100), image.Pt(300, 200), red, 2)

	// 添加文字
	gocv.PutText(&mat, "GoCV ARM64", image.Pt(110, 280), gocv.FontHersheyPlain, 1.5, blue, 2)

	// 保存图像
	filename := "gocv-test.jpg"
	gocv.IMWrite(filename, mat)
	fmt.Printf("已保存测试图像到: %s\n", filename)

	// 尝试显示窗口(如果有GUI)
	flag.Parse()
	if flag.NArg() == 0 || flag.Arg(0) == "gui" {
		window := gocv.NewWindow("GoCV Test")
		defer window.Close()

		fmt.Println("显示图像窗口 - 按ESC键退出")
		for {
			window.IMShow(mat)
			if window.WaitKey(100) == 27 { // ESC键
				break
			}
		}
	}

	fmt.Println("测试完成!")
}
