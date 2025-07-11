
#### etcd 单机版
> 下载: https://github.com/etcd-io/etcd/releases
>
> 目前使用版本: https://github.com/etcd-io/etcd/releases/tag/v3.6.2
```
# 启动
./etcd --listen-client-urls=http://0.0.0.0:2379 --advertise-client-urls=http://0.0.0.0:2379
```