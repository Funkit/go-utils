package utils

import (
	"reflect"
	"testing"
)

func TestValueEqual(t *testing.T) {

	//nil and nil, float64 pointers
	if res := ValueEqual[float64](nil, nil); res != true {
		t.Errorf("ValueEqual() = %v, want %v", res, true)
	}

	//pointer 1 is nil, float64 pointers
	test1Float64Val1 := float64(125.012567891)

	if res := ValueEqual[float64](&test1Float64Val1, nil); res != false {
		t.Errorf("ValueEqual() = %v, want %v", res, false)
	}

	//pointer 2 is nil, float64 pointers
	test2Float64Val2 := float64(125.012567891)

	if res := ValueEqual[float64](nil, &test2Float64Val2); res != false {
		t.Errorf("ValueEqual() = %v, want %v", res, false)
	}

	//equal float64
	float64Val2 := float64(125.012567891)
	float64Val3 := float64Val2

	if res := ValueEqual[float64](&float64Val2, &float64Val3); res != true {
		t.Errorf("ValueEqual() = %v, want %v", res, true)
	}

	//equal strings
	stringVal1 := "this is a test"
	stringVal2 := stringVal1

	if res := ValueEqual[string](&stringVal1, &stringVal2); res != true {
		t.Errorf("ValueEqual() = %v, want %v", res, true)
	}

	//unequal float64
	float64Val4 := float64(-57.15997)
	float64Val5 := float64Val2 + 1
	if res := ValueEqual[float64](&float64Val4, &float64Val5); res != false {
		t.Errorf("ValueEqual() = %v, want %v", res, false)
	}
}

func TestGetFileAsLines(t *testing.T) {
	type args struct {
		filePath      string
		maxLineLength []int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "default working case",
			args: args{
				filePath:      "../samples/sample1.txt",
				maxLineLength: []int{65536},
			},
			want: []string{
				"This is line 1",
				"Here is line 2",
				"Finally line 3",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFileAsLines(tt.args.filePath, tt.args.maxLineLength...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileAsLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFileAsLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}
