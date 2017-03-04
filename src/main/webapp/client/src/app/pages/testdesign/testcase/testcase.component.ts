import {Component} from '@angular/core';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';

@Component({
  selector: 'testcase',
  styles: [require('./testcase.scss')],
  template: require('./testcase.html')
})
export class Testcase {

  contentHeight = Utils.getContainerHeight(100);

  constructor() {
  }

  ngOnInit() {
  }

}
