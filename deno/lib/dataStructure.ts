export type Pt = {
  x: number;
  y: number;
};

export class Grid<T> {
  grid: T[][];
  nrow: number;
  nclo: number;
  constructor(array: T[][]) {
    this.grid = array;
    this.nrow = array.length;
    this.nclo = array[0].length;
  }
  isInside(r: number, c: number): boolean {
    return r < this.nrow && r >= 0 && c < this.nclo && c >= 0;
  }
  getChar(r: number, c: number) {
    return this.grid[r][c];
  }
}
