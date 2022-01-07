package main

import (
	"context"
	"encoding/hex"
	"fmt"
	publicrpcv1 "github.com/certusone/wormhole/node/pkg/proto/publicrpc/v1"
	"github.com/certusone/wormhole/node/pkg/vaa"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"
	"log"
)

func main() {

	cc, err := grpc.Dial("localhost:7070", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while dial: %v", err)
	}
	defer cc.Close()

	c := publicrpcv1.NewPublicRPCServiceClient(cc)

	msg := publicrpcv1.GetSignedVAARequest{
		MessageId: &publicrpcv1.MessageID{
			EmitterChain:   publicrpcv1.ChainID(6),
			EmitterAddress: "000000000000000000000000d734fe34d7904804a64a7f8a91e21bbcdf6603e3",
			Sequence:       1,
		},
	}

	resp, err := c.GetSignedVAA(context.Background(), &msg)
	if err != nil {
		log.Fatalf("failed to run GetSignedVAA RPC: %v", err)
	}

	v, err := vaa.Unmarshal(resp.VaaBytes)
	if err != nil {
		log.Fatalf("failed to decode VAA: %v", err)
	}

	log.Printf("VAA with digest %s: %+v\n", v.HexDigest(), spew.Sdump(v))
	fmt.Println("index: ", v.GuardianSetIndex)
	fmt.Println("signs: ", v.Signatures)
	fmt.Println("index")
	str := hex.EncodeToString(resp.VaaBytes)
	fmt.Println(str)

}
