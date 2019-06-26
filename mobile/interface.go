// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mobile

import (
	"encoding/hex"

	"github.com/lomocoin/HDWallet-Core/core"
)

type Wallet struct {
	mnemonic string
	seed     []byte
	testNet  bool
}

// MnemonicFromEntropy 根据 entropy， 获取对应助记词
func MnemonicFromEntropy(entropy string) (mnemonic string, err error) {
	entropyBytes, err := hex.DecodeString(entropy)
	if err != nil {
		return
	}
	return core.NewMnemonic(entropyBytes)
}

// EntropyFromMnemonic 根据 助记词, 获取 Entropy
// returns the input entropy used to generate the given mnemonic
func EntropyFromMnemonic(mnemonic string) (entropy string, err error) {
	entropyBytes, err := core.EntropyFromMnemonic(mnemonic)
	if err != nil {
		return
	}
	entropy = hex.EncodeToString(entropyBytes)
	return
}

// NewMnemonic 生成助记词
// 默认使用128位密钥，生成12个单词的助记词
func NewMnemonic() (mnemonic string, err error) {
	const bitSize = 128
	entropy, err := core.NewEntropy(bitSize)
	if err != nil {
		return
	}
	return core.NewMnemonic(entropy)
}

// ValidateMnemonic 验证助记词的正确性
func ValidateMnemonic(mnemonic string) (err error) {
	_, err = core.NewSeedFromMnemonic(mnemonic)
	if err != nil {
		return
	}
	return
}

// NewMnemonic 通过助记词得到一个 HD 对象
func NewHDWalletFromMnemonic(mnemonic string, testNet bool) (w *Wallet, err error) {
	seed, err := core.NewSeedFromMnemonic(mnemonic)
	if err != nil {
		return
	}
	w = new(Wallet)
	w.mnemonic = mnemonic
	w.seed = seed
	w.testNet = testNet
	return
}

// DeriveAddress 获取对应币种代号的 地址
func (c Wallet) DeriveAddress(symbol string) (address string, err error) {
	coin, err := c.initCoin(symbol)
	if err != nil {
		return
	}
	return coin.DeriveAddress()
}

// DerivePublicKey 获取对应币种代号的 公钥
func (c Wallet) DerivePublicKey(symbol string) (publicKey string, err error) {
	coin, err := c.initCoin(symbol)
	if err != nil {
		return
	}
	return coin.DerivePublicKey()
}

// DerivePrivateKey 获取对应币种代号的 私钥
func (c Wallet) DerivePrivateKey(symbol string) (privateKey string, err error) {
	coin, err := c.initCoin(symbol)
	if err != nil {
		return
	}
	return coin.DerivePrivateKey()
}

// DecodeTx 解析交易数据
// return: json 数据
func (c Wallet) DecodeTx(symbol, msg string) (tx string, err error) {
	coin, err := c.initCoin(symbol)
	if err != nil {
		return
	}
	return coin.DecodeTx(msg)
}

// Sign 签名交易
func (c Wallet) Sign(symbol, msg string) (sig string, err error) {
	coin, err := c.initCoin(symbol)
	if err != nil {
		return
	}

	privateKey, err := coin.DerivePrivateKey()
	if err != nil {
		return
	}

	return coin.Sign(msg, privateKey)
}