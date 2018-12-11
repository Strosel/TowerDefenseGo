package main

import (
	"errors"
	"io/ioutil"
	"math"
	"strings"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Level Defines a grid-based Level that handles a Track and a Set of Towers
type Level struct {
	*Entity
	EntitySize, Offset  pixel.Vec
	Size, Height, Width int
	Track               Track
	Towers              TowerHandler
}

// NewLevel Creates a new level from a file
func NewLevel(wh float64, h, w int, src string) *Level {
	offset := pixel.V(0, wh-float64(h))
	e := NewEntity(pixel.V(float64(w), float64(h)).Scaled(0.5).Add(offset), pixel.V(float64(w), float64(h)), DrawableRect(pixel.R(0, 0, float64(w), float64(h)), colornames.Red))
	grid := &Level{
		Entity: e,
		Offset: offset,
	}

	Map := fileToMatrix(src)

	for i := 0; i < len(Map); i++ {
		if len(Map) <= 1 {
			Map = append(Map[:i], Map[i+1:]...)
		}
		if i > 0 {
			if len(Map[i]) > len(Map) {
				Map[i] = Map[i][:len(Map)]
			}

			for len(Map[i]) < len(Map) {
				Map[i] = append(Map[i], "X")
			}
		}
	}

	Map = rotate(Map)

	grid.EntitySize = pixel.V(float64(w)/float64(len(Map)), float64(h)/float64(len(Map[0])))
	grid.Size = len(Map)
	grid.Height = h
	grid.Width = w
	grid.Track = NewTrack(h, w, Map, grid.EntitySize)
	grid.Towers = TowerHandler{}

	for _, r := range grid.Track.Path {
		r.Position = r.Position.Add(offset)
	}

	return grid
}

// GridPos Translate a pixel position to a position in the level grid
func (l Level) GridPos(vec pixel.Vec) pixel.Vec {
	if !l.Bounds().Contains(vec) {
		return pixel.V(-1, -1)
	}
	vec = vec.Sub(l.Offset)
	return pixel.V(math.Floor(vec.X/l.EntitySize.X), math.Floor(vec.Y/l.EntitySize.Y))
}

// GetRoad Get the road at position vec
func (l Level) GetRoad(vec pixel.Vec) (*Road, error) {
	pos := l.GridPos(vec)
	out, ok := l.Track.Path[pos]
	if ok {
		return out, nil
	}
	return nil, errors.New("Road Not Found")
}

// GetTower Get the road at position vec
func (l Level) GetTower(vec pixel.Vec) (*Tower, error) {
	pos := l.GridPos(vec)
	out, ok := l.Towers[pos]
	if ok {
		return out, nil
	}
	return nil, errors.New("Tower Not Found")
}

// GridIsEmpty Check if the current position is empty
func (l Level) GridIsEmpty(vec pixel.Vec) bool {
	road, _ := l.GetRoad(vec)
	tower, _ := l.GetTower(vec)
	return road == nil && tower == nil
}

// Draw Draw the Track and Towers in the Level
func (l Level) Draw(win *pixelgl.Window) {
	mat := pixel.IM
	mat = mat.Moved(l.Position)
	mat = mat.ScaledXY(l.Position, l.Scale)
	l.Sprite.Draw(win, mat)

	l.Track.Draw(win)
	l.Towers.Draw(win)
}

// AddTower Add a tower to the level
func (l Level) AddTower(pos pixel.Vec, sprite pixel.Picture, r float64, dmg, frq int) {
	tw := NewTower(pos.ScaledXY(l.EntitySize).Add(l.EntitySize.Scaled(0.5)).Add(l.Offset), l.EntitySize, sprite, r, dmg, frq)
	l.Towers[pos] = tw
}

func fileToMatrix(src string) [][]string {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		panic(err)
	}
	file := strings.Split(string(data), "\n")
	mat := [][]string{}
	for _, row := range file {
		mat = append(mat, strings.Split(row, " "))
	}
	return mat
}

func rotate(mat [][]string) [][]string {
	newMat := [][]string{}

	for _x, x := range mat {
		newMat = append(newMat, []string{})
		for _y := range x {
			newMat[_x] = append(newMat[_x], mat[len(mat)-_y-1][_x])
		}
	}

	return newMat
}
