package moebooru

import (
	"bytes"
	"encoding/json"
	//"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"

	"github.com/ayame-git/gobooru/auth/moebooru"
	"github.com/ayame-git/gobooru/model/moebooru"
)

//TODO: request constatns (e.g. /post.json artist.json)

type API struct {
	Booru       string
	AuthHandler moebooruAuth.AuthHandler
}

//TODO: fix request over https
func (api *API) GetPosts(r moebooruModel.PostListRequest) (posts []moebooruModel.Posts, err error) {
	data, err := json.Marshal(r)
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", api.Booru+"/post.json", bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := api.AuthHandler.Client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&posts)

	return
}

//Handle responses with custom errors i.e. add abstraction
func (api *API) CreatePost(r moebooruModel.PostCreateRequest) error {
	/*if !api.AuthHandler.IsAuth {
		return errors.New("requires authorization!!")
	}*/

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	v := reflect.ValueOf(r)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Interface() != reflect.Zero(field.Type()).Interface() {
			if v.Type().Field(i).Tag.Get("form") == "post[file]" {
				f, err := os.Open(r.File)
				if err != nil {
					return err
				}

				defer f.Close()

				fw, err := w.CreateFormFile("post[file]", r.File)
				if err != nil {
					return err
				}
				if _, err = io.Copy(fw, f); err != nil {
					return err
				}

			} else {
				err := w.WriteField(v.Type().Field(i).Tag.Get("form"), field.Interface().(string))
				if err != nil {
					return err
				}
			}
		}
	}

	w.Close()

	req, err := http.NewRequest("POST", api.Booru+"/post/create.xml", &b)
	if err != nil {
		return err
	}
	//req.Header.Add("Content-Type", w.FormDataContentType())

	resp, err := api.AuthHandler.Client.Do(req)
	fmt.Println(resp.Status)
	return err
}

func NewPostListRequest() (r moebooruModel.PostListRequest) {
	return
}

func NewPostCreateRequest() (r moebooruModel.PostCreateRequest) {
	return
}
