import { Input, Output, EventEmitter, Component, OnInit, AfterViewInit, OnChanges, ElementRef, Inject } from '@angular/core';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';

@Component({
  selector: 'table-tree',
  styleUrls: ['./styles.scss'],
  templateUrl: './table-tree.html'
})
export class TableTreeComponent implements OnInit, AfterViewInit{
  @Input()
  public models: any;
  @Input()
  public maxLevel: number;
  @Input()
  public keywords: string;

  isInit: boolean = false;

  constructor(private _state:GlobalState, private el: ElementRef) {
    let that = this;
  }

  public ngOnInit(): void {
  }

  ngAfterViewInit() {

  }
}
