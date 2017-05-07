import {Component} from "@angular/core";
import { GlobalState } from '../../global.state';

@Component({
  selector: 'autotest',
  styleUrls: ['./autotest.scss'],
  templateUrl: './autotest.html'
})
export class AutoTest {

  constructor(private _state: GlobalState) {

  }

  ngOnInit() {
  }

}
