import {Component, Input, OnInit} from "@angular/core";

import {NgbActiveModal} from "@ng-bootstrap/ng-bootstrap";

import {TreeModel, TreeOptions} from "../../ng2-tree";

import {CaseSelectionService} from "./case-selection.service";
import {TreeService} from "../../ng2-tree/src/tree.service";
import {SuiteService} from "../../../service/suite";

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

  query: any = {keywords: '', status: ''};

  public options: TreeOptions = {
    usage: 'selection',
    isExpanded: false,
    nodeName: '用例',
    folderName: '模块'
  }
  public tree: TreeModel;

  constructor(public activeModal: NgbActiveModal, private _treeService: TreeService,
              public _sutieService: SuiteService, public _caseService: CaseSelectionService,) {
  }

  ngOnInit(): any {
    this.loadData();
  }

  loadData() {
    let that = this;
    that._sutieService.query(that.query).subscribe((json: any) => {
      that.tree = json.data;
      // CONSTANT.CUSTOM_FIELD_FOR_PROJECT = json.customFields;
      // this._state.notifyDataChanged('title.change', '测试用例');
    });
  }

  save(): any {
    this.activeModal.close('save');
  }

  dismiss(): any {
    this.activeModal.dismiss('cancel');
  }
}
