import {Component} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';

@Component({
  selector: 'plan',
  styleUrls: ['./plan.scss'],
  templateUrl: './plan.html'
})
export class Plan {

  constructor(private _route: ActivatedRoute) {

  }

  ngOnInit() {

  }

}
