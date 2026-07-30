# 简介
`sandbox connect`（别名 `cn`）连接到一个已有的沙箱终端。当终端会话结束时，沙箱继续运行。

# 格式
```
qshell sandbox connect <sandboxID> [--retry-max <N>]
qshell sbx cn <sandboxID> [--retry-max <N>]
```

# 帮助文档
```
$ qshell sandbox connect -h
$ qshell sandbox connect --doc
```

# 鉴权
需要配置 `QINIU_API_KEY` 或 `E2B_API_KEY` 环境变量。

# 参数
- `sandboxID`：沙箱 ID（必填）
- `--retry-max`：连接请求的最大自动重试次数；`0` 禁用重试。未传入时优先读取 `SANDBOX_RETRY_MAX`，未设置则默认重试 5 次

# 示例
```
$ qshell sandbox connect sb-xxxxxxxxxxxx
$ qshell sbx cn sb-xxxxxxxxxxxx
```

禁用连接请求自动重试：
```
$ qshell sandbox connect sb-xxxxxxxxxxxx --retry-max 0
```
