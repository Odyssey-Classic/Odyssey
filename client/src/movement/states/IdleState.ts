import { Movement } from "../movement";
import { Direction } from "../../characters/direction";
import { MovementState } from "./MovementState";
import { TurningState } from "./TurningState";
import { MovingState } from "./MovingState";
import { Character } from "../../characters/character";

export class IdleState implements MovementState {
    private movement: Movement;

    constructor(movement: Movement) {
        this.movement = movement;
    }

    update(character: Character, deltaMS: number): void {
        if (this.movement.skipTurn > 0) {
            this.movement.skipTurn--;
        }
    }

    startMove(character: Character, direction: Direction): void {
        if (this.movement.direction !== direction) {
            this.movement.direction = direction;
            if (this.movement.skipTurn <= 0) {
                this.movement.setState(new TurningState(this.movement));
                return;
            }
        }

        if (this.movement.canMove(character, direction)) {
            this.movement.setState(new MovingState(this.movement));
        }
    }

    stopMove(character: Character): void {
        // No specific behavior for stopping in IdleState
    }
}
