package main

// TODO: create string:int mapping here and store ints in events log instead of strings

func (g *game) addEvent(event string) {
	g.currentEvents = append(g.currentEvents, event)
}

func (g *game) cleanCurrentEvents() {
	g.currentEvents = []string{}
}
