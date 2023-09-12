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

### Full Liquid staking - uarch and Full Redemption Workflow

root@swordfish-Lenovo-Y720-15IKB:/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts#
build/archwayd --home
/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/photo1
tx ibc-transfer transfer transfer channel-0
stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 4000000uarch --from pval1 -y

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

\#build/strided --home
/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1
q bank balances stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8

    balances:
    - amount: "4000000"
      denom: ibc/EAEB74E11A7BFFC17E79B802EF01D3307A285F7AA802745037B021609ECFB034
    - amount: "1000000000"
      denom: ustrd
    pagination:
      next_key: null
      total: "0"

\#build/strided --home
/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1
tx stakeibc liquid-stake 1000000 uarch --keyring-backend test --from admin
\--chain-id STRIDE -y

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

\#root@swordfish-Lenovo-Y720-15IKB:/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts#
build/strided --home
/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1
q tx 62E3A222D273428D576FF3E4D201217B34C6D77552F228DA1AE51BC754294361

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

root@swordfish-Lenovo-Y720-15IKB:/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts#
build/strided --home
/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1
q stakeibc list-host-zone

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

root@swordfish-Lenovo-Y720-15IKB:/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts#
build/strided --home
/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1
tx stakeibc redeem-stake 1000 PHOTO
archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m --from admin --keyring-backend
test --chain-id STRIDE -y

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

root@swordfish-Lenovo-Y720-15IKB:/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts#
build/strided --home
/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1
tx stakeibc redeem-stake 1000 localnet
archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m --from admin --keyring-backend
test --chain-id STRIDE -y

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

root@swordfish-Lenovo-Y720-15IKB:/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts#
build/strided --home
/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1
q tx 1A6F1DBA73E3A4A469782B18C2637E5A118E06906A028A5A0A25734DFEDEC462

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
          receiver: archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m
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
