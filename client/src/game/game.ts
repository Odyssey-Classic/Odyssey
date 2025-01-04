import * as Pixi from 'pixi.js'
import { KeyboardHandler } from '../input';
import { Character } from '../characters/character';
import { Direction } from '../characters/direction';

/**
 * Game is our root class for handling all game client activities.
 */
export class Game {
    app: Pixi.Application
    keyboardInput: KeyboardHandler
    player: Character

    constructor() {
        this.app = new Pixi.Application();
        this.keyboardInput = new KeyboardHandler()
        this.player = new Character()

        this.player.position.x = 100
        this.player.position.y = 100
        this.player.direction = Direction.Right
        this.app.stage.addChild(this.player)
    }

    async start() {
        this.keyboardInput.start()

        await this.app.init({ width: 200, height: 200 })
        this.app.ticker.add(this.update.bind(this))
    }

    update(ticker: Pixi.Ticker) {
        let delta = ticker.deltaMS

        let keys = this.keyboardInput.getKeys(true)
        if (keys.length) {
            keys.forEach((k) => {
                console.info(k)
                switch (k) {
                    case 'KeyW':
                        this.player.direction = Direction.Up
                        break
                    case 'KeyS':
                        this.player.direction = Direction.Down
                        break
                    case 'KeyA':
                        this.player.direction = Direction.Left
                        break
                    case 'KeyD':
                        this.player.direction = Direction.Right
                        break
                }
            })
        }
    }
}
