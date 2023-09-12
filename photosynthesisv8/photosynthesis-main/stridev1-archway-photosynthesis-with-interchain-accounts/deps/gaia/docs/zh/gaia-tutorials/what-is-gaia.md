<!-- markdown-link-check-disable -->

# Gaia 是什么

`gaia`是作为 Cosmos SDK 应用程序的 Cosmos Hub 的名称。它有两个主要的入口：

*   `gaiad` : Gaia 的服务进程，运行着`gaia`程序的全节点。
*   `gaiad` : Gaia 的命令行界面，用于同一个 Gaia 的全节点交互。

`gaia`基于 Cosmos SDK 构建，使用了如下模块:

*   `x/auth` : 账户和签名
*   `x/bank` : token 转账
*   `x/staking` : 抵押逻辑
*   `x/mint` : 增发通胀逻辑
*   `x/distribution` : 费用分配逻辑
*   `x/slashing` : 处罚逻辑
*   `x/gov` : 治理逻辑
*   `ibc-go/modules` : 跨链交易
*   `x/params` : 处理应用级别的参数

> 关于 Cosmos Hub : Cosmos Hub 是第一个在 Cosmos Network 中上线的枢纽。枢纽的作
> 用是用以跨链转账。如果区块链通过 IBC 协议连接到枢纽，它会自动获得对其它连接至
> 枢纽的区块链的访问能力。Cosmos Hub 是一个公开的 PoS 区块链。它的权益代币称为
> Atom。

接着，学习如何[安装 Gaia](./installation.md)

<!-- markdown-link-check-enable -->
