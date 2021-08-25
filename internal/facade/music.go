package facade

import (
	"context"
	"io"
	"sync"

	"github.com/pkg/errors"
	pb "github.com/vrealzhou/geekbang_go_final/api"
	"github.com/vrealzhou/geekbang_go_final/internal/log"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type MusicService struct {
	client pb.MusicClient
}

func (s *MusicService) GetArtist(ctx context.Context, arid string) (*Artist, error) {
	a, err := s.client.GetArtist(ctx, &wrapperspb.StringValue{Value: arid})
	if err != nil {
		return nil, errors.Wrapf(err, "failed when fetching artist %s", arid)
	}
	artist := &Artist{}
	convertArtist(a, artist)
	artist.Artworks, err = s.fetchArtworks(ctx, a.Artworks, artist.Artworks)
	return artist, err
}

func (s *MusicService) fetchArtworks(ctx context.Context, artworkIDs []string, list []Artwork) ([]Artwork, error) {
	stream, err := s.client.GetArtwork(ctx)
	if err != nil {
		log.Errorf("failed when fetching artworks: %s", err.Error()) // ignore failed artworks with a log
	} else {
		wg := sync.WaitGroup{}
		go func() {
			i := 0
			for {
				in, err := stream.Recv()
				if err == io.EOF {
					list = list[:i]
					wg.Done()
					return
				}
				if err != nil {
					log.Errorf("failed when fetching artworks: %s", err.Error()) // ignore failed artworks with a log
				}
				convertArtwork(in, &list[i])
				i++
			}
		}()
		wg.Add(1)
		for _, artworkID := range artworkIDs {
			if err := stream.Send(&wrapperspb.StringValue{Value: artworkID}); err != nil {
				log.Errorf("failed when fetching artwork %s: %s", artworkID, err.Error()) // ignore failed artworks with a log
			}
		}
		stream.CloseSend()
		wg.Wait()
	}
	return list, nil
}

func (s *MusicService) GetRecording(ctx context.Context, arid string) (*Recording, error) {
	return nil, nil
}

func (s *MusicService) GetRelease(ctx context.Context, arid string) (*Release, error) {
	return nil, nil
}

type Artwork struct {
	Arid   string `json:"arid,omitempty"`
	URL    string `json:"url,omitempty"`
	Title  string `json:"title,omitempty"`
	Width  uint32 `json:"width,omitempty"`
	Height uint32 `json:"height,omitempty"`
	Size   uint32 `json:"size,omitempty"`
}

func convertArtwork(pbData *pb.Artwork, artwork *Artwork) {
	artwork.Arid = pbData.Arid
	artwork.URL = pbData.Url
	artwork.Title = pbData.Title
	artwork.Width = pbData.Width
	artwork.Height = pbData.Height
	artwork.Size = pbData.Size
}

type Artist struct {
	Arid     string    `json:"arid,omitempty"`
	Name     string    `json:"name,omitempty"`
	Type     string    `json:"type,omitempty"`
	Artworks []Artwork `json:"artworks,omitempty"`
}

func convertArtist(pbData *pb.Artist, artist *Artist) {
	artist.Arid = pbData.Arid
	artist.Name = pbData.Name
	artist.Type = pbData.Type
	artist.Artworks = make([]Artwork, len(pbData.Artworks))
}

type Recording struct {
	Arid     string    `json:"arid,omitempty"`
	Title    string    `json:"title,omitempty"`
	Duration uint32    `json:"duration,omitempty"`
	Metadata string    `json:"metadata,omitempty"`
	Artists  []Artist  `json:"artists,omitempty"`
	Releases []string  `json:"releases,omitempty"`
	Artworks []Artwork `json:"artworks,omitempty"`
}

type Release struct {
	Arid        string    `json:"arid,omitempty"`
	Title       string    `json:"title,omitempty"`
	Year        uint32    `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	RecordLabel string    `protobuf:"bytes,4,opt,name=recordLabel,proto3" json:"recordLabel,omitempty"`
	Artists     []Artist  `json:"artists,omitempty"`
	Recordings  []string  `json:"recordings,omitempty"`
	Artworks    []Artwork `json:"artworks,omitempty"`
}
