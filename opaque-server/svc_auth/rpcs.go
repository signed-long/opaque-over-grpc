package svc_auth

import (
	"io"

	"github.com/cretz/gopaque/gopaque"
	"github.com/google/uuid"
	protos "github.com/signed-long/opaque-over-grpc/opaque-service-protos/protos"
	logrus "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var registeredUsers = map[string]*gopaque.ServerRegisterComplete{}

// Struct that implements interface protos.opaque-over-grpcServiceServer
// Where all rpc endpoints are implemented
type AuthServiceServer struct {
	log logrus.Logger
	protos.UnimplementedOpaqueAuthServiceServer
}

// AuthServiceServer constructor
func NewAuthServiceServer(log logrus.Logger) *AuthServiceServer {
	return &AuthServiceServer{log, protos.UnimplementedOpaqueAuthServiceServer{}}
}

func (svc *AuthServiceServer) OpaqueRegistrationFlowRPC(
	src protos.OpaqueAuthService_OpaqueRegistrationFlowRPCServer) error {

	var registerSessions = map[string]*gopaque.ServerRegister{}
	var username string
	for {
		req, err := src.Recv()
		if err == io.EOF {
			svc.log.Info("Client closed connection")
			return err
		}
		if err != nil {
			svc.log.Error("Unable to read from client: ", err)
			return err
		}
		svc.log.Info("req.GetStep() " + req.GetStep().String())
		svc.log.Info("username" + username)

		switch req.GetStep() {
		case protos.RegistrationFlowSteps(
			protos.RegistrationFlowSteps_value["INIT"]):

			svc.log.Info("INIT request received")
			username = req.GetUserID()

			// Check if user already exists TODO: lookup in db
			if registeredUsers[username] != nil {
				errMsg := "User already Exists"
				svc.log.Error(errMsg)
				err := status.Errorf(
					codes.AlreadyExists,
					errMsg,
				)
				return err
			}

			b64ServerRegInit, serverReg, err := svc.OpaqueRegisterInit(req)
			if err != nil {
				return err
			}

			sessionID := uuid.New()
			registerSessions[sessionID.String()] = &serverReg
			svc.log.Info("Registration session created")

			svc.log.Debug("Sending INIT_ACK")
			res := protos.RegistrationFlowMsg{Step: protos.RegistrationFlowSteps_INIT_ACK, GopaqueTypeGob: b64ServerRegInit, UserID: sessionID.String()}
			err = src.Send(&res)
			if err != nil {
				return err
			}

		case protos.RegistrationFlowSteps(
			protos.RegistrationFlowSteps_value["COMPLETE"]):

			sessionID := req.GetUserID()

			svc.log.Info("sessionID" + sessionID)

			svc.log.Info("COMPLETE request received")
			// Check if there is a session for this user
			if registerSessions[sessionID] == nil {
				svc.log.Error("No registration session found")
				err := status.Errorf(
					codes.FailedPrecondition,
					"No registration session found",
				)
				return err
			}

			serverReg := *registerSessions[sessionID]
			regComplete, err := svc.OpaqueRegisterComplete(req, serverReg)
			if err != nil {
				svc.log.Error("Error completing registration:", err)
				return err
			}

			svc.log.Info("Storing new user")
			registeredUsers[username] = &regComplete
			delete(registerSessions, sessionID)
			svc.log.Info("User stored")

			svc.log.Debug("Sending COMPLETE_ACK")
			res := protos.RegistrationFlowMsg{Step: protos.RegistrationFlowSteps_COMPLETE_ACK}
			err = src.Send(&res)
			if err != nil {
				return err
			}
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
	return nil
}
