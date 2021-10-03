package facebook

import (
	"io"
)

type BinaryData struct {
	Filename    string
	Source      io.Reader
	ContentType string
}

type BinaryFile struct {
	Filename    string
	Path        string
	ContentType string
}

func Data(filename string, source io.Reader) *BinaryData {
	return &BinaryData{
		Filename: filename,
		Source:   source,
	}
}

func DataWithContentType(filename string, source io.Reader, contentType string) *BinaryData {
	return &BinaryData{
		Filename:    filename,
		Source:      source,
		ContentType: contentType,
	}
}

func File(filename string) *BinaryFile {
	return &BinaryFile{
		Filename: filename,
	}
}

func FileAlias(filename, path string) *BinaryFile {
	return &BinaryFile{
		Filename: filename,
		Path:     path,
	}
}

func FileAliasWithContentType(filename, path, contentType string) *BinaryFile {
	if path == "" {
		path = filename
	}

	return &BinaryFile{
		Filename:    filename,
		Path:        path,
		ContentType: contentType,
	}
}
