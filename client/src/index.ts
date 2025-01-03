import { Game } from './game'

let game = new Game()

// puts our game object at the root of our page to make console scripting possible.
window['game'] = game

game.start().then(() => {
    console.log(game)
    document.body.appendChild(game.app.canvas);
})
