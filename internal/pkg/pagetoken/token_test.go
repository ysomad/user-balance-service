package pagetoken

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
)

var (
	testUUID   uuid.UUID
	testTime   time.Time
	testCursor string
)

func generateTestCursor(id uuid.UUID, t time.Time) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s,%s", id, t.Format(time.RFC3339Nano))))
}

func TestMain(m *testing.M) {
	testUUID, _ = uuid.NewRandom()
	testTime = time.Now()
	testCursor = generateTestCursor(testUUID, testTime)

	os.Exit(m.Run())
}

func TestEncode(t *testing.T) {
	type args struct {
		pk uuid.UUID
		t  time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				pk: testUUID,
				t:  testTime,
			},
			want: testCursor,
		},
		{
			name: "empty uuid",
			args: args{
				pk: uuid.UUID{},
				t:  testTime,
			},
			want: generateTestCursor(uuid.UUID{}, testTime),
		},
		{
			name: "empty time",
			args: args{
				pk: testUUID,
				t:  time.Time{},
			},
			want: generateTestCursor(testUUID, time.Time{}),
		},
		{
			name: "empty uuid and time",
			args: args{
				pk: uuid.UUID{},
				t:  time.Time{},
			},
			want: generateTestCursor(uuid.UUID{}, time.Time{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.pk, tt.args.t); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		cursor string
	}
	tests := []struct {
		name     string
		args     args
		wantUUID uuid.UUID
		wantTime time.Time
		wantErr  bool
	}{
		{
			name:     "invalid cursor",
			args:     args{cursor: "invalid"},
			wantUUID: uuid.UUID{},
			wantTime: time.Time{},
			wantErr:  true,
		},
		{
			name:     "success",
			args:     args{cursor: testCursor},
			wantUUID: testUUID,
			wantTime: testTime,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUUID, gotTime, err := Decode(tt.args.cursor)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUUID != tt.wantUUID {
				t.Errorf("decodeCursor() gotUUID = %v wantUUID %v", gotUUID, tt.wantUUID)
			}
			if !gotTime.Equal(tt.wantTime) {
				t.Errorf("decodeCursor() gotTime = %v, wantTime %v", gotTime, tt.wantTime)
			}
		})
	}
}
