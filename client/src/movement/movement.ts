import { Direction } from "../characters/direction";
import { Character } from "../characters/character";
import { Config } from "../config/config";
import { MovementState } from "./states/MovementState";
import { IdleState } from "./states/IdleState";

export class Movement {
    private character: Character;
    private state: MovementState;
    public offset: { x: number, y: number } = { x: 0, y: 0 };
    public direction: Direction = Direction.Right;

    public moveShifted: boolean = true;
    public moveTimeMS: number = 0;
    public turnTimeMS: number = 0;
    public skipTurn: number = 0;

    public speed: number = 150;

    constructor(character: Character) {
        this.character = character;
        this.state = new IdleState(this);
    }

    setState(state: MovementState): void {
        this.state = state;
    }

    update(deltaMS: number): void {
        this.state.update(deltaMS);
    }

    startMove(direction: Direction): void {
        this.state.startMove(direction);
    }

    stopMove(): void {
        this.state.stopMove();
    }

    shiftPosition(d: { x: number; y: number }): void {
        this.offset.x *= -(Math.abs(d.x));
        this.offset.y *= -(Math.abs(d.y));
        this.character.position.x += d.x;
        this.character.position.y += d.y;
    }

    canMove(dir: Direction): boolean {
        const d = this.deltaUnit(dir);
        let next = {
            x: d.x + this.character.position.x,
            y: d.y + this.character.position.y,
        };
        return !(next.x < 0 || next.x >= Config.columns || next.y < 0 || next.y >= Config.rows);
    }

    deltaUnit(d: Direction): { x: number; y: number } {
        return {
            x: ((d - 3) % 2),
            y: ((d - 2) % 2),
        };
    }
}
