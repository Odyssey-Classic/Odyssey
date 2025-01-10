import * as Pixi from "pixi.js"
import { Direction } from "./direction"
import { Position } from "./position"
import { DEG_TO_RAD } from "pixi.js"

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

    public container: Pixi.Container
    private shape: Pixi.Graphics

    constructor() {

        this.container = new Pixi.Container()
        this.container.isRenderGroup = true
        this.shape = new Pixi.Graphics()
        this.container.addChild(this.shape)

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

    update(deltaMS: number) {
        if (this.turnTimeMS > 0) {
            this.turnTimeMS -= deltaMS
            return
        }
        if (this.moving) {
            const d = DeltaUnit(this.direction)
            this.container.position.x += d.x * (deltaMS / 1000) * 32
            this.container.position.y += d.y * (deltaMS / 1000) * 32
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
            x: 16,
            y: 16,
        }
        this.shape.clear()
        this.shape.circle(center.x, center.y, 16)
        this.shape.fill(0xffffff)

        this.shape.moveTo(center.x, center.y)
        let arc = arcs[this.direction]
        this.shape.arc(center.x, center.y, 16, arc.start, arc.end, arc.counterClockwise)
        this.shape.fill(0x0000ff)

        // let d = DeltaUnit(this.direction)
        // this.shape.moveTo(center.x, center.y)
        // this.shape.lineTo(d.x * 24 + 16, d.y * 24 + 16)
        // this.shape.stroke({
        //     width: 3,
        //     color: 0xff0000,
        // })
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
