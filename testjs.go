package main

import (
    "encoding/json"
    "log"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "project/proto"
)

const address = "localhost:50051"

type Category struct {
    Id              string  `json:"id"`
    Name            string  `json:"name"`
}

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewRpcClient(conn)

    cat := Category{Name: "Cat1"}
    b, err := json.Marshal(cat)

    r, err := c.Create(context.Background(), &pb.Request{Data: string(b[:])})
    log.Printf("status: %d response: %s\n", r.Status, r.Data)

    {
        r, _ := c.Get(context.Background(), &pb.Request{Data: "adasdsad"})
        log.Printf("DELETE status: %d response: %s\n", r.Status, r.Data)
    }

    var catt Category
    json.Unmarshal([]byte(r.Data), &catt)

    {
        catt.Name = "Dress"
        b, _ := json.Marshal(catt)
        r, _ := c.Update(context.Background(), &pb.Request{Data: string(b[:])})
        log.Printf("UPDATE status: %d response: %s\n", r.Status, r.Data)
    }

    {
        r, _ := c.Delete(context.Background(), &pb.Request{Data: string(catt.Id)})
        log.Printf("DELETE status: %d response: %s\n", r.Status, r.Data)
    }
}