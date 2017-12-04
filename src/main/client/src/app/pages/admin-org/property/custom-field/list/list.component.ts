import {Component, ViewEncapsulation, OnInit, AfterViewInit, ViewChild, ElementRef} from "@angular/core";
import { FormBuilder, FormGroup } from '@angular/forms';
import {GlobalState} from "../../../../../global.state";

import {CONSTANT} from "../../../../../utils/constant";
import {Utils} from "../../../../../utils/utils";
import {RouteService} from "../../../../../service/route";
import {CustomFieldService} from "../../../../../service/custom-field";

@Component({
  selector: 'custom-field-list',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./list.scss'],
  templateUrl: './list.html'
})
export class CustomFieldList implements OnInit, AfterViewInit {

  models: any[];

  constructor(private _routeService:RouteService, private _state:GlobalState, private fb: FormBuilder, private el: ElementRef,
              private customFieldService: CustomFieldService) {
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

  loadData() {
    let that = this;

    that.customFieldService.list().subscribe((json:any) => {
      that.models = json.data;
    });
  }

  up(item: any) {
    this.customFieldService.changeOrder(item.id, 'up').subscribe((json:any) => {
      if (json.code == 1) {
        this.models = json.data;
      }
    });
  }
  down(item: any) {
    this.customFieldService.changeOrder(item.id, 'down').subscribe((json:any) => {
      if (json.code == 1) {
        this.models = json.data;
      }
    });
  }

}
