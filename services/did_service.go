package services

import (
	"fmt"

	"github.com/your-username/did-example/utils"
)

// DIDService DID服务
type DIDService struct{}

var didService *DIDService

// InitDIDService 初始化DID服务
func InitDIDService() {
	didService = &DIDService{}
}

// GenerateDID 生成DID
func GenerateDID() (string, error) {
	// 生成公私钥对
	pubKey, privKey, err := utils.GenerateKeyPair()
	if err != nil {
		return "", fmt.Errorf("failed to generate key pair: %s", err.Error())
	}

	// 使用公钥生成DID
	did := utils.GenerateDID(pubKey)

	// 将DID和私钥存储在本地文件中
	err = utils.SaveKeyPair(did, pubKey, privKey)
	if err != nil {
		return "", fmt.Errorf("failed to save key pair: %s", err.Error())
	}

	return did, nil
}

// VerifyDID 验证DID是否有效
func VerifyDID(did string) bool {
	// 从本地文件中加载公钥
	pubKey, err := utils.LoadPubKey(did)
	if err != nil {
		return false
	}

	// 验证DID是否合法
	return utils.VerifyDID(did, pubKey)
}
