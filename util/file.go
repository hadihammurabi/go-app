package util

import (
	"encoding/base64"
	"io"
	"os"
	"strings"
)

//SaveBase64StringToFile func to save base64 data to file
func SaveBase64StringToFile(path string, fileNameWithoutType string, encodedBase64 string) (string, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}

	_fileType := encodedBase64[0:strings.Index(encodedBase64, ";")]
	fileType := _fileType[strings.Index(_fileType, "/")+1:]
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
