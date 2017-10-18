package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/render"
	"github.com/go-chi/cors"
	"github.com/pressly/chi"
	"log"
	"math/big"
	"net/http"
)

var key = "b4087f10eacc3a032a2d550c02ae7a3ff88bc62eb0d9f6c02c9d5ef4d1562862"

var clientKey = "d2a99b289915eb11ea50a51247e1cef2c4583ae1d9699a3bb0154c2792bda339"

// Agreement is an agreement
type Agreement struct {
	Account   string `json:"account"`
	Agreement string `json:"agreement"`
}

// Bind binds the request parameters
func (a *Agreement) Bind(r *http.Request) error {
	return nil
}

var nonceCounter int64

func main() {
	privKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
	}
	clientPrivKey, err := crypto.HexToECDSA(clientKey)
	if err != nil {
		log.Fatalf("Failed to convert client private key: %v", err)
	}
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth := bind.NewKeyedTransactor(privKey)
	clientAuth := bind.NewKeyedTransactor(clientPrivKey)
	if err != nil {
		log.Fatalf("Failed to create Transactor: %v", err)
	}
	auth.Nonce = big.NewInt(nonceCounter)

	addr, _, contract, err := DeploySigner(auth, conn)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}
	fmt.Println("Contract Deployed to: ", addr.String())

	r := chi.NewRouter()
	corsOption := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(corsOption.Handler)
	r.Use(middleware.Logger)
	r.Post("/agreement", createAgreementHandler(contract, auth, clientAuth, conn, privKey))

	log.Println("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("tada!")
}

func createAgreementHandler(contract *Signer, auth, clientAuth *bind.TransactOpts, conn *ethclient.Client, privKey *ecdsa.PrivateKey) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		agreement := &Agreement{}
		if err := render.Bind(r, agreement); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid Request, Account and Agreement need to be set"))
			return
		}
		fmt.Println(agreement)
		if agreement.Account == "" || agreement.Agreement == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Account and Agreement need to be set"))
			return
		}
		nonceCounter = nonceCounter + 1
		_, err := contract.CreateAgreement(&bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: big.NewInt(2381623),
			Value:    big.NewInt(0),
			Nonce:    big.NewInt(nonceCounter),
		}, agreement.Agreement, common.HexToAddress(agreement.Account))
		if err != nil {
			log.Fatalf("Failed to create agreement: %v", err)
		}
		fmt.Println("Agreement created: ", agreement.Agreement)
		nonceCounter = nonceCounter + 1
		gasPrice, err := conn.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatalf("Failed to get gas price: %v", err)
		}
		signer := types.HomesteadSigner{}
		tx := types.NewTransaction(uint64(nonceCounter), common.HexToAddress(agreement.Account), big.NewInt(25000000), big.NewInt(21000), gasPrice, nil)
		signed, err := types.SignTx(tx, signer, privKey)
		if err != nil {
			log.Fatalf("Failed to sign transaction: %v", err)
		}
		err = conn.SendTransaction(context.Background(), signed)
		if err != nil {
			log.Fatalf("Failed to send transaction: %v", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}
