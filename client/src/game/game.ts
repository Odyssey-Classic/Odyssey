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
        this.app.stage.addChild(this.player.container)
    }

    async start() {
        this.keyboardInput.start()

        const width: number = 32 * 17
        const height: number = 32 * 17

        await this.app.init({ width: width, height: height })
        this.app.ticker.add(this.update.bind(this))
    }

    update(ticker: Pixi.Ticker) {
        let delta = ticker.deltaMS

        let keys = this.keyboardInput.getKeys(true)
        let moving: boolean = false
        if (keys.length) {
            keys.forEach((k) => {
                console.info(k)
                // TODO checking more keys once a single move key is found causes
                // them to conflict and locks character.
                // Let's see if anyone notices this in their play tests first.
                switch (k) {
                    case 'KeyW':
                        this.player.startMove(Direction.Up)
                        moving = true
                        break
                    case 'KeyS':
                        this.player.startMove(Direction.Down)
                        moving = true
                        break
                    case 'KeyA':
                        this.player.startMove(Direction.Left)
                        moving = true
                        break
                    case 'KeyD':
                        this.player.startMove(Direction.Right)
                        moving = true
                        break
                }
            })
        }

        this.player.update(delta)
        if (!moving) {
            this.player.stopMove()
        }
    }
}
