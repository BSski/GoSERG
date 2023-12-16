<div markdown="1" align="center">    

![GoSERG preview](https://i.imgur.com/rZ3KmWN.gif)

</div>


<p align="center"><h1 align="center">GoSERG: A Natural Ecosystem Simulator</h1></p>

[//]: # ()
[//]: # (<div markdown="1" align="center">)

[//]: # ()
[//]: # (Demo Uptime)

[//]: # (CodeFactor)

[//]: # (Maintainability)

[//]: # (Code style: black)

[//]: # ()
[//]: # (</div>)


## Table of contents

* [Project description](#scroll-project-description)
* [How to use](#game_die-how-to-use)
* [Keybinds](#keyboard-keybinds)
* [Animal traits](#turtle-animal-traits)
* [Achievements](#trophy-achievements)
* [Future improvements](#rocket-future-improvements)
* [Contact](#telephone_receiver-contact)
* [Author](#construction_worker-author)
* [License](#unlock-license)


## :scroll: Project description

GoSERG is a fascinating real-time ecosystem simulation that puts you in the role of a curious observer
of a dynamic world of carnivores, herbivores and herbs. Born out of a Python project, GoSERG was entirely rewritten
in Go, further optimized and compiled into WASM to bring the simulation to life in your web browser at [bsski.github.io](bsski.github.io).

Witness the intricate dance of life in procedurally generated terrains, where animals endowed with
unique traits fight for survival. Analyze the ecosystem through various charts and indicators,
and challenge yourself with a series of achievements designed to add an extra layer of fun and engagement.

Despite some remaining tech debt, the project was launched swiftly to bring this unique experience
to you as soon as possible. Future improvements and features are on the horizon, ready to make GoSERG
even more immersive and captivating.

The goal of this project was to deliver "business" value as soon as possible, so the code has some minor
tech debt at the moment. However, the app is fully functional and ready to be explored.


## :game_die: How to use

Dive right into the world of GoSERG!
Adjust the simulation parameters to steer the course of the ecosystem and achieve your desired outcomes.
Controls include the ability to pause, unpause, time travel forward, reset the board, and slow down the simulation.
Feel free to tweak the settings even during the simulation - the world is at your fingertips.

You can access the web preview at [bsski.github.io](bsski.github.io) or download .exe version for free!


## :keyboard: Keybinds

- Space: Toggle pause
- Up arrow/Down arrow: Cycle through achievements
- R: Reset simulation (new terrain & animals)
- T: Skip 8 days
- 123: Change right panel option in the UI.


## :turtle: Animal traits

Each animal in GoSERG is (almost) unique, possessing distinct traits that influence its survival:

- Speed: Determines how frequently an animal moves.
- Legs Length: Influences the efficiency of an animal's movement, potentially reducing energy cost.
- Fat Limit: Sets the maximum energy an animal can store.
- Bowel Length: Defines how much energy an animal retrieves from its food.

Herbivores feed on the spontaneously growing herbs, while carnivores hunt down the herbivores, creating a captivating circle of life.


## :trophy: Achievements

Choose an achievement you wish to pursue, fulfill its conditions, and then go for another one!
Remember, you have to specify which achievement you're aiming for in order to unlock it.
This feature introduces an element of strategy and provides goals that guide your exploration of the ecosystem.


## :rocket: Future improvements

Although GoSERG is already a fully functional and engaging ecosystem simulation, the journey doesn't stop here.
There are exciting features on the roadmap:
- better sprites of animals,
- clickable animals with their own, personal stats like age or even a family tree,
- death/drowning/eating animations,
- more achievements.


## :telephone_receiver: Contact

- contact.bsski@gmail.com

## :construction_worker: Author

- @BSski

## :unlock: License

MIT
