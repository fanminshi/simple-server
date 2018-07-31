# Generate TLS asset

Generate a self signed CA `ca-key.pen` and `ca.pem`:

`cfssl gencert -initca ca-csr.json | cfssljson -bare ca -`

Create a self signed `server-key.pem` and `server.pem`:

`$ cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=server server.json | cfssljson -bare server`

Create a self signed `client-key.pem` and `client.pem`:

`$ cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=client client.json | cfssljson -bare client`