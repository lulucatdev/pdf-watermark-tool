# PDF Watermark Tool / PDF 水印工具

This tool allows you to add watermarks to PDF files using a command-line interface.

这个工具允许你使用命令行界面为 PDF 文件添加水印。

## Features / 功能

- Add image watermarks to PDF files / 为 PDF 文件添加图片水印
- Adjust watermark opacity and scale / 调整水印的透明度和缩放比例
- Use custom watermark images or a default embedded watermark / 使用自定义水印图片或默认嵌入的水印
- Specify output file or use automatic output folder / 指定输出文件或使用自动输出文件夹

## Prerequisites / 前提条件

- Go 1.16 or later / Go 1.16 或更高版本

## Installation / 安装

### Local Installation / 本地安装

1. Clone this repository / 克隆此仓库:
   ```
   git clone https://github.com/yourusername/pdf-watermark-tool.git
   ```

2. Change to the project directory / 切换到项目目录:
   ```
   cd pdf-watermark-tool
   ```

3. Build the tool / 构建工具:
   ```
   go build -o watermark
   ```

### System-wide Installation / 系统级安装

To make the `watermark` command available system-wide, follow these steps:

要使 `watermark` 命令在系统范围内可用，请按照以下步骤操作：

#### For Unix-like systems (Linux, macOS) / 对于类 Unix 系统（Linux，macOS）：

1. Build the tool / 构建工具:
   ```
   go build -o watermark
   ```

2. Move the binary to a directory in your PATH / 将二进制文件移动到 PATH 中的目录:
   ```
   sudo mv watermark /usr/local/bin/
   ```

3. Ensure the binary is executable / 确保二进制文件可执行:
   ```
   sudo chmod +x /usr/local/bin/watermark
   ```

Now you can run the `watermark` command from any directory.

现在你可以从任何目录运行 `watermark` 命令。

#### For Windows:

1. Build the tool / 构建工具:
   ```
   go build -o watermark.exe
   ```

2. Create a new folder for the tool, e.g., `C:\Program Files\PDFWatermark`

   创建一个新文件夹用于存放工具，例如 `C:\Program Files\PDFWatermark`

3. Move the `watermark.exe` file to this new folder

   将 `watermark.exe` 文件移动到这个新文件夹

4. Add the folder to your system's PATH:
   - Right-click on 'This PC' or 'My Computer' and select 'Properties'
   - Click on 'Advanced system settings'
   - Click on 'Environment Variables'
   - Under 'System variables', find and select 'Path', then click 'Edit'
   - Click 'New' and add the path to your folder (e.g., `C:\Program Files\PDFWatermark`)
   - Click 'OK' to close all dialogs

   将文件夹添加到系统的 PATH 中：
   - 右键点击"此电脑"或"我的电脑"，选择"属性"
   - 点击"高级系统设置"
   - 点击"环境变量"
   - 在"系统变量"下，找到并选择"Path"，然后点击"编辑"
   - 点击"新建"并添加你的文件夹路径（例如 `C:\Program Files\PDFWatermark`）
   - 点击"确定"关闭所有对话框

5. Open a new command prompt to apply the changes

   打开一个新的命令提示符以应用更改

Now you can run the `watermark` command from any directory in Command Prompt.

现在你可以在命令提示符中从任何目录运行 `watermark` 命令。

## Usage / 使用方法

Basic usage / 基本用法:
```
watermark [flags] <input.pdf> [output.pdf]
```

If no output file is specified, the watermarked PDF will be saved in a `watermarked_files` folder with the same filename as the input.

如果未指定输出文件，带水印的 PDF 将保存在 `watermarked_files` 文件夹中，文件名与输入文件相同。

### Flags / 标志

- `-opacity`: Set the watermark opacity (0.0 - 1.0). Default is 0.15. / 设置水印透明度（0.0 - 1.0）。默认值为 0.15。
- `-scale`: Set the watermark scale. Default is 1.0. / 设置水印缩放比例。默认值为 1.0。
- `-watermark`: Specify a custom watermark PNG file path. If not provided, a default embedded watermark will be used. / 指定自定义水印 PNG 文件路径。如果未提供，将使用默认嵌入的水印。

### Examples / 示例

1. Add a watermark using default settings / 使用默认设置添加水印:
   ```
   watermark input.pdf
   ```

2. Specify an output file / 指定输出文件:
   ```
   watermark input.pdf output.pdf
   ```

3. Adjust opacity and scale / 调整透明度和缩放比例:
   ```
   watermark -opacity 0.3 -scale 1.5 input.pdf
   ```

4. Use a custom watermark image / 使用自定义水印图片:
   ```
   watermark -watermark custom_watermark.png input.pdf
   ```

## Notes / 注意事项

- The tool currently supports PNG images for watermarks. / 该工具目前支持 PNG 格式的水印图片。
- Make sure you have the necessary permissions to read the input PDF and write the output PDF. / 确保你有必要的权限来读取输入 PDF 和写入输出 PDF。

## Troubleshooting / 故障排除

If you encounter any issues, please check the error messages displayed in the console. Common issues include:

如果遇到任何问题，请检查控制台显示的错误消息。常见问题包括：

- Invalid input file path / 无效的输入文件路径
- Insufficient permissions to read/write files / 读取/写入文件的权限不足
- Unsupported watermark image format / 不支持的水印图片格式

## Contributing / 贡献

Contributions are welcome! Please feel free to submit a Pull Request.

欢迎贡献！请随时提交 Pull Request。

## License / 许可证

This project is licensed under the MIT License - see the LICENSE file for details.

本项目采用 MIT 许可证 - 有关详细信息，请参阅 LICENSE 文件。