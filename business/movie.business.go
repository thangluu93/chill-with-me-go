package business

import (
	"bytes"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
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

func NewMovie(db *mongo.Database) *Movie {
	var movie = access.NewMovie(db, "movies")
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

func (m *Movie) SplitMovie(folderPath string, file string) (movie *models.Movie, err error) {
	commandArray := []string{"-i", file, "-c:a", "libmp3lame", "-b:a", "128k", "-map", "0:0", "-f", "segment", "-segment_time", "10", "-segment_list", folderPath + "/" + folderPath + ".m3u8", "-segment_format", "mpegts", folderPath + "/output%03d.ts"}
	cmd := exec.Command("ffmpeg", commandArray...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	fmt.Printf("command output: %q\n", out.String())
	return nil, nil
}

func (m *Movie) createDirectory(movieId string) (d string, err error) {
	err = os.Mkdir(movieId, 0777)
	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}
	return movieId, nil
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

	_, err = m.SplitMovie(movieDir, fileDir)
	return err
}
