import { Input, Component, OnInit, ElementRef, Inject } from '@angular/core';
import { FormControl } from '@angular/forms';

import {TreeSettings} from "./type/tree.settings";
import {TreeOptions} from "./type/tree.options";
import { Tree } from './tree';

import { NodeMenuService } from './menu/node-menu.service';
import { TreeService } from './tree.service';

@Component({
  selector: 'tree-toolbar',
  template: `
  <div class="tree-toolbar">
    <form *ngIf="options.usage == 'design'" class="form-inline">
      <label (click)="tree.expandOrNot(options)" class="link no-underline mr-xs-2">
        <span *ngIf="!options.isExpanded">全部展开</span>
        <span *ngIf="options.isExpanded">全部收缩</span>
      </label>
      <input [value]="keywords" [formControl]="keywordsControl" name="keywords" type="search" class="form-control form-control-sm" placeholder="过滤">
    </form>
    <form *ngIf="options.usage == 'selection'" class="form-inline">
      <span (click)="tree.selectAll()" class="link no-underline">
        全选
      </span>
      <span (click)="tree.unselectAll()" class="link no-underline">
        全不选
      </span>
    </form>
  </div>
  `
})
export class TreeToolbarComponent implements OnInit {
  @Input()
  public tree: Tree;

  @Input()
  public options: TreeOptions;

  public keywords: string = '';
  keywordsControl = new FormControl();

  public constructor(@Inject(TreeService) private treeService: TreeService,
                     @Inject(ElementRef) public element: ElementRef) {

  }

  public ngOnInit(): void {
    this.keywordsControl.valueChanges.debounceTime(800).subscribe(values => this.onChange(values));
  }

  onChange(values) {
    this.options['keywords'] = values;
  }

}
