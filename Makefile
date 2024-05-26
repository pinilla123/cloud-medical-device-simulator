build-MedicalDeviceFunction:
	GOOS=linux go build -o main main.go || (echo "Build failed"; exit 1)
	ls -l main || (echo "main binary not found"; exit 1)
	zip function.zip main || (echo "Zip creation failed"; exit 1)
	ls -l function.zip || (echo "Zip file not found"; exit 1)