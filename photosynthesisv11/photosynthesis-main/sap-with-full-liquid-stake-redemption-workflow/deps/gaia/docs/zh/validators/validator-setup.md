<!-- markdown-link-check-disable -->

# 在主网上运行一个验证人

::: 提示加入主网所需的信息(`genesis.json`和种子节点)
在[`lauch` repo](https://github.com/cosmos/launch/tree/master/latest)中可以找到
。 :::

在启动你验证人节点前，确定你已经完成了[启动全节点](../join-mainnet.md)教程。

## 什么是验证人?

[验证人](./overview.md)负责通过投票来向区块链提交新区块。如果验证人不可访问或者
对多个相同高度的区块签名，将会遭受到削减处罚。如果变得不可用或者在同一高度上签名
，则会被削减。请阅读有关 Sentry 节点架构的信息，以保护您的节点免受 DDOS 攻击并确
保高可用性。请阅读[哨兵节点网络架构]()来保护你的节点免于 DDOS 攻击并保证高的可访
问性。

::: 警告如果你想要成为 Cosmos Hub 主网的验证人，你应
该[安全研究](./security.md)。 :::

如果你已经[启动了一个全节点](../join-mainnet.md)，可以跳过下一节的内容。

## 创建你的验证人

你的`cosmosvalconspub`可以用于通过抵押 token 来创建一个新的验证人。你可以通过运
行下面的命令来查看你的验证人公钥：

```bash
gaiad tendermint show-validator
```

使用下面的命令创建你的验证人：

::: 注意不要使用多于你所拥有的`uatom`! :::

```bash
gaiad tx staking create-validator \
  --amount=1000000uatom \
  --pubkey=$(gaiad tendermint show-validator) \
  --moniker="choose a moniker" \
  --chain-id=<chain_id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1000000" \
  --gas="auto" \
  --gas-prices="0.0025uatom" \
  --from=<key_name>
```

::: 提示在指定 commission 参数时，`commission-max-change-rate`用于度
量`commission-rate`的百分比点数的变化。比如，1%到 2%增长了 100%，但反映
到`commission-rate`上只有 1 个百分点。 :::

::: 提示如果没有指定，`consensus_pubkey`将默认
为`gaiad tendermint show-validator`命令的输出。`key_name`是将用于对交易进行签名
的私钥的名称。 :::

你可以在第三方区块链浏览器上确定你是否处于验证人行列。

## 以初始验证人的形式加入到 genesis 文件

::: 警告这一节内容只针对想要在 Cosmos Hub 主网启动前就作为初始验证人身份的节点。
如果主网已经启动，请跳过这一节。 :::

如果你想作为初始验证人被写入到 genesis.json 文件，你需要证明你在创世状态中有一些
权益代币，创建一个（或多个）交易以将股权与你的验证人地址联系起来，并将此交易包含
在 genesis 文件中。

你的`cosmosvalconspub`可以用于通过抵押 token 来创建一个新的验证人。运行如下命令
来获取你的验证人节点公钥：

```bash
gaiad tendermint show-validator
```

然后执行`gaiad gentx`命令:

::: 提示 `gentx`是持有 self-delegation 的 JSON 文件。所有的创世交易会
被`创世协调员`收集起来验证并初始化成一个`genesis.json` :::

::: 注意不要使用多于你所拥有的`uatom`! :::

```bash
gaiad gentx \
  --amount <amount_of_delegation_uatom> \
  --commission-rate <commission_rate> \
  --commission-max-rate <commission_max_rate> \
  --commission-max-change-rate <commission_max_change_rate> \
  --pubkey <consensus_pubkey> \
  --name <key_name>
```

::: 提示在指定佣金相关的参数时，`commission-max-change-rate`用于标
识`commission-rate`每日变动的最大百分点数。比如从 1%到 2%按比率是增长了 100%，但
只增加了 1 个百分点。 :::

你可以提交你的`gentx`到[launch repository](https://github.com/cosmos/launch). 这
些`gentx`将会组成最终的 genesis.json.

## 编辑验证人的描述信息

你可以编辑验证人的公开说明。此信息用于标识你的验证人节点，委托人将根据此信息来决
定要委托的验证人节点。确保为下面的每个标识提供输入，否则该字段将默认为空（
`--moniker`默认为机器名称）。

\<key_name>指定你要编辑的验证人。如果你选择不包含此标识，记住必须要含有--from 标
识来指定你要更新的验证人。

`--identity`可用于验证和 Keybase 或 UPort 这样的系统一起验证身份。与 Keybase 一
起使用时，`--identity`应使用由一个[keybase.io](https://keybase.io/)帐户生成的 16
位字符串。它是一种加密安全的方法，可以跨多个在线网络验证您的身份。 Keybase API
允许我们检索你的 Keybase 头像。这是你可以在验证人配置文件中添加徽标的方法。

```bash
gaiad tx staking edit-validator
  --moniker="choose a moniker" \
  --website="https://cosmos.network" \
  --identity=6A0D65E29A4CBC8E \
  --details="To infinity and beyond!" \
  --chain-id=<chain_id> \
  --gas="auto" \
  --gas-prices="0.0025uatom" \
  --from=<key_name> \
  --commission-rate="0.10"
```

**注意** : `commission-rate`的值必须符合如下的不变量检查：

- 必须在 0 和 验证人的`commission-max-rate` 之间
- 不得超过 验证人的`commission-max-change-rate`, 该参数标识**每日**最大的百分点
  变化数。也就是，一个验证人在`commission-max-change-rate`的界限内每日一次可调整
  的最大佣金变化。

## 查看验证人的描述信息

通过该命令查看验证人的描述信息:

```bash
gaiad query staking validator <account_cosmos>
```

## 跟踪验证人的签名信息

你可以通过`signing-info`命令跟踪过往的验证人签名：

```bash
gaiad query slashing signing-info <validator-pubkey>\
  --chain-id=<chain_id>
```

## unjail 验证人

当验证人因停机而"jailed"(入狱)时，你必须用节点操作人帐户提交一笔`Unjail`交易，使
其再次能够获得区块提交的奖励（奖励多少取决于分区的 fee 分配）。

```bash
gaiad tx slashing unjail \
 --from=<key_name> \
 --chain-id=<chain_id>
```

## 确认你的验证人节点正在运行

如果下面的命令返回有内容就证明你的验证人正处于活跃状态:

```bash
gaiad query tendermint-validator-set | grep "$(gaiad tendermint show-validator)"
```

你必须要在[区块浏览器](https://explorecosmos.network/validators)中看见你的验证人
节点信息。你可以在`~/.gaia/config/priv_validator.json`文件中找到`bech32`编码格式
的`address`。

::: warning 注意为了能进入验证人集合，你的权重必须超过第 100 名的验证人。 :::

## 常见问题

### 问题 #1 : 我的验证人的`voting_power: 0`

你的验证人已经是 jailed 状态。如果验证人在最近`10000`个区块中有超过`500`个区块没
有进行投票，或者被发现双签，就会被 jail 掉。

如果被因为掉线而遭到 jail，你可以重获你的投票股权以重回验证人队伍。首先，如
果`gaiad`没有运行，请再次启动：

```bash
gaiad start
```

等待你的全节点追赶上最新的区块高度。然后，运行如下命令。接着，你可
以[unjail 你的验证人]()。

最后，检查你的验证人看看投票股权是否恢复：

```bash
gaiad status
```

你可能会注意到你的投票权比之前要少。这是由于你的下线受到的削减处罚！

### 问题 #2 : 我的`gaiad`由于`too many open files`而崩溃

Linux 可以打开的默认文件数（每个进程）是 1024。已知`gaiad`可以打开超过 1024 个文
件。这会导致进程崩溃。快速修复运行`ulimit -n 4096`（增加允许的打开文件数）来快速
修复，然后使用`gaiad start`重新启动进程。如果你使用`systemd`或其他进程管理器来启
动`gaiad`，则可能需要在该级别进行一些配置。解决此问题的示例`systemd`文件如下：

```toml
# /etc/systemd/system/gaiad.service
[Unit]
Description=Cosmos Gaia Node
After=network.target

[Service]
Type=simple
User=ubuntu
WorkingDirectory=/home/ubuntu
ExecStart=/home/ubuntu/go/bin/gaiad start
Restart=on-failure
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

<!-- markdown-link-check-enable -->
