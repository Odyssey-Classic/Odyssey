import * as Pixi from 'pixi.js'
import { KeyboardHandler } from '../input';

/**
 * Game is our root class for handling all game client activities.
 */
export class Game {
    app: Pixi.Application
    keyboardInput: KeyboardHandler

    constructor() {
        this.app = new Pixi.Application();

        this.keyboardInput = new KeyboardHandler()
    }

    async start() {
        this.keyboardInput.start()

        await this.app.init({ width: 640, height: 360 })
        this.app.ticker.add(this.update.bind(this))
    }

    update(ticker: Pixi.Ticker) {
        let delta = ticker.deltaMS

        let keys = this.keyboardInput.getKeys(true)
        if (keys.length) {

        }
    }
}
