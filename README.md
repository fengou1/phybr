编译这个项目，生成的可执行文件 phybr 就是 TiDB 物理备份 DEMO 中的 adviser 程序。

执行它，它会监听 127.0.0.1:3379，同时等待 3 个 TiKV 实例上的 recover 程序上报自己的信息。这里两个参数都还没有做成选项。

执行起来之后，就可以到每个 TiKV 上执行 tikv-phybr 程序了，它在 TiKV 仓库中。

tikv-phybr 就是 DEMO 中的 recover 程序，举例说明它接受的参数：tikv-phybr 127.0.0.1:3379 kv1.toml。

第一个参数是上文 adviser 程序监听的地址，第二个参数是一个配置文件的路径。kv1.toml 的内容如下：

```toml
# 必须要有的内容
[server]
addr = "0.0.0.0:40160"
advertise-addr = "127.0.0.1:40160"
status-addr = "0.0.0.0:40180"
advertise-status-addr = "127.0.0.1:40180"

[pd]
endpoints = [ "127.0.0.1:2379" ]

[storage]
data-dir = "/ebs/data/tikv-20177"
reserve-space = "0GB"

# 需要与源集群保持一致的内容
[raft-engine]
enable = false
[server.labels]
host = "logic-host-x1"

# 其他内容
[log]
level = "info"
[log.file]
filename = "kv1.log"
```

配置文件中必须提供全部需要监听的地址以及 PD 的地址，否则默认地址容易与其他实例冲突。另一些配置项需要与源集群保持一致，它们主要涉及数据文件的存储格式。最后是其他一些非必须的配置项，仅仅方便 tikv-phybr 运行中查看日志而已。

运行 tikv-phybr 后，它会上报本地所有 Region 的元信息给 adviser，后者下发命令到 tikv-phybr 上执行。执行完成后，tikv-phybr 是不会退出的，这是因为其他 TiKV 实例上的 tikv-phybr 程序可能需要自身的存活才能推进 Raft 日志的提交。调用者需要综合全部 tikv-phybr 实例的运行状况之后，手动停止它们。

目前的实现有一些需要改进的地方：
1. 可以再在 tikv-phybr 和 adviser 之间增加一个 RPC，用于 tikv-phybr 给 adviser 上报本地的 recover 完成情况。当 adviser 集齐全部上报之后，再命令所有 tikv-phybr 退出，这样就不需要调用者手动结束它们了。
