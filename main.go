package main

import (
	"context"
	"log"
	"net"

	"github.com/rohit907/grpc-service/invoicer"
	"google.golang.org/grpc"
)

type myInvoiceServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (my myInvoiceServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  req.From,
		Docx: req.VAt,
	}, nil

}
func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Error while listening the port %s ", err)
	}
	server := grpc.NewServer()
	service := &myInvoiceServer{}
	invoicer.RegisterInvoicerServer(server, service)
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("server registery failed", err)
	}
}
