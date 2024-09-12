pub mod contract;

pub use crate::contract::{
    instantiate, execute, query, HandleMsg, QueryMsg, InitMsg, StakeResponse,
};

#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::{Addr, coins, Uint128, Empty}; 
    use cw_multi_test::{App, Contract, ContractWrapper, Executor};  

    fn mock_app() -> App {
        App::default()
    }

    fn contract_mission_impossible() -> Box<dyn Contract<Empty>> {
        let contract = ContractWrapper::new(execute, instantiate, query);
        Box::new(contract)
    }

    #[test]
    fn test_instantiate() {
        let mut app = mock_app();
        let code_id = app.store_code(contract_mission_impossible());

        let contract_addr = app.instantiate_contract(
            code_id,
            Addr::unchecked("creator"),
            &InitMsg {},
            &[],
            "Mission_Impossible",
            None,
        ).unwrap();

        assert_eq!(contract_addr, Addr::unchecked("contract0"));
    }

   
}
