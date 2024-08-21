package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	db "simplebank/db/sqlc"
	"simplebank/db/token"

	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountId int64  `json:"from_account_id" binding:"required,min=0"`
	ToAccountId   int64  `json:"to_account_id" binding:"required,min=0"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fromAccount, valid := server.validAccount(ctx, req.FromAccountId, req.Currency)
	if !valid {
		return
	}


	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if fromAccount.Owner != authPayload.Username {
		err := errors.New("from account doesn't belong to this user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_ ,  valid =  server.validAccount(ctx, req.ToAccountId, req.Currency)
	if !valid {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountId,
		ToAccountID:   req.ToAccountId,
		Amount:        req.Amount,
	}

	account, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, account)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Account,bool) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}
	if account.Currency != currency { 
		err := fmt.Errorf("account [%d] currency is not [%s]",accountID,currency)
		ctx.JSON(http.StatusBadRequest,err)
		return account, false
	}
	return account, true
}
