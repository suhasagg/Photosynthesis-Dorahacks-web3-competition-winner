import re

def extract_amount(message):
    match = re.search(r'(\d+)ibc/', message)
    if match:
        return int(match.group(1))
    return 0

# Test the function
if __name__ == "__main__":
    # The command message
    message = "/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd tx --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 bank send archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m archway1mulfsmceu53x0qnwv05d68lxmy5kg430jtnljm 310000ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145 --from pval5 --keyring-backend=test --chain-id localnet --fees 17000uarch -y"
    
    # Extract amount and print it
    extracted_amount = extract_amount(message)
    print(f"Extracted amount: {extracted_amount}")

