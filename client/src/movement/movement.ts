import { Direction } from "../characters/direction";
import { Config } from "../config/config";
import { MovementState } from "./states/MovementState";
import { IdleState } from "./states/IdleState";
import { Character } from "../characters/character";

export class Movement {
    private state: MovementState;
    public offset: { x: number, y: number } = { x: 0, y: 0 };
    public direction: Direction = Direction.Right;

    public moveShifted: boolean = true;
    public moveTimeMS: number = 0;
    public turnTimeMS: number = 0;
    public skipTurn: number = 0;

    public speed: number = 150;

    constructor() {
        this.state = new IdleState(this);
    }

    setState(state: MovementState): void {
        this.state = state;
    }

    update(character: Character, deltaMS: number): void {
        this.state.update(character, deltaMS);
    }

    startMove(character: Character, direction: Direction): void {
        this.state.startMove(character, direction);
    }

    stopMove(character: Character): void {
        this.state.stopMove(character);
    }

    shiftPosition(character: Character, d: { x: number; y: number }): void {
        // flips the offset to be relative to the next tile position
        this.offset.x *= -(Math.abs(d.x));
        this.offset.y *= -(Math.abs(d.y));
        character.position.x += d.x;
        character.position.y += d.y;
    }

    // TODO: canMove can only check for map edges currently.
    // We'll need some way to check for other barriers.
    canMove(character: Character, dir: Direction): boolean {
        const d = this.deltaUnit(dir);
        let next = {
            x: d.x + character.position.x,
            y: d.y + character.position.y,
        };
        return !(next.x < 0 || next.x >= Config.columns || next.y < 0 || next.y >= Config.rows);
    }

    deltaUnit(d: Direction): { x: number; y: number } {
        // Uses math to turn the direction enum values to unit x,y changes
        // because I'm special and I wanted to do it this way rather than
        // with a switch.
        return {
            x: ((d - 3) % 2),
            y: ((d - 2) % 2),
        };
    }
}
