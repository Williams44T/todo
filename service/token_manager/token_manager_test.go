package token_manager

import (
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func Test_TokenManager_IssueToken(t *testing.T) {
	type fields struct {
		secret []byte
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				secret: []byte("secret"),
			},
			args: args{
				userID: "user1234",
			},
			wantErr: false,
		},
		{
			name: "empty user id",
			fields: fields{
				secret: []byte("secret"),
			},
			args: args{
				userID: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TokenManager{
				Secret: tt.fields.secret,
			}
			_, err := tm.IssueToken(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenManager.IssueToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_TokenManager_VerifyToken(t *testing.T) {
	secret := []byte("secret")
	userID := "user1234"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
	})

	signed, err := token.SignedString(secret)
	if err != nil {
		t.Errorf("TokenManager.VerifyToken() error = %v", err)
	}

	type fields struct {
		secret []byte
	}
	type args struct {
		tokenStr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				secret: secret,
			},
			args: args{
				tokenStr: signed,
			},
			want:    userID,
			wantErr: false,
		},
		{
			name: "invalid token",
			fields: fields{
				secret: secret,
			},
			args: args{
				tokenStr: "invalid_token",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TokenManager{
				Secret: tt.fields.secret,
			}
			got, err := tm.VerifyToken(tt.args.tokenStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenManager.VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TokenManager.VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenManager_verifySigningMethod(t *testing.T) {
	type fields struct {
		Secret []byte
	}
	type args struct {
		token *jwt.Token
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				Secret: []byte("secret"),
			},
			args: args{
				token: jwt.New(jwt.SigningMethodHS256),
			},
			want:    []byte("secret"),
			wantErr: false,
		},
		{
			name: "invalid signing method",
			fields: fields{
				Secret: []byte("secret"),
			},
			args: args{
				token: jwt.New(jwt.SigningMethodES384),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TokenManager{
				Secret: tt.fields.Secret,
			}
			got, err := tm.verifySigningMethod(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenManager.verifySigningMethod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenManager.verifySigningMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
