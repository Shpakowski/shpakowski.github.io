package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"time"

	"math/big"
	"mcp-coop-chain/internal/types"

	"crypto/aes"
	"crypto/cipher"
	"errors"

	"github.com/btcsuite/btcutil/base58"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/pbkdf2"
)

// CreateWallet генерирует новый ECDSA-ключ, строит адрес и возвращает приватный кошелёк (PrivateWallet).
func CreateWallet() (*types.PrivateWallet, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	privBytes, _ := x509.MarshalECPrivateKey(priv)
	pubBytes := elliptic.MarshalCompressed(priv.Curve, priv.PublicKey.X, priv.PublicKey.Y)
	pubBase58 := base58.Encode(pubBytes)
	address := GetWalletAddress(pubBase58)
	return &types.PrivateWallet{
		PrivateKey: base64.StdEncoding.EncodeToString(privBytes),
		PublicKey:  pubBase58,
		Address:    address,
		CreatedAt:  time.Now().UTC(),
	}, nil
}

// ToPublicWallet возвращает публичную часть кошелька для блокчейна/снапшота.
func ToPublicWallet(priv *types.PrivateWallet) types.Wallet {
	return types.Wallet{
		Address:   priv.Address,
		PublicKey: priv.PublicKey,
	}
}

// SignTransactionWithWallet подписывает транзакцию приватным ключом кошелька.
func SignTransactionWithWallet(tx *types.Transaction, privWallet *types.PrivateWallet) error {
	privBytes, err := base64.StdEncoding.DecodeString(privWallet.PrivateKey)
	if err != nil {
		return err
	}
	priv, err := x509.ParseECPrivateKey(privBytes)
	if err != nil {
		return err
	}
	hash := sha256.Sum256(tx.Payload)
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		return err
	}
	sig := append(r.Bytes(), s.Bytes()...)
	tx.Signature = base64.StdEncoding.EncodeToString(sig)
	return nil
}

// VerifyWallet проверяет подпись транзакции по публичному ключу.
func VerifyWallet(tx *types.Transaction, pubKey string) bool {
	pubBytes := base58.Decode(pubKey)
	if len(pubBytes) == 0 {
		return false
	}
	x, y := elliptic.UnmarshalCompressed(elliptic.P256(), pubBytes)
	if x == nil || y == nil {
		return false
	}
	pub := &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}
	hash := sha256.Sum256(tx.Payload)
	sig, err := base64.StdEncoding.DecodeString(tx.Signature)
	if err != nil || len(sig) < 64 {
		return false
	}
	r := new(big.Int).SetBytes(sig[:len(sig)/2])
	s := new(big.Int).SetBytes(sig[len(sig)/2:])
	return ecdsa.Verify(pub, hash[:], r, s)
}

// SaveWallet сериализует приватный кошелёк в JSON и сохраняет в файл с шифрованием приватного ключа.
// Требует пароль пользователя. Без пароля приватный ключ не может быть восстановлен.
func SaveWallet(privWallet *types.PrivateWallet, path string, password string) error {
	privBytes, err := base64.StdEncoding.DecodeString(privWallet.PrivateKey)
	if err != nil {
		return err
	}
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return err
	}
	key := deriveKey(password, salt)
	ciphertext, nonce, err := encrypt(privBytes, key)
	if err != nil {
		return err
	}
	privWallet.PrivateKey = base64.StdEncoding.EncodeToString(ciphertext)
	privWallet.Salt = base64.StdEncoding.EncodeToString(salt)
	privWallet.AESNonce = base64.StdEncoding.EncodeToString(nonce)
	data, err := json.MarshalIndent(privWallet, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0600)
}

// LoadWallet загружает приватный кошелёк из файла JSON и расшифровывает приватный ключ по паролю.
// При ошибке пароля или повреждении файла возвращает понятную ошибку.
func LoadWallet(path string, password string) (*types.PrivateWallet, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var privWallet types.PrivateWallet
	if err := json.Unmarshal(data, &privWallet); err != nil {
		return nil, err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(privWallet.PrivateKey)
	if err != nil {
		return nil, errors.New("ошибка декодирования приватного ключа")
	}
	salt, err := base64.StdEncoding.DecodeString(privWallet.Salt)
	if err != nil {
		return nil, errors.New("ошибка декодирования соли")
	}
	nonce, err := base64.StdEncoding.DecodeString(privWallet.AESNonce)
	if err != nil {
		return nil, errors.New("ошибка декодирования nonce")
	}
	key := deriveKey(password, salt)
	plain, err := decrypt(ciphertext, key, nonce)
	if err != nil {
		return nil, errors.New("неверный пароль или повреждённый файл кошелька")
	}
	privWallet.PrivateKey = base64.StdEncoding.EncodeToString(plain)
	// Валидация: адрес должен совпадать с пересчитанным из публичного ключа
	addr := GetWalletAddress(privWallet.PublicKey)
	if privWallet.Address != addr {
		return nil, errors.New("адрес кошелька не совпадает с публичным ключом (возможно, файл повреждён)")
	}
	return &privWallet, nil
}

// GetWalletAddress строит адрес кошелька из публичного ключа (SHA256 + base58).
func GetWalletAddress(pubKey string) string {
	hash := sha256.Sum256([]byte(pubKey))
	return base58.Encode(hash[:])
}

// GetBalance — заглушка, возвращает фиксированный баланс.
// В будущем будет обращаться к базе данных или состоянию блокчейна.
func GetBalance(address string) float64 {
	// TODO: реализовать обращение к базе данных или хранилищу балансов
	return 0
}

// deriveKey генерирует ключ для AES-GCM из пароля и соли
func deriveKey(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, 100_000, 32, sha256.New)
}

// encrypt шифрует данные с помощью AES-GCM
func encrypt(plain, key []byte) (ciphertext, nonce []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}
	nonce = make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, nil, err
	}
	ciphertext = gcm.Seal(nil, nonce, plain, nil)
	return ciphertext, nonce, nil
}

// decrypt расшифровывает данные с помощью AES-GCM
func decrypt(ciphertext, key, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plain, nil
}

// IsValidAddress проверяет, что адрес валиден (base58, длина 44 символа, не пустой)
func IsValidAddress(address string) bool {
	if address == "" {
		return false
	}
	decoded := base58.Decode(address)
	// SHA256 = 32 байта, base58 = 44 символа
	if len(decoded) != 32 {
		return false
	}
	return true
}

// CreateWalletWithMnemonic создаёт приватный кошелёк по сид-фразе (или генерирует новую).
// Если mnemonic пустой — генерируется новая фраза, иначе используется переданная.
func CreateWalletWithMnemonic(mnemonic string) (*types.PrivateWallet, string, error) {
	var err error
	if mnemonic == "" {
		entropy, err := bip39.NewEntropy(128)
		if err != nil {
			return nil, "", err
		}
		mnemonic, err = bip39.NewMnemonic(entropy)
		if err != nil {
			return nil, "", err
		}
	}
	seed := bip39.NewSeed(mnemonic, "")
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, "", err
	}
	// Используем masterKey.Key как приватный ключ для ECDSA
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	priv.D.SetBytes(masterKey.Key)
	priv.PublicKey.X, priv.PublicKey.Y = elliptic.P256().ScalarBaseMult(masterKey.Key)
	privBytes, _ := x509.MarshalECPrivateKey(priv)
	pubBytes := elliptic.MarshalCompressed(priv.Curve, priv.PublicKey.X, priv.PublicKey.Y)
	pubBase58 := base58.Encode(pubBytes)
	address := GetWalletAddress(pubBase58)
	return &types.PrivateWallet{
		PrivateKey: base64.StdEncoding.EncodeToString(privBytes),
		PublicKey:  pubBase58,
		Address:    address,
		CreatedAt:  time.Now().UTC(),
	}, mnemonic, nil
}
