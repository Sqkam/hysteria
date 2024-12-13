gen:
	docker rmi sqkam/hysteria:latest || true
	python hyperbole.py build -r
	chmod +x build/hysteria-linux-amd64
	docker build -t sqkam/hysteria:latest .
