import {Component} from "@angular/core";
import { Router, ActivatedRoute, Params } from '@angular/router';
import { GlobalState } from '../../global.state';

@Component({
  selector: 'implement',
  styleUrls: ['./implement.scss'],
  templateUrl: './implement.html'
})
export class Implement {
  prjId: number;

  constructor(private _route: ActivatedRoute) {

  }

  ngOnInit() {

  }

}
