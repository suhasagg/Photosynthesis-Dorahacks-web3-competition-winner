<!-- markdown-link-check-disable -->

# Ledger Nano 支持

## 关于 HD 钱包

HD 钱包（分层确定性钱包）, 最初是在比特币
的[BIP32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)提案中
提出, 是一种特殊的钱包类型，可以让用户从单个种子派生任意数量的账户。为了更好的理
解, 让我们定义以下术语：

*   **钱包**: 从一个给定的 seed 中获得的一组账户。
*   **账户**: 一组公钥/私钥对。
*   **私钥**: 私钥是用于签署消息的私密信息。在区块链领域, 一个私钥就是一个账户的所
    有者。永远不要想他人透露用户的私钥。
*   **公钥**: 公钥是通过对私钥上使用不可逆的加密函数而获得的一条信息。从公钥中可以
    导出地址。但无法从公钥中逆向获得私钥。
*   **地址**: 地址是一个公开的用于标识账户的，带着可读前缀的字符串。通过对公钥使用
    加密函数获得。
*   **数字签名**: 数字签名是一段加密信息，证明了指定私钥的所有者在不泄露其私钥的情
    况下，批准了指定消息。
*   **种子**: 同助记词。
*   **助记词**: 助记符是一串单词，用作种子来派生私钥。助记符是每个钱包的核心。永远
    不要丢失你的助记词。把它写在一张纸上，然后把它存放在安全的地方。如果你失去了它
    ，那就没有办法去重新获取它了。如果某人获得了助记词访问权限，他们将获得所有相关
    帐户的访问权限。

HD 钱包的核心是种子。用这个种子用户可以确定地生成子帐户。要从种子生成子帐户，使
用了单向的数学转换。要决定生成哪个帐户，用户指定`path`，通常
是`integer`（`0`，`1`，`2`，...）。

例如，通过将`path`指定为`0`，钱包将从种子生成`0号私钥`。然后，可以从`0号私钥`生
成“`号公钥`。最后，可以从`0号公钥`生成`0号地址`。所有这些步骤都是单向的，这意味
着`公钥`无法从`地址`中获得，`私钥`无法从`公钥`中获得，......

         Account 0                         Account 1                         Account 2

    +------------------+              +------------------+               +------------------+
    |                  |              |                  |               |                  |
    |    Address 0     |              |    Address 1     |               |    Address 2     |
    |        ^         |              |        ^         |               |        ^         |
    |        |         |              |        |         |               |        |         |
    |        |         |              |        |         |               |        |         |
    |        |         |              |        |         |               |        |         |
    |        +         |              |        +         |               |        +         |
    |  Public key 0    |              |  Public key 1    |               |  Public key 2    |
    |        ^         |              |        ^         |               |        ^         |
    |        |         |              |        |         |               |        |         |
    |        |         |              |        |         |               |        |         |
    |        |         |              |        |         |               |        |         |
    |        +         |              |        +         |               |        +         |
    |  Private key 0   |              |  Private key 1   |               |  Private key 2   |
    |        ^         |              |        ^         |               |        ^         |
    +------------------+              +------------------+               +------------------+
             |                                 |                                  |
             |                                 |                                  |
             |                                 |                                  |
             +--------------------------------------------------------------------+
                                               |
                                               |
                                     +---------+---------+
                                     |                   |
                                     |  Mnemonic (Seed)  |
                                     |                   |
                                     +-------------------+

从种子中推导出帐户的过程是确定性的。这意味着给定相同的路径，派生出私钥将始终相同
。

存储在帐户中的资金由私钥控制。此私钥对助记词使用单向函数生成的。如果丢失了私钥，
则可以使用助记词恢复它。但是，如果丢失了助记词，则将无法访问所有派生的私钥。同样
，如果有人获得了你的助记词访问权限，他们就可以访问所有相关帐户。

## Ledger 支持 HD 钱包

在 Ledger 钱包的内核，有一个用于生成私钥的助记词。初始化 Ledger 时，会生成助记词
。

::: 危险 **不要遗失或与任何人分享你的 12/24 个单词。为防止盗窃或资金损失，最好确
保备份多个助记词拷贝，并将其存放在安全可靠的地方，并且只有你知道如何访问。如果有
人能够访问你的助记词，他们将能够访问你的私钥并控制与其关联的帐户** :::

助记词与 Cosmos 帐户体系兼容。用于在 Cosmos Hub 网络上生成地址和交易的工具称
为`gaiad`，它支持从 Ledger 种子派生帐户私钥。请注意，Ledger 钱包充当种子和私钥的
沙盒，并且对交易进行签名的过程完全在内部进行。没有任何私人信息会离开 Ledger 钱包
。

要将`gaiad`与 Ledger 钱包一起使用，你需要具备以下条件：

*   [在 Ledger Nano 中安装`COSMOS`应用并生成账户](./delegator-guide-cli.md#using-a-ledger-device)
*   [有一个你打算连接的可访问的并处于运行状态的`gaiad`实例](./delegator-guide-cli.md#accessing-the-cosmos-hub-network)
*   [一个同你所选的`gaiad`实例相连接的`gaiad`实例](./delegator-guide-cli.md#setting-up-gaiad)

现在，你都准备好
去[发送交易到网络](./delegator-guide-cli.md#sending-transactions).

<!-- markdown-link-check-enable -->
