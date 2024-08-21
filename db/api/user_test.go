package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"

	// "reflect"
	mockdb "simplebank/db/mock"
	db "simplebank/db/sqlc"
	"simplebank/db/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	// "golang.org/x/crypto/bcrypt"
)

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, e.arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
	// fmt.Print("there is the print",(e.arg.HashedPassword),"pass:",(e.password))
	// err := bcrypt.CompareHashAndPassword([]byte(e.arg.HashedPassword),[]byte(e.password))
	// if err != nil{
	// 	return false
	// }
	// return true
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestPostUser(t *testing.T) {
	user, password := randomUser(t)
	hashed, err := util.HashPassword(password)
	require.NoError(t, err)
	testcases := []struct {
		name          string
		body          gin.H
		buildstubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "first test",
			body: gin.H{
				"username":  user.Username,
				"password":  hashed,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildstubs: func(store *mockdb.MockStore) {
				arg := db.CreateUserParams{
					Username:       user.Username,
					FullName:       user.FullName,
					HashedPassword: hashed,
					Email:          user.Email,
				}
				// arg = db.CreateUserParams{}
				store.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).Times(1).Return(user, nil)
				// EqCreateUserParams(arg,password)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				// match the body of the retu user here
			},
		},
	}

	for i := range testcases {
		tc := testcases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)

			tc.buildstubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/user")

			// data to json
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)

		})
	}

}

func randomUser(t *testing.T) (user db.User, password string) {
	user.Username = util.RandomString(10)
	user.Email = util.RandomEmail()
	user.FullName = util.RandomString(10)
	user.HashedPassword = ""
	password = util.RandomString(10)
	return
}
