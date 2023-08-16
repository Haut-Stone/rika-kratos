# 系统监控的使用
## pprof
### 监控 off-cpu 耗时
go tool pprof --http=:6061 http://localhost:8000/debug/fgprof\?seconds\=10
### 监控 cpu 耗时
go tool pprof http://localhost:8000/debug/pprof/profile\?seconds\=60
### 观察 trace
- 得到文件：curl -so trace_30s2 http://localhost:8000/debug/pprof/trace\?seconds\=30
- 打开 trace：go tool trace trace_30s2