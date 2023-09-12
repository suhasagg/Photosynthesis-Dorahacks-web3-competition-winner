# Photosynthesis-Archway <-> Stride integrations

![Photosynthesis](https://i.imgur.com/Tn1GUZnb.jpg)

1\)START RELAYER

2\)CREATING CONNECTIONS WITH THE GO RELAYER

3\)Create connections and channels

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

i)Update delegations (just submit this query and confirm the ICQ callback
displays in the stride logs)

# Must be submitted in ICQ window

j)Restore closed channel

# Photosynthesis-Archway IBC integrations.

1\)evmos

2\)gaia

3\)juno

4\)osmosis

5\)stargaze

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

# Relayer registry

Cosmos relayer 1)https://github.com/cosmos/relayer.git

Hermes relayer 2)https://github.com/informalsystems/ibc-rs

# Record Keeping/Queue Creation for Liquid Staking Workflow in Photosynthesis-Archway

# Interchain Accounts fully Integrated in Photosynthesis-Archway

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
