package guardiand

import (
	"context"
	"fmt"
	"github.com/certusone/wormhole/node/pkg/proto/publicrpc/v1"
	"github.com/certusone/wormhole/node/pkg/supervisor"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go.uber.org/zap"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"strings"
)

func allowCORSWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				corsPreflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func corsPreflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{
		"content-type",
		"accept",
		"x-user-agent",
		"x-grpc-web",
		"grpc-status",
		"grpc-message",
		"authorization",
	}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}

func publicwebServiceRunnable(
	logger *zap.Logger,
	listenAddr string,
	upstreamAddr string,
	grpcServer *grpc.Server,
	tlsHostname string,
	tlsProd bool,
	tlsCacheDir string,
) (supervisor.Runnable, error) {
	logger.Info("Tmp -------28213123----------- Start Webservice")
	return func(ctx context.Context) error {
		logger.Info("Start Webservice 00000000000000000000000000000")
		conn, err := grpc.DialContext(
			ctx,
			fmt.Sprintf("unix://%s", upstreamAddr),
			grpc.WithBlock(),
			grpc.WithInsecure())
		logger.Info("SConn: " + fmt.Sprint(conn))
		if err != nil {
			logger.Info("Dial eror: " + err.Error() + "$$$$$$$$$$$$$$$$$$$$$$$")
			return fmt.Errorf("failed to dial upstream: %s", err)
		}
		logger.Info("22222222222222222222222222222222222222xxxxxxxxxxx")
		gwmux := runtime.NewServeMux()
		err = publicrpcv1.RegisterPublicRPCServiceHandler(ctx, gwmux, conn)
		if err != nil {
			logger.Info("Dial eror: " + err.Error() + "$$$$$$$$$$$$$$$$$$$$$$$")
			panic(err)
		}
		logger.Info("33333333333333333333333qqqqqqqqqqqqqqqqqqq")
		mux := http.NewServeMux()
		grpcWebServer := grpcweb.WrapServer(grpcServer)
		mux.Handle("/", allowCORSWrapper(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			if grpcWebServer.IsGrpcWebRequest(req) {
				grpcWebServer.ServeHTTP(resp, req)
			} else {
				gwmux.ServeHTTP(resp, req)
			}
		})))

		srv := &http.Server{
			Handler: mux,
		}
		
		logger.Info("111111111111111111")
		// TLS setup
		if tlsHostname != "" {
			logger.Info("provisioning Let's Encrypt certificate", zap.String("hostname", tlsHostname))

			var acmeApi string
			if tlsProd {
				logger.Info("using production Let's Encrypt server")
				acmeApi = autocert.DefaultACMEDirectory
			} else {
				logger.Info("using staging Let's Encrypt server")
				acmeApi = "https://acme-staging-v02.api.letsencrypt.org/directory"
			}

			certManager := autocert.Manager{
				Prompt:     autocert.AcceptTOS,
				HostPolicy: autocert.HostWhitelist(tlsHostname),
				Cache:      autocert.DirCache(tlsCacheDir),
				Client:     &acme.Client{DirectoryURL: acmeApi},
			}

			srv.TLSConfig = certManager.TLSConfig()
			logger.Info("certificate provisioning configured")
		}

		var listener net.Listener

		// If listenAddr is prefixed by "sd:", look for a matching systemd socket.
		if strings.HasPrefix(listenAddr, "sd:") {
			logger.Info("22222222222222222")
			listeners, err := getSDListeners()
			if err != nil {
				return fmt.Errorf("failed to get systemd listeners: %w", err)
			}

			addr := listenAddr[3:]
			for _, v := range listeners {
				logger.Debug("found systemd socket", zap.String("addr", v.Addr().String()))
				if v.Addr().String() == addr {
					listener = v
				}
			}

			if listener == nil {
				all := make([]string, len(listeners))
				for i := range listeners {
					all[i] = listeners[i].Addr().String()
				}
				return fmt.Errorf("no valid systemd listeners, got: %s", strings.Join(all, ","))
			}
		} else {
			logger.Info("333333333333333333333")
			listener, err = net.Listen("tcp", listenAddr)
			if err != nil {
				return fmt.Errorf("failed to listen: %v", err)
			}
		}

		supervisor.Signal(ctx, supervisor.SignalHealthy)
		errC := make(chan error)
		go func() {
			logger.Info("publicweb server listening", zap.String("addr", listener.Addr().String()))
			if tlsHostname != "" {
				errC <- srv.ServeTLS(listener, "", "")
			} else {
				errC <- srv.Serve(listener)
			}
			logger.Info("Enddddddddddd")
		}()
		select {
		case <-ctx.Done():
			// non-graceful shutdown
			if err := srv.Close(); err != nil {
				return err
			}
			return ctx.Err()
		case err := <-errC:
			logger.Info("Error: " + err.Error())
			return err
		}
	}, nil
}
