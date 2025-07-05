package exercise

import "fmt"

type Player struct {
	Name      string
	Inventory []Item
}
type Item struct {
	Name string
	Type string
}

// Pick up an item (modifies the player's inventory)
func (p *Player) PickUpItem(item Item) {
	//TODO: Implement function to add an item to inventory
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up %s!\n", p.Name, item.Name)
}

// Drop an item (removes from inventory)
func (p *Player) DropItem(itemName string) {
	//TODO: Implement function to remove an item from inventory
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s dropped %s.\n", p.Name, itemName)
			return
		}
	}
	fmt.Printf("%s does not have %s in inventory.\n", p.Name, itemName)
}

// use an item (if potion, remove it after use)
func (p *Player) UseItem(itemName string) {
	//TODO: Implement function to use an item
	for i, item := range p.Inventory {
		if item.Name == itemName {
			if item.Type == "potion" {
				fmt.Printf("%s used %s and feels rejuvenated!\n", p.Name, itemName)
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			} else {
				fmt.Printf("%s used %s.\n", p.Name, itemName)

			}
			return

		}
	}
	fmt.Printf("%s does not have %s in inventory.\n", p.Name, itemName)
}
