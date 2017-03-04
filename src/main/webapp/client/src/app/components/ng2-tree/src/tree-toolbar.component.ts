import { Input, Component, OnInit, ElementRef, Inject } from '@angular/core';
import { Ng2TreeSettings, Ng2TreeOptions } from './tree.types';
import { Subject, Observable } from 'rxjs/Rx';
import { Tree } from './tree';
import {
  NodeRemovedEvent,

  NodeRenamedEvent,

  NodeCreatedEvent,

  NodeMovedEvent

} from './tree.events';

import { NodeMenuService } from './menu/node-menu.service';
import { TreeService } from './tree.service';

@Component({
  selector: 'tree-toolbar',
  template: `
  <div class="tree-toolbar">
    <div class="buttons">
      <a (click)="tree.expandOrNot(options)" href="javascript:void(0);">
        <span *ngIf="!options.isExpanded">展开全部</span>
        <span *ngIf="options.isExpanded">收缩全部</span>
      </a>
    </div>
  </div>
  `
})
export class TreeToolbarComponent implements OnInit {
  @Input()
  public tree: Tree;

  @Input()
  public settings: Ng2TreeSettings;

    @Input()
    public options: Ng2TreeOptions;


  public constructor(@Inject(NodeMenuService) private nodeMenuService: NodeMenuService,
                     @Inject(TreeService) private treeService: TreeService,
                     @Inject(ElementRef) public element: ElementRef) {

  }

  public ngOnInit(): void {

  }

}
