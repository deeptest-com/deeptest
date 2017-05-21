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
    <form class="form-inline">
      <label (click)="tree.expandOrNot(options)" class="link no-underline mr-xs-2">
        <span *ngIf="!options.isExpanded">展开全部</span>
        <span *ngIf="options.isExpanded">收缩全部</span>
      </label>
      <input [value]="keywords" [formControl]="keywordsControl" name="keywords" type="search" class="form-control form-control-sm" placeholder="过滤">
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
