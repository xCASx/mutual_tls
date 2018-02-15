#!/usr/bin/env bash

if [ ! -f keys/client.pem ]; then
    cd keys

    openssl genrsa -out client.key 2048
    openssl req -new -key client.key -out client.csr -subj '/CN=localhost'

    openssl x509 -req -in client.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out client.pem -days 1024 -sha256

    openssl pkcs12 -export -inkey client.key -in client.pem -out client.pkcs12 -passout pass:password
    env pass=password keytool -import -file rootCA.pem -keystore truststore.jks -storepass:env pass -noprompt

    cd ..
fi

javac Client.java
java -Djavax.net.debug=ssl \
     -Djavax.net.ssl.keyStoreType=pkcs12 \
     -Djavax.net.ssl.keyStore=keys/client.pkcs12 \
     -Djavax.net.ssl.keyStorePassword=password \
     -Djavax.net.ssl.trustStoreType=jks \
     -Djavax.net.ssl.trustStore=keys/truststore.jks \
     -Djavax.net.ssl.trustStorePassword=password \
     -Dhttps.cipherSuites=TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA \
     -Djava.util.logging.config.file=logging.properties \
     Client
