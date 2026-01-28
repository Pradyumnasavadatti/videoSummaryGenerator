package internals

import (
	"net/http"
	"os"
	"strings"
)

func CheckFileMimeType(mediaType string, fileName string) (bool,error){
	filePtr, err := os.Open(fileName);
	if err != nil{
		return false,err
	}
	defer filePtr.Close()
	buf := make([]byte, 512)
	_,err = filePtr.Read(buf)
	if err!=nil{
		return false,err
	}
	mimeType := http.DetectContentType(buf)
	if mediaType == VideoType && strings.HasPrefix(mimeType, "video/"){
		return true, nil
	}else if mediaType == AudioType && strings.HasPrefix(mimeType, "audio/"){
		return true, nil
	}
	return false,nil;
}