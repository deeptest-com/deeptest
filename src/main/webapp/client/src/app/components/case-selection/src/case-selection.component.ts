import * as _ from "lodash";

import {Component, Input, OnInit, AfterViewInit} from "@angular/core";
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import {NgbActiveModal} from "@ng-bootstrap/ng-bootstrap";

import {CONSTANT} from "../../../utils/constant";

import {ZtreeService} from "../../ztree/src/ztree.service";
import {SuiteService} from "../../../service/suite";
import {CaseService} from "../../../service/case";

import {CaseSelectionService} from "./case-selection.service";

@Component({
  selector: 'case-selection',
  templateUrl: './case-selection.html',
  styleUrls: ['./styles.scss']
})
export class CaseSelectionComponent implements OnInit {

  @Input() treeModel: any;
  @Input() treeSettings: any = {};
  @Input() users: any[] = [];

  @Input() progress: string = '0';
  @Input() color: string = '#209e91';
  @Input() height: string = '1px';
  @Input() show: boolean = true;

  form: FormGroup;
  _queryModel: any = {type: {}, priority: {}, createUsers: [], updateUsers: []};
  queryModel: any;

  public cases: string[];

  private value:any = ['Athens'];
  private _disabledV:string = '0';
  private disabled:boolean = false;

  constructor(public activeModal: NgbActiveModal, private fb: FormBuilder, private _treeService: ZtreeService,
              public _sutieService: SuiteService, public _caseService: CaseService,) {
    this.queryModel = this._queryModel;


  }

  ngOnInit(): any {
    this.buildForm();

    this.loadData();
  }

  loadData() {
    this._caseService.query(3).subscribe((json:any) => {
      this.cases = json.data;
    });
  }

  save(): any {
    console.log('queryModel', this.queryModel);
    // this.activeModal.close('save');
  }

  reset() {
    console.log('reset');
  }

  dismiss(): any {
    this.activeModal.dismiss('cancel');
  }

  onModuleSelected(event: any) {
    console.log('onNodeSelected', event);
  }
  onCaseSelected(item: any) {
    console.log('onCaseSelected', item);
  }

  buildForm(): void {
    this.form = this.fb.group(
      {
        'type': ['', []],
        'priority': ['', []],
        'estimate': ['', []],
        'createTime': ['', []],
        'updateTime': ['', []],
        'createUser': ['', []],
        'updateUser': ['', []]
      }, {}
    );

    this.form.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(data => this.query(data));
  }
  formErrors = [];
  validateMsg = {

  };

  query(data?: any) {
    let ztree = jQuery.fn.zTree.getZTreeObj('tree');
    console.log(this.queryModel, this.treeModel, this.queryModel);

    let nodes = ztree.getNodesByParam("isHidden", true);
    ztree.showNodes(nodes);

    let typeFilter: string[] = this.validFilter(this.queryModel.type);
    let priorityFilter: string[] = this.validFilter(this.queryModel.priority);
    let estimateFilter: string[] = this.queryModel.estimate?this.queryModel.estimate.split('-'): [];

    let createTimeFilter: number = this.queryModel.createTime * 24 * 60 * 60 * 1000;
    let updateTimeFilter: number = this.queryModel.updateTime * 24 * 60 * 60 * 1000;

    let createByFilter: string[] = this.queryModel.createUsers;
    let updateByFilter: string[] = this.queryModel.updateUsers;

    nodes = ztree.getNodesByFilter((node) => {
      console.log('===', createTimeFilter);
      console.log('===', new Date().getTime() - node.createTime);

      return !node.isParent && (
          ( typeFilter.length > 0 && _.indexOf(typeFilter, node.type) < 0 )
          || ( priorityFilter.length > 0 && _.indexOf(priorityFilter, node.priority) < 0 )
          || ( estimateFilter.length > 0 && (parseInt(node.estimate) < parseInt(estimateFilter[0]) || parseInt(node.estimate) > parseInt(estimateFilter[1])) )

          || ( createTimeFilter && (new Date().getTime() - node.createTime) > createTimeFilter )
          || ( updateTimeFilter && (new Date().getTime() - node.updateTime) > updateTimeFilter )

          || ( createByFilter.length > 0 && _.indexOf(createByFilter, node.createById) < 0 )
          || ( updateByFilter.length > 0 && _.indexOf(updateByFilter, node.updateById) < 0 )
        );
    });
    ztree.hideNodes(nodes);
  }

  resetFilters() {
    this.queryModel = this._queryModel;
  }

  private get disabledV():string {
    return this._disabledV;
  }

  private set disabledV(value:string) {
    this._disabledV = value;
    this.disabled = this._disabledV === '1';
  }

  public selected(item:any, type: string):void {
    this.queryModel[type+'Users'].push(item.id);
    this.query();
  }

  public removed(item:any, type: string):void {
    this.queryModel[type+'Users'].splice(this.queryModel[type+'Users'].indexOf(item.id), 1);
    this.query();
  }

  public refreshValue(value:any):void {
    this.value = value;
  }

  public itemsToString(value:Array<any> = []):string {
    return value
      .map((item:any) => {
        return item.text;
      }).join(',');
  }

  public validFilter(obj: any): string[] {
    let arr:string[] = [];

    for(var i in obj){
      if (obj[i]) {
        arr.push(i);
      }
    }
    return arr;
  }

}
