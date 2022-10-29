package domain

import (
	"reflect"
	"testing"
)

func TestNewAmount(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		want    Amount
		wantErr bool
	}{
		{
			name: "invalid format 1",
			args: args{
				raw: ".1",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "invalid format 2",
			args: args{
				raw: "0.",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "invalid format 3",
			args: args{
				raw: "0",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "invalid format 4",
			args: args{
				raw: "0.555",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "invalid format 5",
			args: args{
				raw: "0.555",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty input",
			args: args{
				raw: "",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "only minor units with one 0",
			args: args{
				raw: "0.32",
			},
			want:    Amount(32),
			wantErr: false,
		},
		{
			name: "only minor units with two 0",
			args: args{
				raw: "00.77",
			},
			want:    Amount(77),
			wantErr: false,
		},
		{
			name: "only major units with one 0",
			args: args{
				raw: "36.0",
			},
			want:    Amount(3600),
			wantErr: false,
		},
		{
			name: "only major units with two 0",
			args: args{
				raw: "36.00",
			},
			want:    Amount(3600),
			wantErr: false,
		},

		{
			name: "1 major unit, 1 minor unit",
			args: args{
				raw: "7.6",
			},
			want:    Amount(706),
			wantErr: false,
		},
		{
			name: "1 major unit, 1 minor unit with 0",
			args: args{
				raw: "7.06",
			},
			want:    Amount(706),
			wantErr: false,
		},
		{
			name: "2 major units, 2 minor units",
			args: args{
				raw: "43.54",
			},
			want:    Amount(4354),
			wantErr: false,
		},
		{
			name: "3 major units, 1 minor unit",
			args: args{
				raw: "519.4",
			},
			want:    Amount(51904),
			wantErr: false,
		},
		{
			name: "3 major units, 1 minor unit with 0",
			args: args{
				raw: "519.04",
			},
			want:    Amount(51904),
			wantErr: false,
		},
		{
			name: "4 major units, 2 minor units",
			args: args{
				raw: "3454.55",
			},
			want:    Amount(345455),
			wantErr: false,
		},
		{
			name: "4 major units, 1 minor unit with 0",
			args: args{
				raw: "3454.06",
			},
			want:    Amount(345406),
			wantErr: false,
		},
		{
			name: "5 major units, 2 minor units",
			args: args{
				raw: "55676.77",
			},
			want:    Amount(5567677),
			wantErr: false,
		},
		{
			name: "6 major units, 1 minor unit",
			args: args{
				raw: "762234.1",
			},
			want:    Amount(76223401),
			wantErr: false,
		},
		{
			name: "6 major units, 1 minor unit with 0",
			args: args{
				raw: "762234.01",
			},
			want:    Amount(76223401),
			wantErr: false,
		},
		{
			name: "6 major units, 2 minor units",
			args: args{
				raw: "762234.11",
			},
			want:    Amount(76223411),
			wantErr: false,
		},
		{
			name: "7 major units, 2 minor units",
			args: args{
				raw: "9653321.17",
			},
			want:    Amount(965332117),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAmount(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAmount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_UInt64(t *testing.T) {
	tests := []struct {
		name string
		a    Amount
		want uint64
	}{
		{
			name: "why?!",
			a:    Amount(1337),
			want: 1337,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.UInt64(); got != tt.want {
				t.Errorf("amount.UInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_LargeUnit(t *testing.T) {
	tests := []struct {
		name string
		a    Amount
		want uint64
	}{
		{
			name: "23",
			a:    Amount(23),
			want: 0,
		},
		{
			name: "100",
			a:    Amount(100),
			want: 1,
		},
		{
			name: "3425",
			a:    Amount(3425),
			want: 34,
		},
		{
			name: "76533",
			a:    Amount(76543),
			want: 765,
		},
		{
			name: "984217",
			a:    Amount(984217),
			want: 9842,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.LargeUnit(); got != tt.want {
				t.Errorf("amount.MajorUnit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_SmallUnit(t *testing.T) {
	tests := []struct {
		name string
		a    Amount
		want uint64
	}{
		{
			name: "1",
			a:    Amount(1),
			want: 1,
		},
		{
			name: "17",
			a:    Amount(17),
			want: 17,
		},
		{
			name: "99",
			a:    Amount(99),
			want: 99,
		},
		{
			name: "100",
			a:    Amount(100),
			want: 0,
		},
		{
			name: "200",
			a:    Amount(200),
			want: 0,
		},
		{
			name: "3122",
			a:    Amount(3122),
			want: 22,
		},
		{
			name: "73312",
			a:    Amount(73312),
			want: 12,
		},
		{
			name: "31235",
			a:    Amount(31235),
			want: 35,
		},
		{
			name: "1234519",
			a:    Amount(1234519),
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SmallUnit(); got != tt.want {
				t.Errorf("amount.MinorUnit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_MajorUnits(t *testing.T) {
	tests := []struct {
		name string
		a    Amount
		want string
	}{
		{
			name: "0",
			a:    Amount(0),
			want: "0.0",
		},
		{
			name: "11.75",
			a:    Amount(1175),
			want: "11.75",
		},
		{
			name: "1.75",
			a:    Amount(175),
			want: "1.75",
		},
		{
			name: "1.5",
			a:    Amount(105),
			want: "1.5",
		},
		{
			name: "10.5",
			a:    Amount(1005),
			want: "10.5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("amount.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_IsZero(t *testing.T) {
	tests := []struct {
		name string
		a    Amount
		want bool
	}{
		{
			name: "true",
			a:    Amount(0),
			want: true,
		},
		{
			name: "false",
			a:    Amount(124123),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.IsZero(); got != tt.want {
				t.Errorf("Amount.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
