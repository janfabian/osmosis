package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cosmos/cosmos-sdk/codec/types"
	query "github.com/cosmos/cosmos-sdk/types/query"

	"github.com/osmosis-labs/osmosis/v26/app"
	balancer "github.com/osmosis-labs/osmosis/v26/x/gamm/pool-models/balancer"
	stableswap "github.com/osmosis-labs/osmosis/v26/x/gamm/pool-models/stableswap"
	gamm "github.com/osmosis-labs/osmosis/v26/x/gamm/types"
	poolmanager "github.com/osmosis-labs/osmosis/v26/x/poolmanager/types"
)

func main() {
	// Replace with the actual Osmosis gRPC endpoint
	// For example, using a public endpoint or your own node
	grpcEndpoint := "178.63.142.152:9090"

	// Set up a connection to the gRPC server
	conn, err := grpc.Dial(grpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a new client for the Query service
	client := gamm.NewQueryClient(conn)

	// Prepare context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	encodingConfig := app.MakeEncodingConfig()

	cdc := encodingConfig.Marshaler

	// Initialize variables for pagination
	var allPools []*types.Any
	var nextKey []byte
	page := 1

	for {
		// Create the request with pagination
		req := &gamm.QueryPoolsRequest{
			Pagination: &query.PageRequest{
				Key:   nextKey,
				Limit: 100, // Adjust the limit as needed
			},
		}

		// Make the gRPC call to fetch pools
		res, err := client.Pools(ctx, req)
		if err != nil {
			log.Fatalf("Failed to query pools: %v", err)
		}

		for _, anyPool := range res.Pools {
			var pool poolmanager.PoolI
			err := cdc.UnpackAny(anyPool, &pool)
			if err != nil {
				log.Printf("Failed to unpack pool: %v", err)
				continue
			}

			// Type assert to the specific pool type
			switch p := pool.(type) {
			case *balancer.Pool:
				fmt.Printf("Balancer Pool ID: %d\n", p.GetId())
				// Additional handling for balancer pools
			case *stableswap.Pool:
				stableswapPool, ok := pool.(*stableswap.Pool)
				if !ok {
					log.Fatalf("Pool is not a stableswap pool")
				}

				if stableswapPool.GetId() == 0 {
					log.Fatalf("Pool ID is 0")
				}

				fmt.Printf("Stableswap Pool ID: %d\n", p.GetId())
				// Additional handling for stableswap pools
			// case *concentratedliquidity.Pool:
			// 	fmt.Printf("Concentrated Liquidity Pool ID: %d\n", p.GetId())
			// 	// Additional handling for concentrated liquidity pools
			default:
				fmt.Printf("Unknown pool type: %T\n", p)
			}
		}

		// Append the fetched pools to the list
		allPools = append(allPools, res.Pools...)

		fmt.Printf("Fetched page %d: %d pools\n", page, len(res.Pools))
		page++

		// Check if there's a next page
		if len(res.Pagination.NextKey) == 0 {
			break
		}

		nextKey = res.Pagination.NextKey
	}

	fmt.Printf("Total pools fetched: %d\n", len(allPools))

	// // Example: Print pool IDs
	// for _, pool := range allPools {
	// 	fmt.Printf("Pool ID: %d\n", pool.Id)
	// }

	// Further processing of pools can be done here
}
