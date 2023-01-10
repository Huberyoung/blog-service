package service

import (
	"errors"
	"mime/multipart"
	"os"
	"service/global"
	"service/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (s *Service) UploadFile(fileType upload.FileType, file multipart.File, header *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(header.Filename)
	savePath := upload.GetSavePath()
	dst := savePath + "/" + fileName

	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix not supported")
	}

	if upload.CheckSavePath(savePath) {
		err := upload.CreateSavePath(savePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}

	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	if upload.CheckPermission(savePath) {
		return nil, errors.New("insufficient file permissions")
	}

	if err := upload.SaveFile(header, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
