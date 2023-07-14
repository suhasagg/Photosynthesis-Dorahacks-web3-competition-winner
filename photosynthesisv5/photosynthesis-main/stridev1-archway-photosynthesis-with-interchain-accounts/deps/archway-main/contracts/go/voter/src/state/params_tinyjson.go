// Code generated by tinyjson for marshaling/unmarshaling. DO NOT EDIT.

package state

import (
	tinyjson "github.com/CosmWasm/tinyjson"
	jlexer "github.com/CosmWasm/tinyjson/jlexer"
	jwriter "github.com/CosmWasm/tinyjson/jwriter"
)

// suppress unused package warning
var (
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ tinyjson.Marshaler
)

func tinyjson54f20b6DecodeGithubComArchwayNetworkVoterSrcState(in *jlexer.Lexer, out *Params) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "owner_addr":
			if in.IsNull() {
				in.Skip()
				out.OwnerAddr = nil
			} else {
				out.OwnerAddr = in.Bytes()
			}
		case "new_voting_cost":
			(out.NewVotingCost).UnmarshalTinyJSON(in)
		case "vote_cost":
			(out.VoteCost).UnmarshalTinyJSON(in)
		case "ibc_send_timeout":
			out.IBCSendTimeout = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson54f20b6EncodeGithubComArchwayNetworkVoterSrcState(out *jwriter.Writer, in Params) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"owner_addr\":"
		out.RawString(prefix[1:])
		out.Base64Bytes(in.OwnerAddr)
	}
	{
		const prefix string = ",\"new_voting_cost\":"
		out.RawString(prefix)
		(in.NewVotingCost).MarshalTinyJSON(out)
	}
	{
		const prefix string = ",\"vote_cost\":"
		out.RawString(prefix)
		(in.VoteCost).MarshalTinyJSON(out)
	}
	{
		const prefix string = ",\"ibc_send_timeout\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.IBCSendTimeout))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Params) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson54f20b6EncodeGithubComArchwayNetworkVoterSrcState(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v Params) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson54f20b6EncodeGithubComArchwayNetworkVoterSrcState(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Params) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson54f20b6DecodeGithubComArchwayNetworkVoterSrcState(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *Params) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson54f20b6DecodeGithubComArchwayNetworkVoterSrcState(l, v)
}
