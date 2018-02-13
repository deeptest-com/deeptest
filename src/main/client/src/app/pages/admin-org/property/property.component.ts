import {Component, OnInit, AfterViewInit,OnDestroy, ViewChild} from "@angular/core";
import {Location} from '@angular/common';
import {Subscription} from 'rxjs/Rx';
import { RouteService } from '../../../service/route';

@Component({
  selector: 'property',
  styleUrls: ['./property.scss'],
  templateUrl: './property.html'
})
export class Property implements OnInit, AfterViewInit, OnDestroy {
  tab: string = 'case-type';
  status: string = 'list';
  @ViewChild('tabset') tabset;

  constructor(private location: Location, private _routeService: RouteService) {
    let path = this.location.path();
    let arr = path.split('property/')[1].split('/');
    this.tab = arr[0];
    this.status = arr[1];
    console.log('===', this.tabset, this.tab);
  }

  ngOnInit() {
  }
  ngAfterViewInit() {
    this.tabset.select(this.tab);
  }

  ngOnDestroy(): void {
  }

  tabChange(event: any) {
    this.tab = event.nextId;
    this._routeService.navTo("/pages/org-admin/property/" + this.tab + "/list");
  }

  create() {
    this.status = 'edit';
    this._routeService.navTo("/pages/org-admin/property/" + this.tab + "/edit/null");
  }
  back() {
    this.status = 'list';
    this._routeService.navTo("/pages/org-admin/property/" + this.tab + "/list");
  }

}
