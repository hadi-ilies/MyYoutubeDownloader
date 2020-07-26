package main

import (
	"fmt"
	"os"
	"context"
	"net/http"
	"github.com/rylio/ytdl"
	"github.com/rs/zerolog/log"
	"encoding/json"
)

type ytDownloader struct {
	client ytdl.Client
	ctx context.Context
	vidInfo *ytdl.VideoInfo
}

func newYtDownloader() ytDownloader {
	return ytDownloader {
		client: ytdl.Client{
			HTTPClient: http.DefaultClient,
			Logger:     log.Logger,
		 },
		ctx: context.Background(),
	}
}

func (ytd *ytDownloader) loadVideoInfo(youtubeURL string) error {
	var err error

	ytd.vidInfo, err = ytd.client.GetVideoInfo(ytd.ctx, youtubeURL)
	if err != nil {
		fmt.Println("Failed to get video info")
		return err
	}
	return nil
}

//directory must have / at the end
func (ytd *ytDownloader) download(directory string, dlFormat uint) (string, error) {
	if int(dlFormat) > len(ytd.vidInfo.Formats) {
		return "", os.ErrExist
	}
	fileExtension := ".mp4"
	if ytd.vidInfo.Formats[dlFormat].Itag.Resolution == "" {
		fileExtension = ".mp3"
	}
	os.Mkdir(directory, 0755)
	filename := ytd.vidInfo.Title + fileExtension
	filePath := directory + filename
	file, err := os.Create(filePath)

	defer file.Close()
	err = ytd.client.Download(ytd.ctx, ytd.vidInfo, ytd.vidInfo.Formats[dlFormat], file)
	if err != nil {
		println("Failed to download the video")
		return "", err
	}
	return filename, nil
}

func (ytd *ytDownloader) printVideoInfo()  {
	println("\n__________________________AVAILABLE FORMAT VIDEO__________________________\n")

	for index, format := range(ytd.vidInfo.Formats) {
		print("-", index, "-")
		print(" {")

		if format.Itag.Resolution == "" {
			print("EXTENSION = ", ".mp3")
			print(", RESOLUTION = ", "Audio only")
		} else {
			print("EXTENSION = ", format.Itag.Extension)
			print(", RESOLUTION = ", format.Itag.Resolution)
		}

		print(", Video Encoding = ")
		if format.Itag.VideoEncoding == "" {
			print("None")
		} else {
			print(format.Itag.VideoEncoding)
		}

		print(", Audio bitRate = ")
		if format.Itag.AudioEncoding == "" {
			print("None")
		} else {
			print(format.Itag.AudioEncoding)
		}
		print(", Audio bitRate = ", format.Itag.AudioBitrate)
		print(", FPS = ", format.Itag.FPS)
		println("}")
	}
}

func (ytd *ytDownloader) getVideoInfo() []byte {
	b, err := json.Marshal(ytd.vidInfo.Formats)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return nil;
    }
	fmt.Println(string(b))
	return b
}
