package test

import (
	"context"
	"testing"
	"time"

	"github.com/flametest/wallet-demo/proto"
	"github.com/labstack/gommon/random"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcService(t *testing.T) {
	conn, err := grpc.NewClient("localhost:9002",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := proto.NewWalletDemoServiceClient(conn)

	var wallet1ID string
	t.Run("CreateWallet", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		wallet1Name := "wallet-1-" + random.String(16, random.Alphanumeric)
		req := &proto.CreateWalletReq{
			Name: wallet1Name,
		}

		resp, err := client.CreateWallet(ctx, req)
		if err != nil {
			t.Fatalf("CreateWallet failed: %v", err)
		}

		if resp.Name != wallet1Name {
			t.Errorf("Expected name 'TestWallet', got '%s'", resp.Name)
		}

		if resp.DisplayId == "" {
			t.Error("Expected non-empty display_id")
		}

		if resp.Balance != "0" {
			t.Errorf("Expected balance '0', got '%s'", resp.Balance)
		}

		wallet1ID = resp.DisplayId

		t.Logf("Created wallet: name=%s, display_id=%s, balance=%s",
			resp.Name, resp.DisplayId, resp.Balance)
	})

	var wallet2ID string
	var wallet2Name string
	t.Run("CreateSecondWallet", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		wallet2Name = "wallet-2-" + random.String(16, random.Alphanumeric)
		req := &proto.CreateWalletReq{
			Name: wallet2Name,
		}

		resp, err := client.CreateWallet(ctx, req)
		if err != nil {
			t.Fatalf("CreateSecondWallet failed: %v", err)
		}

		wallet2ID = resp.DisplayId
		t.Logf("Created second wallet: display_id=%s", wallet2ID)
	})

	t.Run("GetWalletDetail", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		req := &proto.GetWalletDetailReq{
			DisplayId: wallet2ID,
		}

		resp, err := client.GetWalletDetail(ctx, req)
		if err != nil {
			t.Fatalf("GetWalletDetail failed: %v", err)
		}

		if resp.Name != wallet2Name {
			t.Errorf("Expected name %s, got '%s'", wallet2Name, resp.Name)
		}

		if resp.DisplayId != wallet2ID {
			t.Errorf("Expected display_id '%s', got '%s'", wallet2ID, resp.DisplayId)
		}

		t.Logf("Got wallet detail: name=%s, display_id=%s, balance=%s",
			resp.Name, resp.DisplayId, resp.Balance)
	})

	t.Run("WalletTransfer", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		req := &proto.WalletTransferReq{
			FromDisplayId: wallet1ID,
			ToDisplayId:   wallet2ID,
			Amount:        "10",
		}

		resp, err := client.WalletTransfer(ctx, req)
		if err != nil {
			t.Logf("WalletTransfer failed (expected): %v", err)
			return
		}

		if resp.Message == "" {
			t.Error("Expected non-empty message")
		}

		t.Logf("Transfer response: %s", resp.Message)
	})
}

func BenchmarkCreateWallet(b *testing.B) {
	conn, err := grpc.NewClient("localhost:9002",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		b.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewWalletDemoServiceClient(conn)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		req := &proto.CreateWalletReq{
			Name: random.String(16, random.Alphanumeric),
		}
		_, err := client.CreateWallet(ctx, req)
		cancel()
		if err != nil {
			b.Errorf("CreateWallet failed: %v", err)
		}
	}
}
