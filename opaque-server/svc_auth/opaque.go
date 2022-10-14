package svc_auth

import (
	"encoding/base64"

	"github.com/cretz/gopaque/gopaque"
	protos "github.com/signed-long/opaque-over-grpc/opaque-service-protos/protos"
)

var crypto = gopaque.CryptoDefault

// Handles the logic of the OPAQUE protocol through the gopaque library.
// Steps 1 - 3 of the Server Registration flow from the gopaque documentation.
// 	 https://pkg.go.dev/github.com/cretz/gopaque/gopaque#hdr-Registration_Flow
// returns:
//	 b64ServerRegisterInit - base64 encoded ServerRegisterInit to send back to client.
// 	 serverReg - session to save for steps 4 - 5.
func (svc *AuthServiceServer) OpaqueRegisterInit(req *protos.RegistrationFlowMsg) (
	b64ServerRegisterInit string,
	serverReg gopaque.ServerRegister,
	err error) {

	// 1 - Receive the user's UserRegisterInit
	var userRegInit gopaque.UserRegisterInit
	bytes, err := base64.StdEncoding.DecodeString(req.GetGopaqueTypeGob())
	if err != nil {
		svc.log.Error("Error base64 decoding req.GopaqueTypeGob")
		return "", gopaque.ServerRegister{}, err
	}
	err = userRegInit.FromBytes(crypto, bytes)
	if err != nil {
		svc.log.Error("Error creating userRegInit FromBytes")
		return "", gopaque.ServerRegister{}, err
	}

	// 2 - Create a NewServerRegister with a private key
	serverReg = *gopaque.NewServerRegister(crypto, crypto.NewKey(nil))

	// 3 - Call Init with the user's UserRegisterInit and send the resulting ServerRegisterInit to the user
	serverRegInit := serverReg.Init(&userRegInit)
	serializedServerRegInit, err := serverRegInit.ToBytes()
	if err != nil {
		svc.log.Error("Error serializing serverRegInit")
		return "", gopaque.ServerRegister{}, err
	}
	b64ServerRegisterInit = base64.StdEncoding.EncodeToString(serializedServerRegInit)

	return b64ServerRegisterInit, serverReg, nil
}

// Handles the logic of the OPAQUE protocol through the gopaque library.
// Steps 4 - 5 of the Server Registration flow from the gopaque documentation.
// 	 https://pkg.go.dev/github.com/cretz/gopaque/gopaque#hdr-Registration_Flow
// returns:
// 	 regComplete - object to store for future authentication of the user.
func (svc *AuthServiceServer) OpaqueRegisterComplete(req *protos.RegistrationFlowMsg, serverReg gopaque.ServerRegister) (
	regComplete gopaque.ServerRegisterComplete,
	err error) {

	// 4 - Receive the user's UserRegisterComplete
	var userRegComplete gopaque.UserRegisterComplete
	bytes, err := base64.StdEncoding.DecodeString(req.GetGopaqueTypeGob())
	if err != nil {
		svc.log.Error("Error base64 decoding req.GopaqueTypeGob")
		return gopaque.ServerRegisterComplete{}, err
	}
	err = userRegComplete.FromBytes(crypto, bytes)
	if err != nil {
		svc.log.Error("Error creating userRegComplete FromBytes")
		return gopaque.ServerRegisterComplete{}, err
	}

	// 5 - Call Complete with the user's UserRegisterComplete and persist the resulting ServerRegisterComplete
	regComplete = *serverReg.Complete(&userRegComplete)

	return regComplete, nil
}
