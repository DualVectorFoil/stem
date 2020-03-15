package account

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Authing/authing-go-sdk"
	"github.com/DualVectorFoil/stem/app/conf"
	"github.com/DualVectorFoil/stem/httpHelper"
	"github.com/DualVectorFoil/stem/model"
	"github.com/DualVectorFoil/stem/model/graphqlModel"
	"github.com/DualVectorFoil/stem/util/encryptUtil"
	"github.com/hokaccha/go-prettyjson"
	"github.com/kelvinji2009/graphql"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
)

var accountClient *authing.Client
var accountClientOnce sync.Once

func AccountClientInstance() *authing.Client {
	accountClientOnce.Do(func() {
		isRelease := os.Getenv(conf.IS_RELEASE_ENV_KEY)
		isDev := false
		if isRelease == "true" {
			isDev = true
		}
		accountClient = authing.NewClient(conf.USER_POOL_ID, conf.USER_SECRET_KEY, isDev)
		if isDev {
			AccountClientInstance().Client.Log = func(s string) {
				b := []byte(s)
				pj, _ := prettyjson.Format(b)
				fmt.Println(string(pj))
			}
		}
	})
	return accountClient
}

func Register(phoneNum string, pwd string) (*authing.UserRegisterMutation, error) {
	input := &authing.UserRegisterInput{
		Phone:            graphql.String(phoneNum),
		Password:         graphql.String(pwd),
		RegisterInClient: graphql.String(conf.USER_POOL_ID),
	}

	m, err := AccountClientInstance().Register(input)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func LoginWithPwd(phoneNum string, userName string, pwd string, phoneCode string, email string) (*graphqlModel.LoginMutation, error) {
	var m graphqlModel.LoginMutation
	code, err := strconv.Atoi(phoneCode)
	if err != nil {
		return nil, err
	}
	variables := map[string]interface{}{
		"registerInClient": graphql.String(conf.USER_POOL_ID),
		"username":         graphql.String(phoneNum),
		"password":         graphql.String(encryptUtil.EncryptPassword([]byte(pwd))),
		"email":            graphql.String(email),
		"phone":            graphql.String(phoneNum),
		"phoneCode":        graphql.Int(code),
	}

	err = AccountClientInstance().Client.Mutate(context.Background(), &m, variables)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func QueryAccount(id string) (*authing.UserQuery, error) {
	input := &authing.UserQueryParameter{
		ID:               graphql.String(id),
		RegisterInClient: graphql.String(conf.USER_POOL_ID),
	}

	m, err := AccountClientInstance().User(input)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func DeleteAccounts(ids []string) (*authing.RemoveUsersMutation, error) {
	removeIds := []graphql.String{}
	for _, id := range ids {
		removeIds = append(removeIds, graphql.String(id))
	}
	input := &authing.RemoveUsersInput{
		IDs:              removeIds,
		RegisterInClient: graphql.String(conf.USER_POOL_ID),
	}

	re := regexp.MustCompile("^[0-9a-fA-F]{24}$")
	for _, id := range input.IDs {
		if !re.MatchString(string(id)) {
			logrus.WithFields(logrus.Fields{
				"id": id,
			}).Error("user id is invalid, cannot remove")
		}
	}

	m, err := AccountClientInstance().RemoveUsers(input)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Should not update id
func UpdateAccount(id string, userName string, phoneNum string) (*authing.UpdateUserMutation, error) {
	input := authing.UserUpdateInput{
		ID:       graphql.String(id),
		Username: graphql.String(userName),
		Phone:    graphql.String(phoneNum),
	}

	m, err := AccountClientInstance().UpdateUser(&input)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func VerifyToken(token string) (string, error) {
	resp, err := httpHelper.HttpClientInstance().Get("https://users.authing.cn/authing/token?access_token=" + token)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	verifyResp := &model.VerifyTokenRespModel{}
	err = json.Unmarshal(body, verifyResp)
	if err != nil {
		return "", err
	}
	if !verifyResp.Status || verifyResp.Code != http.StatusOK {
		return "", errors.New(verifyResp.Message)
	}
	return verifyResp.Token.Data.ID, nil
}
