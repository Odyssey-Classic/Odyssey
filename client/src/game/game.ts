import * as Pixi from 'pixi.js'
import { KeyboardHandler, KeyState } from '../input';

/**
 * Game is our root class for handling all game client activities.
 */
export class Game {
    app: Pixi.Application
    keyState: KeyState
    keyboardInput: KeyboardHandler

    constructor() {
        this.app = new Pixi.Application();

        this.keyState = new KeyState()
        this.keyboardInput = new KeyboardHandler(this.keyState)
    }

    async start() {
        this.keyboardInput.start()

        await this.app.init({ width: 640, height: 360 })
        this.app.ticker.add(this.update.bind(this))
    }

    update(ticker: Pixi.Ticker) {
        let delta = ticker.deltaMS

        let keys = this.keyState.getKeys(true)
        if (keys.length) {
            console.log(keys[0])
        }
    }
}
