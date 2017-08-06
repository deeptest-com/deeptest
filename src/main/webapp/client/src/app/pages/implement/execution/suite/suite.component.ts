import {Component, ViewEncapsulation, OnInit, AfterViewInit} from "@angular/core";

import {
  NodeEvent,
  NodeMovedRemoteEvent,
  NodeRemovedRemoteEvent,
  NodeCreatedEvent,
  NodeRenamedEvent,
  NodeSelectedEvent,
  TreeModel,
  TreeOptions
} from "../../../../components/ng2-tree";

import {GlobalState} from "../../../../global.state";
import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {SlimLoadingBarService} from "../../../../components/ng2-loading-bar";
import {TreeService} from "../../../../components/ng2-tree/src/tree.service";
import {SuiteService} from "../../../../service/suite";

@Component({
  selector: 'execution-suite',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./suite.scss', '../../../../components/ng2-tree/src/styles.scss'],
  templateUrl: './suite.html'
})
export class ExecutionSuite implements OnInit, AfterViewInit {
  query:any = {keywords: '', status: ''};

  public options: TreeOptions = {
    usage: 'exe',
    isExpanded: true,
    nodeName: '用例',
    folderName: '模块'
  }
  public tree:TreeModel;

  constructor(private _routeService:RouteService, private _state:GlobalState,
              private _treeService:TreeService, private _sutieService: SuiteService,
              private slimLoadingBarService:SlimLoadingBarService) {
  }

  ngOnInit() {
    let that = this;
    that.loadData();
  }

  ngAfterViewInit() {

  }

  loadData() {
    let that = this;
    this.startLoading();
    that._sutieService.query(that.query).subscribe((json:any) => {
      that.tree = json.data;
      CONSTANT.CUSTOM_FIELD_FOR_PROJECT = json.customFields;

      that.completeLoading();
      this._state.notifyDataChanged('title.change', '测试用例');
    });
  }

  public onNodeSelected(e:NodeSelectedEvent):void {
    console.log('===', e);
    this._state.notifyDataChanged('exe.suite.change', {id: e.node.node.id, tm: new Date().getTime()});
  }

  startLoading() {
    this.slimLoadingBarService.start(() => {
      console.log('Loading complete');
    });
  }

  completeLoading() {
    let that = this;
    setTimeout(function () {
      that.slimLoadingBarService.complete();
    }, 500);
  }

}

