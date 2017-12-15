import {Component, ViewEncapsulation, OnInit, AfterViewInit, ViewChild, ElementRef} from "@angular/core";
import { FormBuilder, FormGroup } from '@angular/forms';

import {GlobalState} from "../../../../global.state";

import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {MsgService} from "../../../../service/msg";
import {AccountService} from "../../../../service/account";

@Component({
  selector: 'msg-list',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./list.scss'],
  templateUrl: './list.html'
})
export class MsgList implements OnInit, AfterViewInit {

  queryForm: FormGroup;
  queryModel:any = {keywords: '', disabled: 'false'};
  readMap: Array<any> = CONSTANT.EntityRead;

  models: any;
  collectionSize:number = 0;
  page:number = 1;
  pageSize:number = 6;

  constructor(private _routeService:RouteService, private _state:GlobalState, private fb: FormBuilder, private el: ElementRef,
              private msgService: MsgService, private accountService: AccountService) {
  }

  ngOnInit() {
    let that = this;

    that.queryForm = that.fb.group(
      {
        'isRead': ['', []],
        'keywords': ['', []]
      }, {}
    );

    that.loadData();
  }

  ngAfterViewInit() {
    this.queryForm.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(values => this.queryChange(values));
  }

  create():void {
    this._routeService.navTo('/pages/msg-admin/msg/edit/null');
  }

  markRead(item: any, index: number):void {
    this.msgService.markRead(item.id).subscribe((json:any) => {
      this.loadData();
    });
  }
  delete(item: any, index: number):void {
    this.msgService.delete(item.id).subscribe((json:any) => {
      this.loadData();
    });
  }

  loadData() {
    let that = this;

    that.msgService.list(that.queryModel, that.page, that.pageSize).subscribe((json:any) => {
      that.models = json.data;
      that.collectionSize = json.collectionSize;
    });
  }

  queryChange(values:any):void {
    this.loadData();
  }
  pageChange(event:any):void {
    this.page = event.page;
    this.loadData();
  }
}
