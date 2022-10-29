package domain

import (
	"reflect"
	"testing"
)

func TestNewDepositTransaction(t *testing.T) {
	type args struct {
		userID string
		a      Amount
	}
	tests := []struct {
		name    string
		args    args
		want    DepositTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDepositTransaction(tt.args.userID, tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDepositTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDepositTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
