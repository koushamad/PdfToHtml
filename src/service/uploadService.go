package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	_ "mime/multipart"
	"os"
)

func (s *Service) Upload(c *gin.Context) (string, error) {
	file, _ := c.FormFile("file")
	path := s.getTmpPath(file.Filename)
	log.Println(path)
	if err := c.SaveUploadedFile(file, path); err != nil {
		return "", err
	} else {
		return s.MoveToQueue(path)
	}
}

func (s *Service) MoveToQueue(path string) (string, error) {
	if hash, err := s.makeHash(path); err != nil {
		return "", err
	} else {
		if queuePath := s.getQueuePath(hash); err != nil {
			return "", err
		} else {
			return hash, os.Rename(path, queuePath)
		}
	}
}

func (s *Service) getTmpPath(fileName string) string {
	return "./resource/tmp/" + fileName
}

func (s *Service) getQueuePath(hash string) string {
	return "./resource/queue/" + hash
}

func (s *Service) makeHash(path string) (string, error) {
	sha := sha256.New()

	if buffer, err := ioutil.ReadFile(path); err != nil {
		return "", err
	} else {
		sha.Write(buffer)
	}

	return hex.EncodeToString(sha.Sum(nil)), nil
}
