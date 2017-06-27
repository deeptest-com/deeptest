import { Input, Component, OnInit, EventEmitter, Output, Inject, OnChanges, SimpleChanges } from '@angular/core';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';

import { Prop } from './field.prop';
import { FieldChangedEvent } from './field.events';

import { CustomFieldService } from './custom-field.service';

@Component({
  selector: 'custom-field',
  templateUrl: './custom-field.html',
  styleUrls: ['./styles.scss'],
  providers: [CustomFieldService]
})
export class CustomFieldComponent implements OnInit, OnChanges {

  @Input()
  public prop: Prop;

  @Output()
  public fieldChanged: EventEmitter<any> = new EventEmitter();

  // public tree: Tree;

  public constructor(@Inject(CustomFieldService) private customFieldService: CustomFieldService) {

  }

  public ngOnChanges(changes: SimpleChanges): void {
    // if (!this.treeModel) {
    //   this.tree = TreeComponent.EMPTY_TREE;
    // } else {
    //   this.tree = Tree.buildTreeFromModel(this.treeModel);
    // }
  }

  public ngOnInit(): void {
    // this.treeService.nodeRenamedRemote$.subscribe((e: NodeEvent) => {
    //   this.nodeRenamedRemote.emit(e);
    // });

  }
}
