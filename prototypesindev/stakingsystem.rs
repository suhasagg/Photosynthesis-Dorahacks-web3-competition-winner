struct StakingSystem {
    scheduler: Scheduler,
    blockchain_interface: BlockchainInterface,
}

impl StakingSystem {
    fn new() -> Self {
        Self {
            scheduler: Scheduler::new(),
            blockchain_interface: BlockchainInterface::new(),
        }
    }

    fn run(&mut self) {
        // Schedule the periodic checks and operations
        self.scheduler.every_day_at("00:00", || self.check_and_stake());
        self.scheduler.every_day_at("12:00", || self.check_and_redeem());
        self.scheduler.start(); // Starts the scheduler in a loop
    }

    fn check_and_stake(&self) {
        let rate = self.blockchain_interface.fetch_redemption_rate();
        if self.should_stake(rate) {
            self.blockchain_interface.stake(uarch_amount);
        }
    }

    fn check_and_redeem(&self) {
        let rate = self.blockchain_interface.fetch_redemption_rate();
        if self.should_redeem(rate) {
           self.blockchain_interface.redeem(stuarch_amount);
        }
    }

    fn should_stake(rate: f64) -> bool {
        // Implement the staking logic based on the rate
        rate < 1.45 && rate > 1.05
    }

    fn should_redeem(rate: f64) -> bool {
        // Implement the redemption logic based on the rate
        rate >= 1.45
    }
}

struct Scheduler {
    // Task scheduling logic
}

struct BlockchainInterface {
    // Interface to interact with the blockchain
    fn fetch_redemption_rate(&self) -> f64 {
        // API call to fetch the current rate
    }

    fn stake(&self, amount: u64) {
        // API call to perform staking
    }

    fn redeem(&self, amount: u64) {
        // API call to perform redemption
    }
}

