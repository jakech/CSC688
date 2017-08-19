package server

import (
    "log"
    "errors"
    "encoding/json"
    "golang.org/x/net/context"
    "github.com/rs/xid"
    pb "project/proto"
)

type Product struct {
    Id              string  `json:"id"`
    Name            string  `json:"name"`
    Description     string  `json:"desc"`
    Price           float64 `json:"price"`
    CatId           string  `json:"cat_id"`
}

type Server struct{
    products map[string]*Product
}

func New() *Server {
    return &Server{make(map[string]*Product)}
}

func (s *Server) Create(ctx context.Context, req *pb.Request) (*pb.Response, error) {
    log.Printf("Create: %s\n", req.Data)
    var product Product
    guid := xid.New().String()

    err := json.Unmarshal([]byte(req.Data), &product)

    if err != nil {
        return &pb.Response{ Status: -1 }, errors.New("Data Error")
    }

    product.Id = guid
    s.products[guid] = &product

    b, err := json.Marshal(product)

    if err != nil {
        return &pb.Response{ Status: -1 }, errors.New("json built Error")
    }

    return &pb.Response{ Status: 1, Data: string(b[:]) }, nil
}

func (s *Server) Get(ctx context.Context, req *pb.Request) (*pb.Response, error) {
    log.Printf("Get: %s\n", req.Data)
    if product, ok := s.products[string(req.Data)]; ok {
        b, err := json.Marshal(product)

        if err != nil {
            return &pb.Response{ Status: -1 }, errors.New("json built Error")
        }

        return &pb.Response{ Status: 1, Data: string(b[:]) }, nil

    } else {
        return &pb.Response{ Status: 0, Data: "" }, nil
    }
}

func (s *Server) Update(ctx context.Context, req *pb.Request) (*pb.Response, error) {
    log.Printf("Update: %s\n", req.Data)
    var product Product
    err := json.Unmarshal([]byte(req.Data), &product)

    if err != nil {
        return &pb.Response{ Status: -1 }, errors.New("Data Error")
    }

    if _, ok := s.products[product.Id]; ok {
        // update
        s.products[product.Id] = &product

        b, err := json.Marshal(s.products[product.Id])

        if err != nil {
            return &pb.Response{ Status: -1 }, errors.New("json built Error")
        }

        return &pb.Response{ Status: 1, Data: string(b[:]) }, nil
    } else {
        return &pb.Response{ Status: 0, Data: "" }, nil
    }
}

func (s *Server) Delete(ctx context.Context, req *pb.Request) (*pb.Response, error) {
    log.Printf("Delete: %s\n", req.Data)
    if _, ok := s.products[string(req.Data)]; ok {
        delete(s.products, string(req.Data))
        return &pb.Response{ Status: 1, Data: "" }, nil
    } else {
        return &pb.Response{ Status: 0, Data: "" }, nil
    }
}

