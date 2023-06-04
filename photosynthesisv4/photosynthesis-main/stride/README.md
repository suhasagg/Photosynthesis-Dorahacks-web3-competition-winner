# Photosynthesis-Archway  <-> Stride integrations


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
