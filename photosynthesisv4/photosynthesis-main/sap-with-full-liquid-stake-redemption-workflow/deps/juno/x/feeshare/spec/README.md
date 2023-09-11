<!--
order: 0
title: "FeeShare Overview"
parent:
  title: "feeshare"
-->

# `feeshare`

## Abstract

This document specifies the internal `x/feeshare` module of Juno Network.

The `x/feeshare` module enables the Juno to support splitting transaction fees
between the community and smart contract deployer. This aims to increase the
adoption of Juno by offering a new way of income for CosmWasm smart contract
developers. Developers can register their smart contracts and every time someone
interacts with a registered smart contract, the contract deployer or their
assigned withdrawal account receives a part of the transaction fees.

## Contents

1. **[Concepts](01\_concepts.md)**
2. **[State](02\_state.md)**
3. **[State Transitions](03\_state_transitions.md)**
4. **[Transactions](04\_transactions.md)**
5. **[Hooks](05\_hooks.md)**
6. **[Events](06\_events.md)**
7. **[Parameters](07\_parameters.md)**
8. **[Clients](08\_clients.md)**
