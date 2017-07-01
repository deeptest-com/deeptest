import {Component} from "@angular/core";
import { GlobalState } from '../../global.state';

@Component({
  selector: 'Analysis',
  styleUrls: ['./analysis.scss'],
  templateUrl: './analysis.html'
})
export class Analysis {

  constructor(private _state: GlobalState) {

  }

  ngOnInit() {
  }

}
