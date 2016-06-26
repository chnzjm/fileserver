package main

import (
    "flag"
    "fmt"
    "io"
    "net"
    "os"

    "google.golang.org/grpc"
    pb "github.com/chnzjm/fileserver/fileserver"
)

const (
    BUFSIZE = 1024
)

var (
    testFile = flag.String("test_file", "testdata/testdata", "test fiel")
    port = flag.Int("port", 10000, "The server prot")
)

type fileServer struct {}

func (fs *fileServer) GetFile(fd *pb.FileDescriptor, stream pb.File_GetFileServer) error {
    filename := fd.Filename

    fmt.Println(filename)
    // make sure file existed.
    fo, err := os.Open(filename)
    if err != nil {
        return err
    }

    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()

    content := &pb.FileContent {
        Content: make([]byte, BUFSIZE),
    }

    for {
        n, err := fo.Read(content.Content)
        fmt.Println(n)
        if err != nil && err != io.EOF {
            return err
        }
        if n == 0 {
            break
        }

        if err:= stream.Send(content); err != nil {
            return err
        }
    }

    //indicate the end of stream
    return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

    fileServer := new(fileServer)

    pb.RegisterFileServer(grpcServer, fileServer)
	grpcServer.Serve(lis)
}
