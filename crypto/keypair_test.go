package crypto

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeypair_Sign_Verify(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	msg := []byte("hello world")
	sig, err := privateKey.Sign(msg)
	require.NoError(t, err)
	require.True(t, sig.Verify(publicKey, msg))
	require.False(t, sig.Verify(publicKey, []byte("bye bye")))

	otherPrivateKey := GeneratePrivateKey()
	otherPublicKey := otherPrivateKey.PublicKey()
	require.False(t, sig.Verify(otherPublicKey, msg))
}
