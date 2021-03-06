package auth

import (
	"errors"
	"testing"

	"github.com/huypher/talent-acquisition-management/domain"
	"github.com/stretchr/testify/require"
)

func Test_genJWT(t *testing.T) {
	t.Parallel()

	acc := &domain.User{
		ID:       1,
		UserName: "username",
		Password: "password",
		Name:     "name",
	}

	secretJWT := "jwtToken"

	testCases := []struct {
		name string
		acc  *domain.User
		err  error
	}{
		{"test", acc, nil},
		{"test", nil, errors.New("user empty")},
	}
	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := genJWT(tc.acc, secretJWT)
			require.Equal(t, tc.err, err)
			if tc.err == nil {
				require.NotEmpty(t, got)
			}
		})
	}
}

func Test_genHash(t *testing.T) {
	type args struct {
		pwd []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test",
			args:    args{[]byte("test")},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := genHash(tt.args.pwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("genHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("genHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}
