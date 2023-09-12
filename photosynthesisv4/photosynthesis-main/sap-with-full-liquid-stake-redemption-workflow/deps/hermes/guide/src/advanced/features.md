# Features

This section includes a summary of the supported and planned features. It also
includes a feature matrix which compares `hermes` to the [cosmos-go-relayer].

> **Cosmos SDK & IBC compatibility:** Hermes supports Cosmos SDK chains
> implementing the [IBC protocol v1][ibcv1-proto] protocol specification. Cosmos
> SDK versions `0.41.3` through `0.45.x` are officially supported. IBC-go
> versions `1.1.*` thorough `3.*` are officially supported. In case Hermes finds
> an incompatible SDK or IBC-go version, it will output a log warning upon
> initialization as part of the `start` command or upon `health-check` command.

***

## Supported Features

*   Basic features
    *   Create and update clients.
    *   Refresh clients to prevent expiration.
    *   Establish connections with new or existing clients.
    *   Establish channels with new or existing connection.
    *   Channel closing handshake.
    *   Relay packets, acknowledgments, timeout and timeout-on-close packets, with
        zero or non-zero delay.
    *   Queries for all objects.
*   Packet relaying over:
    *   multiple paths, for the chains in `config.toml`.
*   Restart support:
    *   Clear packets.
    *   Resume channel handshake if configured to relay `all`.
    *   Resume connection handshake if configured to relay `all`.
*   Client upgrade:
    *   Upgrading clients after a counterparty chain has performed an upgrade for
        IBC breaking changes.
*   Packet delay:
    *   Establish path over non-zero delay connection.
    *   Relay all packets with the specified delay.
*   Interchain Accounts & Interchain Security
*   Monitor and submit misbehaviour for clients
    *   Monitor client updates for misbehaviour (fork and BFT time violation).
    *   Submit misbehaviour evidence to the on-chain IBC client.
        > Misbehaviour submission to full node not yet supported.
*   Individual commands that build and send transactions for:
    *   Creating and updating IBC Tendermint light clients.
    *   Sending connection open handshake messages.
    *   Sending channel open handshake messages.
    *   Sending channel closing handshake messages.
    *   Initiating a cross chain transfer (mainly for testing).
    *   Relaying sent packets, acknowledgments and timeouts.
    *   Automatically generate a configuration file from the
        [chain-registry](https://github.com/cosmos/chain-registry)
    *   Client upgrade.
*   Channel handshake for existing channel that is not in `Open` state.
*   Connection handshake for existing connection that is not in `Open` state.
*   Telemetry support.

## Upcoming / Unsupported Features

Planned features:

*   Interchain Queries
*   Non-SDK support
*   Relay from all IBC events, including governance upgrade proposal
*   Dynamic & automatic configuration management

## Features matrix

***

**Legend**:

| Term              | Description                                                                                      |
| ----------------- | ------------------------------------------------------------------------------------------------ |
| ❌                | feature not supported                                                                            |
| ✅                | feature is supported                                                                             |
| `Chain`           | chain related                                                                                    |
| `Cl`              | client related                                                                                   |
| `Conn`            | connection related                                                                               |
| `Chan`            | channel related                                                                                  |
| `Cfg`             | config related                                                                                   |
| `.._Handshake_..` | can execute all transactions required to finish a handshake from a single command                |
| `.._<msg>_A`      | building and sending `msg` from a command that scans chain state                                 |
| `.._<msg>_P`      | building and sending `msg` from IBC event; doesn't apply to `.._Init` and `FT_Transfer` features |

***

**Feature comparison between Hermes and the Go relayer**

| Features \ Status              | Hermes | Cosmos Go | Feature Details                                                |
| ------------------------------ | :----: | :-------: | :------------------------------------------------------------- |
| Restart                        |   ✅   |    ✅     | replays any IBC events that happened before restart            |
| Multiple\_Paths                 |   ✅   |    ✅     | relays on multiple paths concurrently                          |
|                                |        |           |
| Connection Delay               |   ✅   |    ❌     |
| Cl\_Misbehavior                 |   ✅   |    ❌     | monitors and submits IBC client misbehavior                    |
| Cl\_Refresh                     |   ✅   |    ❌     | periodically refresh an on-chain client to prevent expiration  |
| Packet Delay                   |   ✅   |    ❌     |
|                                |        |           |
| Chan\_Unordered                 |   ✅   |    ✅     |
| Chan\_Ordered                   |   ✅   |    ❓     |
|                                |        |           |
| Cl\_Tendermint\_Create           |   ✅   |    ✅     | tendermint light client creation                               |
| Cl\_Tendermint\_Update           |   ✅   |    ✅     | tendermint light client update                                 |
| Cl\_Tendermint\_Upgrade          |   ✅   |    ✅     | tendermint light client upgrade                                |
|                                |        |           |
| Conn\_Open\_Handshake\_A          |   ✅   |    ✅     |
| Conn\_Open\_Handshake\_P          |   ✅   |    ❌     |
|                                |        |           |
| Chan\_Open\_Handshake\_A          |   ✅   |    ✅     |
| Chan\_Open\_Handshake\_P          |   ✅   |    ❌     |
| Chan\_Open\_Handshake\_Optimistic |   ❌   |    ❌     | open a channel on a non-Open connection                        |
|                                |        |           |
| Chan\_Close\_Handshake\_P         |   ✅   |    ✅     |
| Chan\_Close\_Handshake\_A         |   ✅   |    ❌     |
|                                |        |           |
| FT\_Transfer                    |   ✅   |    ✅     | can submit an ICS-20 fungible token transfer message           |
| ICA\_Relay                      |   ✅   |    ❌     | can relay ICS-27 Interchain account packets                    |
| Packet\_Recv\_A                  |   ✅   |    ✅     |
| Packet\_Recv\_P                  |   ✅   |    ✅     |
| Packet\_Timeout\_A               |   ✅   |    ✅     |
| Packet\_Timeout\_P               |   ✅   |    ✅     |
| Packet\_TimeoutClose\_A          |   ✅   |    ❓     |
| Packet\_TimeoutClose\_P          |   ✅   |    ❓     |
| Packet\_Optimistic              |   ❌   |    ❓     | relay packets over non-Open channels                           |
|                                |        |           |
| Cl\_Non\_Tendermint              |   ❌   |    ❌     | supports non tendermint IBC light clients                      |
| Chain\_Non\_Cosmos               |   ❌   |    ❌     | supports non cosmos-SDK chains                                 |
|                                |        |           |
| Cfg\_Static                     |   ✅   |    ✅     | provides means for configuration prior to being started        |
| Cfg\_Dynamic                    |   ❌   |    ❌     | provides means for configuration and monitoring during runtime |
| Cfg\_Download\_Config            |   ✅   |    ✅     | provides means for downloading recommended configuration       |
| Cfg\_Edit\_Config                |   ❌   |    ✅     | provides means for editing the configuration from the CLI      |
| Cfg\_Validation                 |   ✅   |    ✅     | provides means to validate the current configuration           |
| Telemetry                      |   ✅   |    ❌     | telemetry server to collect metrics                            |
| REST API                       |   ✅   |    ❌     | REST API to interact with the relayer                          |

[cosmos-go-relayer]: https://github.com/cosmos/relayer

[ibcv1-proto]: https://github.com/cosmos/ibc
