// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.
package funcmap

import (
	"testing"
)

func TestAppend(t *testing.T) {
	type args struct {
		s      interface{}
		append interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "append string",
			args: args{
				s:      "foo",
				append: "bar",
			},
			want:    "foobar",
			wantErr: false,
		},
		{
			name: "append number",
			args: args{
				s:      6,
				append: 9,
			},
			want:    "69",
			wantErr: false,
		},
		{
			name: "append string",
			args: args{
				s:      nil,
				append: "bar",
			},
			want:    "bar",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Append(tt.args.s, tt.args.append)
			if (err != nil) != tt.wantErr {
				t.Errorf("Append() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNIndent(t *testing.T) {
	type args struct {
		s interface{}
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "indent string",
			args: args{
				s: "foo",
				n: 2,
			},
			want:    "\n  foo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NIndent(tt.args.s, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NIndent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NIndent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrepend(t *testing.T) {
	type args struct {
		s       interface{}
		prepend interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "prepend string",
			args: args{
				s:       "foo",
				prepend: "bar",
			},
			want:    "barfoo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Prepend(tt.args.s, tt.args.prepend)
			if (err != nil) != tt.wantErr {
				t.Errorf("Prepend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Prepend() = %v, want %v", got, tt.want)
			}
		})
	}
}
