package Config

import "os"

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// MinIO Configuration
func GetMinIOConfig() (endpoint, accessKey, secretKey string, useSSL bool) {
	endpoint = getEnv("MINIO_ENDPOINT", "localhost:9000")
	accessKey = getEnv("MINIO_ACCESS_KEY", "minioadmin")
	secretKey = getEnv("MINIO_SECRET_KEY", "minioadmin")
	useSSL = getEnv("MINIO_USE_SSL", "false") == "true"
	return
}
