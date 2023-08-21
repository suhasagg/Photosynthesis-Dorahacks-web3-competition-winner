#!/bin/bash

set -eu 
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

source $SCRIPT_DIR/../config.sh

CHAIN="$1"
KEYS_LOGS=$DOCKERNET_HOME/logs/keys.log

CHAIN_ID=$(GET_VAR_VALUE    ${CHAIN}_CHAIN_ID)
BINARY=$(GET_VAR_VALUE      ${CHAIN}_BINARY)
MAIN_CMD=$(GET_VAR_VALUE    ${CHAIN}_MAIN_CMD)
DENOM=$(GET_VAR_VALUE       ${CHAIN}_DENOM)
RPC_PORT=$(GET_VAR_VALUE    ${CHAIN}_RPC_PORT)
NUM_NODES=$(GET_VAR_VALUE   ${CHAIN}_NUM_NODES)
NODE_PREFIX=$(GET_VAR_VALUE ${CHAIN}_NODE_PREFIX)
VAL_PREFIX=$(GET_VAR_VALUE  ${CHAIN}_VAL_PREFIX)

# THe host zone can optionally specify additional the micro-denom granularity
# If they don't specify the ${CHAIN}_MICRO_DENOM_UNITS variable,
# EXTRA_MICRO_DENOM_UNITS will include 6 0's
MICRO_DENOM_UNITS_VAR_NAME=${CHAIN}_MICRO_DENOM_UNITS
MICRO_DENOM_UNITS="${!MICRO_DENOM_UNITS_VAR_NAME:-000000}"

VAL_TOKENS=${VAL_TOKENS}${MICRO_DENOM_UNITS}
STAKE_TOKENS=${STAKE_TOKENS}${MICRO_DENOM_UNITS}
ADMIN_TOKENS=${ADMIN_TOKENS}${MICRO_DENOM_UNITS}

set_stride_genesis() {
    genesis_config=$1

    # update params
    jq '(.app_state.epochs.epochs[] | select(.identifier=="day") ).duration = $epochLen' --arg epochLen $STRIDE_DAY_EPOCH_DURATION $genesis_config > json.tmp && mv json.tmp $genesis_config
    jq '(.app_state.epochs.epochs[] | select(.identifier=="hour") ).duration = $epochLen' --arg epochLen $STRIDE_HOUR_EPOCH_DURATION $genesis_config > json.tmp && mv json.tmp $genesis_config
    jq '(.app_state.epochs.epochs[] | select(.identifier=="stride_epoch") ).duration = $epochLen' --arg epochLen $STRIDE_EPOCH_EPOCH_DURATION $genesis_config > json.tmp && mv json.tmp $genesis_config
    jq '(.app_state.epochs.epochs[] | select(.identifier=="mint") ).duration = $epochLen' --arg epochLen $STRIDE_MINT_EPOCH_DURATION $genesis_config > json.tmp && mv json.tmp $genesis_config
    jq '.app_state.staking.params.unbonding_time = $newVal' --arg newVal "$UNBONDING_TIME" $genesis_config > json.tmp && mv json.tmp $genesis_config
    jq '.app_state.gov.deposit_params.max_deposit_period = $newVal' --arg newVal "$MAX_DEPOSIT_PERIOD" $genesis_config > json.tmp && mv json.tmp $genesis_config
    jq '.app_state.gov.voting_params.voting_period = $newVal' --arg newVal "$VOTING_PERIOD" $genesis_config > json.tmp && mv json.tmp $genesis_config

    # enable stride as an interchain accounts controller
    jq "del(.app_state.interchain_accounts)" $genesis_config > json.tmp && mv json.tmp $genesis_config
    interchain_accts=$(cat $DOCKERNET_HOME/config/ica_controller.json)
    jq ".app_state += $interchain_accts" $genesis_config > json.tmp && mv json.tmp $genesis_config
}

set_host_genesis() {
    genesis_config=$1

    # Shorten epochs and unbonding time
    jq '(.app_state.epochs.epochs[]? | select(.identifier=="day") ).duration = $epochLen' --arg epochLen $HOST_DAY_EPOCH_DURATION $genesis_config > json.tmp && mv json.tmp $genesis_config
    jq '(.app_state.epochs.epochs[]? | select(.identifier=="hour") ).duration = $epochLen' --arg epochLen $HOST_HOUR_EPOCH_DURATION $genesis_config > json.tmp && mv json.tmp $genesis_config
    jq '(.app_state.epochs.epochs[]? | select(.identifier=="week") ).duration = $epochLen' --arg epochLen $HOST_WEEK_EPOCH_DURATION $genesis_config > json.tmp && mv json.tmp $genesis_config
    jq '(.app_state.epochs.epochs[]? | select(.identifier=="mint") ).duration = $epochLen' --arg epochLen $HOST_MINT_EPOCH_DURATION $genesis_config > json.tmp && mv json.tmp $genesis_config
    jq '.app_state.staking.params.unbonding_time = $newVal' --arg newVal "$UNBONDING_TIME" $genesis_config > json.tmp && mv json.tmp $genesis_config

    # Set the mint start time to the genesis time if the chain configures inflation at the block level (e.g. stars)
    # also reduce the number of initial annual provisions so the inflation rate is not too high
    genesis_time=$(jq .genesis_time $genesis_config | tr -d '"')
    jq 'if .app_state.mint.params.start_time? then .app_state.mint.params.start_time=$newVal else . end' --arg newVal "$genesis_time" $genesis_config > json.tmp && mv json.tmp $genesis_config
    jq 'if .app_state.mint.params.initial_annual_provisions? then .app_state.mint.params.initial_annual_provisions=$newVal else . end' --arg newVal "$INITIAL_ANNUAL_PROVISIONS" $genesis_config > json.tmp && mv json.tmp $genesis_config

    # Add interchain accounts to the genesis set
    jq "del(.app_state.interchain_accounts)" $genesis_config > json.tmp && mv json.tmp $genesis_config
    interchain_accts=$(cat $DOCKERNET_HOME/config/ica_host.json)
    jq ".app_state += $interchain_accts" $genesis_config > json.tmp && mv json.tmp $genesis_config

    # Slightly harshen slashing parameters (if 5 blocks are missed, the validator will be slashed)
    # This makes it easier to test updating weights after a host zone validator is slashed
    sed -i -E 's|"signed_blocks_window": "100"|"signed_blocks_window": "10"|g' $genesis_config
    sed -i -E 's|"downtime_jail_duration": "600s"|"downtime_jail_duration": "10s"|g' $genesis_config
    sed -i -E 's|"slash_fraction_downtime": "0.010000000000000000"|"slash_fraction_downtime": "0.050000000000000000"|g' $genesis_config
}

MAIN_ID=1 # Node responsible for genesis and persistent_peers
MAIN_NODE_NAME=""
MAIN_NODE_ID=""
MAIN_CONFIG=""
MAIN_GENESIS=""
echo "Initializing $CHAIN chain..."
for (( i=1; i <= $NUM_NODES; i++ )); do
    # Node names will be of the form: "stride1"
    node_name="${NODE_PREFIX}${i}"
    # Moniker is of the form: STRIDE_1
    moniker=$(printf "${NODE_PREFIX}_${i}" | awk '{ print toupper($0) }')

    # Create a state directory for the current node and initialize the chain
    mkdir -p $STATE/$node_name
    
    # If the chains commands are run only from docker, grab the command from the config
    # Otherwise, if they're run locally, append the home directory
    if [[ $BINARY == docker-compose* ]]; then
        cmd=$BINARY
    else
        cmd="$BINARY --home ${STATE}/$node_name"
    fi

# Condition to check if chain id is 'localnet'
    if [ "$CHAIN_ID" == "localnet" ]; then
       moniker="archway-id"
    fi

    # Initialize the chain
    $cmd init $moniker --chain-id $CHAIN_ID --overwrite &> /dev/null
    chmod -R 777 $STATE/$node_name

    # Update node networking configuration 
    config_toml="${STATE}/${node_name}/config/config.toml"
    client_toml="${STATE}/${node_name}/config/client.toml"
    app_toml="${STATE}/${node_name}/config/app.toml"
    genesis_json="${STATE}/${node_name}/config/genesis.json"

    sed -i -E "s|cors_allowed_origins = \[\]|cors_allowed_origins = [\"\*\"]|g" $config_toml
    sed -i -E "s|127.0.0.1|0.0.0.0|g" $config_toml
    sed -i -E "s|timeout_commit = \"5s\"|timeout_commit = \"${BLOCK_TIME}\"|g" $config_toml
    sed -i -E "s|prometheus = false|prometheus = true|g" $config_toml

    sed -i -E "s|minimum-gas-prices = \".*\"|minimum-gas-prices = \"0${DENOM}\"|g" $app_toml
    sed -i -E '/\[api\]/,/^enable = .*$/ s/^enable = .*$/enable = true/' $app_toml
    sed -i -E 's|unsafe-cors = .*|unsafe-cors = true|g' $app_toml
    sed -i -E "s|snapshot-interval = 0|snapshot-interval = 300|g" $app_toml

    sed -i -E "s|chain-id = \"\"|chain-id = \"${CHAIN_ID}\"|g" $client_toml
    sed -i -E "s|keyring-backend = \"os\"|keyring-backend = \"test\"|g" $client_toml
    sed -i -E "s|node = \".*\"|node = \"tcp://localhost:$RPC_PORT\"|g" $client_toml

    sed -i -E "s|\"stake\"|\"${DENOM}\"|g" $genesis_json 
    sed -i -E "s|\"aphoton\"|\"${DENOM}\"|g" $genesis_json # ethermint default

    # Get the endpoint and node ID
    node_id=$($cmd tendermint show-node-id)@$node_name:$PEER_PORT
    echo "Node #$i ID: $node_id"

 # add a validator account
    val_acct="${VAL_PREFIX}${i}"
    val_acct2="pval2"
    val_acct3="pval3"
    val_acct4="pval4"
    val_acct5="pval5"
    echo $val_acct
    if [[ $CHAIN_ID == "localnet" ]]; then
        val_mnemonic="lucky alien buyer stumble genius scatter fantasy parent amazing february aisle cargo welcome frog breeze mesh half oak enter endless paper fury gas learn"
    else
         val_mnemonic="${VAL_MNEMONICS[((i-1))]}"
    fi
    echo $val_mnemonic
    echo "$val_mnemonic" | $cmd keys add $val_acct --recover --keyring-backend=test
    val_addr=$($cmd keys show $val_acct --keyring-backend test -a | tr -cd '[:alnum:]._-')
    echo $val_addr
    # Add this account to the current node
    $cmd add-genesis-account ${val_addr} ${VAL_TOKENS}${DENOM}
    echo $val_mnemonic
    if [[ "$CHAIN_ID" == "localnet" ]]; then
        val_mnemonic1="axis spirit huge divert unlock awful melody ill spread little chief youth custom access evidence iron scan trust series champion absorb stable mistake artist"
        echo "$val_mnemonic1" | $cmd keys add $val_acct2 --recover --keyring-backend=test
        val_addr2=$($cmd keys show $val_acct2 --keyring-backend test -a | tr -cd '[:alnum:]._-')
        echo $val_addr2
        $cmd add-genesis-account ${val_addr2} ${VAL_TOKENS}${DENOM}
    fi
   # Add this account to the current node

    if [[ "$CHAIN_ID" == "localnet" ]]; then
        val_mnemonic2="joke drastic chalk vacant hood flock corn appear neither gate garlic bacon credit dice leader away ship fat situate screen task address bike bind"
        echo "$val_mnemonic2" | $cmd keys add $val_acct3 --recover --keyring-backend=test
        val_addr3=$($cmd keys show $val_acct3 --keyring-backend test -a | tr -cd '[:alnum:]._-')
        echo $val_addr3
       # Add this account to the current node
        $cmd add-genesis-account ${val_addr3} ${VAL_TOKENS}${DENOM}
    fi
    # actually set this account as a validator on the current node
    if [[ "$CHAIN_ID" == "localnet" ]]; then
               val_mnemonic3="culture matter ivory forward stuff drama current face soft sell noodle render gasp ski sick common physical jazz climb confirm festival bomb world pluck"
               echo "$val_mnemonic3" | $cmd keys add $val_acct4 --recover --keyring-backend=test
               val_addr4=$($cmd keys show $val_acct4 --keyring-backend test -a | tr -cd '[:alnum:]._-')
               echo $val_addr4
              # Add this account to the current node
               $cmd add-genesis-account ${val_addr4} 10000000000${DENOM}
    fi
    if [[ "$CHAIN_ID" == "localnet" ]]; then
              val_mnemonic4="tape foster admit swing picture quarter release thought pluck multiply grab humble radio wool minimum comic document crush borrow spike cactus gym pupil husband"
              echo "$val_mnemonic4" | $cmd keys add $val_acct5 --recover --keyring-backend=test
              val_addr5=$($cmd keys show $val_acct5 --keyring-backend test -a | tr -cd '[:alnum:]._-')
              echo $val_addr5
                  # Add this account to the current node
              $cmd add-genesis-account ${val_addr5} 10000000000${DENOM}
    fi

    $cmd gentx $val_acct ${STAKE_TOKENS}${DENOM} --chain-id $CHAIN_ID --keyring-backend test

    #if [[ "$CHAIN_ID" == "localnet" ]]; then
    #    $cmd gentx $val_acct2 ${STAKE_TOKENS}${DENOM} --chain-id $CHAIN_ID --keyring-backend test
    #    $cmd gentx $val_acct3 ${STAKE_TOKENS}${DENOM} --chain-id $CHAIN_ID --keyring-backend test
    #    echo $cmd
    #fi

    # Cleanup from seds
    rm -rf ${client_toml}-E
    rm -rf ${genesis_json}-E
    rm -rf ${app_toml}-E

    if [ $i -eq $MAIN_ID ]; then
        MAIN_NODE_NAME=$node_name
        MAIN_NODE_ID=$node_id
        MAIN_CONFIG=$config_toml
        echo $MAIN_CONFIG
        MAIN_GENESIS=$genesis_json
        echo $MAIN_GENESIS
    else
        # also add this account and it's genesis tx to the main node
        $MAIN_CMD add-genesis-account ${val_addr} ${VAL_TOKENS}${DENOM}
        cp ${STATE}/${node_name}/config/gentx/*.json ${STATE}/${MAIN_NODE_NAME}/config/gentx/

        # and add each validator's keys to the first state directory
        echo "$val_mnemonic" | $MAIN_CMD keys add $val_acct --recover --keyring-backend=test &> /dev/null
    fi
done

if [ "$CHAIN" == "STRIDE" ]; then 
    # add the stride admin account
    echo "$STRIDE_ADMIN_MNEMONIC" | $MAIN_CMD keys add $STRIDE_ADMIN_ACCT --recover --keyring-backend=test >> $KEYS_LOGS 2>&1
    STRIDE_ADMIN_ADDRESS=$($MAIN_CMD keys show $STRIDE_ADMIN_ACCT --keyring-backend test -a)
    $MAIN_CMD add-genesis-account ${STRIDE_ADMIN_ADDRESS} ${ADMIN_TOKENS}${DENOM}
    # add relayer accounts
    for i in "${!RELAYER_ACCTS[@]}"; do
        RELAYER_ACCT="${RELAYER_ACCTS[i]}"
        RELAYER_MNEMONIC="${RELAYER_MNEMONICS[i]}"

        echo "$RELAYER_MNEMONIC" | $MAIN_CMD keys add $RELAYER_ACCT --recover --keyring-backend=test >> $KEYS_LOGS 2>&1
        RELAYER_ADDRESS=$($MAIN_CMD keys show $RELAYER_ACCT --keyring-backend test -a)
        $MAIN_CMD add-genesis-account ${RELAYER_ADDRESS} ${VAL_TOKENS}${DENOM}
    done
else 
    # add a revenue account
    REV_ACCT_VAR=${CHAIN}_REV_ACCT
    REV_ACCT=${!REV_ACCT_VAR}
    echo $REV_MNEMONIC | $MAIN_CMD keys add $REV_ACCT --recover --keyring-backend=test >> $KEYS_LOGS 2>&1

    # add a relayer account
    RELAYER_ACCT=$(GET_VAR_VALUE     RELAYER_${CHAIN}_ACCT)
    RELAYER_MNEMONIC=$(GET_VAR_VALUE RELAYER_${CHAIN}_MNEMONIC)

    echo "$RELAYER_MNEMONIC" | $MAIN_CMD keys add $RELAYER_ACCT --recover --keyring-backend=test >> $KEYS_LOGS 2>&1
    RELAYER_ADDRESS=$($MAIN_CMD keys show $RELAYER_ACCT --keyring-backend test -a | tr -cd '[:alnum:]._-')
    $MAIN_CMD add-genesis-account ${RELAYER_ADDRESS} ${VAL_TOKENS}${DENOM}
fi

# now we process gentx txs on the main node
$MAIN_CMD collect-gentxs &> /dev/null

# wipe out the persistent peers for the main node (these are incorrectly autogenerated for each validator during collect-gentxs)
sed -i -E "s|persistent_peers = .*|persistent_peers = \"\"|g" $MAIN_CONFIG

# update chain-specific genesis settings
if [ "$CHAIN" == "STRIDE" ]; then 
    set_stride_genesis $MAIN_GENESIS
else
    set_host_genesis $MAIN_GENESIS
fi

# for all peer nodes....
for (( i=2; i <= $NUM_NODES; i++ )); do
    node_name="${NODE_PREFIX}${i}"
    config_toml="${STATE}/${node_name}/config/config.toml"
    genesis_json="${STATE}/${node_name}/config/genesis.json"

    # add the main node as a persistent peer
    sed -i -E "s|persistent_peers = .*|persistent_peers = \"${MAIN_NODE_ID}\"|g" $config_toml
    # copy the main node's genesis to the peer nodes to ensure they all have the same genesis
    cp $MAIN_GENESIS $genesis_json

    rm -rf ${config_toml}-E
done

# Cleanup from seds
rm -rf ${MAIN_CONFIG}-E
rm -rf ${MAIN_GENESIS}-E