package business

import (
	"bytes"
	"cloud.google.com/go/storage"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"io/ioutil"
	"main/access"
	"main/core"
	"main/data"
	"main/models"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
)

type Movie struct {
	MovieAccess *access.Movie
	Util        core.Utility
}

func NewMovie(db *mongo.Database, storage *storage.BucketHandle) *Movie {
	var movie = access.NewMovie(db, "movies", storage)
	return &Movie{
		MovieAccess: (*access.Movie)(movie),
		Util:        *core.UseUtil(),
	}
}

func (m *Movie) MovieList(page string, noRecord string, genre string) (movies []*models.Movie, err error) {

	pageInt, errorParse := m.Util.StringToInt(page)
	if errorParse != nil {
		pageInt = 1
	}

	noRecordInt, errorParse := m.Util.StringToInt(noRecord)
	if errorParse != nil {
		noRecordInt = data.DEFAULT_PAGE_SIZE
	}

	limit, offset := m.Util.GetLimitOffset(pageInt, noRecordInt)
	movies, err = m.MovieAccess.GetListMovies(limit, offset, genre)

	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (m *Movie) CreateMovie(movie *models.Movie) (newMovie *models.Movie, err error) {
	newMovie, err = m.MovieAccess.CreateMovie(movie)
	if err != nil {
		return nil, err
	}
	return newMovie, nil

}

func (m *Movie) UpdateMovie(movie *models.Movie) (updatedMovie *models.Movie, err error) {
	updatedMovie, err = m.MovieAccess.UpdateMovie(movie)
	if err != nil {
		return nil, err
	}
	return updatedMovie, nil
}

func (m *Movie) DeleteMovie(movie *models.Movie) (success bool, err error) {
	_, err = m.MovieAccess.UpdateMovie(movie)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *Movie) SplitMovie(folderPath string, file string) (err error) {
	commandArray := []string{"-i", file, "-c:a", "libmp3lame", "-b:a", "128k", "-map", "0:0", "-f", "segment", "-segment_time", "10", "-segment_list", folderPath + "/" + folderPath + ".m3u8", "-segment_format", "mpegts", folderPath + "/output%03d.ts"}
	cmd := exec.Command("ffmpeg", commandArray...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Printf("command output: %q\n", out.String())
	return nil
}

func (m *Movie) createDirectory(movieId string) (d string, err error) {
	err = os.Mkdir(movieId, 0777)
	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}
	return movieId, nil
}

func (m *Movie) uploadAllFileInFolder(directory string) error {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Println("Uploading file: ", info.IsDir(), info.Name())
			// read file as byte
			file, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			err = m.MovieAccess.UploadMovieToStorage(info.Name(), file)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *Movie) UploadMovie(movieId string, movie *multipart.FileHeader) error {
	movieDir, err := m.createDirectory(movieId)
	if err != nil {
		return err
	}
	src, err := movie.Open()
	if err != nil {
		return err
	}
	// write file and close then get directory
	dst, err := os.Create(filepath.Join(movieDir, filepath.Base(movieId+".mp4")))
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			return
		}
	}(dst)
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// get file directory when create file
	fileDir := dst.Name()

	err = m.SplitMovie(movieDir, fileDir)
	if err != nil {
		return err
	}
	return err
}

func (m *Movie) PlayMovie(id string) error {
	err := m.MovieAccess.DownloadMovieFromStorage(id)
	if err != nil {
		return err
	}
	return nil
}
