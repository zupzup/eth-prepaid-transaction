package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/pressly/chi"
	"log"
	"math/big"
	"net/http"
)

var key = "b4087f10eacc3a032a2d550c02ae7a3ff88bc62eb0d9f6c02c9d5ef4d1562862"

// Agreement is an agreement
type Agreement struct {
	Account   string `json:"account"`
	Agreement string `json:"agreement"`
}

// Bind binds the request parameters
func (a *Agreement) Bind(r *http.Request) error {
	return nil
}

func main() {
	privKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
	}
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth := bind.NewKeyedTransactor(privKey)
	if err != nil {
		log.Fatalf("Failed to create Transactor: %v", err)
	}
	fmt.Println(auth.From.String(), auth.Nonce.String())
	auth.Nonce = big.NewInt(4)

	addr, _, contract, err := DeploySigner(auth, conn)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}
	fmt.Println("Contract Deployed to: ", addr.String())

	createTrans, err := contract.CreateAgreement(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(2381623),
		Value:    big.NewInt(0),
		Nonce:    big.NewInt(5),
	}, "test project", common.HexToAddress("446b69e95817842b6c27de11cdcba498b6a35c1c"))
	if err != nil {
		log.Fatalf("Failed to create agreement: %v", err)
	}
	fmt.Println("Agreement created: ", createTrans.String())

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
	r.Post("/agreement", createAgreementHandler())

	log.Println("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("tada!")
}

func createAgreementHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		agreement := &Agreement{}
		if err := agreement.Bind(r); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid Request, Account and Agreement need to be set"))
			return
		}
		if agreement.Account == "" || agreement.Agreement == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Account and Agreement need to be set"))
			return
		}
		fmt.Println("save to contract")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}
