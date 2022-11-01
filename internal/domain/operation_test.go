package domain

import (
	"reflect"
	"testing"
)

func TestNewOperation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Op
		wantErr bool
	}{
		{
			name: "lowercase withdraw",
			args: args{
				s: "withdraw",
			},
			want:    OpWithdraw,
			wantErr: false,
		},
		{
			name: "uppercase withdraw",
			args: args{
				s: "WITHDRAW",
			},
			want:    OpWithdraw,
			wantErr: false,
		},
		{
			name: "mixed case withdraw",
			args: args{
				s: "WiThDraW",
			},
			want:    OpWithdraw,
			wantErr: false,
		},
		{
			name: "lowercase deposit",
			args: args{
				s: "deposit",
			},
			want:    OpDeposit,
			wantErr: false,
		},
		{
			name: "uppercase deposit",
			args: args{
				s: "DEPOSIT",
			},
			want:    OpDeposit,
			wantErr: false,
		},
		{
			name: "mixed case deposit",
			args: args{
				s: "DePoSiT",
			},
			want:    OpDeposit,
			wantErr: false,
		},
		{
			name: "invalid operation",
			args: args{
				s: "invalid_oepration",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOp(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOperation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_operation_String(t *testing.T) {
	tests := []struct {
		name string
		op   Op
		want string
	}{
		{
			name: "success deposit",
			op:   OpDeposit,
			want: string(OpDeposit),
		},
		{
			name: "success withdraw",
			op:   OpDeposit,
			want: string(OpDeposit),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.op.String(); got != tt.want {
				t.Errorf("operation.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
