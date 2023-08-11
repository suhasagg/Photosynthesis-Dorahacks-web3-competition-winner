import json
import re

log_lines = """dockernet-photo1-1  | 12:13PM INF Processing ARCH_CENTRAL_LIQUID_STAKE_INTERVAL_EPOCH: {arch-central-liquid-stake-interval-epoch 2023-07-28 12:04:08.993680513 +0000 UTC 3m20s 3 2023-07-28 12:10:48.993680513 +0000 UTC true 294} module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Retrieved EpochInfo for epochstypes.ARCH_CENTRAL_LIQUID_STAKE_INTERVAL_EPOCH: {arch-central-liquid-stake-interval-epoch 2023-07-28 12:04:08.993680513 +0000 UTC 3m20s 3 2023-07-28 12:10:48.993680513 +0000 UTC true 294} module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Retrieved EpochInfo for epochstypes.LIQUID_STAKING_DApp_Rewards_EPOCH: {liquid-staking-epoch 2023-07-28 12:04:08.993680513 +0000 UTC 2m30s 4 2023-07-28 12:11:38.993680513 +0000 UTC true 330} module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF MinimumRewardAmount is greater than 0: 100 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF CurrentEpoch 3 is not 0 and is a multiple of 1 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Start Get Total Liquid Stake----------------: 4 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Get Contract Liquid Stake Deposits Till Epoch----------------: 6172636877617931756A743035616330367A666474726C6C67397061357539773478366B7770737971613267346E 4 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Set Contract Liquid Stake Deposits Till Epoch----------------: 6172636877617931756A743035616330367A666474726C6C67397061357539773478366B7770737971613267346E 4 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Reward Address archway1ujt05ac06zfdtrllg9pa5u9w4x6kwpsyqa2g4n, Epoch 4 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Total liquid stake 0 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Finish Get Total Liquid Stake----------------: 4 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Retrieved TotalLiquidStake: 0 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Start Liquid Stake----------------: 3 4 0 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Liquid stake amount: 0 
dockernet-photo1-1  |  module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Liquid stake amount: 0 
dockernet-photo1-1  |  module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Finish Liquid Stake----------------: 3 4 0 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Retrieved TotalLiquidStake: 0 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF LiquidStakeScheduled: 1 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Start Distribute Liquidity----------------: 4 0 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Get Contract Liquid Stake Deposits Till Epoch----------------: 6172636877617931756A743035616330367A666474726C6C67397061357539773478366B7770737971613267346E 4 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Set Contract Liquid Stake Deposits Till Epoch----------------: 6172636877617931756A743035616330367A666474726C6C67397061357539773478366B7770737971613267346E 4 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Reward Address archway1ujt05ac06zfdtrllg9pa5u9w4x6kwpsyqa2g4n, Till Epoch considered for distribution 4 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Deposit Record Stake ratio determination for liquidity distribution [] module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF --------------- Start Delete Liquid Stake DepositRecord----------------: 6172636877617931756A743035616330367A666474726C6C67397061357539773478366B7770737971613267346E module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Deleted Liquid stake deposit Record module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Finish Delete Liquid Stake DepositRecord----------------: 6172636877617931756A743035616330367A666474726C6C67397061357539773478366B7770737971613267346E module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Finish Distribute Liquidity----------------: 4 0 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Distributed Liquidity for epoch 4 and liquid stake 0 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Checking epoch info: {day 2023-07-28 12:04:08.993680513 +0000 UTC 1m0s 10 2023-07-28 12:13:08.993680513 +0000 UTC true 397} module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Checking epoch info: {liquid-staking-epoch 2023-07-28 12:04:08.993680513 +0000 UTC 2m30s 4 2023-07-28 12:11:38.993680513 +0000 UTC true 330} module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Processing LiquidStakeDappRewards epoch: {liquid-staking-epoch 2023-07-28 12:04:08.993680513 +0000 UTC 2m30s 4 2023-07-28 12:11:38.993680513 +0000 UTC true 330} module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Retrieved EpochInfo: {liquid-staking-epoch 2023-07-28 12:04:08.993680513 +0000 UTC 2m30s 4 2023-07-28 12:11:38.993680513 +0000 UTC true 330} module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF MinimumRewardAmount is greater than 0: 100 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF CurrentEpoch 4 is not 0 and is a multiple of LiquidStakeInterval 1 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF RewardsAddress is not empty: archway1ujt05ac06zfdtrllg9pa5u9w4x6kwpsyqa2g4n module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Start Get Cumulative Reward Amount: 4 6172636877617931756A743035616330367A666474726C6C67397061357539773478366B7770737971613267346E module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Get Cumulative Reward Amount:  module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Get Total Enqueued Rewards----------------: 4 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Get Contract Liquid Stake Deposits Till Epoch----------------: 6172636877617931756A743035616330367A666474726C6C67397061357539773478366B7770737971613267346E 4 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Set Contract Liquid Stake Deposits Till Epoch----------------: 6172636877617931756A743035616330367A666474726C6C67397061357539773478366B7770737971613267346E 4 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF ---------------Set Total Enqueued Rewards----------------: 4 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Get Cumulative Reward Amount Enqueued: 0 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Finish Get Cumulative Reward Amount: 4 6172636877617931756A743035616330367A666474726C6C67397061357539773478366B7770737971613267346E module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Retrieved CumulativeRewardAmount: 0 module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Checking epoch info: {mint 2023-07-28 12:04:08.993680513 +0000 UTC 1m0s 10 2023-07-28 12:13:08.993680513 +0000 UTC true 397} module=x/photosynthesis
dockernet-photo1-1  | 12:13PM INF Checking epoch info: {redemption-rate-query-epoch 2023-07-28 12:04:08.993680513 +0000 UTC 5m0s 2 2023-07-28 12:09:08.993680513 +0000 UTC true 199} module=x/photosynthesis""".split("\n")



for log_line in log_lines:
    # Split the log line into components
    line_parts = log_line.split()
    if line_parts:
        timestamp = line_parts[1]
        log_type = line_parts[2]
        message = " ".join(line_parts[3:]).split(":")[0].strip()
        module = line_parts[-1]

        # Extract data using regular expressions
        data_match = re.search(r":\s+{(.+)}", log_line)
        data_dict = {}

        if data_match:
            data_str = "{" + data_match.group(1) + "}"
            data_dict = json.loads(data_str)

        # Create a dictionary from the components
        log_dict = {
            "timestamp": timestamp,
            "log_type": log_type,
            "message": message,
            "data": data_dict,
            "module": module,
        }

        # Add the dictionary to the list of parsed logs
        parsed_logs.append(log_dict)

# Convert the list of parsed logs into a JSON string
json_logs = json.dumps(parsed_logs, indent=4)

print(json_logs)

