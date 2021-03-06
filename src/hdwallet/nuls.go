package hdwallet

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/hdkeychain"
	"hdwallet/nuls"
)

//Nuls转账
func NulsTransfer(privateKey, toAddress, amount, price, remark string) (signedParam string, err error) {
	var signRawTx string
	if (len(privateKey) == 0) || (len(toAddress) == 0) {
		return signRawTx, errors.New("some params is empty!!!")
	}
	child, err := hdkeychain.NewKeyFromString(privateKey)
	if err != nil {
		return signRawTx, err
	}
	fromAddress := nuls.Address(child)

	fmt.Println("The NULS send address is ", fromAddress)

	//Sign
	signRawTx, err = nuls.Transfer(fromAddress, toAddress, amount, price, remark, child)
	return signRawTx, err
}

func NulsTransferFee(fromAddress, toAddress, amount, price, remark string) (fee string, err error) {
	//Sign
	fee, err = nuls.TransferFee(fromAddress, toAddress, amount, price, remark)
	return fee, err
}

func NulsBroadcast(txHex string) (string, error) {
	return nuls.Broadcast(txHex)
}

func NulsBalance(address string) (balance string, err error) {
	tempAmount, err := nuls.GetBalance(address)
	if err != nil {
		return
	}
	balanceByte, err := json.Marshal(tempAmount)
	balance = string(balanceByte)
	return balance, err
}
