package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type sourceDir string
type Extension []string

const SourceDir = sourceDir(".")
const TorrentDir = sourceDir("torrents")
const CompressedDir = sourceDir("compressed")
const ImageDir = sourceDir("images")
const MusicDir = sourceDir("musics")
const ProgramDir = sourceDir("programs")
const VideoDir = sourceDir("videos")
const DocsDir = sourceDir("docs")

func main() {
	for _, dir := range []string{string(CompressedDir), string(ImageDir), string(MusicDir), string(ProgramDir), string(DocsDir), string(SourceDir), string(TorrentDir)} {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("failed to create directory %s: %v\n", dir, err)
			// return
		}
	}

	compressedExtensions := Extension{".zip", ".7z", ".rar", ".tar"}
	torrentExtensions := Extension{".torrent"}
	imageExtensions := Extension{".jpg", ".jpeg", ".png", ".bmp", ".tiff", ".avif", ".jfif", ".gif", ".webp"}
	musicExtensions := Extension{".mp3"}
	videosExtensions := Extension{".mp4"}
	programExtensions := Extension{".exe", ".msi"}
	docsExtensions := Extension{".epub", ".pdf", ".doc", ".docx", ".txt", ".csv", ".xls", ".xlsx", ".html", ".md", ".mobi"}
	files, err := ioutil.ReadDir(string(SourceDir))
	if err != nil {
		fmt.Printf("error reading source directory : %v \n", err)
		return
	}

	moveFile := func(file os.FileInfo, targetDir string) {
		oldPath := filepath.Join(string(SourceDir), file.Name())
		newPath := filepath.Join(targetDir, file.Name())
		if err := os.Rename(oldPath, newPath); err != nil {
			fmt.Printf("error moving file %s to %s: %v\n", oldPath, newPath, err)
			return
		}
		fmt.Printf("Moved %s to %s\n", oldPath, newPath)
	}

	for _, file := range files {
		if !file.IsDir() {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			for _, f := range compressedExtensions {
				if ext == f {
					moveFile(file, string(CompressedDir))
					break
				}
			}

			for _, i := range videosExtensions {
				if ext == i {
					moveFile(file, string(VideoDir))
					break
				}
			}

			for _, i := range torrentExtensions {
				if ext == i {
					moveFile(file, string(TorrentDir))
					break
				}
			}

			for _, i := range musicExtensions {
				if ext == i {
					moveFile(file, string(MusicDir))
					break
				}
			}

			for _, i := range docsExtensions {
				if ext == i {
					moveFile(file, string(DocsDir))
					break
				}
			}

			for _, i := range imageExtensions {
				if ext == i {
					moveFile(file, string(ImageDir))
					break
				}
			}

			for _, i := range programExtensions {
				if ext == i && file.Name() != "sort.exe" {
					moveFile(file, string(ProgramDir))
					break
				}
			}

		}
	}
	fmt.Println("file sorting complete sir")

}
