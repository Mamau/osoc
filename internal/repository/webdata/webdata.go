package webdata

import (
	"context"
	"crypto/tls"
	"encoding/csv"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"osoc/internal/entity"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/log"
	"strconv"
	"strings"
)

type WebData struct {
	repo userinfo.UserRepo
	log  log.Logger
}

func NewWebData(repo userinfo.UserRepo, log log.Logger) *WebData {
	return &WebData{
		repo: repo,
		log:  log,
	}
}
func (w *WebData) InsertMillion(ctx context.Context) {
	data, err := w.GetData(ctx)
	if err != nil {
		w.log.Err(err).Msg("error while get data")
		return
	}

	var collectionPack [][]entity.SecureUser
	maxPack := 10

	for i := 0; i < len(data); i += maxPack {
		end := i + maxPack

		if end > len(data) {
			end = len(data)
		}

		collectionPack = append(collectionPack, data[i:end])
	}

	for _, v := range collectionPack {
		if err := w.repo.MultiCreateUser(ctx, v); err != nil {
			w.log.Err(err).Msg("error while multi insert")
			break
		}
	}
}

func (w *WebData) GetData(ctx context.Context) ([]entity.SecureUser, error) {
	addr := "https://raw.githubusercontent.com/OtusTeam/highload/master/homework/people.csv"
	response, err := w.doRequest(ctx, addr)
	if err != nil {
		return nil, err
	}

	defer func() { _ = response.Body.Close() }()
	reader := csv.NewReader(response.Body)
	users := make([]entity.SecureUser, 0, 1000000)
	pass := defaultPassword()
	for {
		line, erro := reader.Read()

		if erro == io.EOF {
			break
		} else if erro != nil {
			w.log.Err(erro).Msg("error while read")
			break
		}

		func(record []string) {
			if len(record) < 2 {
				w.log.Warn().Msg("small data")
				return
			}
			userFIO := strings.Split(record[0], " ")
			age, err := strconv.Atoi(record[1])
			if err != nil {
				w.log.Err(erro).Msg("error while atoi")
				return
			}

			users = append(users, entity.SecureUser{
				User: entity.User{
					FirstName: userFIO[0],
					LastName:  userFIO[1],
					Age:       age,
					Sex:       "male",
					Interests: "no",
				},
				Password: pass,
			})
		}(line)
	}
	return users, nil
}

func (w *WebData) doRequest(ctx context.Context, addr string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, addr, nil)

	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusInternalServerError && resp.StatusCode <= http.StatusNetworkAuthenticationRequired {
		return nil, fmt.Errorf("probably remote server is down, status code: %d", resp.StatusCode)
	}

	return resp, nil
}

func defaultPassword() string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
	return string(bytes)
}
