import { Component, Output, EventEmitter, Input } from '@angular/core';

import { Cell } from '../../../lib/data-set/cell';

export class DefaultEditor implements Editor {
  @Input() cell: Cell;
}

export interface Editor {
  cell: Cell;
}
