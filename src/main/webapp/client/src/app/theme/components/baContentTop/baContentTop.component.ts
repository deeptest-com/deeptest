import {Component} from '@angular/core';

import {GlobalState} from '../../../global.state';

@Component({
  selector: 'ba-content-top',
  styles: [require('./baContentTop.scss')],
  template: require('./baContentTop.html'),
})
export class BaContentTop {

  public currentPath:string = '';
  public currentTitle:string = '';

  constructor(private _state:GlobalState) {
    this._state.subscribe('menu.change', (path) => {
      if (path) {
        this.currentPath = path;
      }
    });

    this._state.subscribe('title.change', (title) => {
      if (title) {
        this.currentTitle = title;
      }
    });
  }
}
