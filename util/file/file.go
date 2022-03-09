package file

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

// GetFileTypeFromBase64String func to get file type from base64 data url
func GetFileTypeFromBase64String(encodedBase64 string) string {
	_fileType := encodedBase64[0:strings.Index(encodedBase64, ";")]
	fileType := _fileType[strings.Index(_fileType, "/")+1:]
	return fileType
}

// GetMimeTypeFromBase64String func to get mime type from base64 data url
func GetMimeTypeFromBase64String(encodedBase64 string) string {
	return MimeTypes[fmt.Sprintf(".%s", GetFileTypeFromBase64String(encodedBase64))]
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

var MimeTypes = map[string]string{
	".au":     "audio/basic",
	".avi":    "video/msvideo, video/avi, video/x-msvideo",
	".bmp":    "image/bmp",
	".bz2":    "application/x-bzip2",
	".css":    "text/css",
	".dtd":    "application/xml-dtd",
	".doc":    "application/msword",
	".exe":    "application/octet-stream",
	".gz":     "application/x-gzip",
	".gif":    "image/gif",
	".hqx":    "application/mac-binhex40",
	".html":   "text/html",
	".jar":    "application/java-archive",
	".jpg":    "image/jpeg",
	".js":     "application/x-javascript",
	".midi":   "audio/x-midi",
	".mp3":    "audio/mpeg",
	".mpeg":   "video/mpeg",
	".ogg":    "audio/vorbis, application/ogg",
	".pdf":    "application/pdf",
	".pl":     "application/x-perl",
	".png":    "image/png",
	".ppt":    "application/vnd.ms-powerpoint",
	".ps":     "application/postscript",
	".qt":     "video/quicktime",
	".ra":     "audio/x-pn-realaudio, audio/vnd.rn-realaudio",
	".ram":    "audio/x-pn-realaudio, audio/vnd.rn-realaudio",
	".rdf":    "application/rdf, application/rdf+xml",
	".rtf":    "application/rtf",
	".sgml":   "text/sgml",
	".sit":    "application/x-stuffit",
	".svg":    "image/svg+xml",
	".swf":    "application/x-shockwave-flash",
	".tar.gz": "application/x-tar",
	".tgz":    "application/x-tar",
	".tiff":   "image/tiff",
	".tsv":    "text/tab-separated-values",
	".txt":    "text/plain",
	".wav":    "audio/wav, audio/x-wav",
	".xls":    "application/vnd.ms-excel",
	".xml":    "application/xml",
	".zip":    "application/zip, application/x-compressed-zip",
}
