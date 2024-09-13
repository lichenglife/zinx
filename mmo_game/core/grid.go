package core

import (
	"fmt"
	"sync"
)

// 定义地图中的格子
type Grid struct {
	GID       int // 格子ID
	MinX      int // x轴坐标
	MaxX      int
	Miny      int
	MaxY      int
	playerIDs map[int]bool // 格子中的玩家的ID信息
	pIDLock   sync.RWMutex // 读写锁
}

// 创建一个格子对象

func NewGrid(gid, minx, maxx, miny, maxy int) *Grid {

	return &Grid{
		GID:       gid,
		MinX:      minx,
		Miny:      miny,
		MaxY:      maxy,
		playerIDs: make(map[int]bool),
	}
}

//向当前格子中添加一个玩家

func (r *Grid) Add(playerID int) {
	r.pIDLock.Lock()
	defer r.pIDLock.Unlock()

	r.playerIDs[playerID] = true
}

//从格子中删除一个玩家

func (r *Grid) Remove(playerID int) {
	r.pIDLock.Lock()
	defer r.pIDLock.Unlock()

	delete(r.playerIDs, playerID)
}

//得到当前格子中所有的玩家

func (r *Grid) GetPlyerIDs() (playerIDs []int) {

	r.pIDLock.Lock()
	defer r.pIDLock.Unlock()

	for id := range r.playerIDs {
		playerIDs = append(playerIDs, id)
	}

	return
}

//打印信息方法

func (r *Grid) String() string {

	// 打印当前格子所有信息
	return fmt.Sprintf("gid %d,minX %d,maxX %d,Miny %d,maxy %d,playerIDs %d", r.GID, r.MinX, r.MaxX, r.Miny, r.MaxY, r.GetPlyerIDs())

}
