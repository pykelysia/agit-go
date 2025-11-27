<div align="center">
    <h1>A-GIT</h1>
    <img alt="GitHub License" src="https://img.shields.io/github/license/pykelysia/agit-go">
    <img alt="Backend" src="https://img.shields.io/badge/Go-1.24.5-blue?logo=go">
    <img alt="Model" src="https://img.shields.io/badge/command-brightgreen">
    <br/>
    <p>auto-git rebuild with golang</p>
    <p>原项目来源：
        <a href="https://github.com/pykelysia/auto-git">pykelysia/auto-git </a>
    </p>
    
</div>

## Quick Start

### 从源码开始

```bash
git clone github.com.pykelysia/agit-go
cd agit-go
go build
```

即可生成可执行文件 `agit.exe`

### 直接下载编译好的可执行文件

前往 `Releases` 页面下载对应平台的可执行文件即可。

### 使用方法

```bash
# 在任意文件夹下运行
# 默认模式下会自动为当前目录下的本地仓库提交 push
# PowerShell 需要在命令前添加 "./"
agit.exe
```

#### 参数说明：

- `--path`, `-p`: 指定需要操作的 git 仓库路径，默认为当前路径。
- `--mode`, `-m`: 指定模式，目前可选有 `normal`, `fast`两种模式。
- `--default`, `-d`: 使用默认模式，在当前仓库进行提交 push。
