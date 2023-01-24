package file

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gowok/gowok/file"
)

// GetFileTypeFromBase64String func to get file type from base64 data url
func GetFileTypeFromBase64String(encodedBase64 string) string {
	_fileType := encodedBase64[0:strings.Index(encodedBase64, ";")]
	fileType := _fileType[strings.Index(_fileType, "/")+1:]
	return fileType
}

// GetMimeTypeFromBase64String func to get mime type from base64 data url
func GetMimeTypeFromBase64String(encodedBase64 string) string {
	return file.MimeTypes[fmt.Sprintf(".%s", GetFileTypeFromBase64String(encodedBase64))]
}

// SaveBase64StringToFile func to save base64 data to file
func SaveBase64StringToFile(path string, fileNameWithoutType string, encodedBase64 string) (string, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}

	fileType := GetFileTypeFromBase64String(encodedBase64)
	picturePath := path + "/" + fileNameWithoutType + "." + fileType
	encodedFileData := encodedBase64[strings.Index(encodedBase64, ",")+1:]

	pictureBase64Decoded := base64.NewDecoder(base64.StdEncoding, strings.NewReader(encodedFileData))

	profilePictureFile, err := os.Create(picturePath)
	if err != nil {
		return "", err
	}

	defer profilePictureFile.Close()

	_, err = io.Copy(profilePictureFile, pictureBase64Decoded)
	if err != nil {
		return "", err
	}

	return picturePath, nil
}
