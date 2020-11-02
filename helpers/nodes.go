package helpers

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/img"
    "io/ioutil"
    "math"
    "bytes"
)

type Node struct {
    X int32
    Y int32
    Type NodeType
    Right *Node
    Left *Node
    Up *Node
    Down *Node
    Image *sdl.Surface
}

var Nodes []*Node

func linkNodes() {

    for _, n1 := range Nodes {
        for _, n2 := range Nodes {
            if n1.Y == n2.Y {
                if n2.X == n1.X + 20 {
                    n1.Right = n2
                } else if n2.X == n1.X - 20 {
                    n1.Left = n2
                }
            } else if n1.X == n2.X {
                if n2.Y == n1.Y + 20 {
                    n1.Down = n2
                } else if n2.Y == n1.Y - 20 {
                    n1.Up = n2
                }
            }
        }
    }
}

func LoadImage(file string) *sdl.Surface {
    image, err := img.Load(file)
    if err != nil {
        panic(err)
    }
    //defer image.Free()

    return image
}

func CreatePath() {
    
    b, err := ioutil.ReadFile("./helpers/map_sim.txt")
    if err != nil {
        panic(err)
    }

    b = bytes.Trim(b, "\n")

    for i, byte := range b {
        switch string(byte) {
            case "O": // Ordinary node
                n := Node{int32((i%25)*20), int32((math.Floor(float64(i/25)))*20), Normal, nil, nil, nil, nil, LoadImage(Stone_image)}
                Nodes = append(Nodes, &n)
                break
            case "$": // Treasure node
                n := Node{int32((i%25)*20), int32((math.Floor(float64(i/25)))*20), Treasure, nil, nil, nil, nil, LoadImage(Treasure_image)}
                Nodes = append(Nodes, &n)
                break
            case "D": // Door
                var n Node
                if int32((i%25)*20) == 0 {
                    n = Node{int32((i%25)*20), int32((math.Floor(float64(i/25)))*20), Door, nil, nil, nil, nil, LoadImage(Door_left_image)}
                } else {
                    n = Node{int32((i%25)*20), int32((math.Floor(float64(i/25)))*20), Door, nil, nil, nil, nil, LoadImage(Door_right_image)}
                }
                Nodes = append(Nodes, &n)
                break
            case "I":
                n := Node{int32((i%25)*20), int32((math.Floor(float64(i/25)))*20), InventorySpot, nil, nil, nil, nil, LoadImage(Inventory_spot_empty_image)}
                Nodes = append(Nodes, &n)
                break
            case "#": // Wall
                break

        }
    }

    linkNodes()

}
