package main

import (
	"io/ioutil"
	"log"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Querying all cars...")

	id, err := newIdentity()
	if err != nil {
		log.Fatalln(err)
	}

	sign, err := newSign()
	if err != nil {
		log.Fatalln(err)
	}

	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection, err := grpc.Dial("sampleorgpeer-api.127-0-0-1.nip.io:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	// Create a Gateway connection for a specific client identity
	gateway, err := client.Connect(id, client.WithSign(sign), client.WithClientConnection(clientConnection))
	if err != nil {
		log.Fatalln(err)
	}

	network := gateway.GetNetwork("mychannel")
	contract := network.GetContract("fabcar")

	result, err := contract.EvaluateTransaction("QueryAllCars")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(result))

	gateway.Close()
}

// NewIdentity creates a client identity for this Gateway connection using an X.509 certificate
func newIdentity() (*identity.X509Identity, error) {
	certificatePEM, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		return nil, err
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		return nil, err
	}

	return identity.NewX509Identity("SampleOrgMSP", certificate)
}

// NewSign creates a function that generates a digital signature from a message digest using a private key
func newSign() (identity.Sign, error) {
	privateKeyPEM, err := ioutil.ReadFile("key.pem")
	if err != nil {
		return nil, err
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		return nil, err
	}

	return identity.NewPrivateKeySign(privateKey)
}
