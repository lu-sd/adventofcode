export type Pt = {
  r: number;
  c: number;
};
export class Point {
  r: number;
  c: number;

  constructor(r: number, c: number) {
    this.r = r;
    this.c = c;
  }

  dist(other: Point): [number, number] {
    return [this.r - other.r, this.c - other.c];
  }
  move(dir: Point) {
    return new Point(this.r + dir.r, this.c + dir.c);
  }

  equals(other: Point): boolean {
    return this.r === other.r && this.c === other.c;
  }
  get id() {
    return `${this.r}:${this.c}`;
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
    return p.r < this.nrow && p.r >= 0 && p.c < this.ncol && p.c >= 0;
  }
  getVal(r: number, c: number): T {
    return this.grid[r][c];
  }
  getPVal(p: Point): T {
    return this.grid[p.r][p.c];
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
