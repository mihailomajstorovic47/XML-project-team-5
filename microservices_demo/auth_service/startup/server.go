package startup

import (
	"context"
	"fmt"
	"github.com/mihailomajstorovic47/XML-project-team-5/microservices_demo/auth_service/application"
	"github.com/mihailomajstorovic47/XML-project-team-5/microservices_demo/auth_service/domain"
	"github.com/mihailomajstorovic47/XML-project-team-5/microservices_demo/auth_service/infrastructure/api"
	"github.com/mihailomajstorovic47/XML-project-team-5/microservices_demo/auth_service/infrastructure/persistence"
	"github.com/mihailomajstorovic47/XML-project-team-5/microservices_demo/auth_service/startup/config"
	auth "github.com/tamararankovic/microservices_demo/common/proto/auth_service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	credentialStore := server.initCredentialStore(mongoClient)

	authService := server.initAuthService(credentialStore)

	authHandler := server.initAuthHandler(authService)

	server.startGrpcServer(authHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AuthDBHost, server.config.AuthDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initCredentialStore(client *mongo.Client) domain.UsersStore {
	store := persistence.NewUserMongoDBStore(client)
	store.DeleteAll(context.TODO())
	//for _, user := range users {
	//	err, _ := store.Insert(context.TODO(), user)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	return store
}

func (server *Server) initAuthService(store domain.UsersStore) *application.AuthService {
	//profileServiceEndpoint := fmt.Sprintf("%s:%s", server.config.ProfileServiceHost, server.config.ProfileServicePort)
	return application.NewAuthService(store)
}

func (server *Server) initAuthHandler(service *application.AuthService) *api.UserHandler {
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(authHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, authHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}