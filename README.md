# go_scan

For gofence is terrible  and i want to create something cool my own

需求

能否提供两种方式 一种是基于编译的 一种是不用编译的

https://pkg.go.dev/golang.org/x/tools@v0.1.0/go/callgraph

聚焦于典型问题，而不是覆盖率，及尽量降低误报率

典型问题列表

空指针

整数安全

数组越界

...

1.工具提供默认扫描功能

2.工具提供对外接口可以支持新增规则

3.

A golang tool for statistic check

思路：

以函数为力度，建立调用关系图

2个线程一起跑，一个是添加task，一个是处理task，控制运行时间，是否需要手动GC？

核心问题：

用什么方式处理ssa获得的各种数据

数论？图论？数学建模

有限状态机？https://www.docin.com/p-1144107348.html



Reference：

https://github.com/golang/tools
