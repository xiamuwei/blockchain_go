package crypto

import (
	"blockchain_go/types"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

type Privatkey struct {
	key *ecdsa.PrivateKey
}

func (p Privatkey) Sign(date []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, p.key, date)
	if err != nil {
		return nil, err
	}
	return &Signature{r, s}, nil
}

func GeneratePrivateKey() Privatkey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return Privatkey{
		key: key,
	}
}

func (k Privatkey) PublicKey() PublicKey {
	return PublicKey{
		key: &k.key.PublicKey,
	}
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (k PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(k.key, k.key.X, k.key.Y)
}

func (k PublicKey) Address() types.Address {
	h := sha256.Sum256(k.ToSlice())
	return types.AddressFromBytes(h[len(h)-20:])
}

type Signature struct {
	r, s *big.Int
}

func (sign Signature) Verify(publicKey PublicKey, date []byte) bool {
	return ecdsa.Verify(publicKey.key, date, sign.r, sign.s)
}
