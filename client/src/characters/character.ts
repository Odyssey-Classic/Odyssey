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

enum State {
    Idle = 0,
    Turning = 1,
    Moving = 2,
    Attacking = 3,
}

export class Character {
    public position: Position = new Position()
    private _direction: Direction = Direction.Right
    private offset: { x: number, y: number } = { x: 0, y: 0 }

    private _state: State = State.Idle

    private moveShifted: boolean = true
    private moveTimeMS: number = 0
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

    get state() {
        return this._state
    }
    set state(s: State) {
        this._state = s
        console.log("State Set:", s)
    }

    // Start Move Called
    // Check current direction vs new
    // Can Move check TODO
    // Set Moving True
    // if (moving), when we actually start moving, flag that we need to complete our move.
    //
    // Graphic Position = Game Pos * 32
    // Shapes drawn add the 16,16 to center the shape currently

    update(deltaMS: number) {
        switch (this.state) {
            case State.Turning:
                if (this.turnTimeMS > 0) {
                    this.turnTimeMS -= deltaMS
                }

                if (this.turnTimeMS <= 0) {
                    this.state = State.Idle
                }
                return
            case State.Moving:
                const d = DeltaUnit(this.direction)

                this.moveTimeMS -= deltaMS
                if (this.moveTimeMS <= 500 && !this.moveShifted) {
                    this.shiftPosition(d)
                    this.moveShifted = true
                }

                this.offset.x += (d.x * (deltaMS / 1000) * Config.size)
                this.offset.y += (d.y * (deltaMS / 1000) * Config.size)

                if (this.moveTimeMS <= 0) {
                    this.state = State.Idle
                    this.offset.x = 0
                    this.offset.y = 0
                }
                break
        }

        this.shape.position.x = this.position.x * 32 + this.offset.x
        this.shape.position.y = this.position.y * 32 + this.offset.y
    }

    protected shiftPosition(d: { x: number, y: number }) {
        console.info("Shift Position", d)
        this.offset.x *= -(Math.abs(d.x))
        this.offset.y *= -(Math.abs(d.y))
        this.position.x += d.x
        this.position.y += d.y
    }

    startMove(d: Direction) {
        if (this.state != State.Idle) {
            // If character isn't Idle, we aren't moving
            return
        }

        if (this.direction != d) {
            // We're going to turn character first
            this.turnTimeMS = turnTimeMS
            this.direction = d
            this.state = State.Turning
            return
        }

        this.state = State.Moving
        this.moveShifted = false
        this.moveTimeMS = 1000
    }

    stopMove() {
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
