<!-- markdown-link-check-disable -->

## 加入主网

::: 提示请查看[launch repo](https://github.com/cosmos/launch)获取主网信息，包含
了所使用的 Cosmos-SDK 的正确版本和 genesis 文件。 :::

::: 警告 **在更进一步之前你需要[安装 gaia](./installation.md)** :::

## 创建一个新节点

这些指令适用于从头开始设置一个全节点。

首先，初始化节点并创建必要的配置文件：

```bash
gaiad init <your_custom_moniker>
```

::: 注意 moniker 只能包含 ASCII 字符。使用 Unicode 字符会使得你的节点不可访问
:::

你可以稍后在`~/.gaia/config/config.toml`文件中编辑`moniker`:

```toml
# A custom human readable name for this node
moniker = "<your_custom_moniker>"
```

你可以编辑`~/.gaia/config/config.toml`文件来开启垃圾交易过滤机制以拒绝收到的手续
费过低的交易：

    # This is a TOML config file.
    # For more information, see https://github.com/toml-lang/toml

    ##### main base config options #####

    # The minimum gas prices a validator is willing to accept for processing a
    # transaction. A transaction's fees must meet the minimum of any denomination
    # specified in this config (e.g. 10uatom).

    minimum-gas-prices = ""

你的全节点已经初始化成功！

## Genesis & Seeds

### 复制 genesis 文件

将主网的`genesis.json`文件放置在`gaiad`的配置文件夹中

```bash
mkdir -p $HOME/.gaia/config
curl https://raw.githubusercontent.com/cosmos/launch/master/genesis.json > $HOME/.gaia/config/genesis.json
```

注意我们使用了[launch repo](https://github.com/cosmos/launch)中的`latest`文件夹
，该文件夹包含了最新版本主网的详细信息。

::: 提示如果你想加入的是公共测试网，点击[这里](./join-testnet.md) :::

运行命令验证配置的正确性:

```bash
gaiad start
```

### 添加种子节点

你的节点需要知道如何寻找伙伴节点。你需要添加有用的种子节点
到`$HOME/.gaia/config/config.toml`文件中
。[`launch`](https://github.com/cosmos/launch) repo 包含了一些种子节点的链接。

如果这些种子节点不再运行，你可以在 Cosmos Hub 浏览器(可以
在[launch page](https://cosmos.network/launch)中找到)发现种子节点和持久节点。

你还可以
到[验证人 Riot 聊天室](https://riot.im/app/#/room/#cosmos-validators:matrix.org)里
询问可用节点。

你可以阅
读[这里](https://github.com/tendermint/tendermint/blob/develop/docs/tendermint-core/using-tendermint.md#peers)了
解更多伙伴节点和种子节点的信息。

::: 警告在 Cosmos Hub 主网中，可接受的币种是`uatom`,`1atom = 1.000.000uatom` :::

Cosmos Hub 网络中的交易需要支付一笔交易手续费以得到处理。手续费支付执行交易所消
耗的 gas。计算公式如下：

    fees = gas * gasPrices

`gas`由交易本身决定。不同的交易需要不同数量的`gas`。一笔交易的`gas`数量在它被执
行时计算，但有一种方式可以提前估算，那就是把标识`gas` 的值设置为`auto`。当然，这
只是给出一个预估值。如果你想要确保为交易提供足够的 gas，你可以使
用`--gas-adjustment`标识来调整预估值(默认是`1.0`)。

`gasPrice`是每个单位`gas`的单价。每个验证人节点可以设置`min-gas-price`，只会把那
些`gasPrice`高于自己设置的`min-gas-price`的交易打包。

交易的`fees`是`gas`与`gasPrice`的结果。作为一个用户，你必须输入三者中的两者。更
高的`gasPrice`/`fees`，将提高你的交易被打包的机会。

::: 提示主网中推荐的`gas-prices`是`0.0025uatom` :::

## 设置`minimum-gas-prices`

你的全节点可以在交易池中放入未确认的交易。为了保护其免受 Spam 攻击，最好设置一
个`minimum-gas-prices`来过滤交易以决定是否要放入交易池。这个参数可以
在`~/.gaia/config/gaiad.toml`文件中配置。

推荐的初始`minimum-gas-prices`是`0.0025uatom`，如果你愿意可以稍后再修改它。

## 运行一个全节点

通过这条命令开始运行全节点：

```bash
gaiad start
```

检查一切是否平稳运行中:

```bash
gaiad status
```

使用[Cosmos Explorer](https://cosmos.network/launch)查看网络状态。

## 导出状态

Gaia 能够将整个应用程序的状态转存到一个 JSON 文件中，该文件可以用于分析还可以用
作一个新网络的 genesis 文件。

导出状态:

```bash
gaiad export > [filename].json
```

你还可以导出指定高度的状态(处理完指定高度后的状态):

```bash
gaiad export --height [height] > [filename].json
```

如果你计划使用导出的状态文件启动一个新网络，导出时要加上`--for-zero-height`标识:

```bash
gaiad export --height [height] --for-zero-height > [filename].json
```

## 升级成为验证人节点

你现在有了一个运行状态的全节点。接下来，你可以升级你的全节点，成为一个 Cosmos 验
证人。排名前 100 的验证人节点可以向 Cosmos Hub 提议新的区块。请查
看[创建验证人节点](./validators/validator-setup.md)。

<!-- markdown-link-check-enable -->
