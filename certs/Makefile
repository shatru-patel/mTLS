cert:
	sudo openssl req -nodes -new -newkey rsa:2048 -keyout ca.key -x509 -sha256 -days 365 -out ca.crt -subj \
	"/C=IN/ST=Hyderabad/O=veltrisradius/CN=VeltrisCA/emailAddress=master@veltris.com"
	sudo openssl genrsa -out server.key 2048
	sudo openssl req -new -key server.key -out server.csr -config server.cnf
	sudo openssl req -noout -text -in server.csr
	sudo openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -sha256 -extfile server.cnf -extensions v3_ext
	sudo chmod 777 ca.crt
	sudo chmod 777 server.crt server.key
