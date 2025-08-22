import * as Pixi from "pixi.js";
import { Direction } from "./direction";
import { Position } from "./position";
import { Config } from "../config/config";
import { Movement } from "../movement/movement";

const arcs = {
    [Direction.Up]: { start: 225 * Pixi.DEG_TO_RAD, end: 315 * Pixi.DEG_TO_RAD, counterClockwise: false },
    [Direction.Down]: { start: 45 * Pixi.DEG_TO_RAD, end: 135 * Pixi.DEG_TO_RAD, counterClockwise: false },
    [Direction.Left]: { start: 135 * Pixi.DEG_TO_RAD, end: 225 * Pixi.DEG_TO_RAD, counterClockwise: false },
    [Direction.Right]: { start: 315 * Pixi.DEG_TO_RAD, end: 45 * Pixi.DEG_TO_RAD, counterClockwise: false },
};

export class Character {
    public position: Position = new Position();
    public shape: Pixi.Graphics;
    public movement: Movement;

    constructor() {
        this.shape = new Pixi.Graphics();
        this.movement = new Movement();
        this.makeShape();
    }

    update(deltaMS: number) {
        this.movement.update(this, deltaMS);

        const d = this.movement.deltaUnit(this.movement.direction);
        if (this.movement.moveShifted) {
            this.movement.shiftPosition(this, d);
            this.movement.offset.x = 0;
            this.movement.offset.y = 0;
        } else {
            this.movement.offset.x += d.x * (deltaMS / this.movement.speed) * Config.tileSize;
            this.movement.offset.y += d.y * (deltaMS / this.movement.speed) * Config.tileSize;
        }

        this.makeShape();
        this.shape.position.x = this.position.x * Config.tileSize + this.movement.offset.x;
        this.shape.position.y = this.position.y * Config.tileSize + this.movement.offset.y;
    }

    protected makeShape() {
        let center = {
            x: Config.tileSize / 2,
            y: Config.tileSize / 2,
        };
        this.shape.clear();
        this.shape.circle(center.x, center.y, Config.tileSize / 2);
        this.shape.fill(0xffffff);

        this.shape.moveTo(center.x, center.y);
        let arc = arcs[this.movement.direction];
        this.shape.arc(center.x, center.y, Config.tileSize / 2, arc.start, arc.end, arc.counterClockwise);
        this.shape.fill(0x00ffff);
    }
}
