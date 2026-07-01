# BDPulse Commons

BDPulse Commons 是 [BDPulse (原 Drone)](https://github.com/open-beagle) CI/CD 引擎的基础工具与公共组件包合集。

该仓库将原本离散的各个基础微组件合并收拢在一起，降低了基础组件库的管理与维护成本。

## 包含组件

本仓库包含以下常用的核心模块：

### 1. `envsubst`

环境变量替换引擎。负责在 BDPulse 执行时，将字符串（通常是 YAML 配置文件或 Shell 脚本内容）中的环境变量占位符（如 `${VAR}` 或 `$VAR`）替换为实际的环境变量值，并且支持 Bash 风格的高级字符串替换功能（如默认值、字符串截取等）。

### 2. `funcmap`

用于 Go Template 执行的常用函数映射表（FuncMap）。封装了一系列常用的字符串操作、加解密、类型转换等函数，便于在各种模板引擎解析（如 `.beagle.yml` 解析）时调用。

### 3. `signal`

OS 信号捕获与处理库。提供了优雅的信号拦截与上下文 (Context) 取消机制，确保在引擎、Agent 及各个服务被杀死（如 `SIGINT`, `SIGTERM`）时，能够妥善中断并清理执行中的流水线资源。

## 安装与引入

```bash
go get github.com/open-beagle/bdpulse-commons
```

在代码中按需引入对应包：

```go
import (
	"github.com/open-beagle/bdpulse-commons/envsubst"
	"github.com/open-beagle/bdpulse-commons/funcmap"
	"github.com/open-beagle/bdpulse-commons/signal"
)
```

## 测试

执行全量测试：

```bash
go test ./...
```
