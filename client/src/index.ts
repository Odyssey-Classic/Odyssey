import { Game } from './game'

let game = new Game()

// puts our game object at the root of our page to make console scripting possible.
window['game'] = game

game.start().then(() => {
    console.log(game)

    const gameView = document.getElementById("game-view")
    if (gameView) {
        gameView.appendChild(game.app.canvas)
    } else {
        document.body.appendChild(game.app.canvas)
    }
})
