import { Movement } from "../movement";
import { MovementState } from "./MovementState";
import { IdleState } from "./IdleState";
import { Config } from "../../config/config";
import { Direction } from "../../characters/direction";
import { Character } from "../../characters/character";

export class MovingState implements MovementState {
    private movement: Movement;

    constructor(movement: Movement) {
        this.movement = movement;
        this.movement.moveShifted = false;
        this.movement.moveTimeMS = this.movement.speed;
    }

    update(character: Character, deltaMS: number): void {
        const d = this.movement.deltaUnit(this.movement.direction);

        this.movement.moveTimeMS -= deltaMS;
        if (this.movement.moveTimeMS <= this.movement.speed / 2 && !this.movement.moveShifted) {
            this.movement.shiftPosition(character, d);
            this.movement.moveShifted = true;
        }

        this.movement.offset.x += d.x * (deltaMS / this.movement.speed) * Config.tileSize;
        this.movement.offset.y += d.y * (deltaMS / this.movement.speed) * Config.tileSize;

        if (this.movement.moveTimeMS <= 0) {
            this.movement.setState(new IdleState(this.movement));
            this.movement.skipTurn = 2;
            this.movement.offset.x = 0;
            this.movement.offset.y = 0;
        }
    }

    startMove(character: Character, direction: Direction): void {
        // No specific behavior for starting a move in MovingState
    }

    stopMove(character: Character): void {
        // No specific behavior for stopping in MovingState
    }
}
