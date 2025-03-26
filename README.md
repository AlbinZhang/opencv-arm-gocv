# OpenCV ARM Windows 环境配置

这是一个为 ARM 架构 Windows 系统配置的 OpenCV 环境，专门用于 Go 语言开发（使用 GoCV）。

## 环境要求

- Windows ARM 系统
- Go 语言环境
- LLVM-MinGW 工具链
- OpenCV 4.11.0 (ARM64 版本)

## 目录结构

```
.
├── opencv/           # OpenCV 库文件
│   ├── include/     # 头文件
│   ├── lib/        # 库文件
│   └── bin/        # 动态链接库
└── gocv/           # Go 项目目录
    ├── main.go     # 主程序
    └── script.ps1  # 环境配置脚本
```

## 安装步骤

1. 确保已安装 LLVM-MinGW 工具链
2. 克隆或下载本项目
3. 运行环境配置脚本：
   ```powershell
   cd gocv
   .\script.ps1
   ```

## 环境变量说明

脚本会自动设置以下环境变量：

- `CGO_CXXFLAGS`: 设置 C++ 标准为 C++11
- `CGO_CPPFLAGS`: 设置 OpenCV 头文件路径
- `CGO_LDFLAGS`: 设置 OpenCV 库文件路径
- `PATH`: 添加 OpenCV 动态库路径

## 使用说明

1. 确保在 `gocv` 目录下运行程序
2. 使用以下命令运行程序：
   ```powershell
   go run -tags customenv .\main.go
   ```

## 注意事项

- 本项目使用的是 OpenCV 4.11.0 版本
- 需要确保 LLVM-MinGW 工具链正确安装并配置
- 如果遇到 DLL 加载问题，请检查 PATH 环境变量是否正确设置

## 常见问题

1. 如果遇到 DLL 加载失败，请检查：
   - OpenCV 的 bin 目录是否已添加到 PATH
   - 是否使用了正确的 ARM 版本 OpenCV

2. 如果编译失败，请检查：
   - LLVM-MinGW 是否正确安装
   - 环境变量是否正确设置

## 许可证

本项目遵循 MIT 许可证
