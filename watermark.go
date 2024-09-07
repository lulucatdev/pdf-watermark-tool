package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

//go:embed watermark.png
var embeddedFiles embed.FS

// 主函数
func main() {
	// 定义命令行标志
	opacity := flag.Float64("opacity", 0.15, "水印透明度 (0.0 - 1.0)")
	scale := flag.Float64("scale", 1.0, "水印缩放比例")
	watermarkFile := flag.String("watermark", "", "水印PNG文件路径（可选）")

	// 解析命令行标志
	flag.Parse()

	// 检查必需的参数
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("用法: watermark [flags] <input.pdf> [output.pdf]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// 获取输入文件路径
	inputFile := args[0]
	var outputFile string
	if len(args) > 1 {
		// 如果提供了输出文件名或路径，则使用它
		outputFile = args[1]
		// 检查是否提供了完整路径
		if !filepath.IsAbs(outputFile) && !strings.Contains(outputFile, string(os.PathSeparator)) {
			// 如果只提供了文件名，则保存在当前目录
			outputFile = filepath.Join(".", outputFile)
		}
	} else {
		// 如果没有提供输出文件路径，则创建一个默认的输出文件夹
		watermarkFolder := "watermarked_files"
		if err := os.MkdirAll(watermarkFolder, os.ModePerm); err != nil {
			fmt.Printf("创建水印文件夹时出错: %v\n", err)
			os.Exit(1)
		}

		// 在水印文件夹中使用原始文件名
		outputFile = filepath.Join(watermarkFolder, filepath.Base(inputFile))
	}

	// 如果未指定水印文件，则使用嵌入的水印
	var watermarkPath string
	if *watermarkFile == "" {
		// 将嵌入的水印提取到临时文件中
		embeddedWatermark, err := embeddedFiles.ReadFile("watermark.png")
		if err != nil {
			fmt.Printf("读取嵌入的水印时出错: %v\n", err)
			os.Exit(1)
		}

		tempFile, err := os.CreateTemp("", "watermark*.png")
		if err != nil {
			fmt.Printf("创建临时文件时出错: %v\n", err)
			os.Exit(1)
		}
		defer os.Remove(tempFile.Name())

		if _, err := tempFile.Write(embeddedWatermark); err != nil {
			fmt.Printf("写入临时文件时出错: %v\n", err)
			os.Exit(1)
		}
		tempFile.Close()

		watermarkPath = tempFile.Name()
	} else {
		// 使用用户指定的水印文件
		watermarkPath = *watermarkFile
	}

	// 创建水印配置
	wm, err := pdfcpu.ParseImageWatermarkDetails(watermarkPath, "", true, types.POINTS)
	if err != nil {
		fmt.Printf("创建水印配置时出错: %v\n", err)
		os.Exit(1)
	}

	// 设置水印的透明度和缩放比例
	wm.Opacity = *opacity
	wm.Scale = *scale

	// 将水印添加到PDF文件
	err = api.AddWatermarksFile(inputFile, outputFile, nil, wm, nil)
	if err != nil {
		fmt.Printf("添加水印时出错: %v\n", err)
		os.Exit(1)
	}

	// 输出成功消息
	fmt.Printf("水印添加成功！输出文件: %s\n", outputFile)
}
