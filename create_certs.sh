#!/bin/bash


# create CA's key,csr,cert
openssl genrsa -out ca.key 2048
openssl req -new -sha256 -key ca.key -out ca.csr -config openssl.cnf
openssl x509 -in ca.csr -days 365 -req -signkey ca.key -sha256 -out ca.crt -extfile ./openssl.cnf -extensions CA


# create Server's key,csr,cert
openssl genrsa -out server.key 2048
openssl req -new -nodes -sha256 -key server.key -out server.csr -config openssl.cnf
openssl x509 -req -days 365 -in server.csr -sha256 -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions Server


openssl genrsa -out client.key 2048
openssl req -new -nodes -sha256 -key client.key -out client.csr -config openssl.cnf
openssl x509 -req -days 365 -in client.csr -sha256 -out client.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions Client