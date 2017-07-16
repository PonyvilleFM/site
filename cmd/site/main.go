package main

import (
	"errors"
	"flag"
	"net"
	"net/http"

	"github.com/PonyvilleFM/site/db"
	"github.com/PonyvilleFM/site/db/mockdb"
	"github.com/PonyvilleFM/site/schema"
	"github.com/Xe/ln"
	"github.com/facebookgo/flagenv"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	grpcPort = flag.String("grpc-port", "31337", "tcp port to listen on for grpc connections")
	port     = flag.String("port", "9090", "tcp port to listen the web server on")
)

func main() {
	flagenv.Parse()
	flag.Parse()

	gs := grpc.NewServer()

	schema.RegisterUsersServer(gs, &usersService{
		ud: mockdb.NewUserDao(),
	})

	l, err := net.Listen("tcp", ":"+*grpcPort)
	if err != nil {
		ln.FatalErr(err, ln.F{"action": "bind", "to": *grpcPort})
	}

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = schema.RegisterUsersHandlerFromEndpoint(context.Background(), mux, "127.0.0.1:"+*grpcPort, opts)
	if err != nil {
		ln.FatalErr(err, ln.F{"action": "registerAppsHandler"})
	}

	go http.ListenAndServe(":"+*port, mux)

	err = gs.Serve(l)
	if err != nil {
		ln.FatalErr(err)
	}
}

type usersService struct {
	ud db.UserDAOer
}

func (us *usersService) Register(ctx context.Context, rc *schema.RegisterCall) (*schema.User, error) {
	if rc.Password != rc.PasswordConfirm {
		return nil, errors.New("password not repeated correctly")
	}

	pretty.Println(rc)

	du, err := us.ud.Register(db.User{
		Username:      rc.User.Username,
		Email:         rc.User.Email,
		IsAdmin:       rc.User.IsAdmin,
		IsDJ:          rc.User.IsDj,
		TwitterHandle: rc.User.TwitterHandle,
	}, rc.Password)
	if err != nil {
		return nil, err
	}

	return &schema.User{
		Username:      du.Username,
		Email:         du.Email,
		IsAdmin:       du.IsAdmin,
		IsDj:          du.IsDJ,
		TwitterHandle: du.TwitterHandle,
	}, nil
}

func (us *usersService) Login(ctx context.Context, lc *schema.LoginCall) (*schema.User, error) {
	ok, du, err := us.ud.Login(lc.Username, lc.Password)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("invalid credentials")
	}

	return &schema.User{
		Username:      du.Username,
		Email:         du.Email,
		IsAdmin:       du.IsAdmin,
		IsDj:          du.IsDJ,
		TwitterHandle: du.TwitterHandle,
	}, nil
}

func (us *usersService) Info(ctx context.Context, ic *schema.InfoCall) (*schema.User, error) {
	du, err := us.ud.Info(ic.Username)
	if err != nil {
		return nil, err
	}

	return &schema.User{
		Username:      du.Username,
		Email:         du.Email,
		IsAdmin:       du.IsAdmin,
		IsDj:          du.IsDJ,
		TwitterHandle: du.TwitterHandle,
	}, nil
}
