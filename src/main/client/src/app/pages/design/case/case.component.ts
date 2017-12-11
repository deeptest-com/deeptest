import {Component} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';

@Component({
  selector: 'case',
  styleUrls: ['./case.scss'],
  templateUrl: './case.html'
})
export class Case {
  projectId: number;
  key: number;

  contentHeight = Utils.getContainerHeight(110);

  constructor(private _route: ActivatedRoute) {

  }

  ngOnInit() {

  }

}
