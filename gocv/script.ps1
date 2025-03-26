# 设置GoCV所需的环境变量

# 获取脚本所在目录的绝对路径
$scriptPath = $PSScriptRoot
$projectRoot = Split-Path -Parent $scriptPath
$opencvPath = Join-Path $projectRoot "opencv"

# 设置环境变量
$env:CGO_CXXFLAGS = "--std=c++11"
$env:CGO_CPPFLAGS = "-I$opencvPath\include"
$env:CGO_LDFLAGS = "-L$opencvPath\lib -lopencv_world4110"
$env:PATH = "$opencvPath\bin;$env:PATH"

# 列出所有以 CGO_ 开头的环境变量
Get-ChildItem env:CGO_*

# 运行主程序
go run -tags customenv .\main.go