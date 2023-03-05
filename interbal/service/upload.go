package service

import (
	"blog-service/global"
	"blog-service/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (srv *Service) UploadFile(fileType upload.FileType, file multipart.File, header *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(header.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}

	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	savePath := upload.GetSavePath()
	if upload.CheckSavePath(savePath) {
		if err := upload.CreateSavePath(savePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}

	if upload.CheckPermission(savePath) {
		return nil, errors.New("insufficient file permissions")
	}

	dst := savePath + "/" + fileName
	if err := upload.SaveFile(header, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
