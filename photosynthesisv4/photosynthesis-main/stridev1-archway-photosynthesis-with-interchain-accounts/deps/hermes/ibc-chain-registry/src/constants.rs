pub const PROTOCOL: &str = "https";
pub const HOST: &str = "raw.githubusercontent.com";
pub const REGISTRY_PATH: &str = "/cosmos/chain-registry";
pub const DEFAULT_REF: &str = "master";
pub const ALL_CHAINS: &[&str] = &[
    "agoric",
    "aioz",
    "akash",
    "arkh",
    "assetmantle",
    "axelar",
    "bandchain",
    "bitcanna",
    "bitsong",
    "bostrom",
    "carbon",
    "cerberus",
    "cheqd",
    "chihuahua",
    "chronicnetwork",
    "comdex",
    "cosmoshub",
    "crescent",
    "cronos",
    "cryptoorgchain",
    "cudos",
    "decentr",
    "desmos",
    "dig",
    "echelon",
    "emoney",
    "ethos",
    "evmos",
    "fetchhub",
    "firmachain",
    "galaxy",
    "genesisl1",
    "gravitybridge",
    "idep",
    "impacthub",
    "injective",
    "irisnet",
    "juno",
    "kava",
    "kichain",
    "konstellation",
    "kujira",
    "likecoin",
    "logos",
    "lumenx",
    "lumnetwork",
    "meme",
    "microtick",
    "mythos",
    //"nomic", Does not have assetlist.json
    "octa",
    "odin",
    "omniflixhub",
    "oraichain",
    "osmosis",
    "panacea",
    "persistence",
    "provenance",
    "regen",
    "rizon",
    "secretnetwork",
    "sentinel",
    "shentu",
    "sifchain",
    "sommelier",
    "stargaze",
    "starname",
    "terra",
    "terra2",
    "tgrade",
    //"thorchain", Does not have assetlist.json
    "umee",
    "vidulum",
];
pub const ALL_PATHS: &[&str] = &[
    "akash-cosmoshub.json",
    "akash-cryptoorgchain.json",
    "akash-irisnet.json",
    "akash-juno.json",
    "akash-osmosis.json",
    "akash-persistence.json",
    "akash-regen.json",
    "akash-secretnetwork.json",
    "akash-sentinel.json",
    "akash-sifchain.json",
    "akash-starname.json",
    "assetmantle-juno.json",
    "assetmantle-osmosis.json",
    "axelar-crescent.json",
    "axelar-juno.json",
    "axelar-osmosis.json",
    "axelar-secretnetwork.json",
    "bandchain-osmosis.json",
    "bitcanna-juno.json",
    "bitcanna-osmosis.json",
    "bitsong-juno.json",
    "bitsong-osmosis.json",
    "bostrom-osmosis.json",
    "carbon-osmosis.json",
    "cerberus-osmosis.json",
    "cheqd-osmosis.json",
    "chihuahua-juno.json",
    "chihuahua-osmosis.json",
    "chihuahua-secretnetwork.json",
    "comdex-juno.json",
    "comdex-osmosis.json",
    "cosmoshub-crescent.json",
    "cosmoshub-cryptoorgchain.json",
    "cosmoshub-emoney.json",
    "cosmoshub-impacthub.json",
    "cosmoshub-irisnet.json",
    "cosmoshub-juno.json",
    "cosmoshub-likecoin.json",
    "cosmoshub-osmosis.json",
    "cosmoshub-persistence.json",
    "cosmoshub-regen.json",
    "cosmoshub-secretnetwork.json",
    "cosmoshub-sentinel.json",
    "cosmoshub-sifchain.json",
    "cosmoshub-starname.json",
    "cosmoshub-umee.json",
    "crescent-gravitybridge.json",
    "crescent-osmosis.json",
    "crescent-secretnetwork.json",
    "cryptoorgchain-irisnet.json",
    "cryptoorgchain-osmosis.json",
    "cryptoorgchain-persistence.json",
    "cryptoorgchain-regen.json",
    "cryptoorgchain-sentinel.json",
    "cryptoorgchain-sifchain.json",
    "cryptoorgchain-starname.json",
    "decentr-osmosis.json",
    "desmos-osmosis.json",
    "dig-juno.json",
    "dig-osmosis.json",
    "emoney-irisnet.json",
    "emoney-juno.json",
    "emoney-osmosis.json",
    "evmos-osmosis.json",
    "evmos-secretnetwork.json",
    "fetchhub-osmosis.json",
    "galaxy-osmosis.json",
    "genesisl1-osmosis.json",
    "gravitybridge-osmosis.json",
    "gravitybridge-secretnetwork.json",
    "impacthub-osmosis.json",
    "impacthub-sifchain.json",
    "injective-osmosis.json",
    "injective-secretnetwork.json",
    "irisnet-osmosis.json",
    "irisnet-persistence.json",
    "irisnet-regen.json",
    "irisnet-sentinel.json",
    "irisnet-sifchain.json",
    "irisnet-starname.json",
    "juno-osmosis.json",
    "juno-persistence.json",
    "juno-secretnetwork.json",
    "juno-sifchain.json",
    "juno-stargaze.json",
    "juno-terra.json",
    "kava-osmosis.json",
    "kichain-osmosis.json",
    "konstellation-osmosis.json",
    "kujira-osmosis.json",
    "likecoin-osmosis.json",
    "lumenx-osmosis.json",
    "lumnetwork-osmosis.json",
    "meme-osmosis.json",
    "microtick-osmosis.json",
    "oraichain-osmosis.json",
    "osmosis-panacea.json",
    "osmosis-persistence.json",
    "osmosis-provenance.json",
    "osmosis-regen.json",
    "osmosis-rizon.json",
    "osmosis-secretnetwork.json",
    "osmosis-sentinel.json",
    "osmosis-shentu.json",
    "osmosis-sifchain.json",
    "osmosis-sommelier.json",
    "osmosis-stargaze.json",
    "osmosis-starname.json",
    "osmosis-terra.json",
    "osmosis-tgrade.json",
    "osmosis-umee.json",
    "osmosis-vidulum.json",
    "persistence-regen.json",
    "persistence-sentinel.json",
    "persistence-sifchain.json",
    "persistence-starname.json",
    "regen-sentinel.json",
    "regen-sifchain.json",
    "regen-starname.json",
    "secretnetwork-sentinel.json",
    "secretnetwork-sifchain.json",
    "secretnetwork-stargaze.json",
    "secretnetwork-terra.json",
    "secretnetwork-terra2.json",
    "sentinel-sifchain.json",
    "sentinel-starname.json",
];
