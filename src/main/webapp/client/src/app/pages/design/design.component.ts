import {Component} from "@angular/core";
import { Router, ActivatedRoute, Params } from '@angular/router';
import { GlobalState } from '../../global.state';

@Component({
  selector: 'design',
  styleUrls: ['./design.scss'],
  templateUrl: './design.html'
})
export class Design {
  prjId: number;

  constructor(private _route: ActivatedRoute) {

  }

  ngOnInit() {

  }

}
