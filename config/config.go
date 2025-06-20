package config

import (
    "os"
)

func LoadEnv() {
    // Normally you'd load from a .env file
    os.Setenv("PORT", "8080")
    os.Setenv("JWT_SECRET", "supersecret")
    os.Setenv("DB_URL", "host=localhost user=postgres password=postgres dbname=clinic port=5432 sslmode=disable")
}

func EnvPort() string {
    return os.Getenv("PORT")
}

func EnvDB() string {
    return os.Getenv("DB_URL")
}

func EnvJWTSecret() string {
    return os.Getenv("JWT_SECRET")
}