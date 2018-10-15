k := kubectl -n yeo

tls:
	echo "Pre-copied tls.key and tls.crt into this directory"
	$(k) delete secret bd-tls || true
	$(k) create secret tls bd-tls --key betterdev.link.key --cert fullchain.cer

configmap:
	$(k) delete configmap members || true
	$(k) create configmap members --from-file=../content/prod.csv --from-file=../content/dev.csv

secret:
	./secret.sh | $(k) apply -f -
