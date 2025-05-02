import { Direction } from "../characters/direction"
import { Character } from "../characters/character"
import { Config } from "../config/config"

const turnTimeMS: number = 100

enum State {
    Idle = 0,
    Turning = 1,
    Moving = 2,
}

export class Movement {
    private character: Character
    private _state: State = State.Idle
    private _direction: Direction = Direction.Right
    public offset: { x: number, y: number } = { x: 0, y: 0 }

    private moveShifted: boolean = true
    private moveTimeMS: number = 0
    private turnTimeMS: number = 0
    private skipTurn: number = 0

    public speed: number = 150

    constructor(character: Character) {
        this.character = character
    }

    get state() {
        return this._state
    }

    set state(s: State) {
        this._state = s
    }

    get direction() {
        return this._direction
    }

    set direction(d: Direction) {
        this._direction = d
    }

    update(deltaMS: number) {
        switch (this.state) {
            case State.Idle:
                if (this.skipTurn != 0) {
                    this.skipTurn--
                }
                break
            case State.Turning:
                if (this.turnTimeMS > 0) {
                    this.turnTimeMS -= deltaMS
                }

                if (this.turnTimeMS <= 0) {
                    this.state = State.Idle
                    this.offset.x = 0
                    this.offset.y = 0
                }
                break
            case State.Moving:
                const d = this.deltaUnit(this.direction)

                this.moveTimeMS -= deltaMS
                if (this.moveTimeMS <= this.speed / 2 && !this.moveShifted) {
                    this.shiftPosition(d)
                    this.moveShifted = true
                }

                this.offset.x += d.x * (deltaMS / this.speed) * Config.tileSize
                this.offset.y += d.y * (deltaMS / this.speed) * Config.tileSize

                if (this.moveTimeMS <= 0) {
                    this.state = State.Idle
                    this.skipTurn = 2
                    this.offset.x = 0
                    this.offset.y = 0
                }
                break
        }
    }

    startMove(d: Direction) {
        if (this.state != State.Idle) {
            return
        }

        if (this.direction != d) {
            this.direction = d

            if (this.skipTurn <= 0) {
                this.turnTimeMS = turnTimeMS
                this.state = State.Turning
                return
            }
        }

        if (this.canMove(this.direction)) {
            this.state = State.Moving
            this.moveShifted = false
            this.moveTimeMS = this.speed
        }
    }

    stopMove() { }

    private shiftPosition(d: { x: number; y: number }) {
        this.offset.x *= -(Math.abs(d.x))
        this.offset.y *= -(Math.abs(d.y))
        this.character.position.x += d.x
        this.character.position.y += d.y
    }

    private canMove(dir: Direction) {
        const d = this.deltaUnit(dir)
        let next = {
            x: d.x + this.character.position.x,
            y: d.y + this.character.position.y,
        }
        return !(next.x < 0 || next.x >= Config.columns || next.y < 0 || next.y >= Config.rows)
    }

    private deltaUnit(d: Direction): { x: number; y: number } {
        return {
            x: ((d - 3) % 2),
            y: ((d - 2) % 2),
        }
    }
}
