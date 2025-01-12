import * as Pixi from "pixi.js"
import { Direction } from "./direction"
import { Position } from "./position"
import { DEG_TO_RAD } from "pixi.js"
import { Config } from "../config/config"

const arcs = {
    [Direction.Up]: { start: 225 * DEG_TO_RAD, end: 315 * DEG_TO_RAD, counterClockwise: false },
    [Direction.Down]: { start: 45 * DEG_TO_RAD, end: 135 * DEG_TO_RAD, counterClockwise: false },
    [Direction.Left]: { start: 135 * DEG_TO_RAD, end: 225 * DEG_TO_RAD, counterClockwise: false },
    [Direction.Right]: { start: 315 * DEG_TO_RAD, end: 45 * DEG_TO_RAD, counterClockwise: false },
}

const turnTimeMS: number = 100

export class Character {
    public position: Position = new Position()
    private _direction: Direction = Direction.Right
    private offset: { x: number, y: number } = { x: 0, y: 0 }

    private moving: boolean = false
    private turnTimeMS: number = 0

    private speed: number = 1

    public shape: Pixi.Graphics

    constructor() {
        this.shape = new Pixi.Graphics()
        this.makeShape()
    }

    get direction() {
        return this._direction
    }
    set direction(d: Direction) {
        this._direction = d
        this.makeShape()
    }

    // Start Move Called
    // Check current direction vs new
    // Can Move check TODO
    // Set Moving True
    // 
    // Graphic Position = Game Pos * 32 + 16

    update(deltaMS: number) {
        if (this.turnTimeMS > 0) {
            this.turnTimeMS -= deltaMS
            return
        }
        if (this.moving) {
            const d = DeltaUnit(this.direction)
            this.shape.position.x += d.x * (deltaMS / 1000) * Config.size
            this.shape.position.y += d.y * (deltaMS / 1000) * Config.size
        }
    }

    startMove(d: Direction) {
        if (this.direction != d) {
            this.turnTimeMS = turnTimeMS
        }
        this.direction = d
        this.moving = true
    }

    stopMove() {
        this.moving = false
    }

    protected makeShape() {
        let center = {
            x: Config.size / 2,
            y: Config.size / 2,
        }
        this.shape.clear()
        this.shape.circle(center.x, center.y, Config.size / 2)
        this.shape.fill(0xffffff)

        this.shape.moveTo(center.x, center.y)
        let arc = arcs[this.direction]
        this.shape.arc(center.x, center.y, Config.size / 2, arc.start, arc.end, arc.counterClockwise)
        this.shape.fill(0x00ffff)
    }
}

// Yes, we could do this with a conditional block, but I like being special.
function getDeltaX(d: Direction): number {
    return ((d - 3) % 2)
}

// Yes, we could do this with a conditional block, but I like being special.
function getDeltaY(d: Direction): number {
    return ((d - 2) % 2)
}

function DeltaUnit(d: Direction): { x: number, y: number } {
    return {
        x: getDeltaX(d),
        y: getDeltaY(d),
    }
}
