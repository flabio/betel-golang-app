package patters

type PasswordProtector struct {
	password                string
	cryptoAlgorithmPassword CryptoAlgorithmPassword
}

type CryptoAlgorithmPassword interface {
	Hash(pwd []byte)
}

func NewPasswordProtector(password string, hash CryptoAlgorithmPassword) *PasswordProtector {
	return &PasswordProtector{
		password:                password,
		cryptoAlgorithmPassword: hash,
	}
}
