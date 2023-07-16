// Copyright 2022 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package types

// Allow Allow specifies supported Operation status, Operation types, and all possible error
// statuses. This Allow object is used by clients to validate the correctness of a Rosetta Server
// implementation. It is expected that these clients will error if they receive some response that
// contains any of the above information that is not specified here.
type Allow struct {
	// All Operation.Status this implementation supports. Any status that is returned during parsing
	// that is not listed here will cause client validation to error.
	OperationStatuses []*OperationStatus `json:"operation_statuses"`
	// All Operation.Type this implementation supports. Any type that is returned during parsing
	// that is not listed here will cause client validation to error.
	OperationTypes []string `json:"operation_types"`
	// All Errors that this implementation could return. Any error that is returned during parsing
	// that is not listed here will cause client validation to error.
	Errors []*Error `json:"errors"`
	// Any Rosetta implementation that supports querying the balance of an account at any height in
	// the past should set this to true.
	HistoricalBalanceLookup bool `json:"historical_balance_lookup"`
	// If populated, `timestamp_start_index` indicates the first block index where block timestamps
	// are considered valid (i.e. all blocks less than `timestamp_start_index` could have invalid
	// timestamps). This is useful when the genesis block (or blocks) of a network have timestamp 0.
	// If not populated, block timestamps are assumed to be valid for all available blocks.
	TimestampStartIndex *int64 `json:"timestamp_start_index,omitempty"`
	// All methods that are supported by the /call endpoint. Communicating which parameters should
	// be provided to /call is the responsibility of the implementer (this is en lieu of defining an
	// entire type system and requiring the implementer to define that in Allow).
	CallMethods []string `json:"call_methods"`
	// BalanceExemptions is an array of BalanceExemption indicating which account balances could
	// change without a corresponding Operation. BalanceExemptions should be used sparingly as they
	// may introduce significant complexity for integrators that attempt to reconcile all account
	// balance changes. If your implementation relies on any BalanceExemptions, you MUST implement
	// historical balance lookup (the ability to query an account balance at any BlockIdentifier).
	BalanceExemptions []*BalanceExemption `json:"balance_exemptions"`
	// Any Rosetta implementation that can update an AccountIdentifier's unspent coins based on the
	// contents of the mempool should populate this field as true. If false, requests to
	// `/account/coins` that set `include_mempool` as true will be automatically rejected.
	MempoolCoins        bool `json:"mempool_coins"`
	BlockHashCase       Case `json:"block_hash_case,omitempty"`
	TransactionHashCase Case `json:"transaction_hash_case,omitempty"`
}
