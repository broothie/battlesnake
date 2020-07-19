package game

type Sectors struct {
	counter int
	sectors map[*Cell]int
}

func newSectors() *Sectors {
	return &Sectors{sectors: make(map[*Cell]int)}
}

func (s *Sectors) NewSectorFor(cell *Cell) {
	s.sectors[cell] = s.counter
	s.counter++
}

func (s *Sectors) AddToSector(existing, new *Cell) {
	s.sectors[new] = s.sectors[existing]
}

func (s *Sectors) HasSectorFor(cell *Cell) bool {
	_, exists := s.sectors[cell]
	return exists
}

func (g *Grid) FindCellSectorSizes() map[*Cell]int {
	sectors := newSectors()

	for _, row := range g.rows {
		for _, cell := range row {
			if cell.HasSnake() {
				continue
			}

			if sectors.HasSectorFor(cell) {
				continue
			}

			sectors.NewSectorFor(cell)
			cell.expand(sectors)
		}
	}

	sectorCounts := make([]int, sectors.counter)
	for _, id := range sectors.sectors {
		sectorCounts[id]++
	}

	cellSectorSizes := make(map[*Cell]int)
	for cell, id := range sectors.sectors {
		cellSectorSizes[cell] = sectorCounts[id]
	}

	return cellSectorSizes
}

func (c *Cell) expand(sectors *Sectors) {
	for _, neighbor := range c.Neighbors() {
		if neighbor.HasSnake() {
			continue
		}

		if sectors.HasSectorFor(neighbor) {
			continue
		}

		sectors.AddToSector(c, neighbor)
		neighbor.expand(sectors)
	}
}
