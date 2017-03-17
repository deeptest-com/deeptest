import {Component} from '@angular/core';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';

@Component({
  selector: 'case',
  styles: [require('./case.scss')],
  template: require('./case.html')
})
export class Case {

  contentHeight = Utils.getContainerHeight(105);

  constructor() {
  }

  ngOnInit() {
  }

}
