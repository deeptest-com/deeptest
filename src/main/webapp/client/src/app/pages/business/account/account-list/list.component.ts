import {Component, ViewEncapsulation} from '@angular/core';

import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import { CONSTANT } from '../../../../utils/constant';
import { RouteService } from '../../../../service/route';
import { AccountService } from '../../../../service/account';

@Component({
  selector: 'account-list',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./list.scss')],
  template: require('./list.html')
})
export class AccountList implements OnInit, AfterViewInit {
  totalItems:number = 0;
  currentPage:number = 1;
  itemsPerPage:number = 6;

  me: any = {};
  model: any = {status: ''};
  items: Array<any> = [];

  constructor(private _routeService: RouteService,
              private _accountService: AccountService) {

  }
  ngOnInit() {
    let that = this;
    that.loadData();
  }

  ngAfterViewInit() {

  }

  pageChanged(account:any):void {
    let that = this;
    that.currentPage = account.page;
    that.loadData();
  }
  create():void {
    let that = this;

    that._routeService.navTo("/pages/account/edit/null");
  }
  disable(accountId: number):void {
    let that = this;
    // that._accountService.disable(accountId).subscribe((json:any) => {
    //   that.loadData();
    // });
  }
  remove(accountId: number):void {
    let that = this;
    // that._accountService.remove(accountId).subscribe((json:any) => {
    //   that.loadData();
    // });
  }

  loadData() {
    let that = this;

    // that._accountService.list(that.itemsPerPage, that.currentPage, that.model.status).subscribe((json:any) => {
    //   that.totalItems = json.totalItems;
    //   that.items = json.data;
    // });
  }
}
