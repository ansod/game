package helpers

type NodeType string
type EnemyType string
type PlayerDirection string
type TreasureObject string
type InventoryItem string

const (
    Looter EnemyType = "Looter"
    Madman = "Madman"
    DungeonCrawler = "DungeonCrawler"
)

const (
    Normal NodeType = "Normal"
    Treasure = "Treasure"
    Door = "Door"
    Wall = "Wall"
    InventorySpot = "InventorySpot"
)

const (
    Right PlayerDirection = "Right"
    Down = "Down"
    Left = "Left"
    Up = "Up"
)

// /Users/Anton/go/src/github.com/ansod/dungeon

const (
    Stone_image = "./assets/stone.png"
    Treasure_image = "./assets/treasure.png"
    Treasure_open_image = "./assets/treasure_open.png"
    Door_left_image = "./assets/door_left.png"
    Door_right_image = "./assets/door_right.png"
    Player_left = "./assets/player_left.png"
    Player_down = "./assets/player_down.png"
    Player_right = "./assets/player_right.png"
    Player_up = "./assets/player_up.png"
    DungeonCrawler_down = "./assets/dungeon_crawler_down.png"
    Inventory_spot_empty_image = "./assets/inventory_spot_empty.png"
    Inventory_spot_starlight_image = "./assets/inventory_spot_light.png"
)

const (
    Empty InventoryItem = "Empty"
    StarLight = "StarLight"
)

var Treasure_objects = [2]InventoryItem{Empty, StarLight}
