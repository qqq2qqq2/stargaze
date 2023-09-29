package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"
	"github.com/public-awesome/stargaze/v12/x/globalfee/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) SetCodeAuthorization(goCtx context.Context, msg *types.MsgSetCodeAuthorization) (*types.MsgSetCodeAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !k.IsPrivilegedAddress(ctx, msg.Sender) {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "sender address is not privileged address")
	}
	err = k.Keeper.SetCodeAuthorization(ctx, *msg.CodeAuthorization)
	if err != nil {
		return nil, err
	}
	return &types.MsgSetCodeAuthorizationResponse{}, nil
}

func (k msgServer) RemoveCodeAuthorization(goCtx context.Context, msg *types.MsgRemoveCodeAuthorization) (*types.MsgRemoveCodeAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !k.IsPrivilegedAddress(ctx, msg.Sender) {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "sender address is not privileged address")
	}
	k.Keeper.DeleteCodeAuthorization(ctx, msg.GetCodeID())
	return &types.MsgRemoveCodeAuthorizationResponse{}, nil
}

func (k msgServer) SetContractAuthorization(goCtx context.Context, msg *types.MsgSetContractAuthorization) (*types.MsgSetContractAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !k.IsPrivilegedAddress(ctx, msg.Sender) {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "sender address is not privileged address")
	}
	err = k.Keeper.SetContractAuthorization(ctx, *msg.ContractAuthorization)
	if err != nil {
		return nil, err
	}
	return &types.MsgSetContractAuthorizationResponse{}, nil
}

func (k msgServer) RemoveContractAuthorization(goCtx context.Context, msg *types.MsgRemoveContractAuthorization) (*types.MsgRemoveContractAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !k.IsPrivilegedAddress(ctx, msg.Sender) {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "sender address is not privileged address")
	}
	contractAddr, err := sdk.AccAddressFromBech32(msg.ContractAddress)
	if err != nil {
		return nil, err
	}
	k.Keeper.DeleteContractAuthorization(ctx, contractAddr)
	return &types.MsgRemoveContractAuthorizationResponse{}, nil
}

func (k msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	if msg.Sender != k.Keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "sender address is not authorized address to update module params")
	}

	err = msg.GetParams().Validate() // need to explicitly validate as x/gov invokes this msg and it does not validate
	if err != nil {
		return nil, err
	}

	k.SetParams(ctx, msg.GetParams())

	return &types.MsgUpdateParamsResponse{}, nil
}

// func (m msgServer) isAuthorized(ctx sdk.Context, actor string) bool {
// 	if actor == m.Keeper.GetAuthority() {
// 		return true
// 	}
// 	if m.Keeper.IsAdminAddress(ctx, actor) {
// 		return true
// 	}
// 	return false
// }
