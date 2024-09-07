pub mod contract;

pub use crate::contract::{
    instantiate, execute, query, ExecuteMsg, QueryMsg, Config, User, Campaign, Message, InitMsg,
};

// Testing framework
#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::{Addr, Empty};
    use cw_multi_test::{App, Contract, ContractWrapper, Executor};

    // Helper function to set up the mock app
    fn mock_app() -> App {
        App::default()
    }

    fn contract_campaign() -> Box<dyn Contract<Empty>> {
        let contract = ContractWrapper::new(execute, instantiate, query);
        Box::new(contract)
    }

    // Test to verify contract instantiation
    #[test]
    fn test_instantiate_contract() {
        let mut app = mock_app();

        // Define the InitMsg with owner address
        let owner = "cosmos1address".to_string();
        let init_msg = InitMsg { owner };

        // Store the contract and instantiate it
        let code_id = app.store_code(contract_campaign());

        let contract_addr = app
            .instantiate_contract(
                code_id,
                Addr::unchecked("creator"),  // Sender's address
                &init_msg,  
                &[],
                "CampaignContract",
                None,
            )
            .unwrap();

        // Query the contract's configuration to ensure it's properly instantiated
        let config: Config = app
            .wrap()
            .query_wasm_smart(contract_addr.clone(), &QueryMsg::GetConfig {})
            .unwrap();

        // Assert the owner is correctly set
        assert_eq!(config.owner, Addr::unchecked("cosmos1address"));
    }
}

