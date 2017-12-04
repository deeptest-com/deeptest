import {Component, ViewEncapsulation, OnInit, AfterViewInit, ViewChild, ElementRef} from "@angular/core";
import { FormBuilder, FormGroup } from '@angular/forms';
import {GlobalState} from "../../../../../global.state";

import {CONSTANT} from "../../../../../utils/constant";
import {Utils} from "../../../../../utils/utils";
import {RouteService} from "../../../../../service/route";
import {CasePriorityService} from "../../../../../service/case-priority";

@Component({
  selector: 'case-priority-list',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./list.scss'],
  templateUrl: './list.html'
})
export class CasePriorityList implements OnInit, AfterViewInit {

  models: any;

  constructor(private _routeService:RouteService, private _state:GlobalState, private fb: FormBuilder, private el: ElementRef,
              private casePriorityService: CasePriorityService) {
  }

  ngOnInit() {
    let that = this;

    that.loadData();
  }

  ngAfterViewInit() {
  }

  edit($event: any):void {
    let that = this;

    console.log($event);
  }
  delete($event: any):void {
    let that = this;

    console.log($event);
  }
  setDefault(item: any):void {
    this.casePriorityService.setDefault(item.id).subscribe((json:any) => {
      if (json.code == 1) {
        this.models = json.data;
      }
    });
  }

  loadData() {
    let that = this;

    that.casePriorityService.list().subscribe((json:any) => {
      that.models = json.data;
    });
  }

  up(item: any) {
    this.casePriorityService.changeOrder(item.id, 'up').subscribe((json:any) => {
      if (json.code == 1) {
        this.models = json.data;
      }
    });
  }
  down(item: any) {
    this.casePriorityService.changeOrder(item.id, 'down').subscribe((json:any) => {
      if (json.code == 1) {
        this.models = json.data;
      }
    });
  }

}
