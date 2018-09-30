k := kubectl -n yeo

tls:
	echo "Pre-copied tls.key and tls.crt into this directory"
	$(k) create secret tls bd-tls --key tls.key --cert tls.crt

secret:
	./secret.sh | $(k) apply -f -
