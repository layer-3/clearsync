package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/layer-3/clearsync/pkg/userop"
)

func main() {
	client, err := userop.NewClient(config)
	if err != nil {
		panic(fmt.Errorf("failed to create userop client: %w", err))
	}

	op, err := client.NewUserOp(context.Background(), sender, receiver, token, amount)
	if err != nil {
		panic(fmt.Errorf("failed to build userop: %w", err))
	}

	callback := func() {}
	if err := client.SendUserOp(context.Background(), op, callback); err != nil {
		panic(fmt.Errorf("failed to send userop: %w", err))
	}
}

func signUserOp(op userop.UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
	slog.Info("signing user operation")

	// hash := op.UserOpHash(entryPoint, chainID)
	// privateKey, err := crypto.HexToECDSA("9be4e88f3d84ff8b58ec9f11f047245cad2f8e2c6cf43ca4087331ffc0fea6c8")
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to parse private key: %w", err)
	// }

	// sigBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to sign user operation: %w", err)
	// }

	// signatureStr := hexutil.Encode(sigBytes)
	// slog.Info("signed", "signatureStr", signatureStr)

	// slog.Info("user operation signed:", "userOpHash", hash, "signature", signatureStr)
	// return []byte(signatureStr), nil

	return []byte("0x00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000001c5b32f37f5bea87bdd5374eb2ac54ea8e000000000000000000000000000000000000000000000000000000000000004118c5324f4228fd989fcc67cd7299572d687d9f22bf2ff47d208acda055dcae282117add286111ca21455d03467ccf5b1c70e31366a5feaa8c426ec501ed97e921c00000000000000000000000000000000000000000000000000000000000000"), nil
}
