package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/ansod/dungeon/helpers"
    "math"
    "math/rand"
    //"fmt"
)

type Player struct {
    rect *sdl.Rect
    curr_node *helpers.Node
    direction helpers.PlayerDirection
    right_image *sdl.Surface
    down_image *sdl.Surface
    left_image *sdl.Surface
    up_image *sdl.Surface
    inventory [3]helpers.InventoryItem
}

const (
    W_HEIGHT = 500
    W_WIDTH = 500
    R_HEIGHT = 15
    R_WIDTH = 15
)

// TODO:
// - Preloada imgs och textures i början och Free() och Destroy() när programmet stängs av
// - Lägg till så att ljuset "fadear" runt karaktären
// - Lägg till Enemies
// - Lägg till healthbar och inventory
// - Fixa varningar när man kör programmet


func drawMap(renderer *sdl.Renderer, p *Player) {

    p.rect.X = p.curr_node.X+2
    p.rect.Y = p.curr_node.Y+2

    renderer.SetDrawColor(0,0,0,250)
    renderer.Clear()


    for _, node := range helpers.Nodes {
        if true || (math.Abs(float64((p.curr_node.X+10) - (node.X+10))) < 60 && math.Abs(float64((p.curr_node.Y+10) - (node.Y+10))) < 60) {
            r := sdl.Rect{node.X, node.Y, 20, 20}
            switch node.Type {
            case helpers.Normal:
                    texture, err := renderer.CreateTextureFromSurface(node.Image)
                    if err != nil {
                        panic(err)
                    }
                    defer texture.Destroy()
                    src := sdl.Rect{0,0,20,20}
                    renderer.Copy(texture, &src, &r)
                    break
                case helpers.Treasure:
                    texture, err := renderer.CreateTextureFromSurface(node.Image)
                    if err != nil {
                        panic(err)
                    }
                    defer texture.Destroy()
                    src := sdl.Rect{0,0,20,20}
                    renderer.Copy(texture, &src, &r)
                    break
                case helpers.Door:
                    texture, err := renderer.CreateTextureFromSurface(node.Image)
                    if err != nil {
                        panic(err)
                    }
                    defer texture.Destroy()
                    src := sdl.Rect{0,0,20,20}
                    renderer.Copy(texture, &src, &r)
                    break
                case helpers.InventorySpot:
                    texture, err := renderer.CreateTextureFromSurface(node.Image)
                    if err != nil {
                        panic(err)
                    }
                    defer texture.Destroy()
                    src := sdl.Rect{0,0,20,20}
                    renderer.Copy(texture, &src, &r)
                case helpers.Wall:
                    renderer.SetDrawColor(0,0,0,250)
                    renderer.FillRect(&r)
                    break
            }
        }

    }

    // TODO: Fix this
    for _, enemy := range helpers.Enemies {
        if true || (math.Abs(float64((p.curr_node.X+10) - (enemy.Curr_node.X+10))) < 60 && math.Abs(float64((p.curr_node.Y+10) - (enemy.Curr_node.Y+10))) < 60) {
            r := sdl.Rect{helpers.Nodes[10].X+2, helpers.Nodes[10].Y+2, R_WIDTH, R_HEIGHT}
            switch enemy.Type {
                case helpers.Looter:
                    renderer.SetDrawColor(203, 72, 77,250)
                    break
                case helpers.DungeonCrawler:
                    texture, err := renderer.CreateTextureFromSurface(enemy.Image)
                    if err != nil {
                        panic(err)
                    }
                    defer texture.Destroy()
                    src := sdl.Rect{0,0,20,20}
                    renderer.Copy(texture, &src, &r)
                    break
            }
        }
    }

    switch p.direction {
    case helpers.Right:
        texture, err := renderer.CreateTextureFromSurface(p.right_image)
        if err != nil {
            panic(err)
        }
        defer texture.Destroy()
        src := sdl.Rect{0,0,15,15}
        renderer.Copy(texture, &src, p.rect)
        break
    case helpers.Down:
        texture, err := renderer.CreateTextureFromSurface(p.down_image)
        if err != nil {
            panic(err)
        }
        defer texture.Destroy()
        src := sdl.Rect{0,0,15,15}
        renderer.Copy(texture, &src, p.rect)
        break
    case helpers.Left:
        texture, err := renderer.CreateTextureFromSurface(p.left_image)
        if err != nil {
            panic(err)
        }
        defer texture.Destroy()
        src := sdl.Rect{0,0,15,15}
        renderer.Copy(texture, &src, p.rect)
        break
    case helpers.Up:
        texture, err := renderer.CreateTextureFromSurface(p.up_image)
        if err != nil {
            panic(err)
        }
        defer texture.Destroy()
        src := sdl.Rect{0,0,15,15}
        renderer.Copy(texture, &src, p.rect)
        break

    }

    renderer.Present()
}

func main() {

    helpers.CreatePath()

    var player Player
    player.inventory = [3]helpers.InventoryItem{helpers.Empty, helpers.Empty, helpers.Empty}
    player.curr_node = helpers.Nodes[3]

    helpers.CreateEnemies()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		W_WIDTH, W_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

    renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
    if err != nil {
        panic(err)
    }
    defer renderer.Destroy()

    renderer.Clear()

    rect := sdl.Rect{player.curr_node.X+10, player.curr_node.Y+10, R_WIDTH, R_HEIGHT}
    player.rect = &rect
    player.direction = helpers.Right
    player.right_image = helpers.LoadImage(helpers.Player_right)
    player.down_image = helpers.LoadImage(helpers.Player_down)
    player.left_image = helpers.LoadImage(helpers.Player_left)
    player.up_image = helpers.LoadImage(helpers.Player_up)

    drawMap(renderer, &player)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
            case *sdl.KeyboardEvent:
                if t.Type == sdl.KEYDOWN {
                    switch t.Keysym.Sym {
                    case sdl.K_RIGHT:
                        if player.curr_node.Right != nil && player.curr_node.Right.Type == helpers.Normal{
                            player.curr_node = player.curr_node.Right
                        }
                        player.direction = helpers.Right
                        break
                    case sdl.K_LEFT:
                        if player.curr_node.Left != nil && player.curr_node.Left.Type == helpers.Normal{
                            player.curr_node = player.curr_node.Left
                        }
                        player.direction = helpers.Left
                        break
                    case sdl.K_UP:
                        if player.curr_node.Up != nil && player.curr_node.Up.Type == helpers.Normal {
                            player.curr_node = player.curr_node.Up
                        }
                        player.direction = helpers.Up
                        break
                    case sdl.K_DOWN:
                        if player.curr_node.Down != nil && player.curr_node.Down.Type == helpers.Normal{
                            player.curr_node = player.curr_node.Down

                        }
                        player.direction = helpers.Down
                        break
                    case sdl.K_e:
                        if player.curr_node.Up != nil {
                            if player.curr_node.Up.Type == helpers.Treasure {
                                player.curr_node.Up.Image = helpers.LoadImage(helpers.Treasure_open_image)
                                randomIndex := rand.Intn(len(helpers.Treasure_objects))
                                pick := helpers.Treasure_objects[randomIndex]
                                if pick != helpers.Empty {
                                    for i, spot := range player.inventory {
                                        if spot == helpers.Empty {
                                            player.inventory[i] = pick
                                            helpers.Nodes[i].Image = helpers.LoadImage(helpers.Inventory_spot_starlight_image)
                                            break
                                        }
                                    }
                                }
                            }

                        }
                        break
                    }
                }
                drawMap(renderer, &player)
                break
			}
		}
	}
}
