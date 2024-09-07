pub mod contract;

pub use crate::contract::{
    instantiate, execute, query, ExecuteMsg, QueryMsg, Config, User};

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

   
}

