package core

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewGrid(t *testing.T) {
	type args struct {
		gid  int
		minx int
		maxx int
		miny int
		maxy int
	}
	tests := []struct {
		name string
		args args
		want *Grid
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGrid(tt.args.gid, tt.args.minx, tt.args.maxx, tt.args.miny, tt.args.maxy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_Add(t *testing.T) {
	type fields struct {
		GID       int
		MinX      int
		MaxX      int
		Miny      int
		MaxY      int
		playerIDs map[int]bool
		pIDLock   sync.RWMutex
	}
	type args struct {
		playerID int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Grid{
				GID:       tt.fields.GID,
				MinX:      tt.fields.MinX,
				MaxX:      tt.fields.MaxX,
				Miny:      tt.fields.Miny,
				MaxY:      tt.fields.MaxY,
				playerIDs: tt.fields.playerIDs,
				pIDLock:   tt.fields.pIDLock,
			}
			r.Add(tt.args.playerID)
		})
	}
}

func TestGrid_Remove(t *testing.T) {
	type fields struct {
		GID       int
		MinX      int
		MaxX      int
		Miny      int
		MaxY      int
		playerIDs map[int]bool
		pIDLock   sync.RWMutex
	}
	type args struct {
		playerID int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Grid{
				GID:       tt.fields.GID,
				MinX:      tt.fields.MinX,
				MaxX:      tt.fields.MaxX,
				Miny:      tt.fields.Miny,
				MaxY:      tt.fields.MaxY,
				playerIDs: tt.fields.playerIDs,
				pIDLock:   tt.fields.pIDLock,
			}
			r.Remove(tt.args.playerID)
		})
	}
}

func TestGrid_GetPlyerIDs(t *testing.T) {
	type fields struct {
		GID       int
		MinX      int
		MaxX      int
		Miny      int
		MaxY      int
		playerIDs map[int]bool
		pIDLock   sync.RWMutex
	}
	tests := []struct {
		name          string
		fields        fields
		wantPlayerIDs []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Grid{
				GID:       tt.fields.GID,
				MinX:      tt.fields.MinX,
				MaxX:      tt.fields.MaxX,
				Miny:      tt.fields.Miny,
				MaxY:      tt.fields.MaxY,
				playerIDs: tt.fields.playerIDs,
				pIDLock:   tt.fields.pIDLock,
			}
			if gotPlayerIDs := r.GetPlyerIDs(); !reflect.DeepEqual(gotPlayerIDs, tt.wantPlayerIDs) {
				t.Errorf("Grid.GetPlyerIDs() = %v, want %v", gotPlayerIDs, tt.wantPlayerIDs)
			}
		})
	}
}

func TestGrid_String(t *testing.T) {
	type fields struct {
		GID       int
		MinX      int
		MaxX      int
		Miny      int
		MaxY      int
		playerIDs map[int]bool
		pIDLock   sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Grid{
				GID:       tt.fields.GID,
				MinX:      tt.fields.MinX,
				MaxX:      tt.fields.MaxX,
				Miny:      tt.fields.Miny,
				MaxY:      tt.fields.MaxY,
				playerIDs: tt.fields.playerIDs,
				pIDLock:   tt.fields.pIDLock,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("Grid.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
