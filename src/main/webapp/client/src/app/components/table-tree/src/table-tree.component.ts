import { Input, Output, EventEmitter, Component, OnInit, AfterViewInit, OnChanges, ElementRef, Inject } from '@angular/core';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';

@Component({
  selector: 'table-tree',
  styles: [require('./styles.scss')],
  template: require('./table-tree.html')
})
export class TableTreeComponent implements OnInit, AfterViewInit, OnChanges{
  @Input()
  public models: any;
  @Input()
  public maxLevel: number;

  isInit: boolean = false;

  constructor(private _state:GlobalState, private el: ElementRef) {
    let that = this;
  }

  public ngOnInit(): void {
  }

  ngAfterViewInit() {

  }

  ngOnChanges() {
    let that = this;
    
  }

}
