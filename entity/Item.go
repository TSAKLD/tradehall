package entity

type Rarity string

const (
	Common    Rarity = "Common"
	Uncommon  Rarity = "Uncommon"
	Rare      Rarity = "Rare"
	Mythical  Rarity = "Mythical"
	Legendary Rarity = "Legendary"
	Immortal  Rarity = "Immortal"
	Arcana    Rarity = "Arcana"
	Ancient   Rarity = "Ancient"
	Persona   Rarity = "Persona"
)

type Item struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Rarity  Rarity `json:"rarity"`
	ForHero string `json:"hero"`
	Cost    int    `json:"cost,"`
}

type ItemChanger struct {
	Name    *string
	Rarity  *Rarity
	ForHero *string
	Cost    *int
}

type ItemFinder struct {
	ID           *int
	Name         *string
	Rarity       *string
	Hero         *string
	LowestPrice  *int
	HighestPrice *int
}
