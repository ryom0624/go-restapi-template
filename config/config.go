package config

import "github.com/caarlos0/env/v6"

type Config struct {
	GoEnv                  string `env:"GO_ENV" envDefault:"localhost"`
	DBHost                 string `env:"DB_HOST" envDefault:"mysql"`
	DBName                 string `env:"DB_NAME" envDefault:"webapp_localhost"`
	DBUser                 string `env:"DB_USER" envDefault:"user"`
	DBPass                 string `env:"DB_PASS" envDefault:"pass"`
	GoogleProjectId        string `env:"GOOGLE_PROJECT_ID" envDefault:""`
	FirebaseCredentialJSON []byte `env:"FIREBASE_CREDENTIAL_FILE_PATH" envDefault:""`
	RecaptchaSecretKey     string `env:"RECAPTCHA_SECRET_KEY" envDefault:""`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
