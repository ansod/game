package helpers

import (
    "github.com/veandco/go-sdl2/sdl"
)

type Enemy struct {
    Curr_node *Node
    Type EnemyType
    Image *sdl.Surface
}

var Enemies []*Enemy

func CreateEnemies() {

    var e1 = Enemy{Nodes[10], DungeonCrawler, LoadImage(DungeonCrawler_down)}

    Enemies = append(Enemies, &e1)
}
