import * as Pixi from 'pixi.js'
import { KeyboardHandler } from '../input';
import { Character } from '../characters/character';
import { Direction } from '../characters/direction';
import { Config } from '../config/config';

/**
 * Game is our root class for handling all game client activities.
 */
export class Game {
    app: Pixi.Application
    keyboardInput: KeyboardHandler
    player: Character

    characterLayer: Pixi.Container

    constructor() {
        this.app = new Pixi.Application();
        this.keyboardInput = new KeyboardHandler()
        this.player = new Character()

        this.app.stage.addChild(drawGrid())

        this.player.position.x = 8
        this.player.position.y = 8
        this.player.direction = Direction.Right


        this.characterLayer = new Pixi.Container()
        this.characterLayer.isRenderGroup = true
        this.characterLayer.addChild(this.player.shape)
        this.app.stage.addChild(this.characterLayer)

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
                console.info("Game Update: ", k)
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

    valueInput(name: string, value: string) {
        switch (name) {
            case "speed":
                this.player.speed = parseInt(value)
                break;
        }
    }
}

function drawGrid(): Pixi.Graphics {
    const grid = new Pixi.Graphics()
    grid.setStrokeStyle({
        width: 1,
        color: 0x333333,
    })

    for (let c = 0; c <= 17; c++) {
        grid.moveTo(c * Config.size, 0)
        grid.stroke(0x333333)
        grid.lineTo(c * Config.size, Config.rows * Config.size)
    }
    for (let r = 0; r <= 17; r++) {
        grid.moveTo(0, r * Config.size)
        grid.stroke(0x333333)
        grid.lineTo(Config.size * Config.columns, r * Config.size)
    }

    return grid
}
