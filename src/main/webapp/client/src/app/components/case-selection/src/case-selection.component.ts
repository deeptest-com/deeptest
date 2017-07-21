import {Component, Input, OnInit} from "@angular/core";
import {NgbActiveModal} from "@ng-bootstrap/ng-bootstrap";

import {TreeModel, TreeOptions} from "../../ng2-tree";

import {TreeService} from "../../ng2-tree/src/tree.service";
import {SuiteService} from "../../../service/suite";
import {CaseService} from "../../../service/case";

import {CaseSelectionService} from "./case-selection.service";

@Component({
  selector: 'case-selection',
  templateUrl: './case-selection.html',
  styleUrls: ['./styles.scss']
})
export class CaseSelectionComponent implements OnInit {

  @Input() progress: string = '0';
  @Input() color: string = '#209e91';
  @Input() height: string = '1px';
  @Input() show: boolean = true;

  queryModel: any = {type: 'functional', priority: '0', estimate: '', createBy: '', createOn: '', updateBy: '', updateOn: '' };

  public options: TreeOptions = {
    usage: 'selection',
    isExpanded: false,
    nodeName: '用例',
    folderName: '模块'
  }
  public tree: TreeModel;
  public cases: string[];

  constructor(public activeModal: NgbActiveModal, private _treeService: TreeService,
              public _sutieService: SuiteService, public _caseService: CaseService,) {
  }

  ngOnInit(): any {
    this.loadData();
  }

  loadData() {
    let that = this;
    that._sutieService.query(that.queryModel).subscribe((json: any) => {
      that.tree = json.data;
      // CONSTANT.CUSTOM_FIELD_FOR_PROJECT = json.customFields;
      // this._state.notifyDataChanged('title.change', '测试用例');
    });

    that._caseService.query(3).subscribe((json:any) => {
      that.cases = json.data;
    });
  }

  save(): any {
    this.activeModal.close('save');
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

}
