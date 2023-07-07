# Photosynthesis-Archway  <-> Stride integrations

![Photosynthesis](https://i.imgur.com/Tn1GUZnb.jpg)


1)START RELAYER

2)CREATING CONNECTIONS WITH THE GO RELAYER 

3)Create connections and channels


a)Get channel ID created on the photosynthesis-archway

b)Start Go Relayer 

c)IBC Transfer from Photosynthesis-archway to stride (from relayer account)

d)Confirm funds were recieved on stride and get IBC denom

e)Register host zone

f)Add validator

g)Confirm ICA channels were registered


FLOW

Go Through Flow


a)Liquid stake (then wait and LS again)

b)Confirm stTokens, StakedBal, and Redemption Rate

c)Redeem

d)Confirm stTokens and StakedBal

e)Add another validator

f)Liquid stake and confirm the stake was split 50/50 between the validators

g)Change validator weights

h)LS and confirm delegation aligned with new weights

i)Update delegations (just submit this query and confirm the ICQ callback displays in the stride logs)


# Must be submitted in ICQ window

j)Restore closed channel



# Photosynthesis-Archway IBC integrations. 

1)evmos

2)gaia

3)juno

4)osmosis

5)stargaze

```
global:
  api-listen-addr: :5183
  timeout: 10s
  memo: ""
  light-cache-size: 20
chains:
  photo:
    type: cosmos
    value:
      key: rly0
      chain-id: localnet
      rpc-addr: http://photo1:26657
      account-prefix: archway
      keyring-backend: test
      gas-adjustment: 1.3
      gas-prices: 100uatom
      coin-type: 118
      debug: false
      timeout: 20s
      output-format: json
      sign-mode: direct  
  stride:
    type: cosmos
    value:
      key: rly1
      chain-id: STRIDE
      rpc-addr: http://stride1:26657
      account-prefix: stride
      keyring-backend: test
      gas-adjustment: 1.3
      gas-prices: 0.01ustrd
      coin-type: 118
      debug: false
      timeout: 20s
      output-format: json
      sign-mode: direct
  gaia:
    type: cosmos
    value:
      key: rly2
      chain-id: GAIA
      rpc-addr: http://gaia1:26657
      account-prefix: cosmos
      keyring-backend: test
      gas-adjustment: 1.2
      gas-prices: 0.01uatom
      coin-type: 118
      debug: false
      timeout: 20s
      output-format: json
      sign-mode: direct
  juno:
    type: cosmos
    value:
      key: rly3
      chain-id: JUNO
      rpc-addr: http://juno1:26657
      account-prefix: juno
      keyring-backend: test
      gas-adjustment: 1.2
      gas-prices: 0.01ujuno
      coin-type: 118
      debug: false
      timeout: 20s
      output-format: json
      sign-mode: direct
  osmo:
    type: cosmos
    value:
      key: rly4
      chain-id: OSMO
      rpc-addr: http://osmo1:26657
      account-prefix: osmo
      keyring-backend: test
      gas-adjustment: 1.2
      gas-prices: 0.01uosmo
      coin-type: 118
      debug: false
      timeout: 20s
      output-format: json
      sign-mode: direct
  stars:
    type: cosmos
    value:
      key: rly5
      chain-id: STARS
      rpc-addr: http://stars1:26657
      account-prefix: stars
      keyring-backend: test
      gas-adjustment: 1.2
      gas-prices: 0.01ustars
      coin-type: 118
      debug: false
      timeout: 20s
      output-format: json
      sign-mode: direct
  host:
    type: cosmos
    value:
      key: rly6
      chain-id: HOST
      rpc-addr: http://host1:26657
      account-prefix: stride
      keyring-backend: test
      gas-adjustment: 1.3
      gas-prices: 0.01uwalk
      coin-type: 118
      debug: false
      timeout: 20s
      output-format: json
      sign-mode: direct
  evmos:
    type: cosmos
    value:
      key: rly7
      chain-id: evmos_9001-2
      rpc-addr: http://evmos1:26657
      account-prefix: evmos
      keyring-backend: test
      gas-adjustment: 1.2
      gas-prices: 0.01aevmos
      coin-type: 60
      debug: false
      timeout: 20s
      output-format: json
      sign-mode: direct
      extra-codecs:
        - ethermint
```

# Relayer registry 

Cosmos relayer
1)https://github.com/cosmos/relayer.git

Hermes relayer 
2)https://github.com/informalsystems/ibc-rs 


# Record Keeping/Queue Creation for Liquid Staking Workflow in Photosynthesis-Archway
# Interchain Accounts fully Integrated in Photosynthesis-Archway

```
Liquid staking workflow in Photosynthesis-Archway
 STRIDE @ 343 | 1 VALS
PHOTO   @ 337 | 1 VALS
GAIA   @ 337 | 1 VALS

LIST-HOST-ZONES STRIDE
host_zone:
- address: stride1755g4dkhpw73gz9h9nwhlcefc6sdf8kcmvcwrk4rxfrz8xpxxjms7savm8
  bech32prefix: cosmos
  blacklisted_validators: []
  chain_id: GAIA
  connection_id: connection-1
  delegation_account:
    address: cosmos1rkth5ywkueewvs29xkalckhhl3w5esg9jwn03ywp30k3dh2ys6aqffvtav
    target: DELEGATION
  fee_account:
    address: cosmos1a8ue6w4x9yv2rxq2m3l7urua39ffzed457x3yeetk94yl59w322qpavyph
    target: FEE
  halted: false
  host_denom: uatom
  ibc_denom: ibc/C4CFF46FD6DE35CA4CF4CE031E643C8FDC9BA4B99AE598E9B0ED98FE3A2319F9
  last_redemption_rate: "1.000000000000000000"
  max_redemption_rate: "1.500000000000000000"
  min_redemption_rate: "0.900000000000000000"
  redemption_account:
    address: cosmos19mtvgh3mezrqttxnddtqx3hnfhnkt7xqjylladj7qmsl4edn23tqcv5dyz
    target: REDEMPTION
  redemption_rate: "1.000000000000000000"
  staked_bal: "0"
  transfer_channel_id: channel-1
  unbonding_frequency: "1"
  validators:
  - address: cosmosvaloper1uk4ze0x4nvh4fk0xm4jdud58eqn4yxhrdt795pcosmosvaloper1uk4ze0x4nvh4fk0xm4jdud58eqn4yxhrdt795p
    delegation_amt: "0"
    internal_exchange_rate: null
    name: gval1
    weight: "5"
  withdrawal_account:
    address: cosmos1svjuhjlw8mea66tj2phnmtg050dljpx4ku3qzcgehdpxslxnen6sc68j0u
    target: WITHDRAWAL
- address: stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag
  bech32prefix: archway
  blacklisted_validators: []
  chain_id: localnet
  connection_id: connection-0
  delegation_account:
    address: archway1c9zxssf4u9rcmlx9pdsfsqzgvhdgjykjckn9apevr9ht46kputmq9v738w
    target: DELEGATION
  fee_account:
    address: archway1ds0m5f2mp4j9jr5hpzayaefjguxc8l92zvfurd34gdpz3zpu4fuqmzvznt
    target: FEE
  halted: false
  host_denom: uarch
  ibc_denom: ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2
  last_redemption_rate: "1.000000000000000000"
  max_redemption_rate: "1.500000000000000000"
  min_redemption_rate: "0.900000000000000000"
  redemption_account:
    address: archway1q8v4vlmwlextv4mm5eyh203hv98dey2fzwuc3y6d8evhhuhp3zvsdmsc7x
    target: REDEMPTION
  redemption_rate: "1.000000000000000000"
  staked_bal: "0"
  transfer_channel_id: channel-0
  unbonding_frequency: "1"
  validators:
  - address: cosmosvaloper1uk4ze0x4nvh4fk0xm4jdud58eqn4yxhrdt795p
    delegation_amt: "0"
    internal_exchange_rate: null
    name: pval1
    weight: "5"
  withdrawal_account:
    address: archway16mxc3u8cy0lz9w98ep5xecn8m7n8uegvj402u03099t7k792gxhqqtlsrk
    target: WITHDRAWAL
pagination:
  next_key: null
  total: "0"

LIST-DEPOSIT-RECORDS
deposit_record:
- amount: "0"
  denom: uarch
  deposit_epoch_number: "10"
  host_zone_id: localnet
  id: "1"
  source: STRIDE
  status: TRANSFER_QUEUE
- amount: "0"
  denom: uatom
  deposit_epoch_number: "10"
  host_zone_id: GAIA
  id: "2"
  source: STRIDE
  status: TRANSFER_QUEUE
pagination:
  next_key: null
  total: "0"

LIST-EPOCH-UNBONDING-RECORDS
epoch_unbonding_record:
- epoch_number: "4"
  host_zone_unbondings:
  - denom: uarch
    host_zone_id: localnet
    native_token_amount: "0"
    st_token_amount: "0"
    status: UNBONDING_QUEUE
    unbonding_time: "0"
    user_redemption_records: []
  - denom: uatom
    host_zone_id: GAIA
    native_token_amount: "0"
    st_token_amount: "0"
    status: UNBONDING_QUEUE
    unbonding_time: "0"
    user_redemption_records: []
pagination:
  next_key: null
  total: "0"

LIST-USER-REDEMPTION-RECORDS
pagination:
  next_key: null
  total: "0"
user_redemption_record: []
```

### Full Liquid staking - uarch and Full Redemption Workflow

root@swordfish-Lenovo-Y720-15IKB:/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts# build/archwayd --home /media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/photo1 tx ibc-transfer transfer transfer channel-0 stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 4000000uarch --from pval1 -y
```
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 54BBCC0D3B6F1EDA3ECB396654AD10C2E8269D937AAED90061F245D3DE0E0640
```

#build/strided --home /media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1 q bank balances stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
```
balances:
- amount: "4000000"
  denom: ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
- amount: "1000000000"
  denom: ustrd
pagination:
  next_key: null
  total: "0"
```

#build/strided --home /media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1 tx stakeibc liquid-stake 1000000 uarch --keyring-backend test --from admin --chain-id STRIDE -y
```
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 62E3A222D273428D576FF3E4D201217B34C6D77552F228DA1AE51BC754294361
```

#root@swordfish-Lenovo-Y720-15IKB:/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts# build/strided --home /media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1 q tx 62E3A222D273428D576FF3E4D201217B34C6D77552F228DA1AE51BC754294361
```
code: 0
codespace: ""
data: 12290A272F7374726964652E7374616B656962632E4D73674C69717569645374616B65526573706F6E7365
events:
- attributes:
  - index: true
    key: fee
    value: ""
  - index: true
    key: fee_payer
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
  type: tx
- attributes:
  - index: true
    key: acc_seq
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8/5
  type: tx
- attributes:
  - index: true
    key: signature
    value: LnosCKm4Ae5/YpHs15OcVJNJu5zZsJYg6asmLgx+6P5x89+vx2XaTe+l3nk6LfmyK4O2azjIuBU6I3UnxcpJMA==
  type: tx
- attributes:
  - index: true
    key: action
    value: /stride.stakeibc.MsgLiquidStake
  - index: true
    key: sender
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
  type: message
- attributes:
  - index: true
    key: spender
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
  - index: true
    key: amount
    value: 1000000ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
  type: coin_spent
- attributes:
  - index: true
    key: receiver
    value: stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag
  - index: true
    key: amount
    value: 1000000ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
  type: coin_received
- attributes:
  - index: true
    key: recipient
    value: stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag
  - index: true
    key: sender
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
  - index: true
    key: amount
    value: 1000000ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
  type: transfer
- attributes:
  - index: true
    key: sender
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
  type: message
- attributes:
  - index: true
    key: receiver
    value: stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza
  - index: true
    key: amount
    value: 1000000stuarch
  type: coin_received
- attributes:
  - index: true
    key: minter
    value: stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza
  - index: true
    key: amount
    value: 1000000stuarch
  type: coinbase
- attributes:
  - index: true
    key: spender
    value: stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza
  - index: true
    key: amount
    value: 1000000stuarch
  type: coin_spent
- attributes:
  - index: true
    key: receiver
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
  - index: true
    key: amount
    value: 1000000stuarch
  type: coin_received
- attributes:
  - index: true
    key: recipient
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
  - index: true
    key: sender
    value: stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza
  - index: true
    key: amount
    value: 1000000stuarch
  type: transfer
- attributes:
  - index: true
    key: sender
    value: stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza
  type: message
- attributes:
  - index: true
    key: module
    value: stakeibc
  - index: true
    key: liquid_staker
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
  - index: true
    key: host_zone
    value: localnet
  - index: true
    key: native_base_denom
    value: uarch
  - index: true
    key: native_ibc_denom
    value: ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
  - index: true
    key: native_amount
    value: "1000000"
  - index: true
    key: sttoken_amount
    value: "1000000"
  type: liquid_stake
gas_used: "103371"
gas_wanted: "200000"
height: "486"
info: ""
logs:
- events:
  - attributes:
    - key: action
      value: /stride.stakeibc.MsgLiquidStake
    - key: sender
      value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
    type: message
  - attributes:
    - key: spender
      value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
    - key: amount
      value: 1000000ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
    type: coin_spent
  - attributes:
    - key: receiver
      value: stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag
    - key: amount
      value: 1000000ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
    type: coin_received
  - attributes:
    - key: recipient
      value: stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag
    - key: sender
      value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
    - key: amount
      value: 1000000ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
    type: transfer
  - attributes:
    - key: sender
      value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
    type: message
  - attributes:
    - key: receiver
      value: stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza
    - key: amount
      value: 1000000stuarch
    type: coin_received
  - attributes:
    - key: minter
      value: stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza
    - key: amount
      value: 1000000stuarch
    type: coinbase
  - attributes:
    - key: spender
      value: stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza
    - key: amount
      value: 1000000stuarch
    type: coin_spent
  - attributes:
    - key: receiver
      value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
    - key: amount
      value: 1000000stuarch
    type: coin_received
  - attributes:
    - key: recipient
      value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
    - key: sender
      value: stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza
    - key: amount
      value: 1000000stuarch
    type: transfer
  - attributes:
    - key: sender
      value: stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza
    type: message
  - attributes:
    - key: module
      value: stakeibc
    - key: liquid_staker
      value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
    - key: host_zone
      value: localnet
    - key: native_base_denom
      value: uarch
    - key: native_ibc_denom
      value: ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
    - key: native_amount
      value: "1000000"
    - key: sttoken_amount
      value: "1000000"
    type: liquid_stake
  log: ""
  msg_index: 0
raw_log: '[{"msg_index":0,"events":[{"type":"message","attributes":[{"key":"action","value":"/stride.stakeibc.MsgLiquidStake"},{"key":"sender","value":"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8"}]},{"type":"coin_spent","attributes":[{"key":"spender","value":"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8"},{"key":"amount","value":"1000000ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034"}]},{"type":"coin_received","attributes":[{"key":"receiver","value":"stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag"},{"key":"amount","value":"1000000ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag"},{"key":"sender","value":"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8"},{"key":"amount","value":"1000000ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034"}]},{"type":"message","attributes":[{"key":"sender","value":"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8"}]},{"type":"coin_received","attributes":[{"key":"receiver","value":"stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza"},{"key":"amount","value":"1000000stuarch"}]},{"type":"coinbase","attributes":[{"key":"minter","value":"stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza"},{"key":"amount","value":"1000000stuarch"}]},{"type":"coin_spent","attributes":[{"key":"spender","value":"stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza"},{"key":"amount","value":"1000000stuarch"}]},{"type":"coin_received","attributes":[{"key":"receiver","value":"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8"},{"key":"amount","value":"1000000stuarch"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8"},{"key":"sender","value":"stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza"},{"key":"amount","value":"1000000stuarch"}]},{"type":"message","attributes":[{"key":"sender","value":"stride1mvdq4nlupl39243qjz7sds5ez3rl9mnx253lza"}]},{"type":"liquid_stake","attributes":[{"key":"module","value":"stakeibc"},{"key":"liquid_staker","value":"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8"},{"key":"host_zone","value":"localnet"},{"key":"native_base_denom","value":"uarch"},{"key":"native_ibc_denom","value":"ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034"},{"key":"native_amount","value":"1000000"},{"key":"sttoken_amount","value":"1000000"}]}]}]'
timestamp: "2023-06-12T19:24:13Z"
tx:
  '@type': /cosmos.tx.v1beta1.Tx
  auth_info:
    fee:
      amount: []
      gas_limit: "200000"
      granter: ""
      payer: ""
    signer_infos:
    - mode_info:
        single:
          mode: SIGN_MODE_DIRECT
      public_key:
        '@type': /cosmos.crypto.secp256k1.PubKey
        key: A3ZqkrrLNDVlI97RP4onAgAmjb+j3K8MqN2KcxJjrwXk
      sequence: "5"
    tip: null
  body:
    extension_options: []
    memo: ""
    messages:
    - '@type': /stride.stakeibc.MsgLiquidStake
      amount: "1000000"
      creator: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
      host_denom: uarch
    non_critical_extension_options: []
    timeout_height: "0"
  signatures:
  - LnosCKm4Ae5/YpHs15OcVJNJu5zZsJYg6asmLgx+6P5x89+vx2XaTe+l3nk6LfmyK4O2azjIuBU6I3UnxcpJMA==
txhash: 62E3A222D273428D576FF3E4D201217B34C6D77552F228DA1AE51BC754294361
```

root@swordfish-Lenovo-Y720-15IKB:/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts# build/strided --home /media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1 q stakeibc list-host-zone
```
host_zone:
- address: stride1755g4dkhpw73gz9h9nwhlcefc6sdf8kcmvcwrk4rxfrz8xpxxjms7savm8
  bech32prefix: cosmos
  blacklisted_validators: []
  chain_id: GAIA
  connection_id: connection-0
  delegation_account:
    address: cosmos18kf4ehs3pvqw9pp2pzq7q9fhurzrkememc0gg4zvjuk4xav5vurq8ra6pk
    target: DELEGATION
  fee_account:
    address: cosmos1v4k7xa4a8h3d90zawznhltuyvqpj640y47azpq9g5qzzjsdw4m4ste46vl
    target: FEE
  halted: false
  host_denom: uatom
  ibc_denom: ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2
  last_redemption_rate: "1.000000000000000000"
  max_redemption_rate: "1.500000000000000000"
  min_redemption_rate: "0.900000000000000000"
  redemption_account:
    address: cosmos19d9kj4n3mcdlvj8l5rrqh0qmk35u7qltu8nrmfgxpu0q6e9jc8msdq7hgv
    target: REDEMPTION
  redemption_rate: "1.000000000000000000"
  staked_bal: "0"
  transfer_channel_id: channel-0
  unbonding_frequency: "1"
  validators:
  - address: cosmosvaloper1uk4ze0x4nvh4fk0xm4jdud58eqn4yxhrdt795pcosmosvaloper1uk4ze0x4nvh4fk0xm4jdud58eqn4yxhrdt795p
    delegation_amt: "0"
    internal_exchange_rate: null
    name: gval1
    weight: "5"
  withdrawal_account:
    address: cosmos15ytq6r77chp4edvzjvdt75xt2x6cststypu94tnknjhzwv6vqqksjz5r0j
    target: WITHDRAWAL
- address: stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag
  bech32prefix: archway
  blacklisted_validators: []
  chain_id: localnet
  connection_id: connection-1
  delegation_account:
    address: archway1j8k542t5tjmnlc4f7lw4p4xnjg6hz6hxxx59yt2cfjwu2053w3wssyh9pd
    target: DELEGATION
  fee_account:
    address: archway1942qt84f0wj9e3hes9ncg4rqasp67xe4f73nz456rchw6u99trjq2m8udd
    target: FEE
  halted: false
  host_denom: uarch
  ibc_denom: ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
  last_redemption_rate: "1.000000000000000000"
  max_redemption_rate: "1.500000000000000000"
  min_redemption_rate: "0.900000000000000000"
  redemption_account:
    address: archway18jhfzqj250rsvdwyvmcyju6c8vxarvfyhghw4xj3ez7p5urly2vqsgsg0t
    target: REDEMPTION
  redemption_rate: "1.000000000000000000"
  staked_bal: "1000000"
  transfer_channel_id: channel-1
  unbonding_frequency: "1"
  validators:
  - address: archwayvaloper15js809uedxqs2wl0lyt58httasr5rlplj45fqw
    delegation_amt: "1000000"
    internal_exchange_rate: null
    name: pval1
    weight: "5"
  withdrawal_account:
    address: archway1r27yqfceelgjlxl8rd5kvlxl2nv67c7ell99hvyfnwhjfky7e63ssklwvv
    target: WITHDRAWAL
pagination:
  next_key: null
  total: "0"
```

root@swordfish-Lenovo-Y720-15IKB:/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts# build/strided --home /media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1 tx stakeibc redeem-stake 1000 PHOTO archway15js809uedxqs2wl0lyt58httasr5rlplj3atd8 --from admin --keyring-backend test --chain-id STRIDE -y
```
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 4ECC6062BE5B794005A1FAF5D9CC4F32CBEA18F6586EEA9B15D7DB342204608B
```

root@swordfish-Lenovo-Y720-15IKB:/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts# build/strided --home /media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1 tx stakeibc redeem-stake 1000 localnet archway15js809uedxqs2wl0lyt58httasr5rlplj3atd8 --from admin --keyring-backend test --chain-id STRIDE -y
```
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 1A6F1DBA73E3A4A469782B18C2637E5A118E06906A028A5A0A25734DFEDEC462
```

root@swordfish-Lenovo-Y720-15IKB:/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts# build/strided --home /media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1 q tx 1A6F1DBA73E3A4A469782B18C2637E5A118E06906A028A5A0A25734DFEDEC462
```
code: 0
codespace: ""
data: 12290A272F7374726964652E7374616B656962632E4D736752656465656D5374616B65526573706F6E7365
events:
- attributes:
  - index: true
    key: fee
    value: ""
  - index: true
    key: fee_payer
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
  type: tx
- attributes:
  - index: true
    key: acc_seq
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8/7
  type: tx
- attributes:
  - index: true
    key: signature
    value: T7oOOGG7SGYR0lKDJjMtz00mvxOtvb0ymBCbVeUmNI8rDY0XW2AFKcEvQXJaLb6i2jZGRPwExL8pluli0IummA==
  type: tx
- attributes:
  - index: true
    key: action
    value: /stride.stakeibc.MsgRedeemStake
  - index: true
    key: sender
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
  - index: true
    key: module
    value: stakeibc
  type: message
- attributes:
  - index: true
    key: spender
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
  - index: true
    key: amount
    value: 1000stuarch
  type: coin_spent
- attributes:
  - index: true
    key: receiver
    value: stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag
  - index: true
    key: amount
    value: 1000stuarch
  type: coin_received
- attributes:
  - index: true
    key: recipient
    value: stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag
  - index: true
    key: sender
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
  - index: true
    key: amount
    value: 1000stuarch
  type: transfer
- attributes:
  - index: true
    key: sender
    value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
  type: message
gas_used: "70867"
gas_wanted: "200000"
height: "925"
info: ""
logs:
- events:
  - attributes:
    - key: action
      value: /stride.stakeibc.MsgRedeemStake
    - key: sender
      value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
    - key: module
      value: stakeibc
    type: message
  - attributes:
    - key: spender
      value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
    - key: amount
      value: 1000stuarch
    type: coin_spent
  - attributes:
    - key: receiver
      value: stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag
    - key: amount
      value: 1000stuarch
    type: coin_received
  - attributes:
    - key: recipient
      value: stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag
    - key: sender
      value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
    - key: amount
      value: 1000stuarch
    type: transfer
  - attributes:
    - key: sender
      value: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
    type: message
  log: ""
  msg_index: 0
raw_log: '[{"msg_index":0,"events":[{"type":"message","attributes":[{"key":"action","value":"/stride.stakeibc.MsgRedeemStake"},{"key":"sender","value":"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8"},{"key":"module","value":"stakeibc"}]},{"type":"coin_spent","attributes":[{"key":"spender","value":"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8"},{"key":"amount","value":"1000stuarch"}]},{"type":"coin_received","attributes":[{"key":"receiver","value":"stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag"},{"key":"amount","value":"1000stuarch"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag"},{"key":"sender","value":"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8"},{"key":"amount","value":"1000stuarch"}]},{"type":"message","attributes":[{"key":"sender","value":"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8"}]}]}]'
timestamp: "2023-06-12T19:35:38Z"
tx:
  '@type': /cosmos.tx.v1beta1.Tx
  auth_info:
    fee:
      amount: []
      gas_limit: "200000"
      granter: ""
      payer: ""
    signer_infos:
    - mode_info:
        single:
          mode: SIGN_MODE_DIRECT
      public_key:
        '@type': /cosmos.crypto.secp256k1.PubKey
        key: A3ZqkrrLNDVlI97RP4onAgAmjb+j3K8MqN2KcxJjrwXk
      sequence: "7"
    tip: null
  body:
    extension_options: []
    memo: ""
    messages:
    - '@type': /stride.stakeibc.MsgRedeemStake
      amount: "1000"
      creator: stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
      host_zone: localnet
      receiver: archway15js809uedxqs2wl0lyt58httasr5rlplj3atd8
    non_critical_extension_options: []
    timeout_height: "0"
  signatures:
  - T7oOOGG7SGYR0lKDJjMtz00mvxOtvb0ymBCbVeUmNI8rDY0XW2AFKcEvQXJaLb6i2jZGRPwExL8pluli0IummA==
txhash: 1A6F1DBA73E3A4A469782B18C2637E5A118E06906A028A5A0A25734DFEDEC462
root@swordfish-Lenovo-Y720-15IKB:/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts# build/strided --home /media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1 q bank balances stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
balances:
- amount: "3000000"
  denom: ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
- amount: "999000"
  denom: stuarch
- amount: "1000000000"
  denom: ustrd
pagination:
  next_key: null
  total: "0"
```

Epoch based Central liquid stake of Archway Rewards and Redemption at maximum redemption rate of host zone, here Photosynthesis-Archway chain. 

```
/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/cronenable exists.
archwayd ibc transfer
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 1DFC0DA65103B032E57884FAC7F78E29643CFD4E9A4D4FAB9BFDF9B195A4EB97
strided track balance
balances:
- amount: "2100"
  denom: ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
- amount: "692"
  denom: stuarch
- amount: "1000000000"
  denom: ustrd
pagination:
  next_key: null
  total: "0"
strided liquid stake
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 9EE7B25286398E6DDB335AF0370A30932DC3698F14070D19A7A6EE63FDC40814
strided track balance
balances:
- amount: "2400"
  denom: ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
- amount: "788"
  denom: stuarch
- amount: "1000000000"
  denom: ustrd
pagination:
  next_key: null
  total: "0"
strided binary
archwayd binary
/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/cronenable exists.
archwayd ibc transfer
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: D8DDAB63FBC74CBC0DC1E83D420A9285B59F4AC3B3C8254DFCF4315BD190EA87
strided track balance
balances:
- amount: "2400"
  denom: ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
- amount: "788"
  denom: stuarch
- amount: "1000000000"
  denom: ustrd
pagination:
  next_key: null
  total: "0"
strided liquid stake
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 58E1F9E5E47134C38B058FE397967BF7ACA326C624476D10E1B1F4DA0EEFCE3F
strided track balance
balances:
- amount: "2700"
  denom: ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
- amount: "884"
  denom: stuarch
- amount: "1000000000"
  denom: ustrd
pagination:
  next_key: null
  total: "0"
```





Photosynthesis-Archway chain run Log snapshot





```
dockernet-photo1-1  | 10:58PM INF Iterating over contract metadata: airdrop_recipient_address: archway1wmuuy0eqvhq5s3j9e80p8judf55c8v8mgwfytr
dockernet-photo1-1  | airdrop_vesting_period: 6000
dockernet-photo1-1  | archway_reward_funds_transfer_address: archway1gnvac03v6xgtz3vt00p25j2nq28j9c55jlfntt
dockernet-photo1-1  | liquid_stake_interval: 1
dockernet-photo1-1  | liquidity_provider_address: archway1wmuuy0eqvhq5s3j9e80p8judf55c8v8mgwfytr
dockernet-photo1-1  | liquidity_provider_commission: 2
dockernet-photo1-1  | liquidity_token_address: archway1smd403gckfc4m3upzxfuwxkree5lr9854u4un9
dockernet-photo1-1  | maximum_threshold: 4
dockernet-photo1-1  | minimum_reward_amount: 100
dockernet-photo1-1  | owner_address: archway1qygx0pxuttycdddzz5lre5rlxcxjemthwmlh63
dockernet-photo1-1  | redemption_address: archway18kpsdc76xg5884ey3qnesqtw8l9n06yw0u898p
dockernet-photo1-1  | redemption_interval: 1
dockernet-photo1-1  | redemption_interval_threshold: 1
dockernet-photo1-1  | redemption_rate_threshold: 1
dockernet-photo1-1  | rewards_address: archway1p74uyn42qnktc50mxflx6frzs4luqwtyjq3cwp
dockernet-photo1-1  | rewards_withdrawal_interval: 10
dockernet-photo1-1  |  module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Checking epoch info: {arch-central-liquid-stake-interval-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 56s 78 2023-07-06 22:58:13.355016608 +0000 UTC true 4209} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Processing ARCH_CENTRAL_LIQUID_STAKE_INTERVAL_EPOCH: {arch-central-liquid-stake-interval-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 56s 78 2023-07-06 22:58:13.355016608 +0000 UTC true 4209} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Retrieved EpochInfo for epochstypes.ARCH_CENTRAL_LIQUID_STAKE_INTERVAL_EPOCH: {arch-central-liquid-stake-interval-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 56s 78 2023-07-06 22:58:13.355016608 +0000 UTC true 4209} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Retrieved EpochInfo for epochstypes.LIQUID_STAKING_DApp_Rewards_EPOCH: {liquid-staking-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 7s 617 2023-07-06 22:58:13.355016608 +0000 UTC true 4209} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF MinimumRewardAmount is greater than 0: 100 module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF CurrentEpoch 78 is not 0 and is a multiple of 1 module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Retrieved TotalLiquidStake: 0 module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Command ran successfully module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Computed LiquidStake: 0 module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Contract Address: archway1p74uyn42qnktc50mxflx6frzs4luqwtyjq3cwp, Liquid Token Amount: %!d(float64=0)
dockernet-photo1-1  |  module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Distributed Liquidity for epoch 617 and liquid stake 0 module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Checking epoch info: {day 2023-07-06 21:46:21.355016608 +0000 UTC 1m0s 72 2023-07-06 22:57:21.355016608 +0000 UTC true 4158} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Checking epoch info: {liquid-staking-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 7s 617 2023-07-06 22:58:13.355016608 +0000 UTC true 4209} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Processing LiquidStakeDappRewards epoch: {liquid-staking-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 7s 617 2023-07-06 22:58:13.355016608 +0000 UTC true 4209} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Retrieved EpochInfo: {liquid-staking-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 7s 617 2023-07-06 22:58:13.355016608 +0000 UTC true 4209} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF MinimumRewardAmount is greater than 0: 100 module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF CurrentEpoch 617 is not 0 and is a multiple of LiquidStakeInterval 1 module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF RewardsAddress is not empty: archway1p74uyn42qnktc50mxflx6frzs4luqwtyjq3cwp module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Retrieved CumulativeRewardAmount:  module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Checking epoch info: {mint 2023-07-06 21:46:21.355016608 +0000 UTC 1m0s 72 2023-07-06 22:57:21.355016608 +0000 UTC true 4158} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Checking epoch info: {redemption-rate-query-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 6h0m0s 1 2023-07-06 21:46:21.355016608 +0000 UTC true 1} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Processing REDEMPTION_RATE_QUERY_EPOCH: {redemption-rate-query-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 6h0m0s 1 2023-07-06 21:46:21.355016608 +0000 UTC true 1} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Retrieved EpochInfo for epochstypes.REDEMPTION_RATE_QUERY_EPOCH: {Identifier:redemption-rate-query-epoch StartTime:2023-07-06 21:46:21.355016608 +0000 UTC Duration:6h0m0s CurrentEpoch:1 CurrentEpochStartTime:2023-07-06 21:46:21.355016608 +0000 UTC EpochCountingStarted:true CurrentEpochStartHeight:1} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF CurrentEpoch 1 is a multiple of RedemptionIntervalThreshold 1 module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Using RedemptionRateThreshold: 1 module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF CurrentEpoch 1 is not 0 and is a multiple of RedemptionRateThreshold 1 module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Successfully queried RedemptionRate: 1.2 module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Checking epoch info: {rewards_withdrawal-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 7h0m0s 1 2023-07-06 21:46:21.355016608 +0000 UTC true 1} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Processing REWARDS_WITHDRAWAL_EPOCH: {Identifier:rewards_withdrawal-epoch StartTime:2023-07-06 21:46:21.355016608 +0000 UTC Duration:7h0m0s CurrentEpoch:1 CurrentEpochStartTime:2023-07-06 21:46:21.355016608 +0000 UTC EpochCountingStarted:true CurrentEpochStartHeight:1} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Retrieved EpochInfo for epochstypes.REWARDS_WITHDRAWAL_EPOCH: {Identifier:rewards_withdrawal-epoch StartTime:2023-07-06 21:46:21.355016608 +0000 UTC Duration:7h0m0s CurrentEpoch:1 CurrentEpochStartTime:2023-07-06 21:46:21.355016608 +0000 UTC EpochCountingStarted:true CurrentEpochStartHeight:1} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Deleted 0 reward records module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Checking epoch info: {stride_epoch 2023-07-06 21:46:21.355016608 +0000 UTC 6h0m0s 1 2023-07-06 21:46:21.355016608 +0000 UTC true 1} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Checking epoch info: {week 2023-07-06 21:46:21.355016608 +0000 UTC 1m0s 72 2023-07-06 22:57:21.355016608 +0000 UTC true 4158} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Inflation rewards - block module=x/rewards rewards=82451
dockernet-photo1-1  | 10:58PM INF executed block height=4215 module=state num_invalid_txs=0 num_valid_txs=0
dockernet-photo1-1  | 10:58PM INF commit synced commit=436F6D6D697449447B5B323132203135203732203239203131392032353320322035302036372033203738203532203135203230312038352035392031343820323331203131392031373220313231203936203930203233203138322031393020313133203133352031393920393920323432203231375D3A313037377D
dockernet-photo1-1  | 10:58PM INF committed state app_hash=D40F481D77FD023243034E340FC9553B94E777AC79605A17B6BE7187C763F2D9 height=4215 module=state num_txs=0
dockernet-photo1-1  | 10:58PM INF indexed block exents height=4215 module=txindex
dockernet-photo1-1  | 10:58PM INF client state updated client-id=07-tendermint-0 height=0-2736 module=x/ibc/client
dockernet-photo1-1  | 10:58PM INF packet received dst_channel=channel-6 dst_port=icahost module=x/ibc/channel sequence=191 src_channel=channel-6 src_port=icacontroller-localnet.DELEGATION
dockernet-photo1-1  | 10:58PM INF acknowledgement written dst_channel=channel-6 dst_port=icahost module=x/ibc/channel sequence=191 src_channel=channel-6 src_port=icacontroller-localnet.DELEGATION
dockernet-photo1-1  | 10:58PM INF client state updated client-id=07-tendermint-0 height=0-2736 module=x/ibc/client
dockernet-photo1-1  | 10:58PM INF packet received dst_channel=channel-6 dst_port=icahost module=x/ibc/channel sequence=191 src_channel=channel-6 src_port=icacontroller-localnet.DELEGATION
dockernet-photo1-1  | 10:58PM INF acknowledgement written dst_channel=channel-6 dst_port=icahost module=x/ibc/channel sequence=191 src_channel=channel-6 src_port=icacontroller-localnet.DELEGATION
dockernet-photo1-1  | 10:58PM INF Timed out dur=954.911705 height=4216 module=consensus round=0 step=1
dockernet-photo1-1  | 10:58PM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"36C1DE788C35E7CACAE2556197BF8B4115D9B16DD654B5AEDCFD883877481C4A","parts":{"hash":"9FCDF3794639046F7D96F702F7594189BCD89AC93CDEAB82C5829387748C1CE1","total":1}},"height":4216,"pol_round":-1,"round":0,"signature":"9EpWvBIiDO4942sjX12YBNMe/ac02OxG7FXN1M4MNnD7J5DA1C2ue0s8pwtY7l0piBF71Xlq6NBBtBArBQBdCA==","timestamp":"2023-07-06T22:58:21.677835049Z"}
dockernet-photo1-1  | 10:58PM INF received complete proposal block hash=36C1DE788C35E7CACAE2556197BF8B4115D9B16DD654B5AEDCFD883877481C4A height=4216 module=consensus
dockernet-photo1-1  | 10:58PM INF finalizing commit of block hash={} height=4216 module=consensus num_txs=1 root=D40F481D77FD023243034E340FC9553B94E777AC79605A17B6BE7187C763F2D9
dockernet-photo1-1  | 10:58PM INF minted coins from module account amount=412256uarch from=mint module=x/bank
dockernet-photo1-1  | 10:58PM INF Minimum consensus fee update fee={"amount":"0.082451000000000000","denom":"uarch"} module=x/rewards
dockernet-photo1-1  | 10:58PM INF ending epoch identifier=liquid-staking-epoch module=x/epochs
dockernet-photo1-1  | 10:58PM INF Retrieved state from rewardKeeper: {0xc000eda3d0 0xc000f18d00} module=x/photosynthesis
dockernet-photo1-1  | 10:58PM INF Retrieved contract metadata state: {{0xc00ddb9920 [0]} 0xc000f18d00 {0xc00013a020 {0xc00dd66740 map[0xc000eda2c0:0xc00dd66980 0xc000eda2d0:0xc00dd66840 0xc000eda2e0:0xc00dd66900 0xc000eda2f0:0xc00dd66c00 0xc000eda300:0xc00dd667c0 0xc000eda310:0xc00dd66940 0xc000eda320:0xc00dd66b40 0xc000eda330:0xc00dd66cc0 0xc000eda340:0xc00dd66ac0 0xc000eda350:0xc00dd66bc0 0xc000eda360:0xc00dd66b80 0xc000eda370:0xc00dd66780 0xc000eda380:0xc00dd669c0 0xc000eda390:0xc00dd66b00 0xc000eda3a0:0xc00dd66c80 0xc000eda3b0:0xc00dd66a40 0xc000eda3c0:0xc00dd668c0 0xc000eda3d0:0xc00dd66d00 0xc000eda3f0:0xc00dd66a80 0xc000eda400:0xc00dd66a00 0xc000eda410:0xc00dd66880 0xc000eda610:0xc00dd66800 0xc000eda4b0:0xc00dd66c40] map[acc:0xc000eda2c0 authz:0xc000eda3a0 bank:0xc000eda2d0 capability:0xc000eda380 distribution:0xc000eda300 epochs:0xc000eda3f0 evidence:0xc000eda360 feegrant:0xc000eda390 gov:0xc000eda320 ibc:0xc000eda340 icahost:0xc000eda410 mem_capability:0xc000eda610 mint:0xc000eda2f0 params:0xc000eda330 photosynthesis:0xc000eda400 rewards:0xc000eda3d0 slashing:0xc000eda310 staking:0xc000eda2e0 tracking:0xc000eda3c0 transfer:0xc000eda370 transient_params:0xc000eda4b0 upgrade:0xc000eda350 wasm:0xc000eda3b0] <nil> map[] map[]} {{11 0} localnet 4216 {665438705 63824281100 <nil>} {[10 169 2 219 35 191 229 252 234 254 45 146 182 148 102 202 47 231 112 173 241 37 29 106 157 2 115 233 133 59 245 14] {1 [189 139 38 34 148 73 245 114 205 166 197 168 61 105 65 37 32 46 136 9 244 207 128 130 144 203 246 70 242 104 1 147]}} [143 255 170 228 207 225 57 4 147 188 124 27 195 109 220 231 203 164 110 232 112 206 221 163 217 88 118 43 163 192 179 127] [74 160 184 199 182 63 95 9 4 70 61 34 183 181 237 143 105 43 144 51 18 233 135 58 59 124 166 64 4 101 100 135] [226 234 202 167 156 127 33 65 74 231 182 19 111 37 123 205 40 169 19 224 170 137 108 192 8 72 252 58 38 13 33 48] [226 234 202 167 156 127 33 65 74 231 182 19 111 37 123 205 40 169 19 224 170 137 108 192 8 72 252 58 38 13 33 48] [4 128 145 188 125 220 40 63 119 191 191 145 215 60 68 218 88 195 223 138 156 188 134 116 5 216 183 243 218 173 162 47] [212 15 72 29 119 253 2 50 67 3 78 52 15 201 85 59 148 231 119 172 121 96 90 23 182 190 113 135 199 99 242 217] [227 176 196 66 152 252 28 20 154 251 244 200 153 111 185 36 39 174 65 228 100 155 147 76 164 149 153 27 120 82 184 85] [227 176 196 66 152 252 28 20 154 251 244 200 153 111 185 36 39 174 65 228 100 155 147 76 164 149 153 27 120 82 184 85] [112 19 17 47 145 122 178 15 33 75 187 183 218 233 254 162 143 79 208 50]} [54 193 222 120 140 53 231 202 202 226 85 97 151 191 139 65 21 217 177 109 214 84 181 174 220 253 136 56 119 72 28 74] localnet [] {{{{0xc000138010 false  [] [] [] <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>}} 1 <nil> [123] [{}] false}} [] 0xc00baae4f8 0xc00baae698 false false [] 0xc00e848800 0xc00d041b00}} module=x/photosynthesis




dockernet-photo1-1  | 10:24PM INF Iterating over contract metadata: airdrop_recipient_address: archway1wmuuy0eqvhq5s3j9e80p8judf55c8v8mgwfytr
dockernet-photo1-1  | airdrop_vesting_period: 6000
dockernet-photo1-1  | archway_reward_funds_transfer_address: archway1gnvac03v6xgtz3vt00p25j2nq28j9c55jlfntt
dockernet-photo1-1  | liquid_stake_interval: 1
dockernet-photo1-1  | liquidity_provider_address: archway1wmuuy0eqvhq5s3j9e80p8judf55c8v8mgwfytr
dockernet-photo1-1  | liquidity_provider_commission: 2
dockernet-photo1-1  | liquidity_token_address: archway1smd403gckfc4m3upzxfuwxkree5lr9854u4un9
dockernet-photo1-1  | maximum_threshold: 4
dockernet-photo1-1  | minimum_reward_amount: 100
dockernet-photo1-1  | owner_address: archway1qygx0pxuttycdddzz5lre5rlxcxjemthwmlh63
dockernet-photo1-1  | redemption_address: archway18kpsdc76xg5884ey3qnesqtw8l9n06yw0u898p
dockernet-photo1-1  | redemption_interval: 1
dockernet-photo1-1  | redemption_interval_threshold: 1
dockernet-photo1-1  | redemption_rate_threshold: 1
dockernet-photo1-1  | rewards_address: archway1p74uyn42qnktc50mxflx6frzs4luqwtyjq3cwp
dockernet-photo1-1  | rewards_withdrawal_interval: 10
dockernet-photo1-1  |  module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Checking epoch info: {arch-central-liquid-stake-interval-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 56s 41 2023-07-06 22:23:41.355016608 +0000 UTC true 2175} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Processing ARCH_CENTRAL_LIQUID_STAKE_INTERVAL_EPOCH: {arch-central-liquid-stake-interval-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 56s 41 2023-07-06 22:23:41.355016608 +0000 UTC true 2175} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Retrieved EpochInfo for epochstypes.ARCH_CENTRAL_LIQUID_STAKE_INTERVAL_EPOCH: {arch-central-liquid-stake-interval-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 56s 41 2023-07-06 22:23:41.355016608 +0000 UTC true 2175} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Retrieved EpochInfo for epochstypes.LIQUID_STAKING_DApp_Rewards_EPOCH: {liquid-staking-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 7s 327 2023-07-06 22:24:23.355016608 +0000 UTC true 2216} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF MinimumRewardAmount is greater than 0: 100 module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF CurrentEpoch 41 is not 0 and is a multiple of 1 module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Retrieved TotalLiquidStake: 0 module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Command ran successfully module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Computed LiquidStake: 0 module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Distributed Liquidity for epoch 327 and liquid stake 0 module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Checking epoch info: {day 2023-07-06 21:46:21.355016608 +0000 UTC 1m0s 39 2023-07-06 22:24:21.355016608 +0000 UTC true 2214} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Checking epoch info: {liquid-staking-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 7s 327 2023-07-06 22:24:23.355016608 +0000 UTC true 2216} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Processing LiquidStakeDappRewards epoch: {liquid-staking-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 7s 327 2023-07-06 22:24:23.355016608 +0000 UTC true 2216} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Retrieved EpochInfo: {liquid-staking-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 7s 327 2023-07-06 22:24:23.355016608 +0000 UTC true 2216} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF MinimumRewardAmount is greater than 0: 100 module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF CurrentEpoch 327 is not 0 and is a multiple of LiquidStakeInterval 1 module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF RewardsAddress is not empty: archway1p74uyn42qnktc50mxflx6frzs4luqwtyjq3cwp module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Retrieved CumulativeRewardAmount: 15000uarch module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF CumulativeRewardAmount is greater than or equal to MinimumRewardAmount module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Created ContractLiquidStakeDepositRecordsForEpoch: contract_address:"archway1p74uyn42qnktc50mxflx6frzs4luqwtyjq3cwp" epoch:327 amount:15000 status:"pending"  module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Enqueued LiquidStakeRecord module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF EmitLiquidStakeDepositRecordCreatedEvent for record: contract_address:"archway1p74uyn42qnktc50mxflx6frzs4luqwtyjq3cwp" epoch:327 amount:15000 status:"pending"  and amount: 15000 module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Checking epoch info: {mint 2023-07-06 21:46:21.355016608 +0000 UTC 1m0s 39 2023-07-06 22:24:21.355016608 +0000 UTC true 2214} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Checking epoch info: {redemption-rate-query-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 6h0m0s 1 2023-07-06 21:46:21.355016608 +0000 UTC true 1} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Processing REDEMPTION_RATE_QUERY_EPOCH: {redemption-rate-query-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 6h0m0s 1 2023-07-06 21:46:21.355016608 +0000 UTC true 1} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Retrieved EpochInfo for epochstypes.REDEMPTION_RATE_QUERY_EPOCH: {Identifier:redemption-rate-query-epoch StartTime:2023-07-06 21:46:21.355016608 +0000 UTC Duration:6h0m0s CurrentEpoch:1 CurrentEpochStartTime:2023-07-06 21:46:21.355016608 +0000 UTC EpochCountingStarted:true CurrentEpochStartHeight:1} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF CurrentEpoch 1 is a multiple of RedemptionIntervalThreshold 1 module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Using RedemptionRateThreshold: 1 module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF CurrentEpoch 1 is not 0 and is a multiple of RedemptionRateThreshold 1 module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Successfully queried RedemptionRate: 1.2 module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Checking epoch info: {rewards_withdrawal-epoch 2023-07-06 21:46:21.355016608 +0000 UTC 7h0m0s 1 2023-07-06 21:46:21.355016608 +0000 UTC true 1} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Processing REWARDS_WITHDRAWAL_EPOCH: {Identifier:rewards_withdrawal-epoch StartTime:2023-07-06 21:46:21.355016608 +0000 UTC Duration:7h0m0s CurrentEpoch:1 CurrentEpochStartTime:2023-07-06 21:46:21.355016608 +0000 UTC EpochCountingStarted:true CurrentEpochStartHeight:1} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Retrieved EpochInfo for epochstypes.REWARDS_WITHDRAWAL_EPOCH: {Identifier:rewards_withdrawal-epoch StartTime:2023-07-06 21:46:21.355016608 +0000 UTC Duration:7h0m0s CurrentEpoch:1 CurrentEpochStartTime:2023-07-06 21:46:21.355016608 +0000 UTC EpochCountingStarted:true CurrentEpochStartHeight:1} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Deleted 1 reward records module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Checking epoch info: {stride_epoch 2023-07-06 21:46:21.355016608 +0000 UTC 6h0m0s 1 2023-07-06 21:46:21.355016608 +0000 UTC true 1} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Checking epoch info: {week 2023-07-06 21:46:21.355016608 +0000 UTC 1m0s 39 2023-07-06 22:24:21.355016608 +0000 UTC true 2214} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Inflation rewards - block module=x/rewards rewards=82421
dockernet-photo1-1  | 10:24PM INF executed block height=2222 module=state num_invalid_txs=0 num_valid_txs=0
dockernet-photo1-1  | 10:24PM INF commit synced commit=436F6D6D697449447B5B3536203139392032343220313839203339203134362038382036382031323220313530203135342037322037372031323020323434203135382031353920313938203430203237203133302031333420313637203132302032343620323339203137322031333020313234203138203539203130325D3A3841457D
dockernet-photo1-1  | 10:24PM INF committed state app_hash=38C7F2BD279258447A969A484D78F49E9FC6281B8286A778F6EFAC827C123B66 height=2222 module=state num_txs=0
dockernet-photo1-1  | 10:24PM INF indexed block exents height=2222 module=txindex
dockernet-photo1-1  | 10:24PM INF Timed out dur=971.20038 height=2223 module=consensus round=0 step=1
dockernet-photo1-1  | 10:24PM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"26F56C4A43F792F24F50C771F0463DE2710E6CF1CC98AA37C14CA739DB5015D9","parts":{"hash":"B7B18011224FD0623D1CED5DD68A2C32D3DC613D429AE221923045EF3F226718","total":1}},"height":2223,"pol_round":-1,"round":0,"signature":"DqFCAhzIRbP4o1eHyH5Z3gecUrUWnMprcALQXqMAfvMt1kXjtfuERWnp81h40kU8fUKwanu6Bn16qyS54wiqDw==","timestamp":"2023-07-06T22:24:31.924120153Z"}
dockernet-photo1-1  | 10:24PM INF received complete proposal block hash=26F56C4A43F792F24F50C771F0463DE2710E6CF1CC98AA37C14CA739DB5015D9 height=2223 module=consensus
dockernet-photo1-1  | 10:24PM INF finalizing commit of block hash={} height=2223 module=consensus num_txs=0 root=38C7F2BD279258447A969A484D78F49E9FC6281B8286A778F6EFAC827C123B66
dockernet-photo1-1  | 10:24PM INF minted coins from module account amount=412109uarch from=mint module=x/bank
dockernet-photo1-1  | 10:24PM INF Minimum consensus fee update fee={"amount":"0.082421000000000000","denom":"uarch"} module=x/rewards
dockernet-photo1-1  | 10:24PM INF ending epoch identifier=liquid-staking-epoch module=x/epochs
dockernet-photo1-1  | 10:24PM INF Retrieved state from rewardKeeper: {0xc000eda3d0 0xc000f18d00} module=x/photosynthesis
dockernet-photo1-1  | 10:24PM INF Retrieved contract metadata state: {{0xc0064af620 [0]} 0xc000f18d00 {0xc00013a020 {0xc003a9ba40 map[0xc000eda2c0:0xc003a9bd80 0xc000eda2d0:0xc003a9be40 0xc000eda2e0:0xc003bb8040 0xc000eda2f0:0xc003a9bd00 0xc000eda300:0xc003a9bb40 0xc000eda310:0xc003a9bc00 0xc000eda320:0xc003a9bc40 0xc000eda330:0xc003a9bdc0 0xc000eda340:0xc003a9be80 0xc000eda350:0xc003a9bac0 0xc000eda360:0xc003a9bf00 0xc000eda370:0xc003a9bd40 0xc000eda380:0xc003a9bf40 0xc000eda390:0xc003a9bfc0 0xc000eda3a0:0xc003a9bbc0 0xc000eda3b0:0xc003a9bc80 0xc000eda3c0:0xc003a9bb00 0xc000eda3d0:0xc003a9bf80 0xc000eda3f0:0xc003a9ba80 0xc000eda400:0xc003a9bec0 0xc000eda410:0xc003a9bcc0 0xc000eda610:0xc003a9be00 0xc000eda4b0:0xc003a9bb80] map[acc:0xc000eda2c0 authz:0xc000eda3a0 bank:0xc000eda2d0 capability:0xc000eda380 distribution:0xc000eda300 epochs:0xc000eda3f0 evidence:0xc000eda360 feegrant:0xc000eda390 gov:0xc000eda320 ibc:0xc000eda340 icahost:0xc000eda410 mem_capability:0xc000eda610 mint:0xc000eda2f0 params:0xc000eda330 photosynthesis:0xc000eda400 rewards:0xc000eda3d0 slashing:0xc000eda310 staking:0xc000eda2e0 tracking:0xc000eda3c0 transfer:0xc000eda370 transient_params:0xc000eda4b0 upgrade:0xc000eda350 wasm:0xc000eda3b0] <nil> map[] map[]} {{11 0} localnet 2223 {916820865 63824279070 <nil>} {[242 221 196 24 91 4 172 171 242 253 84 189 109 165 163 150 46 182 98 250 102 234 139 221 197 226 136 44 209 104 131 124] {1 [132 36 87 169 60 73 152 115 68 168 154 35 205 114 161 95 43 133 148 212 58 110 147 140 151 98 123 185 186 255 120 119]}} [251 176 191 44 84 161 8 137 39 202 204 129 116 158 97 194 66 77 150 95 34 4 148 175 53 251 204 143 116 233 162 168] [227 176 196 66 152 252 28 20 154 251 244 200 153 111 185 36 39 174 65 228 100 155 147 76 164 149 153 27 120 82 184 85] [226 234 202 167 156 127 33 65 74 231 182 19 111 37 123 205 40 169 19 224 170 137 108 192 8 72 252 58 38 13 33 48] [226 234 202 167 156 127 33 65 74 231 182 19 111 37 123 205 40 169 19 224 170 137 108 192 8 72 252 58 38 13 33 48] [4 128 145 188 125 220 40 63 119 191 191 145 215 60 68 218 88 195 223 138 156 188 134 116 5 216 183 243 218 173 162 47] [56 199 242 189 39 146 88 68 122 150 154 72 77 120 244 158 159 198 40 27 130 134 167 120 246 239 172 130 124 18 59 102] [227 176 196 66 152 252 28 20 154 251 244 200 153 111 185 36 39 174 65 228 100 155 147 76 164 149 153 27 120 82 184 85] [227 176 196 66 152 252 28 20 154 251 244 200 153 111 185 36 39 174 65 228 100 155 147 76 164 149 153 27 120 82 184 85] [112 19 17 47 145 122 178 15 33 75 187 183 218 233 254 162 143 79 208 50]} [38 245 108 74 67 247 146 242 79 80 199 113 240 70 61 226 113 14 108 241 204 152 170 55 193 76 167 57 219 80 21 217] localnet [] {{{{0xc000138010 false  [] [] [] <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>}} 1 <nil> [123] [{}] false}} [] 0xc00ac11718 0xc00ac11858 false false [] 0xc00aaf7660 0xc006172420}} module=x/photosynthesis
```


Relayer Logs Snapshot

```

STRIDE <> PHOTO - Creating client, connection, and transfer channel...ts=2023-07-06T21:51:11.751658Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet gas_used=109973 fees=11622uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=260 msg_types=/ibc.core.client.v1.MsgCreateClient tx_hash=A1255C3DF7A263822B1B7B15462E33CDBBCEF829034360D1CCCBB0D392E3555A
ts=2023-07-06T21:51:11.751749Z lvl=info msg="Client Created" src_chain_id=localnet src_client_id=07-tendermint-0 dst_chain_id=STRIDE
ts=2023-07-06T21:51:11.804171Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=89234 fees=964ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=166 msg_types=/ibc.core.client.v1.MsgCreateClient tx_hash=BDD2FF79807742ED4CD470704761ED5FAE4CE11708BF68A732F090616259B7B2
ts=2023-07-06T21:51:11.804237Z lvl=info msg="Client Created" src_chain_id=STRIDE src_client_id=07-tendermint-1 dst_chain_id=localnet
ts=2023-07-06T21:51:11.804270Z lvl=info msg="Clients created" src_client_id=07-tendermint-1 src_chain_id=STRIDE dst_client_id=07-tendermint-0 dst_chain_id=localnet
ts=2023-07-06T21:51:36.904018Z lvl=info msg="Starting event processor for connection handshake" src_chain_id=STRIDE src_client_id=07-tendermint-1 dst_chain_id=localnet dst_client_id=07-tendermint-0
ts=2023-07-06T21:51:36.914967Z lvl=info msg="Chain is not yet in sync" chain_name=stride chain_id=STRIDE latest_queried_block=162 latest_height=182
ts=2023-07-06T21:51:36.917441Z lvl=info msg="Chain is not yet in sync" chain_name=photo chain_id=localnet latest_queried_block=264 latest_height=284
ts=2023-07-06T21:51:37.916038Z lvl=info msg="Chain is in sync" chain_name=stride chain_id=STRIDE
ts=2023-07-06T21:51:37.918052Z lvl=info msg="Chain is in sync" chain_name=photo chain_id=localnet
ts=2023-07-06T21:51:39.918214Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=105077 fees=1205ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=184 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.connection.v1.MsgConnectionOpenInit tx_hash=F45AA0D32B0A14A3B711B49266D99ECF4026B56746883F4FEA5482F4BF4E5E79
ts=2023-07-06T21:51:42.323104Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet gas_used=178882 fees=19891uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=290 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.connection.v1.MsgConnectionOpenTry tx_hash=28D796F4F3C19182B6678EDD110DEB611A8C4A88B5B2871E14AF649E2E552D1E
ts=2023-07-06T21:51:44.813300Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=140464 fees=1629ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=187 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.connection.v1.MsgConnectionOpenAck tx_hash=D3F9169E66C219F04324A8597CC3BE1D19C57449EA9E7EC5B88C08EE5DC44DEB
ts=2023-07-06T21:51:47.409909Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet gas_used=139509 fees=15166uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=295 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.connection.v1.MsgConnectionOpenConfirm tx_hash=9CB16B957D2642E5E0834FB82003F514B885D8A3E07DD9B3EBD1FEE591FD5B76
ts=2023-07-06T21:51:47.953033Z lvl=info msg="Connection handshake termination candidate" path_name=stride-photo chain_id=localnet client_id=07-tendermint-0 termination_client_id=07-tendermint-0 observed_client_id=07-tendermint-0 termination_counterparty_client_id=07-tendermint-1 observed_counterparty_client_id=07-tendermint-1
ts=2023-07-06T21:51:47.953059Z lvl=info msg="Found termination condition for connection handshake" path_name=stride-photo chain_id=localnet client_id=07-tendermint-0
ts=2023-07-06T21:52:13.063746Z lvl=info msg="Starting event processor for channel handshake" src_chain_id=STRIDE src_port_id=transfer dst_chain_id=localnet dst_port_id=transfer
ts=2023-07-06T21:52:13.073211Z lvl=info msg="Chain is in sync" chain_name=stride chain_id=STRIDE
ts=2023-07-06T21:52:13.075336Z lvl=info msg="Chain is in sync" chain_name=photo chain_id=localnet
ts=2023-07-06T21:52:15.958452Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=152677 fees=1788ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=207 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgChannelOpenInit tx_hash=8B1E8998E5BD03023AABB35EE00554516A0B29832332A72D6D41D86CDA3B8AC7
ts=2023-07-06T21:52:18.967892Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet gas_used=194684 fees=21787uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=326 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgChannelOpenTry tx_hash=D776B30A8735D0D294D10384A9C231FCFA9EDCD9C91FC35C601C5D559A80478A
ts=2023-07-06T21:52:20.869201Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=116955 fees=1324ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=210 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgChannelOpenAck tx_hash=87ADD58D380BA414A68B6E38D6B21EFA748CD4B32A6ED790D3757AA902C70030
ts=2023-07-06T21:52:24.075272Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet gas_used=146456 fees=16000uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=331 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgChannelOpenConfirm tx_hash=515C6C64FB37FFCEBED0D75B9BAF8BEE8871FC2310AE64929F8F57F5D43954E3


dockernet-relayer-photo-1  | ts=2023-07-06T21:59:40.713487Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=109892 fees=1232ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=488 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=7B57FA733427744F9AE71715F2A6BE9E6C6C083F700A2A9167B3EB21698B10DD
dockernet-relayer-photo-1  | ts=2023-07-06T21:59:42.397312Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=179488 fees=19964uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=758 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=32A50D316A08CDB37FAFA4AA2E0255077EB98F11AB10EF568EFADF89F4F9E4D9
dockernet-relayer-photo-1  | ts=2023-07-06T21:59:45.345232Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=132971 fees=1569ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=491 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=982E7D946A300EB6135C5FBE3A6682ECA1BFFB700CFBF2561EB04B69A797FF11
dockernet-relayer-photo-1  | ts=2023-07-06T22:00:04.589148Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-0 packet_dst_channel=channel-1 gas_used=167173 fees=1977ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=503 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=41F77A12E93327ACD9E83765624307A1610F8E30131896994E79339BDF247AED
dockernet-relayer-photo-1  | ts=2023-07-06T22:00:07.212401Z lvl=warn msg="Error querying block data" chain_name=photo chain_id=localnet error="RPC error -32603 - Internal error: could not find results for height #782"
dockernet-relayer-photo-1  | ts=2023-07-06T22:00:07.308414Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-0 packet_dst_channel=channel-1 gas_used=151861 fees=16648uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=782 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=3E857417118F06EFF769A1332B56CCAD53C8650294B47B5B37F76E30051781EA
dockernet-relayer-photo-1  | ts=2023-07-06T22:00:21.997182Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=107096 fees=1233ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=514 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=97016753B41C49CC44B4EEB166201AC08DE0A56186D6CC04E96C3B5E61CF036E
dockernet-relayer-photo-1  | ts=2023-07-06T22:00:23.009180Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=231618 fees=26219uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=797 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket,/ibc.core.channel.v1.MsgRecvPacket tx_hash=58C704D9FFF8907B4D9A13025F53480529A481762E20E1F24BD314257AE9AFF3
dockernet-relayer-photo-1  | ts=2023-07-06T22:00:24.994560Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-1 packet_dst_channel=channel-0 gas_used=173196 fees=2055ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=516 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=313C15CAB530CBB782CF37E76B0113A81CE4D67414CF5D012C9283336B1548DC
dockernet-relayer-photo-1  | ts=2023-07-06T22:01:01.211277Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=109156 fees=1260ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=539 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=DE15C6EEB9E1AE64CE876B4D88AF06C1FA4CD7A6AAC47A7B0B2CD573BC83D068
dockernet-relayer-photo-1  | ts=2023-07-06T22:01:02.114271Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=179892 fees=20012uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=835 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=58DC4A7642A53C12477CAD39E00E8A4F140AAAEF2E1A4DE95B19A80A4D37DB39
dockernet-relayer-photo-1  | ts=2023-07-06T22:01:04.424801Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=135818 fees=1569ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=541 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=DF5CB9FF6431EE96EFB843FDE9C00E153E76C47FD734DD14E8AC2EAD937C157F
dockernet-relayer-photo-1  | ts=2023-07-06T22:01:06.114507Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=263644 fees=30062uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=839 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=7F6FD7611BE4A270768D752BFC4B357280E3FDB4F60721113D00523A51CA4F5D
dockernet-relayer-photo-1  | ts=2023-07-06T22:01:09.104492Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=168056 fees=1988ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=544 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=3DB47B29D9339F3212425675B3276A206C99F4A907454238CDBA25F5CC51F55C
dockernet-relayer-photo-1  | ts=2023-07-06T22:01:24.245251Z lvl=warn msg="Error querying block data" chain_name=photo chain_id=localnet error="RPC error -32603 - Internal error: could not find results for height #857"
dockernet-relayer-photo-1  | ts=2023-07-06T22:01:41.732306Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=109942 fees=1233ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=565 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=324B6323304F7E3DF3D3503570D927EE1095B6730DFC9173082BBF33C186CB48
dockernet-relayer-photo-1  | ts=2023-07-06T22:01:42.533165Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=179845 fees=20006uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=875 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=8EE159CB28AD96CB058EA2E7ACF3C1A0653E267BE3BD655093923B74BAA54253
dockernet-relayer-photo-1  | ts=2023-07-06T22:01:46.337561Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=135815 fees=1569ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=568 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=C40DD879E5645AD8C9755FECCEC6BEFAC9CFCEE0CF802C55B4B71C18F9EC89B7
dockernet-relayer-photo-1  | ts=2023-07-06T22:02:22.167712Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=107096 fees=1233ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=591 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=9A7E353D34C3F854FDF8BC02F66F6D0C3D271F56201F751267294E3179DCE1CA
dockernet-relayer-photo-1  | ts=2023-07-06T22:02:23.090615Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=179128 fees=19920uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=915 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=78A06250C59F9A225CAB86FC189E4FD2D26E7B24EA6331D92BCE92422FC172A8
dockernet-relayer-photo-1  | ts=2023-07-06T22:02:25.570205Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=136258 fees=1575ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=593 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=2C4A94336B397A700909DE7C330447E1D6E1D04D7B395CC0DF44186C2C764ACF
dockernet-relayer-photo-1  | ts=2023-07-06T22:02:45.290987Z lvl=warn msg="Error querying block data" chain_name=photo chain_id=localnet error="RPC error -32603 - Internal error: could not find results for height #937"
dockernet-relayer-photo-1  | ts=2023-07-06T22:03:01.180371Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=107096 fees=1233ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=616 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=3D0F37721E3BE2DA69FAEE8A118C6C2BFF9C1558F4CACEF03DC01D7E71080750
dockernet-relayer-photo-1  | ts=2023-07-06T22:03:01.682643Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=179768 fees=19997uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=953 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=0104A63E2A1F35133262D2CA844678008D90A5203DBA4AD0CFD90B193D5A9E3D
dockernet-relayer-photo-1  | ts=2023-07-06T22:03:04.178812Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=136258 fees=1575ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=618 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=AC7F861C29558B965399A0711BBA1EEB736D39F67C31EA32D703DC6D4B4E7130
dockernet-relayer-photo-1  | ts=2023-07-06T22:03:05.881087Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-0 packet_dst_channel=channel-1 gas_used=164388 fees=1978ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=619 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=CA74E672EB49449AA8DBCA4825D6F4E361A170ECA91991DA5229AB659FA5562C
dockernet-relayer-photo-1  | ts=2023-07-06T22:03:08.675021Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-0 packet_dst_channel=channel-1 gas_used=152584 fees=16735uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=960 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=D8587885E0E96B653E1A5007A09A98A62668BD5D6D3E435E8B4324899706F8F4
dockernet-relayer-photo-1  | ts=2023-07-06T22:03:41.482002Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=109944 fees=1233ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=642 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=280ACDDA91C80D6BB20673FD1D4B5A4D4416C93324F7E2067570ECACA7FAE9ED
dockernet-relayer-photo-1  | ts=2023-07-06T22:03:43.104313Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=231080 fees=26155uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=994 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket,/ibc.core.channel.v1.MsgRecvPacket tx_hash=63601F3E03C3EDBD00AFED44C4EA4663E2360F9A6010DC7FC1749095982B82AE
dockernet-relayer-photo-1  | ts=2023-07-06T22:03:46.287143Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=171133 fees=2065ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=645 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=9D86635DE899C7ED325533627FDC43EE43DB74EF8C5EE91367FDB89972A7BAD9
dockernet-relayer-photo-1  | ts=2023-07-06T22:04:05.321507Z lvl=warn msg="Error querying block data" chain_name=photo chain_id=localnet error="RPC error -32603 - Internal error: could not find results for height #1016"
dockernet-relayer-photo-1  | ts=2023-07-06T22:04:20.703968Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=109276 fees=1261ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=667 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=AC56D09CDCE65036D538475D9324CEC51FEC347D638DCE5A644724F768ABAF5D
dockernet-relayer-photo-1  | ts=2023-07-06T22:04:21.602172Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=179818 fees=20003uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1032 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=E441E0A3689E1617150083D8F7ECD8DA2181D264617032099E6C9B5F8F9D234F
dockernet-relayer-photo-1  | ts=2023-07-06T22:04:25.507126Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=136402 fees=1577ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=670 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=7361FEE954CFFCFDF5DB5B075DF35AF8338E478336930EA57613DAA4E686ACF2
dockernet-relayer-photo-1  | ts=2023-07-06T22:04:26.622350Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=276704 fees=31630uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1037 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=725D3E1E45F9E77691AD9C648E528CFEA1C96C91361BB20B5421DCB8415ED4E3
dockernet-relayer-photo-1  | ts=2023-07-06T22:04:30.208346Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=168279 fees=1991ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=673 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=4CC379F6BD40223EA78D375CCA50A23BBCBCAE7F670B7D0146362945C66219A6
dockernet-relayer-photo-1  | ts=2023-07-06T22:05:01.322071Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-8 packet_dst_channel=channel-2 gas_used=170547 fees=2058ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=693 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=9BF9C7D909A361B3F3CF36ADA5ED3380D0A4C907B74D67808D92FFD3A6A2F79F
dockernet-relayer-photo-1  | ts=2023-07-06T22:05:02.030656Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=179581 fees=19975uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1072 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=AC026680DB4788743D1906AD90F21C7346F3F061A245B8FAEF38100DE0210CB9
dockernet-relayer-photo-1  | ts=2023-07-06T22:05:04.027368Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-8 packet_dst_channel=channel-2 gas_used=185552 fees=20691uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1074 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=1A1F02E6E7B7E3617D27923DE7C68949C1DC76CDF0050012B1C31EA2E2BAA312
dockernet-relayer-photo-1  | ts=2023-07-06T22:05:04.538295Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=135982 fees=1571ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=695 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=36B77D742993839629B1CD7566880FB138EC0EFB3619F1F0BEFEAB751AF275A6
dockernet-relayer-photo-1  | ts=2023-07-06T22:05:06.019103Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-8 packet_dst_channel=channel-2 gas_used=161308 fees=1937ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=696 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=B315F3CE7C9BC5F64E4438B925F7E6437AF56BF47F1C2C9985AD35EF9CF22FB0
dockernet-relayer-photo-1  | ts=2023-07-06T22:05:07.626071Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=112114 fees=1261ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=697 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=B990BC7299790D4AA58664E38794B9B8E0C3D171BBAF8386FDD6F00112FE2BDE
dockernet-relayer-photo-1  | ts=2023-07-06T22:05:42.244553Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=112135 fees=1261ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=719 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=CF55A480A91CFE7956753CE76A93E195C63ED8B8DA3D97D15B9BC14F6383C42F
dockernet-relayer-photo-1  | ts=2023-07-06T22:05:43.466080Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=179496 fees=19965uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1113 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=3E6BE141538794F02411ECBBC5FF6F468ACF7CAF7789528682BBFB64CAC15067
dockernet-relayer-photo-1  | ts=2023-07-06T22:05:46.946919Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=133561 fees=1577ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=722 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=4C436731B3FCD2DCCE16AA43D0D82DBB353870019770480F9830E7DE9ED402B6
dockernet-relayer-photo-1  | ts=2023-07-06T22:05:48.566369Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=275989 fees=31544uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1118 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=59946A010CA8967A34071B06BC8C26F550D1A27104B391F1E0EA4E13BAE4BE5E
dockernet-relayer-photo-1  | ts=2023-07-06T22:05:51.549846Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=164719 fees=1945ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=725 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=E86631A8D22FFE136D10C66BE97CDFBC99A14A77A73AC2FFBB58052932704C47
dockernet-relayer-photo-1  | ts=2023-07-06T22:06:05.654549Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-0 packet_dst_channel=channel-1 gas_used=167396 fees=1979ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=734 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=19190FE4BD31687CC89138479FB9CF55833994059CFC582F0940AC006037356D
dockernet-relayer-photo-1  | ts=2023-07-06T22:06:07.754401Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-0 packet_dst_channel=channel-1 gas_used=151930 fees=16657uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1137 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=7629CC1D0A9F967E81DB26634945F94BEDF8B73D04B1B90D665B82AEA41C8390
dockernet-relayer-photo-1  | ts=2023-07-06T22:06:21.268977Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-8 packet_dst_channel=channel-2 gas_used=170553 fees=2058ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=744 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=CA8BC10508956325E1E99D2AF04F539A051F51D4A8E9D86F959F148BB99FF52E
dockernet-relayer-photo-1  | ts=2023-07-06T22:06:21.882241Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=231426 fees=26196uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1151 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket,/ibc.core.channel.v1.MsgRecvPacket tx_hash=3C9223FC38C9F6F61A5C33C7A9F9BF7097392F5C9B3CD17AD96FC0EBFC7B52A0
dockernet-relayer-photo-1  | ts=2023-07-06T22:06:23.965697Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-8 packet_dst_channel=channel-2 gas_used=186108 fees=20758uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1153 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=DCE18ECD16F3B7AD1F57AC9EA30EE3F4274110F306AD7F9AC21AD2F99AE3D238
dockernet-relayer-photo-1  | ts=2023-07-06T22:06:24.682778Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-1 packet_dst_channel=channel-0 gas_used=174232 fees=2068ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=746 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=8C5E33916B116A00FF0D6B50B5B6FE92A9AB4FA5E0475043BAAC008C0FAEA49A
dockernet-relayer-photo-1  | ts=2023-07-06T22:06:26.159439Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-8 packet_dst_channel=channel-2 gas_used=157754 fees=1891ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=747 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=930929BB44A452E5AA3B13391256E856E7BF6AA837FA3EEC287B3EBFD1F7AAD5
dockernet-relayer-photo-1  | ts=2023-07-06T22:06:27.768914Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=108150 fees=1209ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=748 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=A5959564177AB7F07E28D6F758531518C472BFDC6783F4E9B6B95BA450C687CB
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:02.096974Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=112118 fees=1261ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=770 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=ACEBECC47CA784D83F9A169DCDDB4DBBBB487E146EEB28B39F3F4DE7C20AE316
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:03.396321Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=179936 fees=20017uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1192 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=E5AC4D6DD7BBE803AE727917E457BB239EC1B00F65AED7BE8F9F9C6A4B60394D
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:06.797270Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=136419 fees=1577ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=773 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=E9836809929BF5AC9F1D1C0942D68806C9DCD3E97F40369DBC40442C74D1F18A
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:08.393014Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=278482 fees=31843uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1197 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=CACB593855E0F46F79C1BA6363A2FEB7B7A34C7081BF3AF9F1461376F246152B
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:11.588123Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=168725 fees=1997ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=776 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=553FF68F72825376A0E32D15ABC19E56636093D7366B9EADC0B8EB617A3D71CB
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:13.422039Z lvl=warn msg="Error querying block data" chain_name=photo chain_id=localnet error="RPC error -32603 - Internal error: could not find results for height #1202"
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:13.504930Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=259337 fees=29545uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1202 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=68B9E1BA461853E1B153514DEE6C9DDFB40F3688AFD7298BC634CCB8FFC598C8
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:16.301400Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=168729 fees=1997ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=779 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=7B487740C74679BBF0203BD295CECD23520E8A082A72E9B6294D11E3590A6A28
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:41.406970Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-8 packet_dst_channel=channel-2 gas_used=170553 fees=2058ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=795 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=DC74AAA9C20755102E64F072A51EAA4A0931165F1D3099DD11D8211C02DED89B
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:41.817164Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=179218 fees=19931uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1230 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=2BBEA31FCEFB9D3C738BFCCC388AEE8EBA446FA818C7C21E34307A9A818B26FD
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:43.806740Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-8 packet_dst_channel=channel-2 gas_used=185992 fees=20744uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1232 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=6539A3D3EDB501EB61A1E4F8283B5DFF85BF2FD99BA031F92747CB716B6C06D5
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:44.519464Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=136418 fees=1577ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=797 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=76746238A79D9E1C2E8D47BCE77497750594A8FD351E6699AF3185C96351EB81
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:47.721521Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-8 packet_dst_channel=channel-2 gas_used=164162 fees=1937ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=799 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=F953D155C3DCDAD65AA5C7FA174A1A289390656DB5452FEED75F8C6F7FF532E7
dockernet-relayer-photo-1  | ts=2023-07-06T22:07:49.503669Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=112130 fees=1261ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=800 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=2C3E01BD9E5EC198431F3EC29202653E9FCAE7F7D5B13D93C511AA500361ABC0
dockernet-relayer-photo-1  | ts=2023-07-06T22:08:20.627299Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE gas_used=109240 fees=1261ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=820 msg_types=/ibc.core.client.v1.MsgUpdateClient,/stride.interchainquery.v1.MsgSubmitQueryResponse tx_hash=572CD1B83A95DE9E21EC305B9FD26FAAE5278261EC4197F5F39DEF687511A495
dockernet-relayer-photo-1  | ts=2023-07-06T22:08:22.230404Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=localnet packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=179892 fees=20012uarch fee_payer=archway1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cdl9u0v height=1270 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgRecvPacket tx_hash=31EC2AD3DD26F7234573A040C46A71CA941623F1FCD8B41A7865FC4A3EE0E0CA
dockernet-relayer-photo-1  | ts=2023-07-06T22:08:25.427595Z lvl=info msg="Successful transaction" provider_type=cosmos chain_id=STRIDE packet_src_channel=channel-6 packet_dst_channel=channel-6 gas_used=136353 fees=1576ustrd fee_payer=stride1jx7ns7v9s3nes9k7p0sez7nq2hvmae5cmley3h height=823 msg_types=/ibc.core.client.v1.MsgUpdateClient,/ibc.core.channel.v1.MsgAcknowledgement tx_hash=54FB496CB5E2F9D575CB4D9D6D1F64218F8983BCDD417B1D27DCAE5527D7B6F3

```


Hostzone Parameters


```

host_zone:
- address: stride1755g4dkhpw73gz9h9nwhlcefc6sdf8kcmvcwrk4rxfrz8xpxxjms7savm8
  bech32prefix: cosmos
  blacklisted_validators: []
  chain_id: GAIA
  connection_id: connection-0
  delegation_account:
    address: cosmos1d48lulv6cad8hhgezvfeglhqds5hf50k295n3dhqvjwdhjjeq9asdalwt5
    target: DELEGATION
  fee_account:
    address: cosmos1sf5a2zl7gr6llllvvywfrq8ec2valgxrjqjjd0qxpj5qr6zexstqqmmur9
    target: FEE
  halted: false
  host_denom: uatom
  ibc_denom: ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2
  last_redemption_rate: "1.000000000000000000"
  max_redemption_rate: "1.500000000000000000"
  min_redemption_rate: "0.900000000000000000"
  redemption_account:
    address: cosmos1tyh4vkp267rtfhfm8xfqvlj2n0fq4xceelr7av2gn9kzs9dpktzqtw3hxm
    target: REDEMPTION
  redemption_rate: "1.000000000000000000"
  staked_bal: "0"
  transfer_channel_id: channel-0
  unbonding_frequency: "1"
  validators:
  - address: cosmosvaloper1uk4ze0x4nvh4fk0xm4jdud58eqn4yxhrdt795pcosmosvaloper1uk4ze0x4nvh4fk0xm4jdud58eqn4yxhrdt795p
    delegation_amt: "0"
    internal_exchange_rate: null
    name: gval1
    weight: "5"
  withdrawal_account:
    address: cosmos1dndts20gfszqhlz2vyal903yd57wrg6jx8xwwntyck2enp3e49dsalcjnq
    target: WITHDRAWAL
- address: stride19467hx6r0qkj5crjff3yr38uzts5hwj7detdw4tr0qdsc5rufelsm88tag
  bech32prefix: archway
  blacklisted_validators: []
  chain_id: localnet
  connection_id: connection-1
  delegation_account:
    address: archway10kgq25m2v2738hxh8ene03cu4n87ruhrh823x2rmuwpa4grhs2zq4gt55n
    target: DELEGATION
  fee_account:
    address: archway1zadujthv3mpwu9wj7l7cmpwk34tcwu7mphzqpyxeyemntsk99fksf3vlx4
    target: FEE
  halted: false
  host_denom: uarch
  ibc_denom: ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
  last_redemption_rate: "1.162616339193381593"
  max_redemption_rate: "1.500000000000000000"
  min_redemption_rate: "0.900000000000000000"
  redemption_account:
    address: archway1d8ellefkcjsvxqwf5ahny0dxwfa0lxl7zkcwcapuju4n4ahspqtqtv3jg4
    target: REDEMPTION
  redemption_rate: "1.162616339193381593"
  staked_bal: "2393"
  transfer_channel_id: channel-1
  unbonding_frequency: "1"
  validators:
  - address: archwayvaloper15js809uedxqs2wl0lyt58httasr5rlplj45fqw
    delegation_amt: "2393"
    internal_exchange_rate: null
    name: pval1
    weight: "5"
  withdrawal_account:
    address: archway1x73qfxvs0dakrn7z9mj3fhpmdfc8xnnxs86cg6zavkgfs8mdwaasgk0hrf
    target: WITHDRAWAL
pagination:
  next_key: null
  total: "0"

```
