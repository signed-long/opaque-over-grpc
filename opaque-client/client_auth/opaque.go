package client_auth

import (
	"encoding/base64"

	"github.com/cretz/gopaque/gopaque"
	protos "github.com/signed-long/opaque-over-grpc/opaque-service-protos/protos"
)

var crypto = gopaque.CryptoDefault

// Handles the logic of the OPAQUE protocol through the gopaque library.
// Steps 1 - 3 of the Client Registration flow from the gopaque documentation.
// 	 https://pkg.go.dev/github.com/cretz/gopaque/gopaque#hdr-Registration_Flow
// returns:
//	 b64ServerRegisterInit - base64 encoded ServerRegisterInit to send back to client.
func OpaqueRegisterInit(email string, password string) (
	b64UserRegInit string,
	userReg gopaque.UserRegister,
	err error) {

	// 1 - Create a NewUserRegister with the user ID
	userReg = *gopaque.NewUserRegister(crypto, []byte(email), nil)

	// 2 - Call Init with the password and send the resulting UserRegisterInit to the server
	userRegInit := userReg.Init([]byte(password))
	serializedUserRegInit, err := userRegInit.ToBytes()
	if err != nil {
		return "", userReg, err
	}
	b64UserRegInit = base64.StdEncoding.EncodeToString(serializedUserRegInit)

	return b64UserRegInit, userReg, nil
}

// Handles the logic of the OPAQUE protocol through the gopaque library.
// Steps 4 - 5 of the Server Registration flow from the gopaque documentation.
// 	 https://pkg.go.dev/github.com/cretz/gopaque/gopaque#hdr-Registration_Flow
// returns:
// 	 regComplete - object to store for future authentication of the user.
func OpaqueRegisterComplete(userReg gopaque.UserRegister, res *protos.RegistrationFlowMsg) (
	b64UserRegComplete string,
	err error) {

	// 3 - Receive the server's ServerRegisterInit
	var serverRegInit gopaque.ServerRegisterInit
	bytes, err := base64.StdEncoding.DecodeString(res.GetGopaqueTypeGob())
	if err != nil {
		return "", err
	}
	err = serverRegInit.FromBytes(crypto, bytes)
	if err != nil {
		return "", err
	}

	// 4 - Call Complete with the server's ServerRegisterInit and send the resulting UserRegisterComplete to the server
	userRegComplete := userReg.Complete(&serverRegInit)
	serializedUserRegComplete, err := userRegComplete.ToBytes()
	if err != nil {
		return "", err
	}
	b64UserRegComplete = base64.StdEncoding.EncodeToString(serializedUserRegComplete)

	return b64UserRegComplete, nil
}
