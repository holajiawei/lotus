package validation

import (
	"golang.org/x/xerrors"
)

// Errors extracted from `ValidateBlock` and `checkBlockMessages`.
// Explicitly split `var` declarations instead of having a single block to avoid `go fmt`
// potentially changing alignment of all errors (producing harder-to-read diffs).
// FIXME: How to graphically express this error hierarchy *without* reflection? (nested packages?)

var ErrInvalidBlock = xerrors.New("invalid block")

var ErrBlockNilSignature = ErrorWrapString(ErrInvalidBlock, "block had nil signature")

var ErrTimestamp = ErrorWrapString(ErrInvalidBlock, "block timestamp error")
var ErrBlockFutureTimestamp = ErrorWrapString(ErrTimestamp, "ahead of current time")

var ErrBlockMinedEarly = ErrorWrapString(ErrInvalidBlock, "block was generated too soon")

var ErrWinner = ErrorWrapString(ErrInvalidBlock, "not a winner block")
var ErrSlashedMiner = ErrorWrapString(ErrWinner, "slashed or invalid miner")
var ErrNoCandidates = ErrorWrapString(ErrWinner, "no candidates")
var ErrDuplicateCandidates = ErrorWrapString(ErrWinner, "duplicate epost candidates")
// FIXME: Might want to include these in some EPost category.

var ErrMiner = ErrorWrapString(ErrInvalidBlock, "invalid miner")

var ErrInvalidMessagesCID = ErrorWrapString(ErrInvalidBlock, "messages CID didn't match message root in header")

var ErrInvalidMessageInBlock = ErrorWrapString(ErrInvalidBlock, "invalid message")
var ErrInvalidBlsSignature = ErrorWrapString(ErrInvalidMessageInBlock, "invalid bls aggregate signature")
var ErrInvalidSecpkSignature = ErrorWrapString(ErrInvalidMessageInBlock, "invalid secpk signature")
var ErrEmptyRecipient = ErrorWrapString(ErrInvalidMessageInBlock, "empty 'To' address")
var ErrWrongNonce = ErrorWrapString(ErrInvalidMessageInBlock, "wrong nonce")
// FIXME: Errors from `checkBlockMessages` are too generic and should probably be extracted, is there
//  another place where we validate them (like the message pool)? In that case we need a different
//  root error like `ErrInvalidMessage`.
