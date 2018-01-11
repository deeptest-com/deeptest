import {Component, OnInit, AfterViewInit} from "@angular/core";

import {Router} from "@angular/router";

import {GlobalState} from "../../../global.state";

import {CONSTANT} from "../../../utils/constant";
import {WS_CONSTANT} from '../../../utils/ws-constant';
import {RouteService} from "../../../service/route";
import {SockService} from "../../../service/sock";

import {OrgService} from "../../../service/org";
import {AccountService} from "../../../service/account";

@Component({
  selector: 'ba-page-top',
  templateUrl: './baPageTop.html',
  styleUrls: ['./baPageTop.scss']
})
export class BaPageTop implements OnInit, AfterViewInit {

  orgId: number;
  prjId: number;

  profile: any = {};
  project: any = {};
  projects: any[] = [];

  orgs: any[] = [];
  keywords: string;

  msgs: any[] = [];
  alerts: any[] = [];

  public isScrolled: boolean = false;
  public isMenuCollapsed: boolean = false;

  constructor(private _router: Router, private _state: GlobalState, private _routeService: RouteService,
              private sockService: SockService, private orgService: OrgService, private accountService: AccountService) {

    this._state.subscribe(CONSTANT.STATE_CHANGE_PROFILE, (profile) => {
      console.log(CONSTANT.STATE_CHANGE_PROFILE, profile);
      this.profile = profile;

      this.wsConnect();
    });

    this._state.subscribe(CONSTANT.STATE_CHANGE_ORGS, (data: any) => {
      console.log(CONSTANT.STATE_CHANGE_ORGS, data);
      if (data.currOrgId) {
        this.orgId = data.currOrgId;
      }
      if (data.orgs) {
        this.orgs = data.orgs;
      }
    });

    this._state.subscribe(CONSTANT.STATE_CHANGE_PROJECTS, (data) => {
      console.log(CONSTANT.STATE_CHANGE_PROJECTS, data);
      this.prjId = CONSTANT.CURR_PRJ_ID;

      if (data.recentProjects) {
        this.projects = data.recentProjects;
      }
    });

    this._state.subscribe('menu.isCollapsed', (isCollapsed) => {
      this.isMenuCollapsed = isCollapsed;
    });
  }

  ngOnInit() {
    this.orgId = CONSTANT.CURR_ORG_ID;
    this.prjId = CONSTANT.CURR_PRJ_ID;

  }
  ngAfterViewInit() {}

  public changeOrg(item: any) {
    this.orgService.setDefault(item.id, {disabled: false}).subscribe((json: any) => {
      if (json.code == 1) {
        this.orgId = item.id;

        CONSTANT.PROFILE.orgPrivilege = json.orgPrivilege;
        CONSTANT.PROFILE.projectPrivilege = json.projectPrivilege;

        this.accountService.changeMyOrgs(null, this.orgId, true);
        this.accountService.changeRecentProjects(json.recentProjects);
        this.accountService.changeCasePropertyMap(json.casePropertyMap);
      }
    });
  }

  public scrolledChanged(isScrolled) {
    this.isScrolled = isScrolled;
  }

  gotoModule(module: string) {
    let url = '';
    if (module == 'design') {
      url = '/pages/org/' + CONSTANT.CURR_ORG_ID + '/prj/' + CONSTANT.CURR_PRJ_ID + '/design/case';
    } else if (module == 'implement') {
      url = '/pages/org/' + CONSTANT.CURR_ORG_ID + '/prj/' + CONSTANT.CURR_PRJ_ID + '/implement/plan/list';
    }

    this._routeService.navTo(url);
  }

  logout() {
    this.accountService.logout();

    this._routeService.navTo('/pages/login');
  }

  onSearchKeywordChanged(e: any) {
    console.log('0-', this._router.url);
    if (this._router.url.indexOf('design/case') < 0) {
      console.log('1-', this.keywords);
      this._routeService.quickJump(this.keywords);
    } else {
      console.log('2-', this.keywords);
      this._state.notifyDataChanged('case.jump', this._routeService.caseIdForJump(this.keywords));
    }
  }

  wsConnect() {
    this.sockService.onMessage((json) => {

      if (json.code != 1) {
        console.log('ws error: ', json.code);
        return;
      }

      if (WS_CONSTANT.WS_MSG_AND_ALERT_LASTEST === json.type) {
        this.alerts = json.alerts;
        this.msgs = json.msgs;
      }

    });
    this.sockService.onOpen((e) => {
      this.sockService.send({
        type: WS_CONSTANT.WS_OPEN
      });
    });
    this.sockService.open();
  }

}
