package client_auth

import (
	"context"
	"io"

	protos "github.com/signed-long/opaque-over-grpc/opaque-service-protos/protos"
	logrus "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Struct that implements interface protos.opaque-over-grpcServiceServer
// Where all rpc endpoints are implemented
type AuthServiceClient struct {
	client protos.OpaqueAuthServiceClient
	log    logrus.Logger
	protos.UnimplementedOpaqueAuthServiceServer
}

// AuthServiceClient constructor
func NewAuthServiceClient(client protos.OpaqueAuthServiceClient, log logrus.Logger) *AuthServiceClient {
	return &AuthServiceClient{client, log, protos.UnimplementedOpaqueAuthServiceServer{}}
}

func (svc *AuthServiceClient) RegisterFlow(
	username string, password string) error {

	svc.log.Debug("Starting register flow")
	src, err := svc.client.OpaqueRegistrationFlowRPC(context.Background())
	if err != nil {
		return err
	}

	svc.log.Debug("Sending INIT")
	b64UserRegInit, userReg, err := OpaqueRegisterInit(username, password)
	if err != nil {
		svc.log.Error("Error creating b64UserRegInit")
		return err
	}
	req := protos.RegistrationFlowMsg{
		Step:           protos.RegistrationFlowSteps_INIT,
		GopaqueTypeGob: b64UserRegInit,
		UserID:         username}
	err = src.Send(&req)
	if err == io.EOF {
		svc.log.Debug("Server closed connection")
		return err
	}
	if err != nil {
		svc.log.Error("Error sending INIT")
		return err
	}

	for {
		svc.log.Debug("Waiting for response from server")
		res, err := src.Recv()
		if err == io.EOF {
			svc.log.Debug("Server closed connection")
			return err
		}
		if err != nil {
			svc.log.Error("Unable to read from server: ", err)
			return err
		}

		switch res.GetStep() {
		case protos.RegistrationFlowSteps(
			protos.RegistrationFlowSteps_value["INIT_ACK"]):

			svc.log.Debug("INIT_ACK received")

			// OpaqueRegisterComplete
			if err != nil {
				return err
			}

			b64UserRegComplete, err := OpaqueRegisterComplete(userReg, res)

			svc.log.Debug("Sending COMPLETE")
			req := protos.RegistrationFlowMsg{Step: protos.RegistrationFlowSteps_COMPLETE, GopaqueTypeGob: b64UserRegComplete, UserID: res.GetUserID()}
			err = src.Send(&req)
			if err != nil {
				svc.log.Error("Error sending COMPLETE: ", err)
				return err
			}

		case protos.RegistrationFlowSteps(
			protos.RegistrationFlowSteps_value["COMPLETE_ACK"]):
			return nil

		default:
			svc.log.Error("Invalid step provided")
			err := status.Errorf(
				codes.FailedPrecondition,
				"Invalid step provided",
			)
			return err
		}
	}
}
