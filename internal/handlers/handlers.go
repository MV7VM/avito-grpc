package handlers

import (
	"avito/proto/proto/pb"
	"context"
	"fmt"
	"log/slog"
	"os"
)

type service interface {
	Create(msg string) (string, error)
	Get(id string) (string, error)
}

type Handlers struct {
	service
	pb.UnimplementedYourServiceServer
}

func New(service service) *Handlers {
	return &Handlers{service, pb.UnimplementedYourServiceServer{}}
}

//func (h *Handlers) Create(ctx context.Context, req *pb.CreateUrlRequest) (*pb.CreateUrlResponse, error) {
//	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
//	originalUrl := req.Url
//	fmt.Println("Post - ", originalUrl)
//	shortUrl, err := h.service.Create(originalUrl)
//	if err != nil {
//		return nil, status.Error(http.StatusBadRequest, "Bad request")
//		log.Error("Inserting problems" + err.Error())
//		//return nil, err
//	}
//	return &pb.CreateUrlResponse{ShortUrl: shortUrl}, nil
//}

func (h *Handlers) Create(ctx context.Context, request *pb.CreateUrlRequest) (*pb.CreateUrlResponse, error) {
	// Read the param noteId
	//link := c.Params("Link")
	//var shortUrl model.Links
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	longUrl := request.Url
	fmt.Println("Post - ", longUrl)
	msg, err := h.service.Create(longUrl)
	if err != nil {
		return nil, err
		log.Error("error in servise: " + err.Error())
	}
	return &pb.CreateUrlResponse{ShortUrl: msg}, nil
	// Find the note with the given id
	// Return the note with the id
}

func (h *Handlers) Get(ctx context.Context, request *pb.GetUrlRequest) (*pb.GetUrlResponse, error) {
	link := request.Url
	fmt.Println("Get - ", link)
	msg, err := h.service.Get(link)
	if err != nil {
		return nil, err
	}
	return &pb.GetUrlResponse{OriginalUrl: msg}, nil
}
