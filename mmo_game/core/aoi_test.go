package core

import (
	"reflect"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	type args struct {
		minx  int
		maxx  int
		cntsx int
		miny  int
		maxy  int
		cntsy int
	}
	tests := []struct {
		name string
		args args
		want *AOIManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAOIManager(tt.args.minx, tt.args.maxx, tt.args.cntsx, tt.args.miny, tt.args.maxy, tt.args.cntsy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAOIManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAOIManager_gridWidth(t *testing.T) {
	type fields struct {
		MinX  int
		MaxX  int
		CntsX int
		MinY  int
		MaxY  int
		CntsY int
		grids map[int]*Grid
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AOIManager{
				MinX:  tt.fields.MinX,
				MaxX:  tt.fields.MaxX,
				CntsX: tt.fields.CntsX,
				MinY:  tt.fields.MinY,
				MaxY:  tt.fields.MaxY,
				CntsY: tt.fields.CntsY,
				grids: tt.fields.grids,
			}
			if got := m.gridWidth(); got != tt.want {
				t.Errorf("AOIManager.gridWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAOIManager_gridLength(t *testing.T) {
	type fields struct {
		MinX  int
		MaxX  int
		CntsX int
		MinY  int
		MaxY  int
		CntsY int
		grids map[int]*Grid
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AOIManager{
				MinX:  tt.fields.MinX,
				MaxX:  tt.fields.MaxX,
				CntsX: tt.fields.CntsX,
				MinY:  tt.fields.MinY,
				MaxY:  tt.fields.MaxY,
				CntsY: tt.fields.CntsY,
				grids: tt.fields.grids,
			}
			if got := m.gridLength(); got != tt.want {
				t.Errorf("AOIManager.gridLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAOIManager_String(t *testing.T) {
	type fields struct {
		MinX  int
		MaxX  int
		CntsX int
		MinY  int
		MaxY  int
		CntsY int
		grids map[int]*Grid
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
			m := &AOIManager{
				MinX:  tt.fields.MinX,
				MaxX:  tt.fields.MaxX,
				CntsX: tt.fields.CntsX,
				MinY:  tt.fields.MinY,
				MaxY:  tt.fields.MaxY,
				CntsY: tt.fields.CntsY,
				grids: tt.fields.grids,
			}
			if got := m.String(); got != tt.want {
				t.Errorf("AOIManager.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAOIManager_GetSurroundGridsByGid(t *testing.T) {
	type fields struct {
		MinX  int
		MaxX  int
		CntsX int
		MinY  int
		MaxY  int
		CntsY int
		grids map[int]*Grid
	}
	type args struct {
		gID int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantGrids []*Grid
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AOIManager{
				MinX:  tt.fields.MinX,
				MaxX:  tt.fields.MaxX,
				CntsX: tt.fields.CntsX,
				MinY:  tt.fields.MinY,
				MaxY:  tt.fields.MaxY,
				CntsY: tt.fields.CntsY,
				grids: tt.fields.grids,
			}
			if gotGrids := m.GetSurroundGridsByGid(tt.args.gID); !reflect.DeepEqual(gotGrids, tt.wantGrids) {
				t.Errorf("AOIManager.GetSurroundGridsByGid() = %v, want %v", gotGrids, tt.wantGrids)
			}
		})
	}
}

func TestAOIManager_GetGIDByPos(t *testing.T) {
	type fields struct {
		MinX  int
		MaxX  int
		CntsX int
		MinY  int
		MaxY  int
		CntsY int
		grids map[int]*Grid
	}
	type args struct {
		x float32
		y float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AOIManager{
				MinX:  tt.fields.MinX,
				MaxX:  tt.fields.MaxX,
				CntsX: tt.fields.CntsX,
				MinY:  tt.fields.MinY,
				MaxY:  tt.fields.MaxY,
				CntsY: tt.fields.CntsY,
				grids: tt.fields.grids,
			}
			if got := m.GetGidByPos(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("AOIManager.GetGIDByPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAOIManager_GetPIDsByPos(t *testing.T) {
	type fields struct {
		MinX  int
		MaxX  int
		CntsX int
		MinY  int
		MaxY  int
		CntsY int
		grids map[int]*Grid
	}
	type args struct {
		x float32
		y float32
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantPlayerIDs []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AOIManager{
				MinX:  tt.fields.MinX,
				MaxX:  tt.fields.MaxX,
				CntsX: tt.fields.CntsX,
				MinY:  tt.fields.MinY,
				MaxY:  tt.fields.MaxY,
				CntsY: tt.fields.CntsY,
				grids: tt.fields.grids,
			}
			if gotPlayerIDs := m.GetPidsByPos(tt.args.x, tt.args.y); !reflect.DeepEqual(gotPlayerIDs, tt.wantPlayerIDs) {
				t.Errorf("AOIManager.GetPIDsByPos() = %v, want %v", gotPlayerIDs, tt.wantPlayerIDs)
			}
		})
	}
}

func TestAOIManager_GetPidsByGid(t *testing.T) {
	type fields struct {
		MinX  int
		MaxX  int
		CntsX int
		MinY  int
		MaxY  int
		CntsY int
		grids map[int]*Grid
	}
	type args struct {
		gID int
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantPlayerIDs []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AOIManager{
				MinX:  tt.fields.MinX,
				MaxX:  tt.fields.MaxX,
				CntsX: tt.fields.CntsX,
				MinY:  tt.fields.MinY,
				MaxY:  tt.fields.MaxY,
				CntsY: tt.fields.CntsY,
				grids: tt.fields.grids,
			}
			if gotPlayerIDs := m.GetPidsByGid(tt.args.gID); !reflect.DeepEqual(gotPlayerIDs, tt.wantPlayerIDs) {
				t.Errorf("AOIManager.GetPidsByGid() = %v, want %v", gotPlayerIDs, tt.wantPlayerIDs)
			}
		})
	}
}

func TestAOIManager_RemovePidFromGrid(t *testing.T) {
	type fields struct {
		MinX  int
		MaxX  int
		CntsX int
		MinY  int
		MaxY  int
		CntsY int
		grids map[int]*Grid
	}
	type args struct {
		pID int
		gID int
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
			m := &AOIManager{
				MinX:  tt.fields.MinX,
				MaxX:  tt.fields.MaxX,
				CntsX: tt.fields.CntsX,
				MinY:  tt.fields.MinY,
				MaxY:  tt.fields.MaxY,
				CntsY: tt.fields.CntsY,
				grids: tt.fields.grids,
			}
			m.RemovePidFromGrid(tt.args.pID, tt.args.gID)
		})
	}
}

func TestAOIManager_AddPidToGrid(t *testing.T) {
	type fields struct {
		MinX  int
		MaxX  int
		CntsX int
		MinY  int
		MaxY  int
		CntsY int
		grids map[int]*Grid
	}
	type args struct {
		pID int
		gID int
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
			m := &AOIManager{
				MinX:  tt.fields.MinX,
				MaxX:  tt.fields.MaxX,
				CntsX: tt.fields.CntsX,
				MinY:  tt.fields.MinY,
				MaxY:  tt.fields.MaxY,
				CntsY: tt.fields.CntsY,
				grids: tt.fields.grids,
			}
			m.AddPidToGrid(tt.args.pID, tt.args.gID)
		})
	}
}

func TestAOIManager_AddToGridByPos(t *testing.T) {
	type fields struct {
		MinX  int
		MaxX  int
		CntsX int
		MinY  int
		MaxY  int
		CntsY int
		grids map[int]*Grid
	}
	type args struct {
		pID int
		x   float32
		y   float32
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
			m := &AOIManager{
				MinX:  tt.fields.MinX,
				MaxX:  tt.fields.MaxX,
				CntsX: tt.fields.CntsX,
				MinY:  tt.fields.MinY,
				MaxY:  tt.fields.MaxY,
				CntsY: tt.fields.CntsY,
				grids: tt.fields.grids,
			}
			m.AddToGridByPos(tt.args.pID, tt.args.x, tt.args.y)
		})
	}
}

func TestAOIManager_RemoveFromGridByPos(t *testing.T) {
	type fields struct {
		MinX  int
		MaxX  int
		CntsX int
		MinY  int
		MaxY  int
		CntsY int
		grids map[int]*Grid
	}
	type args struct {
		pID int
		x   float32
		y   float32
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
			m := &AOIManager{
				MinX:  tt.fields.MinX,
				MaxX:  tt.fields.MaxX,
				CntsX: tt.fields.CntsX,
				MinY:  tt.fields.MinY,
				MaxY:  tt.fields.MaxY,
				CntsY: tt.fields.CntsY,
				grids: tt.fields.grids,
			}
			m.RemoveFromGridByPos(tt.args.pID, tt.args.x, tt.args.y)
		})
	}
}
