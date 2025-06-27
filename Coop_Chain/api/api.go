package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mcp-chain/core/blockchain"
	"mcp-chain/core/services"
	"mcp-chain/core/state"
	"mcp-chain/core/wallet"
	"mcp-chain/internal/config"
	"mcp-chain/types"
	"net/http"
)

// StartAPIServer запускает REST API, принимает chain и cfg как параметры
func StartAPIServer(port int, chain interface{}, cfg interface{}) {
	// Приведение типов (ожидаем *blockchain.Chain и *config.NetworkConfig)
	bc := chain.(*blockchain.Chain)
	conf := cfg.(*config.NetworkConfig)

	http.HandleFunc("/v1/wallet/import", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			Seed string `json:"seed"`
		}
		body, _ := ioutil.ReadAll(r.Body)
		_ = json.Unmarshal(body, &req)
		if req.Seed == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"seed required"}`))
			return
		}
		entry, err := wallet.ImportWalletAndSave(conf, req.Seed)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"` + err.Error() + `"}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entry)
	})

	http.HandleFunc("/v1/wallet/list", func(w http.ResponseWriter, r *http.Request) {
		wallets := wallet.ListWalletsFromState()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(wallets)
	})

	http.HandleFunc("/v1/state/summary", func(w http.ResponseWriter, r *http.Request) {
		st := state.GlobalState
		res := struct {
			Accounts   int `json:"accounts"`
			Validators int `json:"validators"`
			Blocks     int `json:"blocks"`
		}{
			Accounts:   len(st.Accounts),
			Validators: len(st.Validators),
			Blocks:     len(st.Blocks),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/v1/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintln(w, "501 Not Implemented: /v1/status")
	})

	http.HandleFunc("/v1/send", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintln(w, "501 Not Implemented: /v1/send")
	})

	http.HandleFunc("/v1/coop/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintln(w, "501 Not Implemented: /v1/coop/*")
	})

	http.HandleFunc("/v1/tx/send", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			From   string  `json:"from"`
			To     string  `json:"to"`
			Amount float64 `json:"amount"`
			Fee    float64 `json:"fee"`
			Seed   string  `json:"seed"`
		}
		body, _ := ioutil.ReadAll(r.Body)
		_ = json.Unmarshal(body, &req)
		if req.From == "" || req.To == "" || req.Amount <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"from, to, amount required"}`))
			return
		}
		amount := uint64(req.Amount)
		fee := uint64(req.Fee)
		fromInput := req.From
		if req.Seed != "" {
			fromInput = req.Seed
		}
		id, err := services.SendTxAndSave(bc, conf, fromInput, types.Address(req.To), amount, fee)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"` + err.Error() + `"}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"tx_id":"` + id + `"}`))
	})

	fmt.Printf("API сервер запущен на порту %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("[API] Не удалось запустить сервер на порту %d: %v", port, err)
	}
}
