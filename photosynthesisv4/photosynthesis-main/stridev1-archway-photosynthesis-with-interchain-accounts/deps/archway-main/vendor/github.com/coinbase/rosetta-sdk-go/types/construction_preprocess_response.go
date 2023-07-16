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

// ConstructionPreprocessResponse ConstructionPreprocessResponse contains `options` that will be
// sent unmodified to `/construction/metadata`. If it is not necessary to make a request to
// `/construction/metadata`, `options` should be omitted.  Some blockchains require the PublicKey of
// particular AccountIdentifiers to construct a valid transaction. To fetch these PublicKeys,
// populate `required_public_keys` with the AccountIdentifiers associated with the desired
// PublicKeys. If it is not necessary to retrieve any PublicKeys for construction,
// `required_public_keys` should be omitted.
type ConstructionPreprocessResponse struct {
	// The options that will be sent directly to `/construction/metadata` by the caller.
	Options            map[string]interface{} `json:"options,omitempty"`
	RequiredPublicKeys []*AccountIdentifier   `json:"required_public_keys,omitempty"`
}
