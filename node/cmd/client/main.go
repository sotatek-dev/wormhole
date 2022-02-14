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

	cc, err := grpc.Dial("klaytn-api1.sotatek.works/publicrpc.v1.PublicRPCService/GetSignedVAA", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while dial: %v", err)
	}
	defer cc.Close()

	c := publicrpcv1.NewPublicRPCServiceClient(cc)

	msg := publicrpcv1.GetSignedVAARequest{
		MessageId: &publicrpcv1.MessageID{
			EmitterChain:   publicrpcv1.ChainID(4097),
			EmitterAddress: "0000000000000000000000007a0c9a0e9e6c82bd25acf6e3f5437bce64738b3b",
			Sequence:       74,
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
