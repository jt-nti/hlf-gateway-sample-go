# hlf-gateway-sample-go

Basic client for new [Fabric Gateway feature](https://hyperledger.github.io/fabric-rfcs/text/0000-fabric-gateway.html).

There is a [Gateway wiki page](https://github.com/hyperledgendary/hyperledgendary.github.io/wiki/Gateway) with some notes on using the new feature.

## Usage

Requires:
- `fabcar` contract to be deployed
- certificate and private key files named `cert.pem` and `key.pem` respectively

The following commands will retrieve the required certificate and private key when using the microfab environment described in the Wiki above:

```
docker cp microfab:/opt/microfab/data/admin-sampleorg/signcerts/cert.pem ${PWD}/cert.pem
docker cp microfab:/opt/microfab/data/admin-sampleorg/keystore/key.pem ${PWD}/key.pem
```

Run the sample using:

```
go run main.go
```
