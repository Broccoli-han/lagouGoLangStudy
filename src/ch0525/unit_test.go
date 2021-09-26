package ch0525

import (
	"reflect"
	"testing"
)

func TestGetPersonDetail(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *Detail
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPersonDetail(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPersonDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPersonDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkEmail(tt.args.email); got != tt.want {
				t.Errorf("checkEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkUsername(tt.args.username); got != tt.want {
				t.Errorf("checkUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPersonDetailHttp(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *Detail
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPersonDetailHttp(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPersonDetailHttp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPersonDetailHttp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
