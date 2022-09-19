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
)

type Item struct {
	ID      int
	Name    string
	Rarity  Rarity
	ForHero string
	Cost    int
}

type ItemChanger struct {
	Name    *string
	Rarity  *Rarity
	ForHero *string
	Cost    *int
}

type ItemFinder struct {
	Name         *string
	Rarity       *string
	Hero         *string
	LowestPrice  *int
	HighestPrice *int
}
