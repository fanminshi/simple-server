# Generate TLS asset

The first, create a self signed `server-key.pem` and `server.pem`:

`$ openssl req -x509 -nodes -days 730 -newkey rsa:2048 -keyout server-key.pem -out server.pem -config req.conf -extensions 'v3_req'`

Finally, create a TLS secret:

`$ kubectl create secret tls simple-server-tls --key server-key.pem --cert server.pem`