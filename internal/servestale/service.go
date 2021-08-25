/*
	servestale是一种降级手段，当服务端无法正常返回结果的时候，将缓存的旧数据返回给客户使用。
	另外遇到熔断的情况，也可以使用缓存的旧数据返回
	这里servestale的缓存数据放在AWS S3上，增加一个grpc服务从S3上取得结果返回。

	servestale作为sidecar发布并且和上游服务使用相同的api接口，可以比较容易的在上游服务崩溃的时候替换。
	而且在做压力测试的时候可以不用上游团队配合。
	客户端不需要做任何特殊的处理
*/
package servestale

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/vrealzhou/geekbang_go_final/api"
	"github.com/vrealzhou/geekbang_go_final/internal/log"
	api "github.com/vrealzhou/geekbang_go_final/internal/servestale/api"
)

// server is used to implement helloworld.GreeterServer.
type MusicService struct {
	pb.UnimplementedMusicServer
	bucket   *string
	s3client *s3.Client
}

func (s *MusicService) GetArtist(ctx context.Context, arid *wrapperspb.StringValue) (*pb.Artist, error) {
	// fetch from S3 and convert to *pb.Artist
	resp, err := s.s3client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: s.bucket,
		Key:    aws.String("music/artist/" + arid.Value),
	})
	if err != nil {
		log.Errorf("failed when get artist %s", arid.Value)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("failed when read artist %s", arid.Value)
	}
	artist := &pb.Artist{}
	err = proto.Unmarshal(b, artist)
	// need to do more research see if there are otherways to avoid encoding/decoding operation
	return artist, err
}

func (s *MusicService) GetRelease(context.Context, *wrapperspb.StringValue) (*pb.Release, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelease not implemented")
}
func (s *MusicService) GetRecording(context.Context, *wrapperspb.StringValue) (*pb.Recording, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecording not implemented")
}

func (s *MusicService) GetArtwork(stream pb.Music_GetArtworkServer) error {
	return status.Errorf(codes.Unimplemented, "method GetArtwork not implemented")
}

// 使用grpc的stream批量更新
func (s *MusicService) UpdateArtist(stream api.Music_UpdateArtistServer) error {
	for {
		artist, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&emptypb.Empty{})
		}
		if err != nil {
			return err
		}
		if artist != nil {
			b, err := proto.Marshal(artist)
			if err != nil {
				log.Errorf("failed when marshal artist %s", artist.Arid)
			}
			_, err = s.s3client.PutObject(context.Background(), &s3.PutObjectInput{
				Bucket: s.bucket,
				Key:    aws.String("music/artist/" + artist.Arid),
				Body:   bytes.NewReader(b),
			})
			if err != nil {
				log.Errorf("failed when store artist %s", artist.Arid)
			}
		}
	}
}
func (s *MusicService) UpdateRelease(stream api.Music_UpdateReleaseServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateRelease not implemented")
}
func (s *MusicService) UpdateRecording(stream api.Music_UpdateRecordingServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateRecording not implemented")
}
func (s *MusicService) UpdateArtwork(stream api.Music_UpdateArtworkServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateArtwork not implemented")
}
