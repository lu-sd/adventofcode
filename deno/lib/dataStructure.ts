export type Pt = {
  x: number;
  y: number;
};
export class Point {
  x: number;
  y: number;

  constructor(x: number, y: number) {
    this.x = x;
    this.y = y;
  }

  dist(other: Point): [number, number] {
    return [this.x - other.x, this.y - other.y];
  }
  move(dir: Point) {
    return new Point(this.x + dir.x, this.y + dir.y);
  }
  get id() {
    return `${this.x}:${this.y}`;
  }
}

export class Grid<T> {
  grid: T[][];
  nrow: number;
  ncol: number;
  constructor(array: T[][]) {
    this.grid = array;
    this.nrow = array.length;
    this.ncol = array[0].length;
  }
  isInside(r: number, c: number): boolean {
    return r < this.nrow && r >= 0 && c < this.ncol && c >= 0;
  }

  isPInside(p: Point): boolean {
    return p.x < this.nrow && p.x >= 0 && p.y < this.ncol && p.y >= 0;
  }
  getVal(r: number, c: number): T {
    return this.grid[r][c];
  }
  getPVal(p: Point): T {
    return this.grid[p.x][p.y];
  }

  getId(r: number, c: number): string {
    return `${r}:${c}`;
  }
}

export const Dirs4 = [
  new Point(-1, 0),
  new Point(0, 1),
  new Point(1, 0),
  new Point(0, -1),
];
