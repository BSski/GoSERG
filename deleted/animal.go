package deleted

//type Animal struct {
//	name   string
//	energy int
//	pos    *mat.VecDense
//	color  color.NRGBA
//}

//type AnimalMover interface {
//	move(animals *map[float64]map[float64]map[any]struct{})
//}
//
//type AnimalAppearator interface {
//	drawMe(screen *ebiten.Image)
//}

//func makeAMove(animalMover AnimalMover, animals *map[float64]map[float64]map[any]struct{}) {
//	animalMover.move(animals)
//}

//func drawAnimal(animalAppearator AnimalAppearator, screen *ebiten.Image) {
//	animalAppearator.drawMe(screen)
//}

//func (a *Animal) animalInit(name string, c color.NRGBA) {
//	a.name = name
//	a.color = c
//	a.pos = mat.NewVecDense(
//		2,
//		[]float64{
//			float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
//			float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
//		},
//	)
//}

//func (a *Animal) putOnScreen(screen *ebiten.Image, size float64) {
//	ebitenutil.DrawCircle(
//		screen,
//		boardStartX+boardBorderWidth+tileMiddlePx+a.pos.AtVec(0),
//		boardStartY+boardBorderWidth+tileMiddlePx+a.pos.AtVec(1),
//		size,
//		a.color,
//	)
//}

//func (a *Animal) changePos(animalsP *map[float64]map[float64]map[any]struct{}) {
//	x, y := a.pos.AtVec(0), a.pos.AtVec(1)
//	animals := *animalsP
//	delete(animals[y][x], a)
//	direction := mat.NewVecDense(
//		2,
//		[]float64{
//			float64((rand.Intn(3) - 1) * (tileSize + boardTilesGapWidth)),
//			float64((rand.Intn(3) - 1) * (tileSize + boardTilesGapWidth)),
//		},
//	)
//
//	a.pos.AddVec(a.pos, direction)
//	a.pos = teleportAtBoundary(a.pos)
//
//	x, y = a.pos.AtVec(0), a.pos.AtVec(1)
//	animals[y][x][a] = struct{}{}
//}

//func teleportAtBoundary(pos *mat.VecDense) *mat.VecDense {
//	// teleport at X boundaries
//	if pos.AtVec(0) > lastTilePx {
//		pos.SetVec(0, 0)
//	} else if pos.AtVec(0) < 0 {
//		pos.SetVec(0, lastTilePx)
//	}
//
//	// teleport at Y boundaries
//	if pos.AtVec(1) > lastTilePx {
//		pos.SetVec(1, 0)
//	} else if pos.AtVec(1) < 0 {
//		pos.SetVec(1, lastTilePx)
//	}
//	return pos
//}
