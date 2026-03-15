package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/labstack/gommon/random"
)

const (
	baseURL = "http://localhost:8080"
)

func TestHTTPService(t *testing.T) {
	var wallet1ID string
	var wallet1Name string

	t.Run("CreateWallet", func(t *testing.T) {
		wallet1Name = "wallet-1-" + random.String(16, random.Alphanumeric)
		reqBody := map[string]string{
			"name": wallet1Name,
		}
		jsonData, _ := json.Marshal(reqBody)

		resp, err := http.Post(baseURL+"/wallets", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatalf("CreateWallet request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("CreateWallet returned status %d: %s", resp.StatusCode, string(body))
		}

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)

		if result["name"] != wallet1Name {
			t.Errorf("Expected name '%s', got '%v'", wallet1Name, result["name"])
		}

		if result["display_id"] == nil {
			t.Error("Expected non-empty display_id")
		} else {
			wallet1ID = result["display_id"].(string)
		}

		if result["balance"] == nil || result["balance"] != "0" {
			t.Errorf("Expected balance '0', got '%v'", result["balance"])
		}

		t.Logf("Created wallet: name=%s, display_id=%s, balance=%s",
			result["name"], result["display_id"], result["balance"])
	})

	var wallet2ID string
	var wallet2Name string

	t.Run("CreateSecondWallet", func(t *testing.T) {
		wallet2Name = "wallet-2-" + random.String(16, random.Alphanumeric)
		reqBody := map[string]string{
			"name": wallet2Name,
		}
		jsonData, _ := json.Marshal(reqBody)

		resp, err := http.Post(baseURL+"/wallets", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatalf("CreateSecondWallet request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("CreateSecondWallet returned status %d: %s", resp.StatusCode, string(body))
		}

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)

		wallet2ID = result["display_id"].(string)
		t.Logf("Created second wallet: display_id=%s", wallet2ID)
	})

	t.Run("GetWalletDetail", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/wallets/%s", baseURL, wallet2ID))
		if err != nil {
			t.Fatalf("GetWalletDetail request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("GetWalletDetail returned status %d: %s", resp.StatusCode, string(body))
		}

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)

		if result["name"] != wallet2Name {
			t.Errorf("Expected name '%s', got '%v'", wallet2Name, result["name"])
		}

		if result["display_id"] != wallet2ID {
			t.Errorf("Expected display_id '%s', got '%v'", wallet2ID, result["display_id"])
		}

		t.Logf("Got wallet detail: name=%s, display_id=%s, balance=%s",
			result["name"], result["display_id"], result["balance"])
	})

	t.Run("WalletTransfer", func(t *testing.T) {
		reqBody := map[string]string{
			"from_display_id": wallet1ID,
			"to_display_id":   wallet2ID,
			"amount":          "10",
		}
		jsonData, _ := json.Marshal(reqBody)

		resp, err := http.Post(baseURL+"/wallets/transfer", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatalf("WalletTransfer request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Logf("WalletTransfer returned status %d (expected): %s", resp.StatusCode, string(body))
			return
		}

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)

		t.Logf("Transfer response: %v", result)
	})

	t.Run("InvalidRequests", func(t *testing.T) {
		t.Run("EmptyName", func(t *testing.T) {
			reqBody := map[string]string{
				"name": "",
			}
			jsonData, _ := json.Marshal(reqBody)

			resp, err := http.Post(baseURL+"/wallets", "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				t.Fatalf("EmptyName request failed: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusBadRequest {
				t.Errorf("Expected status 400, got %d", resp.StatusCode)
			}
		})

		t.Run("NonExistentWallet", func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("%s/wallets/nonexistent-id", baseURL))
			if err != nil {
				t.Fatalf("NonExistentWallet request failed: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusNotFound {
				t.Errorf("Expected status 404, got %d", resp.StatusCode)
			}
		})
	})
}

func BenchmarkCreateWalletHTTP(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			walletName := random.String(16, random.Alphanumeric)
			reqBody := map[string]string{
				"name": walletName,
			}
			jsonData, _ := json.Marshal(reqBody)

			resp, err := http.Post(baseURL+"/wallets", "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				b.Errorf("CreateWallet request failed: %v", err)
				continue
			}
			resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				b.Errorf("Expected status 200, got %d", resp.StatusCode)
			}
		}
	})
}
