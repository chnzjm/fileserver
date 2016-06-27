package main

import (
    "flag"
    "fmt"
    "io"
    "os"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "github.com/chnzjm/fileserver/fileserver"
)

var (
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func ReadFile(conn pb.FileClient, filename string) error {
  fd := &pb.FileDescriptor { Filename: "testdata/testdata" }
  stream, err := conn.GetFile(context.Background(), fd)
  if err != nil {
    panic(err)
  }

  outFd, err := os.Create("testdata.out")
  if err != nil {
    panic(err)
  }
  defer func() {
      if err := outFd.Close(); err != nil {
        panic(err)
      }
  }()

  for {
    content, err := stream.Recv()
    if err == io.EOF {
      break
    }
    if err != nil {
      panic(err)
    }

    if n, err := outFd.Write(content.Content); err != nil {
      panic(err)
    } else {
      fmt.Println(n)
    }
  }

  return nil
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
  opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		fmt.Printf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewFileClient(conn)

  ReadFile(client, "testdata/testdata")
}

