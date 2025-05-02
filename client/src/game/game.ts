import * as Pixi from 'pixi.js'
import { KeyboardHandler } from '../input';
import { Character } from '../characters/character';
import { Direction } from '../characters/direction';
import { Config } from '../config/config';
import { MovementInput } from '../input/movement_input';

/**
 * Game is our root class for handling all game client activities.
 */
export class Game {
    app: Pixi.Application
    keyboardInput: KeyboardHandler
    player: Character
    movementInput: MovementInput

    characterLayer: Pixi.Container

    constructor() {
        this.app = new Pixi.Application();
        this.keyboardInput = new KeyboardHandler()
        this.player = new Character()
        this.movementInput = new MovementInput(this.player.movement)

        this.app.stage.addChild(drawGrid())

        this.player.position.x = 8
        this.player.position.y = 8
        this.player.movement.direction = Direction.Right

        this.characterLayer = new Pixi.Container()
        this.characterLayer.isRenderGroup = true
        this.characterLayer.addChild(this.player.shape)
        this.app.stage.addChild(this.characterLayer)
    }

    async start() {
        this.keyboardInput.start()
        this.movementInput.start()

        const width: number = 32 * 17
        const height: number = 32 * 17

        await this.app.init({ width: width, height: height })
        this.app.ticker.add(this.update.bind(this))
    }

    update(ticker: Pixi.Ticker) {
        let delta = ticker.deltaMS

        this.movementInput.update()
        this.player.update(delta)
    }

    valueInput(name: string, value: string) {
        switch (name) {
            case "speed":
                this.player.movement.speed = parseInt(value)
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
        grid.moveTo(c * Config.tileSize, 0)
        grid.stroke(0x333333)
        grid.lineTo(c * Config.tileSize, Config.rows * Config.tileSize)
    }
    for (let r = 0; r <= 17; r++) {
        grid.moveTo(0, r * Config.tileSize)
        grid.stroke(0x333333)
        grid.lineTo(Config.tileSize * Config.columns, r * Config.tileSize)
    }

    return grid
}
