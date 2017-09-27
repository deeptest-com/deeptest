import {Component, Input, OnInit} from "@angular/core";
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

  @Input() progress: string = '0';
  @Input() color: string = '#209e91';
  @Input() height: string = '1px';
  @Input() show: boolean = true;

  queryModel: any = {type: 'functional', priority: '0', estimate: '', createBy: '', createOn: '', updateBy: '', updateOn: '' };

  public cases: string[];

  constructor(public activeModal: NgbActiveModal, private _treeService: ZtreeService,
              public _sutieService: SuiteService, public _caseService: CaseService,) {
  }

  ngOnInit(): any {
    this.loadData();
  }

  loadData() {
    this._caseService.query(3).subscribe((json:any) => {
      this.cases = json.data;
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
