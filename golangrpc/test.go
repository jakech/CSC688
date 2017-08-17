package main

import (
    "encoding/json"
    "log"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "project/proto"
)

const address = "localhost:50051"

type Product struct {
    Id              string  `json:"id"`
    Name            string  `json:"name"`
    Description     string  `json:"desc"`
    Price           float64 `json:"price"`
    CatId           string  `json:"cat_id"`
}

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewRpcClient(conn)

    product := Product{Name: "Mug", Description: "nice big mug", Price: 12.99}
    b, err := json.Marshal(product)

    r, err := c.Create(context.Background(), &pb.Request{Data: string(b[:])})
    // if err != nil {
    //     log.Fatalf("could not greet: %v", err)
    // }
    log.Printf("status: %d response: %s\n", r.Status, r.Data)

    var rProduct Product
    json.Unmarshal([]byte(r.Data), &rProduct)
    // log.Printf("product id: %s", rProduct.Id)
    //
    // pr, err := c.Delete(context.Background(), &pb.Request{Data: rProduct.Id})
    
    rProduct.Name = "Tee"
    b, err = json.Marshal(rProduct)
    pr, err := c.Update(context.Background(), &pb.Request{Data: string(b[:])})

    // pr, err := c.Get(context.Background(), &pb.Request{Data: rProduct.Id})
    if err != nil {
        log.Printf(err.Error())
    } else {
        log.Printf("status: %d response: %s\n", pr.Status, pr.Data)
    }
}
