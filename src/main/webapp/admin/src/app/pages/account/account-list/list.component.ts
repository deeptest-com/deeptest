import {Component, ViewEncapsulation} from '@angular/core';

import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import { CONSTANT } from '../../../utils/constant';
import { RouteService } from '../../../service/route';
import { AccountService } from '../../../service/account';

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

  model: any = {status: ''};
  statusMap: Array<any> = CONSTANT.AccountStatus;
  accounts: Array<any> = [];

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

    that._routeService.navTo("/pages/account/edit/null/property");
  }
  statusChange(e: any):void {
    let that = this;
    that.model.status = e;
    that.loadData();
  }
  disable(accountId: string):void {
    let that = this;
    console.log('accountId=' + accountId);
  }
  delete(accountId: string):void {
    let that = this;
    console.log('accountId=' + accountId);
  }

  loadData() {
    let that = this;
    
    that._accountService.list(that.itemsPerPage, that.currentPage, that.model.status).subscribe((json:any) => {
      that.totalItems = json.totalItems;
      that.accounts = json.accounts;
    });
  }
}
